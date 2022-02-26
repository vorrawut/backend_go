package init_data

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"safebsc/features/init_data/models"

	"go.mongodb.org/mongo-driver/bson"

	"safebsc/features/shared"
	"safebsc/mongo"

	"github.com/gin-gonic/gin"
)

var ErrNoDocuments = errors.New("mongo: no documents in result").Error()

type connectionClient struct {
	db mongo.DB
}

func initDataHandler(c *gin.Context) {
	db, err := mongo.Connect()
	conn := connectionClient{
		db: db,
	}

	defer db.Client.Disconnect(db.Ctx)

	supportedTokens, err := conn.getSupporttedTokens()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	supportedFarms, err := conn.getSupporttedFarms()
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	resp := models.InitDataResponse{
		SupportedTokens: supportedTokens,
		SupportedFarms:  supportedFarms,
	}

	contents, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, string(contents))
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
