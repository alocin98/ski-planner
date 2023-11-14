package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,alpha"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,alpha"`
}

type User struct {
	IssuerId string `json:"issuer_id,omitempty" bson:"issuer_id,omitempty" validate: "required"`
	Email    string `json:"email,omitempty" bson:"email,omitempty" validate: "required,email"`
}
