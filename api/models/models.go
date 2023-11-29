package models

import (
	"time"

	"github.com/alocin98/ski-planner-api/strava"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,alpha"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,alpha"`
}

type User struct {
	IssuerId     string      `json:"issuerId,omitempty" bson:"issuerId,omitempty" validate: "required"`
	Email        string      `json:"email,omitempty" bson:"email,omitempty" validate: "required,email"`
	IsFirstLogin bool        `json:"isFirstLogin" bson:"isFirstLogin" validate: ""`
	StravaInfo   *StravaInfo `json:"stravaInfo,omitempty" bson:"stravaInfo,omitempty" validate: ""`
}

type StravaInfo struct {
	StravaConnected       bool                `json:"stravaConnected,omitempty" bson:"stravaConnected,omitempty" validate: ""`
	StravaAthlete         strava.Athlete      `json:"stravaAthlete,omitempty" bson:"stravaAthlete,omitempty" validate: ""`
	StravaTokenDetails    strava.TokenDetails `json:"stravaTokenDetails,omitempty" bson:"stravaTokenDetails,omitempty" validate: ""`
	HasImportedFromStrava bool                `bson:"hasImportedFromStrava", json: "hasImportedFromStrava"`
	StravaImportedUntil   time.Time           `bson:"stravaImportedUntil", json: "stravaImportedUntil"`
}

type Training struct {
	Origin               string    `json:"origin,omitempty" bson:"origin,omitempty" validate: "required"`
	TrainingId           string    `json:"trainingId,omitempty" bson:"trainingId,omitempty" validate: "required"`
	AthleteId            string    `json:"athleteId,omitempty" bson:"athleteId,omitempty" validate: "required"`
	Name                 string    `bson:"name"`
	Distance             float64   `bson:"distance"`
	MovingTime           int       `bson:"moving_time"`
	ElapsedTime          int       `bson:"elapsed_time"`
	TotalElevationGain   float64   `bson:"total_elevation_gain"`
	ElevHigh             float64   `bson:"elev_high"`
	ElevLow              float64   `bson:"elev_low"`
	Type                 string    `bson:"type"`
	SportType            string    `bson:"sport_type"`
	StartDate            time.Time `bson:"start_date"`
	StartDateLocal       time.Time `bson:"start_date_local"`
	Week                 int       `bson:"week"`
	Year                 int       `bson:"year"`
	Month                int       `bson:"month"`
	Timezone             string    `bson:"timezone"`
	StartLatLng          []float64 `bson:"start_latlng" json:"start_latlng"`
	EndLatLng            []float64 `bson:"end_latlng" json:"end_latlng"`
	AverageSpeed         float64   `bson:"average_speed"`
	MaxSpeed             float64   `bson:"max_speed"`
	Kilojoules           float64   `bson:"kilojoules"`
	AverageWatts         float64   `bson:"average_watts"`
	DeviceWatts          bool      `bson:"device_watts"`
	MaxWatts             int       `bson:"max_watts"`
	WeightedAverageWatts int       `bson:"weighted_average_watts"`
	AverageHeartrate     float32   `bson:"average_heartrate`
	MaxHeartrate         float32   `bson:max_heartrate`
}
