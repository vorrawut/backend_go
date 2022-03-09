package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"safebsc/features/scz_price/models"

	"go.mongodb.org/mongo-driver/bson"

	"safebsc/mongo"

	"github.com/gin-gonic/gin"
)

const SCZ = "0x39f1014a88c8ec087cedf1bfc7064d24f507b894"

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

type connectionClient struct {
	db mongo.DB
}

type Controller struct{}

func (u Controller) GetSczPrice(c *gin.Context) {
	getSczPrice(c)
}

func getSczPrice(c *gin.Context) {
	db, err := mongo.Connect()
	conn := connectionClient{
		db: db,
	}

	defer db.Client.Disconnect(db.Ctx)

	sczPrice, err := conn.findSCZPrice()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	contents, err := json.Marshal(sczPrice)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
}

func (dbClient connectionClient) findSCZPrice() (*models.TokenPrice, error) {
	var tokenPrice models.TokenPrice
	collection := dbClient.db.Client.Database(dbClient.db.DatabaseName).Collection("token_prices")

	err := collection.FindOne(dbClient.db.Ctx, bson.M{"token": SCZ}).Decode(&tokenPrice)
	if err != nil && err.Error() != ErrNoDocuments {
		log.Println(err.Error())
		return nil, err
	}
	return &tokenPrice, err
}
