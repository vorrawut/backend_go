package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"safebsc/features/alert_price/models"
	"safebsc/features/alert_price/utils"
	"safebsc/mongo"
	"safebsc/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAlertPrice(c *gin.Context) {
	userId := c.Param("userId")
	log.Printf("!!!!! GetAlertPrice by userId: %s\n", userId)
	if userId == "" {
		log.Println("UserId Invalid Empty")
		c.String(http.StatusBadRequest, "Error")
	}

	tokenAddress := c.Param("tokenAddress")
	log.Printf("!!!!! GetAlertPrice by tokenAddress: %s\n", tokenAddress)
	if userId == "" {
		log.Println("TokenAddress Invalid Empty")
		c.String(http.StatusBadRequest, "Error")
	}

	// Connect DB
	db := postgres.Connect()
	defer db.Close()

	var alertPrice []models.AlertPrice
	if err := db.Model(&alertPrice).
		Where("users_id = ?", userId).
		Where("is_active = ?", true).
		Where("token_address = ?", tokenAddress).
		Select(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	if alertPrice == nil || len(alertPrice) == 0 {
		log.Println("AlertPriceItems not found")
		c.String(http.StatusNoContent, "Error")
	}

	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println("mongo.Connect error ", err.Error())
	}

	collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("token_prices")
	if err != nil {
		log.Println("token_prices Not Found error ", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	apJSON, err := json.Marshal(alertPrice)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	log.Printf("!!!!! GetAlertPrice content: %s\n", string(apJSON))

	var tokenPrice models.TokenPrice
	err = collection.FindOne(mongoDB.Ctx, bson.M{"token": tokenAddress}).Decode(&tokenPrice)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println("TokenAddress Not Found error ", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	var price float64 = 0.0
	if !isNil(tokenPrice) {
		price, err = strconv.ParseFloat(tokenPrice.Price, 64)
		if err != nil {
			log.Println("Token Price error ", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}
	}

	var items []models.GetAlertPriceItem
	symbolName := ""
	for _, ap := range alertPrice {
		symbolName = ap.Symbol

		alertType := utils.MapAlertTypeByName(ap.AlertType)
		if alertType == 0 {
			log.Println("AlertType Invalid value !")
		}

		frequency := utils.MapFrequencyByName(ap.Frequency)
		if frequency == 0 {
			log.Println("Frequency Invalid value !")
		}

		items = append(items, models.GetAlertPriceItem{
			ID:          strconv.Itoa(ap.ID),
			Value:       fmt.Sprintf("%f", ap.Value),
			TargetPrice: fmt.Sprintf("%f", ap.TargetPrice),
			AlertType:   strconv.Itoa(alertType),
			Frequency:   strconv.Itoa(frequency),
		})
	}

	response := models.GetAlertPriceResponse{
		Symbol:          symbolName,
		TokenAddress:    tokenAddress,
		Price:           fmt.Sprintf("%f", price),
		AlertPriceItems: items,
	}

	contents, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	log.Printf("!!!!! GetAlertPrice contents: %s\n", string(contents))
	c.String(http.StatusOK, string(contents))
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
