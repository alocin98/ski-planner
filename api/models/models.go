package models

import (
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
	IssuerId          string              `json:"issuerId,omitempty" bson:"issuerId,omitempty" validate: "required"`
	Email             string              `json:"email,omitempty" bson:"email,omitempty" validate: "required,email"`
	StravaConnected   bool                `json:"stravaConnected,omitempty" bson:"stravaConnected,omitempty" validate: ""`
	StravaAthlete     strava.Athlete      `json:"stravaAthlete,omitempty" bson:"stravaAthlete,omitempty" validate: ""`
	SravaTokenDetails strava.TokenDetails `json:"stravaTokenDetails,omitempty" bson:"stravaTokenDetails,omitempty" validate: ""`
}
