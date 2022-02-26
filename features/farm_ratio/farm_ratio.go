package farm_ratio

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"

	"safebsc/features/shared"
	"safebsc/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

type connectionClient struct {
	db mongo.DB
}

func sumValuePerFarm(userPortfolioLatest []UserPortfolioLatests) []DataFarmRatio {
	var result []DataFarmRatio
	for _, userPort := range userPortfolioLatest {
		isNewFarm := true
		for i := 0; i < len(result); i++ {
			if result[i].FarmId == userPort.FarmId {
				isNewFarm = false
				if s, err := strconv.ParseFloat(userPort.PortValue, 64); err == nil {
					result[i].Value += s
				}
			}
		}
		if isNewFarm {
			var value float64
			if s, err := strconv.ParseFloat(userPort.PortValue, 64); err == nil {
				value = s
			}
			result = append(result, DataFarmRatio{
				FarmId:   userPort.FarmId,
				FarmName: userPort.FarmName,
				Value:    value,
			})
		}
	}
	return result
}

func sortAndGroupData(dataFarmRatios []DataFarmRatio) ([]DataFarmRatio, float64) {
	sort.SliceStable(dataFarmRatios, func(i, j int) bool {
		return dataFarmRatios[i].Value > dataFarmRatios[j].Value
	})
	var total float64
	var result []DataFarmRatio
	for i, data := range dataFarmRatios {
		var value = data.Value
		total += value
		if i < 4 {
			result = append(result, DataFarmRatio{
				FarmName: data.FarmName,
				Value:    value,
			})
		} else if i == 4 {
			result = append(result, DataFarmRatio{
				FarmName: "Other",
				Value:    value,
			})
		} else {
			result[len(result)-1].Value += value
		}
	}
	return result, total
}

func prepareData(supportedFarm []shared.SupportedFarm, dataFarmRatio []DataFarmRatio, total float64) FarmRatioResponse {
	var data []DataFarmRatioResponse
	var totalPercent float64
	for i, dataFR := range dataFarmRatio {
		var farmIcon string
		for _, sf := range supportedFarm {
			if dataFR.FarmId == sf.OrderID {
				farmIcon = sf.FarmIcon
			}
		}
		if i == len(dataFarmRatio)-1 {
			data = append(data, DataFarmRatioResponse{
				FarmName: dataFR.FarmName,
				FarmIcon: farmIcon,
				Value:    fmt.Sprintf("%f", dataFR.Value),
				Percent:  fmt.Sprintf("%.2f", 100-totalPercent),
			})
		} else {
			percent := math.Floor((dataFR.Value/total)*10000) / 100
			data = append(data, DataFarmRatioResponse{
				FarmName: dataFR.FarmName,
				FarmIcon: farmIcon,
				Value:    fmt.Sprintf("%f", dataFR.Value),
				Percent:  fmt.Sprintf("%.2f", percent),
			})
			totalPercent += percent
		}
	}
	return FarmRatioResponse{
		Data: data,
	}
}

func getFarmRatio(c *gin.Context) {
	log.Println(">>>>> Start getFarmRatio")
	walletAddress := c.Param("wallet_address")
	if walletAddress == "" {
		log.Println("Missing wallet address parameter")
		c.String(http.StatusBadRequest, "Error")
	}

	mongoDB, err := mongo.Connect()
	conn := connectionClient{
		db: mongoDB,
	}

	if err != nil {
		log.Println("Mongo Connected Failed", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	log.Println(">>>>> getFarmRatio Mongo user_portfolio_latests")
	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_portfolio_latests")
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
	})

	log.Println(">>>>> After getFarmRatio Mongo user_portfolio_latests")
	var userPortfolioLatests []UserPortfolioLatests
	if err = collectionData.All(mongoDB.Ctx, &userPortfolioLatests); err != nil {
		log.Println("userPortfolioLatests Find Error: ", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	if len(userPortfolioLatests) == 0 {
		log.Println("userPortfolioLatests Empty")
		c.String(http.StatusBadRequest, "Not found")
	}

	supportedFarms, err := conn.getSupporttedFarms()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	log.Println(">> After supportedFarms Empty")
	dataValuePerFarm := sumValuePerFarm(userPortfolioLatests)
	data, total := sortAndGroupData(dataValuePerFarm)
	resp := prepareData(supportedFarms, data, total)

	log.Println(">> Before return contents")
	contents, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func (dbClient connectionClient) getSupporttedFarms() ([]shared.SupportedFarm, error) {
	var sf []shared.SupportedFarm
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("supported_farms")

	cur, err := collection.Find(dbClient.db.Ctx, bson.M{})
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(dbClient.db.Ctx) {
		var elem shared.SupportedFarm
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		sf = append(sf, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(dbClient.db.Ctx)
	return sf, err
}

func Handler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getFarmRatio(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
