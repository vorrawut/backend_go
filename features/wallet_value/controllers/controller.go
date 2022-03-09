package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"safebsc/features/wallet_value/models"
	"safebsc/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct{}

func (u Controller) getWalletValue(c *gin.Context) {
	getWalletValue(c)
}

func sumValueAtTime(userWallets []models.UserWallets) []models.DataWalletValue {
	var result []models.DataWalletValue
	for _, userPort := range userWallets {
		createdAt := userPort.CreatedAt
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), createdAt.Hour(), createdAt.Minute(), 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			if s, err := strconv.ParseFloat(userPort.Value, 64); err == nil {
				result[len(result)-1].SumOfValue += s
			}
		} else {
			var value float64
			if s, err := strconv.ParseFloat(userPort.Value, 64); err == nil {
				value = s
			}
			result = append(result, models.DataWalletValue{
				Date:       date,
				SumOfValue: value,
			})
		}

	}
	return result
}

func sumValuePerDay(dataWalletValues []models.DataWalletValue) ([]models.DataWalletValue, error) {
	var result []models.DataWalletValue
	for i := 0; i < len(dataWalletValues); i++ {
		createdAt := dataWalletValues[i].Date
		date := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, createdAt.Location())
		if len(result) > 0 && result[len(result)-1].Date == date {
			result[len(result)-1].SumOfValue += dataWalletValues[i].SumOfValue
			result[len(result)-1].TotalOfData += 1
			result[len(result)-1].AvgValue = result[len(result)-1].SumOfValue / result[len(result)-1].TotalOfData
		} else {
			result = append(result, models.DataWalletValue{
				Date:        date,
				SumOfValue:  dataWalletValues[i].SumOfValue,
				TotalOfData: 1,
			})
		}

	}
	return result, nil
}

func calculageTotalPercent(dataWalletValue []models.DataWalletValueResponse) float64 {
	var firstValue float64
	var lastValue float64
	if s, err := strconv.ParseFloat(dataWalletValue[0].Value, 64); err == nil {
		firstValue = s
	}
	if s, err := strconv.ParseFloat(dataWalletValue[len(dataWalletValue)-1].Value, 64); err == nil {
		lastValue = s
	}
	var percent float64 = 100
	if firstValue != 0 {
		percent = ((lastValue / firstValue) * 100) - 100
	}
	return percent
}

func prepareData(totalValue float64, dataWalletValues []models.DataWalletValue) models.WalletValueResponse {
	sort.SliceStable(dataWalletValues, func(i, j int) bool {
		return dataWalletValues[i].Date.Before(dataWalletValues[j].Date)
	})
	var data []models.DataWalletValueResponse
	for _, dataWalletValue := range dataWalletValues {
		data = append(data, models.DataWalletValueResponse{
			Date:  dataWalletValue.Date,
			Value: fmt.Sprintf("%.4f", dataWalletValue.SumOfValue),
		})
	}
	return models.WalletValueResponse{
		Total:        fmt.Sprintf("%f", totalValue),
		TotalPercent: fmt.Sprintf("%.2f", calculageTotalPercent(data)),
		Data:         data,
	}
}

func getWalletValue(c *gin.Context) {
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

	apiUrl := os.Getenv("API_URL")
	var url string = apiUrl + "/my-wallet-moralis/" + walletAddress + "?isFetch=false"

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var myWallets models.MyWallet
	err = json.Unmarshal([]byte(responseData), &myWallets)
	if err != nil {
		panic(err)
	}
	var totalValue float64 = 0
	for i := 0; i < len(myWallets.Wallet); i++ {
		if s, err := strconv.ParseFloat(myWallets.Wallet[i].Value, 64); err == nil {
			totalValue = totalValue + s
		}
	}

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("user_wallets")
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"wallet_address": walletAddress,
	})

	var userWallets []models.UserWallets
	if err = collectionData.All(mongoDB.Ctx, &userWallets); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
	if len(userWallets) == 0 {
		log.Println("Not found")
		c.String(http.StatusBadRequest, "Not found")
	}

	dataWalletValueAtTime := sumValueAtTime(userWallets)
	// dataWalletValue, err := sumValuePerDay(dataWalletValueAtTime)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	resp := prepareData(totalValue, dataWalletValueAtTime)

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
		getWalletValue(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
}
