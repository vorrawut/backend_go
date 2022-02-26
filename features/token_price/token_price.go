package token_price

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"safebsc/features/shared"
	"safebsc/mongo"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

type connectionClient struct {
	db mongo.DB
}

func filterSupportedTokenByIsHidden(st []SupportedToken) []SupportedToken {
	result := []SupportedToken{}
	for i := range st {
		if !st[i].IsHidden {
			result = append(result, st[i])
		}
	}
	return result
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

func calculateChange(tokenPrices []tokenPrice) response {
	lastRecord := tokenPrices[len(tokenPrices)-1]
	firstRecord := tokenPrices[0]

	timeOfLastRecord, err := time.Parse(time.RFC3339, lastRecord.CreatedAt.Format(time.RFC3339))
	timeOfFirstRecord, err := time.Parse(time.RFC3339, firstRecord.CreatedAt.Format(time.RFC3339))
	if err != nil {
		log.Print(err.Error())
	}

	result := response{}
	high, low := findHighAndLow(tokenPrices)

	var lastedPrice, firstPrice float64

	if (timeOfLastRecord.Day() > timeOfFirstRecord.Day()) || (timeOfLastRecord.Day() == timeOfFirstRecord.Day() && timeOfLastRecord.Hour() > timeOfFirstRecord.Hour()) {
		lastedPrice, err = strconv.ParseFloat(lastRecord.Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
		firstPrice, err = strconv.ParseFloat(firstRecord.Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
	} else {
		lastedPrice, err = strconv.ParseFloat(firstRecord.Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
		firstPrice, err = strconv.ParseFloat(lastRecord.Price, 64)
		if err != nil {
			log.Print(err.Error())
		}
	}

	change := lastedPrice - firstPrice
	priceMovement := "up"
	if lastedPrice-firstPrice < 0 {
		change = firstPrice - lastedPrice
		priceMovement = "down"
	}
	valueString := strconv.FormatFloat(lastedPrice, 'f', -1, 64)
	changeString := strconv.FormatFloat(change, 'f', -1, 64)
	result = response{
		Symbol:        tokenPrices[0].Symbol,
		Value:         valueString,
		Change:        changeString,
		PriceMovement: priceMovement,
		High:          high,
		Low:           low,
	}
	return result
}

func (dbClient connectionClient) getSupporttedTokens() ([]shared.SupportedToken, error) {
	var st []shared.SupportedToken
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("supported_tokens")

	cur, err := collection.Find(dbClient.db.Ctx, bson.M{})
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}

	for cur.Next(dbClient.db.Ctx) {
		var elem shared.SupportedToken
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		st = append(st, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(dbClient.db.Ctx)
	return st, err
}

func Handler(c *gin.Context) {
	mongoDB, err := mongo.Connect()
	conn := connectionClient{
		db: mongoDB,
	}
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	defer mongoDB.Client.Disconnect(mongoDB.Ctx)

	supportedTokens, err := conn.getSupporttedTokens()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	location, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(location)
	startTime := time.Date(now.Year(), now.Month(), now.Day()-1, now.Hour(), now.Minute(), 0, 0, now.Location())

	startDateTime, err := time.Parse(time.RFC3339, startTime.Format(time.RFC3339))
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, err.Error())
	}

	results := []response{}
	for i := 0; i < len(supportedTokens); i++ {
		collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection(strings.ToLower(supportedTokens[i].Address))
		findOptions := options.Find()
		findOptions.SetSort(bson.D{primitive.E{Key: "createdAt", Value: -1}})
		collectionData, err := collection.Find(mongoDB.Ctx, bson.M{
			"createdAt": bson.M{
				"$gt": startDateTime,
			},
		})
		if err != nil {
			log.Println(err.Error())
		}
		var tokenPrices []tokenPrice
		if err = collectionData.All(mongoDB.Ctx, &tokenPrices); err != nil {
			log.Println(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}

		if len(tokenPrices) > 0 {
			results = append(results, calculateChange(tokenPrices))
		}
	}

	contents, err := json.Marshal(results)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))

}
