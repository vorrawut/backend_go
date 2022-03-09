package services

import (
	"context"
	"encoding/json"
	"log"
	"os"
	models "safebsc/features/apr_current/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOne interface {
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
}

func ConnectMongo(mongoDBUri string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	credential := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUri).SetAuth(credential))
	defer cancel()

	dbName := os.Getenv("MONGO_DATABASE")
	return client.Database(dbName), err
}

func GetAprCurrent(c *gin.Context) {
	id := c.Query("poolId")

	i64, err := strconv.ParseInt(id, 10, 32)
	poolId := int32(i64)

	mongoDBUri := os.Getenv("MONGO_URI")
	mongo, err := ConnectMongo(mongoDBUri)
	if err != nil {
		Response(err.Error(), 500, c)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := mongo.Collection("apr_current")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var apr []bson.M
	if err = cursor.All(ctx, &apr); err != nil {
		log.Fatal(err)
	}

	poolAddress, err := GetOneByID(collection, poolId)
	if err != nil {
		Response(err.Error(), 500, c)
	}

	content, err := json.Marshal(poolAddress)
	if err != nil {
		Response(err.Error(), 500, c)
	}

	Response(string(content), 200, c)
}

func GetOneByID(collection FindOne, poolId int32) (*models.PoolAddress, error) {
	var poolAddress *models.PoolAddress

	filter := bson.M{"pool_id": poolId}
	err := collection.FindOne(context.TODO(), filter).Decode(&poolAddress)

	if err != nil {
		return poolAddress, err
	}
	return poolAddress, nil
}

func Response(body string, statusCode int, c *gin.Context) {
	c.String(statusCode, body)
}
