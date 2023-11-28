package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/alocin98/ski-planner-api/strava"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveTrainingToDb(training models.Training) {
	collection := providers.MongoClient.Database("skiyeti-db").Collection("trainings")

	// Define options to specify upsert behavior
	opts := options.Replace().SetUpsert(true)

	// Create a filter to check if the document already exists based on a unique field
	filter := bson.M{"trainingId": training.TrainingId} // Replace "uniqueField" with your unique field name

	// Convert the training struct to BSON
	trainingBSON, err := bson.Marshal(training)
	if err != nil {
		panic(err)
	}

	// Replace the document in the collection, or insert it if it doesn't exist
	val, err := collection.ReplaceOne(context.TODO(), filter, trainingBSON, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func SaveStravaTrainingToDb(activity strava.SummaryActivity, issuerId string) {
	training := mapStravaTraining(activity)
	training.AthleteId = issuerId
	SaveTrainingToDb(training)
}

func mapStravaTraining(activity strava.SummaryActivity) models.Training {
	startDate, err := time.Parse("2006-01-02T15:04:05Z", activity.StartDate) // Adjust the layout based on the actual date format received
	if err != nil {
		fmt.Println(err)
	}

	year, week := startDate.ISOWeek()
	month := startDate.Month()
	training := models.Training{
		Origin:               "strava",
		TrainingId:           strconv.FormatInt(activity.ID, 10),
		AthleteId:            "",
		Name:                 activity.Name,
		Distance:             activity.Distance,
		MovingTime:           activity.MovingTime,
		ElapsedTime:          activity.ElapsedTime,
		TotalElevationGain:   activity.TotalElevationGain,
		ElevHigh:             activity.ElevHigh,
		ElevLow:              activity.ElevLow,
		Type:                 activity.Type,
		SportType:            activity.SportType,
		StartDate:            startDate,
		StartDateLocal:       startDate,
		Timezone:             activity.Timezone,
		StartLatLng:          activity.StartLatLng,
		EndLatLng:            activity.EndLatLng,
		AverageSpeed:         activity.AverageSpeed,
		MaxSpeed:             activity.MaxSpeed,
		Kilojoules:           activity.Kilojoules,
		AverageWatts:         activity.AverageWatts,
		DeviceWatts:          activity.DeviceWatts,
		MaxWatts:             activity.MaxWatts,
		WeightedAverageWatts: activity.WeightedAverageWatts,
		AverageHeartrate:     activity.AverageHeartrate,
		MaxHeartrate:         activity.MaxHeartrate,
		Week:                 week,
		Year:                 year,
		Month:                int(month),
	}
	return training

}
