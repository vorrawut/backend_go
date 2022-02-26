package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client       *mongo.Client
	Ctx          context.Context
	DatabaseName string
}

func Connect() (DB, error) {
	mongoUri := os.Getenv("MONGO_URI")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoDatabase := os.Getenv("MONGO_DATABASE")
	var db DB
	var err error

	credential := options.Credential{
		AuthSource: mongoDatabase,
		Username:   mongoUsername,
		Password:   mongoPassword,
	}
	db.Client, err = mongo.NewClient(options.Client().ApplyURI(mongoUri).SetAuth(credential))
	if err != nil {
		return db, err
	}

	db.Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Client.Connect(db.Ctx)
	if err != nil {
		return db, err
	}

	db.DatabaseName = mongoDatabase

	return db, nil
}
