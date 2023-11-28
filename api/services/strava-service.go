package service

import (
	"context"
	"fmt"
	"time"

	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/alocin98/ski-planner-api/strava"
	"go.mongodb.org/mongo-driver/bson"
)

func ConnectToStrava(issuerId string, tokenResponse strava.TokenResponse) (models.User, error) {
	var user models.User
	tokenDetails := strava.TokenDetails{
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

func SaveStravaTokenDetails(issuerId string, tokenResponse strava.TokenResponse) {
	tokenDetails := strava.TokenDetails{
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
			{Key: "stravaTokenDetails", Value: tokenDetails},
			{Key: "stravaConnected", Value: true},
			{Key: "stravaAthlete", Value: tokenResponse.Athlete},
		}},
	})
}

func LoadTrainingData(issuerId string) []strava.SummaryActivity {
	accessToken := getStravaAccessToken(issuerId)
	activities := strava.StravaGetAthleteActivitiesLastYear(accessToken)
	return activities

}

func getStravaAccessToken(issuerId string) string {
	var user models.User
	err := providers.MongoClient.Database("skiyeti-db").Collection("users").FindOne(context.TODO(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}).Decode(&user)
	if err != nil {
		panic(err)
	}
	if user.StravaTokenDetails.ExpiresAt < time.Now().Unix() {
		fmt.Println("Refreshing token")
		refreshToken := user.StravaTokenDetails.RefreshToken
		tokenResponse := strava.StravaRefreshToken(refreshToken)
		SaveStravaTokenDetails(issuerId, tokenResponse)
	}

	return user.StravaTokenDetails.AccessToken
}

func saveActivitiesToDb(activities []strava.SummaryActivity) {
	providers.MongoClient.Database("skiyeti-db").Collection("trainings")

}
