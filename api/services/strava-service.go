package userservice

import (
	"context"
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

func LoadTrainingData(issuerId string, refreshToken string) {
	now := time.Now()
	// Last 365 days
	for i := 0; i < 365; i++ {
		date := now.AddDate(0, 0, -i)
		activities := strava.StravaGetAthleteActivitesAtDay(refreshToken, date.Format("2006-01-02"))
		saveActivitiesToDb(activities)

	}
}

func saveActivitiesToDb(activities []strava.SummaryActivity) {
	providers.MongoClient.Database("skiyeti-db").Collection("trainings")

}
