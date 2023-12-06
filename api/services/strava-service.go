package service

import (
	"context"
	"fmt"
	"time"

	"github.com/alocin98/ski-planner-api/models"
	. "github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/alocin98/ski-planner-api/strava"
	"go.mongodb.org/mongo-driver/bson"
)

func ConnectToStrava(issuerId string, tokenResponse StravaTokenResponse) (models.User, error) {
	var user models.User
	tokenDetails := StravaTokenDetails{
		TokenType:    tokenResponse.TokenType,
		ExpiresAt:    tokenResponse.ExpiresAt,
		ExpiresIn:    tokenResponse.ExpiresIn,
		RefreshToken: tokenResponse.RefreshToken,
		AccessToken:  tokenResponse.AccessToken,
	}
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOneAndUpdate(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "stravaConnected", Value: true},
			{Key: "stravaTokenDetails", Value: tokenDetails},
			{Key: "stravaAthlete", Value: tokenResponse.Athlete},
		}},
	}).Decode(&user)
	return user, err

}

func SaveStravaTokenDetails(issuerId string, tokenResponse StravaTokenResponse) {
	tokenDetails := StravaTokenDetails{
		TokenType:    tokenResponse.TokenType,
		ExpiresAt:    tokenResponse.ExpiresAt,
		ExpiresIn:    tokenResponse.ExpiresIn,
		RefreshToken: tokenResponse.RefreshToken,
		AccessToken:  tokenResponse.AccessToken,
	}
	providers.MongoClient.Database("skiyeti-db").Collection("users").FindOneAndUpdate(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "stravaInfo.stravaTokenDetails", Value: tokenDetails},
			{Key: "stravaInfo.stravaConnected", Value: true},
			{Key: "stravaInfo.stravaAthlete", Value: tokenResponse.Athlete},
		}},
	})
}

func LoadTrainingData(issuerId string) ([]StravaSummaryActivity, time.Time) {
	accessToken := getStravaAccessToken(issuerId)
	loadedUntil, activities := strava.StravaGetAthleteActivitiesLastYear(accessToken)
	return loadedUntil, activities

}

func getStravaAccessToken(issuerId string) string {
	var user models.User
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}).Decode(&user)
	if err != nil {
		panic(err)
	}
	if user.StravaInfo.StravaTokenDetails.ExpiresAt < time.Now().Unix() {
		fmt.Println("Refreshing token")
		refreshToken := user.StravaInfo.StravaTokenDetails.RefreshToken
		tokenResponse := strava.StravaRefreshToken(refreshToken)
		SaveStravaTokenDetails(issuerId, tokenResponse)
	}

	return user.StravaInfo.StravaTokenDetails.AccessToken
}

func saveActivitiesToDb(activities []StravaSummaryActivity) {
	providers.MongoClient.Database("skiyeti-db").Collection("trainings")

}
