package investment_report

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

func sumReward(rewards []Reward) float64 {
	var result float64 = 0
	for _, reward := range rewards {
		if value, err := strconv.ParseFloat(reward.Value, 64); err == nil {
			result += value
		}
	}
	return result
}

func sumValuePerFarm(farmId string, portValue string, rewards []Reward, farms []Farm) []Farm {
	isNewFram := true
	for i := 0; i < len(farms); i++ {
		if farms[i].FarmId == farmId {
			isNewFram = false
			if s, err := strconv.ParseFloat(portValue, 64); err == nil {
				farms[i].NetWorth += s
			}
			farms[i].Reward += sumReward(rewards)
		}
	}
	if isNewFram {
		var value float64
		if s, err := strconv.ParseFloat(portValue, 64); err == nil {
			value = s
		}
		farms = append(farms, Farm{
			FarmId:   farmId,
			NetWorth: value,
			Reward:   sumReward(rewards),
		})
	}
	return farms
}

func sumValueAtHour(userPorts []UserPortfolios) []DataInvestmentReport {
	var result []DataInvestmentReport
	for _, userPort := range userPorts {
		farmId := userPort.FarmId
		createdAt := userPort.CreatedAt
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), createdAt.Hour(), createdAt.Minute(), 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			farms := sumValuePerFarm(farmId, userPort.PortValue, userPort.Reward, result[len(result)-1].Farms)
			result[len(result)-1].Farms = farms
		} else {
			var netWorth float64
			if s, err := strconv.ParseFloat(userPort.PortValue, 64); err == nil {
				netWorth = s
			}
			var farms []Farm
			farms = append(farms, Farm{
				FarmId:   userPort.FarmId,
				NetWorth: netWorth,
				Reward:   sumReward(userPort.Reward),
			})
			result = append(result, DataInvestmentReport{
				Date:  date,
				Farms: farms,
			})
		}

	}
	return result
}

func avgValuePerFarm(farms []Farm, farms2 []Farm) ([]Farm, float64, float64) {
	var avgNetWorth float64
	var avgReward float64
	for _, farm2 := range farms2 {
		isNewFram := true
		var newFarm Farm
		for i := 0; i < len(farms); i++ {
			if farms[i].FarmId == farm2.FarmId {
				isNewFram = false
				farms[i].NetWorth += farm2.NetWorth
				farms[i].TotalOfData += 1
				avgNetWorth += (farms[i].NetWorth / farms[i].TotalOfData)
				avgReward += (farms[i].Reward / farms[i].TotalOfData)
			} else {
				newFarm = Farm{
					FarmId:      farm2.FarmId,
					NetWorth:    farm2.NetWorth,
					Reward:      farm2.Reward,
					TotalOfData: 1,
				}
			}
		}
		if isNewFram && len(newFarm.FarmId) > 0 {
			farms = append(farms, newFarm)
			avgNetWorth += (newFarm.NetWorth / newFarm.TotalOfData)
			avgReward += (newFarm.Reward / newFarm.TotalOfData)
		}
	}
	return farms, avgNetWorth, avgReward
}

func avgDataPerDay(dataInvestmentReportPerHour []DataInvestmentReport) []DataInvestmentReport {
	var result []DataInvestmentReport
	for _, dataPerHour := range dataInvestmentReportPerHour {
		createdAt := dataPerHour.Date
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			farms, avgNetWorth, avgReward := avgValuePerFarm(result[len(result)-1].Farms, dataPerHour.Farms)
			result[len(result)-1].Farms = farms
			result[len(result)-1].NetWorth = avgNetWorth
			result[len(result)-1].Reward = avgReward
		} else {
			result = append(result, DataInvestmentReport{
				Date:     date,
				NetWorth: dataPerHour.NetWorth,
				Reward:   dataPerHour.Reward,
				Farms:    dataPerHour.Farms,
			})
		}
	}

	return result
}

func prepareData(dataInvestmentReport []DataInvestmentReport, endDateTime time.Time) InvestmentReportResponse {
	sort.SliceStable(dataInvestmentReport, func(i, j int) bool {
		return dataInvestmentReport[i].Date.Before(dataInvestmentReport[j].Date)
	})
	startDateTime := dataInvestmentReport[0].Date
	if endDateTime.After(time.Now()) {
		endDateTime = time.Now()
	}
	startDate := time.Date(startDateTime.Year(), startDateTime.Month(), startDateTime.Day(), 0, 0, 0, 0, startDateTime.Location())
	endDate := time.Date(endDateTime.Year(), endDateTime.Month(), endDateTime.Day(), 0, 0, 0, 0, endDateTime.Location())
	var data []DataInvestmentReportResponse
	i := 0
	for !startDate.After(endDate) {
		valueNetWorth := "0"
		valueReword := "0"
		if i < len(dataInvestmentReport) && startDate == dataInvestmentReport[i].Date {
			valueNetWorth = fmt.Sprintf("%f", dataInvestmentReport[i].NetWorth)
			valueReword = fmt.Sprintf("%f", dataInvestmentReport[i].Reward)
			i++
		}
		data = append(data, DataInvestmentReportResponse{
			Date:     startDate,
			NetWorth: valueNetWorth,
			Reward:   valueReword,
		})
		startDate = startDate.AddDate(0, 0, 1)
	}
	return InvestmentReportResponse{
		Data: data,
	}
}

func getInvesterReport(c *gin.Context) {
	walletAddress := c.Param("wallet_address")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	if walletAddress == "" {
		log.Println("Missing wallet address parameter")
		c.String(http.StatusBadRequest, "Error")
	}

	startDateTime, err := time.Parse(time.RFC3339, string(startDate))
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	endDateTime, err := time.Parse(time.RFC3339, string(endDate))
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_portfolios")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
		"createdAt": bson.M{
			"$gt": startDateTime,
			"$lt": endDateTime,
		},
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

	dataInvestmentReportPerHour := sumValueAtHour(userPortfolios)
	dataInvestmentReport := avgDataPerDay(dataInvestmentReportPerHour)
	resp := prepareData(dataInvestmentReport, endDateTime)

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
		getInvesterReport(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
