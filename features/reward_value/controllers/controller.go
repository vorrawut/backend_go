package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"safebsc/features/reward_value/models"
	"safebsc/mongo"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct{}

func (u Controller) GetInvesterReport(c *gin.Context) {
	getInvesterReport(c)
}

func sumReward(rewards []models.Reward) float64 {
	var result float64 = 0
	for _, reward := range rewards {
		if value, err := strconv.ParseFloat(reward.Value, 64); err == nil {
			result += value
		}
	}
	return result
}

func sumValueAtHour(userPorts []models.UserPortfolios) []models.DataInvestmentReport {
	var result []models.DataInvestmentReport
	for _, userPort := range userPorts {
		createdAt := userPort.CreatedAt
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), createdAt.Hour(), createdAt.Minute(), 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			result[len(result)-1].Reward += sumReward(userPort.Reward)
		} else {
			result = append(result, models.DataInvestmentReport{
				Date:   date,
				Reward: sumReward(userPort.Reward),
			})
		}

	}
	return result
}

func calculateTotalPercent(dataFarmValue []models.DataInvestmentReportResponse) float64 {
	var firstValue float64
	var lastValue float64
	if s, err := strconv.ParseFloat(dataFarmValue[0].Reward, 64); err == nil {
		firstValue = s
	}
	if s, err := strconv.ParseFloat(dataFarmValue[len(dataFarmValue)-1].Reward, 64); err == nil {
		lastValue = s
	}
	var percent float64 = 100
	if firstValue != 0 {
		percent = ((lastValue / firstValue) * 100) - 100
	}
	return percent
}

func prepareData(totalValue float64, dataInvestmentReports []models.DataInvestmentReport) models.InvestmentReportResponse {
	sort.SliceStable(dataInvestmentReports, func(i, j int) bool {
		return dataInvestmentReports[i].Date.Before(dataInvestmentReports[j].Date)
	})
	var data []models.DataInvestmentReportResponse
	for _, dataInvestmentReport := range dataInvestmentReports {
		data = append(data, models.DataInvestmentReportResponse{
			Date:   dataInvestmentReport.Date,
			Reward: fmt.Sprintf("%.4f", dataInvestmentReport.Reward),
		})
	}
	return models.InvestmentReportResponse{
		Total:        fmt.Sprintf("%f", totalValue),
		TotalPercent: fmt.Sprintf("%.2f", calculateTotalPercent(data)),
		Data:         data,
	}
}

func getInvesterReport(c *gin.Context) {
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

	var userPortfolioLatests []models.UserPortfolioLatests
	if err = collectionDataLatest.All(mongoDB.Ctx, &userPortfolioLatests); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	var totalValue float64 = 0
	for i := 0; i < len(userPortfolioLatests); i++ {
		totalValue += sumReward(userPortfolioLatests[i].Reward)
	}

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_portfolios")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
	})

	var userPortfolios []models.UserPortfolios
	if err = collectionData.All(mongoDB.Ctx, &userPortfolios); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	if len(userPortfolios) == 0 {
		log.Println("Not found")
		c.String(http.StatusBadRequest, "Not found")
	}

	dataInvestmentReportPerHour := sumValueAtHour(userPortfolios)
	resp := prepareData(totalValue, dataInvestmentReportPerHour)
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
