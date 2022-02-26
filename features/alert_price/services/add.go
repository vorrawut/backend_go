package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"safebsc/features/alert_price/constants"
	"safebsc/features/alert_price/models"
	"safebsc/features/alert_price/utils"
	"safebsc/mongo"
	"safebsc/postgres"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

func AddAlertPrice(c *gin.Context) {
	log.Printf("!!!!! AddAlertPrice")
	var alertPriceRequest models.AlertPriceRequest
	jsonData, err1 := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal([]byte(jsonData), &alertPriceRequest)
	if err1 != nil || err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "Error")
	}
	log.Printf("!!!!! AddAlertPrice request: %s\n", alertPriceRequest)

	// Connect DB
	db := postgres.Connect()
	defer db.Close()

	frequency := utils.MapFrequency(alertPriceRequest.Frequency)
	if frequency == "" {
		log.Println("Frequency Invalid value !")
		c.String(http.StatusBadRequest, "Error")
	}

	alertType := utils.MapAlertType(alertPriceRequest.AlertType)
	if alertType == "" {
		log.Println("AlertType Invalid value !")
		c.String(http.StatusBadRequest, "Error")
	}

	commonAlertPrice := alertPriceRequest.CommonAlertPrice

	mongoDB, err := mongo.Connect()
	if err != nil {
		log.Println("mongo.Connect error ", err.Error())
	}

	var targetPrice = 0.0
	if alertType == constants.AlertType.ChangeIsOver || alertType == constants.AlertType.ChangeIsUnder {
		var tokenPrice models.TokenPrice
		collection := mongoDB.Client.Database(mongoDB.DatabaseName).Collection("token_prices")
		if err != nil {
			log.Println("token_prices Not Found error ", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}

		err = collection.FindOne(mongoDB.Ctx, bson.M{"token": commonAlertPrice.TokenAddress}).Decode(&tokenPrice)
		if err != nil && err.Error() != ErrNoDocuments {
			log.Println("TokenAddress Not Found error ", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}

		price, err := strconv.ParseFloat(tokenPrice.Price, 64)
		if err != nil {
			log.Println("Price error ", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
		}

		if alertType == constants.AlertType.ChangeIsOver {
			targetPrice = (commonAlertPrice.Value + 100.0) / 100 * price
		} else if alertType == constants.AlertType.ChangeIsUnder {
			targetPrice = (100.0 - commonAlertPrice.Value) / 100 * price
		}
	} else {
		targetPrice = commonAlertPrice.Value
	}

	ap := models.AlertPrice{
		CommonAlertPrice: commonAlertPrice,
		Frequency:        frequency,
		AlertType:        alertType,
		TargetPrice:      targetPrice,
		AlertDateTime:    time.Now(),
		IsActive:         true,
		CreatedDate:      time.Now(),
		CreatedBy:        "System",
		UpdatedDate:      time.Now(),
		UpdatedBy:        "System",
	}

	if _, err := db.Model(&ap).Insert(); err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusCreated, "Created")
}
