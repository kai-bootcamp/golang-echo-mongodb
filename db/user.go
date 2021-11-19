package db

import (
	"context"
	"fmt"
	"go-echo-mongodb/types/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDB) CreateUser(ctx context.Context, u *model.User) (*mongo.InsertOneResult, error) {
	collection := db.db.Collection("users")
	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Unable to create user")
		return result, err
	}

	return result, nil
}

func (db *MongoDB) ReadUser(username string) (bson.M, error) {
	var result bson.M
	collection := db.db.Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Unable to read user info  %s\n", username)
		return result, nil
	}

	return result, nil
}

func (db *MongoDB) RemoveUser(username string) (*mongo.DeleteResult, error) {
	// var result bson.M
	collection := db.db.Collection("users")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("unable to remove account %s\n", username)
		return result, nil
	}

	return result, nil
}

func (db *MongoDB) UpdateUser(userUpdateInfo *model.User) (*mongo.UpdateResult, error) {
	collection := db.db.Collection("users")
	result, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": userUpdateInfo.Id}, userUpdateInfo)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("unable to remove account", result)
		return result, nil
	}
	return result, nil
}
