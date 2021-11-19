package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	MongoUri = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	DatabaseName = "dore-draft"
)

type mongoDB struct {
	db *mongo.Database
}

func NewMongoDB() (*mongoDB, error) {
	var mongoDB mongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	database := client.Database(DatabaseName)

	mongoDB.db = database
	return &mongoDB, nil
}

