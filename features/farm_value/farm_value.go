package farm_value

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"safebsc/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func sumValuePerFarm(farmId string, portValue string, farms []Farm) ([]Farm, float64) {
	var sumValue float64
	isNewFram := true
	for i := 0; i < len(farms); i++ {
		if farms[i].FarmId == farmId {
			isNewFram = false
			if s, err := strconv.ParseFloat(portValue, 64); err == nil {
				farms[i].SumOfValue += s

			}
		}
		sumValue += farms[i].SumOfValue
	}
	if isNewFram {
		var value float64
		if s, err := strconv.ParseFloat(portValue, 64); err == nil {
			value = s
		}
		farms = append(farms, Farm{
			FarmId:     farmId,
			SumOfValue: value,
		})
	}
	return farms, sumValue
}

func sumValueAtTime(userPorts []UserPortfolios) []DataFarmValue {
	var result []DataFarmValue
	for _, userPort := range userPorts {
		farmId := userPort.FarmId
		portValue := userPort.PortValue
		createdAt := userPort.CreatedAt
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), createdAt.Hour(), createdAt.Minute(), 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			farms, sumValue := sumValuePerFarm(farmId, portValue, result[len(result)-1].Farms)
			result[len(result)-1].Farms = farms
			result[len(result)-1].Value = sumValue
		} else {
			var value float64
			if s, err := strconv.ParseFloat(userPort.PortValue, 64); err == nil {
				value = s
			}
			var farms []Farm
			farms = append(farms, Farm{
				FarmId:      userPort.FarmId,
				SumOfValue:  value,
				TotalOfData: 1,
			})
			result = append(result, DataFarmValue{
				Date:  date,
				Value: value,
				Farms: farms,
			})
		}
	}
	return result
}

func avgValuePerFarm(farms []Farm, farms2 []Farm) ([]Farm, float64) {
	var avgValue float64
	var newFarm Farm
	for _, farm2 := range farms2 {
		isNewFram := true
		newFarm = Farm{}
		for i := 0; i < len(farms); i++ {
			if farms[i].FarmId == farm2.FarmId {
				isNewFram = false
				farms[i].SumOfValue += farm2.AvgValue
				farms[i].TotalOfData += 1
				farms[i].AvgValue = farms[i].SumOfValue / farms[i].TotalOfData
				avgValue += farms[i].AvgValue
			} else {
				newFarm = Farm{
					FarmId:      farm2.FarmId,
					SumOfValue:  farm2.SumOfValue,
					TotalOfData: 1,
				}
			}
		}
		if isNewFram && len(newFarm.FarmId) > 0 {
			farms = append(farms, newFarm)
			avgValue += newFarm.SumOfValue
		}
	}
	return farms, avgValue
}

func sumValuePerDay(dataFarmValues []DataFarmValue) []DataFarmValue {
	var result []DataFarmValue
	for _, dataFarmValue := range dataFarmValues {
		createdAt := dataFarmValue.Date
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			farms, avgValue := avgValuePerFarm(result[len(result)-1].Farms, dataFarmValue.Farms)
			result[len(result)-1].Farms = farms
			result[len(result)-1].Value = avgValue
		} else {
			result = append(result, DataFarmValue{
				Date:  date,
				Value: dataFarmValue.Value,
				Farms: dataFarmValue.Farms,
			})
		}
	}

	return result
}

func calculateTotalPercent(dataFarmValue []DataFarmValueResponse) float64 {
	var firstValue float64
	var lastValue float64
	if s, err := strconv.ParseFloat(dataFarmValue[0].Value, 64); err == nil {
		firstValue = s
	}
	if s, err := strconv.ParseFloat(dataFarmValue[len(dataFarmValue)-1].Value, 64); err == nil {
		lastValue = s
	}
	var percent float64 = 100
	if firstValue != 0 {
		percent = ((lastValue / firstValue) * 100) - 100
	}
	return percent
}

func prepareData(totalValue float64, dataFarmValues []DataFarmValue) FarmValueResponse {
	sort.SliceStable(dataFarmValues, func(i, j int) bool {
		return dataFarmValues[i].Date.Before(dataFarmValues[j].Date)
	})
	var data []DataFarmValueResponse
	for _, dataFarmValue := range dataFarmValues {
		data = append(data, DataFarmValueResponse{
			Date:  dataFarmValue.Date,
			Value: fmt.Sprintf("%.4f", dataFarmValue.Value),
		})
	}
	return FarmValueResponse{
		Total:        fmt.Sprintf("%f", totalValue),
		TotalPercent: fmt.Sprintf("%.2f", calculateTotalPercent(data)),
		Data:         data,
	}
}

func getFarmValue(c *gin.Context) {
	walletAddress := c.Param("wallet_address")
	if walletAddress == "" {
		log.Println("Missing wallet address parameter")
		c.String(http.StatusBadRequest, "Error")
	}

	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	collectionLatest := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_portfolio_latests")
	findOptionsLatest := options.Find()
	findOptionsLatest.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	collectionDataLatest, err := collectionLatest.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
	})

	var userPortfolioLatests []UserPortfolioLatests
	if err = collectionDataLatest.All(mongoDB.Ctx, &userPortfolioLatests); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	var totalValue float64 = 0
	for i := 0; i < len(userPortfolioLatests); i++ {
		if s, err := strconv.ParseFloat(userPortfolioLatests[i].PortValue, 64); err == nil {
			totalValue = totalValue + s
		}
	}

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_portfolios")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
	})

	var userPortfolios []UserPortfolios
	if err = collectionData.All(mongoDB.Ctx, &userPortfolios); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	if len(userPortfolios) == 0 {
		log.Println("Not found")
		c.String(http.StatusBadRequest, "Not found")
	}

	dataFarmValueAtTime := sumValueAtTime(userPortfolios)
	// dataFarmValue := sumValuePerDay(dataFarmValueAtTime)
	resp := prepareData(totalValue, dataFarmValueAtTime)

	contents, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getFarmValue(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
