package db

import (
	"context"
	"go-echo-mongodb/types/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (db *mongoDB) CreateUser(ctx context.Context, u *model.User) (*mongo.InsertOneResult,error) {
	collection := db.db.Collection("user")
	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	return result, nil
}
