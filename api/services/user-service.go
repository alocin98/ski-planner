package userservice

import (
	"context"
	"fmt"

	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertUser(user models.User) (models.User, error) {
	// upsert
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	filter := bson.D{{Key: "issuerId", Value: user.IssuerId}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "email", Value: user.Email},
		}},
	}

	insertError := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&user)
	return user, insertError
}

func GetUser(issuerId string) (models.User, error) {
	fmt.Println(issuerId)
	fmt.Println("issuerId")
	var user models.User
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}).Decode(&user)
	fmt.Println(user)
	return user, err
}
