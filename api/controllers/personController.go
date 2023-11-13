package controllers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alocin98/ski-planner-api/db"
	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/validators"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Dbconnect()

// Auths -> get token
func Auths(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	validToken, err := middlewares.GenerateJWT()
	if err != nil {
		middlewares.ErrorResponse("Failed to generate token", response)
	}

	middlewares.SuccessResponse(string(validToken), response)
}

// CreatePersonEndpoint -> create person
func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var person models.Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(person); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}
	collection := client.Database("golang").Collection("people")
	result, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
}

// GetPeopleEndpoint -> get people
func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var people []*models.Person

	collection := client.Database("golang").Collection("people")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var person models.Person
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, &person)
	}
	if err := cursor.Err(); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessArrRespond(people, response)
}

// GetPersonEndpoint -> get person by id
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, _ := primitive.ObjectIDFromHex(params.ByName("id"))
	var person models.Person

	collection := client.Database("golang").Collection("people")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&person)
	if err != nil {
		middlewares.ErrorResponse("Person does not exist", response)
		return
	}
	middlewares.SuccessRespond(person, response)
}

// DeletePersonEndpoint -> delete person by id
func DeletePersonEndpoint(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, _ := primitive.ObjectIDFromHex(params.ByName("id"))
	var person models.Person

	collection := client.Database("golang").Collection("people")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&person)
	if err != nil {
		middlewares.ErrorResponse("Person does not exist", response)
		return
	}
	_, derr := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if derr != nil {
		middlewares.ServerErrResponse(derr.Error(), response)
		return
	}
	middlewares.SuccessResponse("Deleted", response)
}

// UpdatePersonEndpoint -> update person by id
func UpdatePersonEndpoint(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, _ := primitive.ObjectIDFromHex(params.ByName("id"))
	type fname struct {
		Firstname string `json:"firstname"`
	}
	var fir fname
	json.NewDecoder(request.Body).Decode(&fir)
	collection := client.Database("golang").Collection("people")
	res, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}, bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "firstname", Value: fir.Firstname}}}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if res.MatchedCount == 0 {
		middlewares.ErrorResponse("Person does not exist", response)
		return
	}
	middlewares.SuccessResponse("Updated", response)
}

// UploadFileEndpoint -> upload file
func UploadFileEndpoint(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	file, handler, err := request.FormFile("file")
	// fileName := request.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile("uploaded/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	middlewares.SuccessResponse("Uploaded Successfully", response)
}
