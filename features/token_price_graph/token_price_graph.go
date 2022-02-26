package token_price_graph

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"safebsc/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getStartEndEvery15Minute(currentTime time.Time) (startingTime time.Time, endingTime time.Time) {
	startingTime = currentTime
	endingTime = currentTime.Add(time.Minute * 15)
	return startingTime, endingTime
}

func calculateForGraph(tokenPrices []*tokenPrice) []response {
	var prices [][]tokenPrice
	var price []tokenPrice
	startingTime, endingTime := getStartEndEvery15Minute(tokenPrices[0].CreatedAt)
	for i := 0; i < len(tokenPrices); i++ {
		if (tokenPrices[i].CreatedAt.Equal(startingTime) || tokenPrices[i].CreatedAt.After(startingTime)) &&
			(tokenPrices[i].CreatedAt.Before(endingTime) || tokenPrices[i].CreatedAt.Equal(endingTime)) {
			price = append(price, *tokenPrices[i])
			if i != len(tokenPrices)-1 {
				continue
			}
		}
		prices = append(prices, price)
		price = []tokenPrice{*tokenPrices[i-1], *tokenPrices[i]}
		startingTime, endingTime = getStartEndEvery15Minute(tokenPrices[i-1].CreatedAt)
	}

	var results []response
	for i := 0; i < len(prices); i++ {
		var result response
		result.Time = prices[i][0].CreatedAt
		high, low := findHighAndLow(prices[i])
		result.High = high
		result.Low = low
		openPrice, err := strconv.ParseFloat(prices[i][0].Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
		closePrice, err := strconv.ParseFloat(prices[i][len(prices[i])-1].Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
		openPriceString := strconv.FormatFloat(openPrice, 'f', -1, 64)
		closePriceString := strconv.FormatFloat(closePrice, 'f', -1, 64)
		result.Open = openPriceString
		result.Close = closePriceString
		results = append(results, result)
	}

	return results
}

func findHighAndLow(tokenPrices []tokenPrice) (highString string, lowString string) {
	low, err := strconv.ParseFloat(tokenPrices[0].Price, 64)
	if err != nil {
		log.Print(err.Error())
	}
	high := low
	for _, tokenPrice := range tokenPrices {
		value, err := strconv.ParseFloat(tokenPrice.Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
		if value < low {
			low = value
		}
		if value > high {
			high = value
		}
	}
	highString = strconv.FormatFloat(high, 'f', -1, 64)
	lowString = strconv.FormatFloat(low, 'f', -1, 64)

	return highString, lowString
}

func Handler(c *gin.Context) {
	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	tokenAddress := c.Param("tokenAddress")
	fromDate, err := time.Parse(time.RFC3339, c.Query("from"))
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection(tokenAddress)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: 1}}).SetLimit(6600)
	collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
		"createdAt": bson.M{
			"$gt": fromDate,
		},
	}, findOptions)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var tokenPrices []*tokenPrice
	if err = collectionData.All(mongoDB.Ctx, &tokenPrices); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var result []response
	if len(tokenPrices) > 0 {
		result = calculateForGraph(tokenPrices)
	}

	contents, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}
