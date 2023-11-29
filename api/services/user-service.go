package service

import (
	"context"
	"fmt"
	"time"

	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertUser(user models.User) (models.User, error) {
	// check if user exists
	userFilter := bson.D{{Key: "issuerId", Value: user.IssuerId}}
	var existingUser models.User
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOne(context.TODO(), userFilter).Decode(&existingUser)
	if err != nil {
		user.IsFirstLogin = true
	} else {
		user.IsFirstLogin = false
	}

	// upsert
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	filter := bson.D{{Key: "issuerId", Value: user.IssuerId}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "email", Value: user.Email},
			{Key: "isFirstLogin", Value: user.IsFirstLogin},
		}},
	}

	insertError := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&user)
	return user, insertError
}

func GetUser(issuerId string) (models.User, error) {
	var user models.User
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}).Decode(&user)
	fmt.Println(user)
	return user, err
}

func SetHasImportedFromStrava(issuerId string, until time.Time) {
	providers.MongoClient.Database("skiyeti-db").Collection("users").FindOneAndUpdate(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "stravaInfo.hasImportedFromStrava", Value: true},
			{Key: "stravaInfo.stravaImportedUntil", Value: until},
		}},
	})
}

func FinishTutorial(issuerId string) error {
	_, err := providers.MongoClient.Database("skiyeti-db").Collection("users").UpdateOne(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "isFirstLogin", Value: false},
		}},
	})
	return err
}
