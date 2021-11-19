package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoUri = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	// DatabaseName = "dore-draft"
	DatabaseName = "dorepoc"
)

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB(api_string string) (*MongoDB, error) {
	var mongoDB MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(api_string))
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
	//defer func(client *mongo.Client, ctx context.Context) {
	//	err = client.Disconnect(ctx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(client, ctx)
	//defer client.Disconnect(ctx)

	mongoDB.db = database
	return &mongoDB, nil
}
