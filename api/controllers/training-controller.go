package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTrainings(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	// populate values
	issuerId := request.Header.Get("uid")
	fmt.Println(issuerId)

	filter := bson.D{{"athleteId", issuerId}}

	options := options.Find()
	options.SetSort(bson.D{{"start_date", -1}}) // 1 for ascending, -1 for descending

	cursor, err := providers.MongoClient.Database("skiyeti-db").Collection("trainings").Find(context.TODO(), filter, options)
	if err != nil {
		panic(err)
	}

	var trainings []models.Training
	if err = cursor.All(context.TODO(), &trainings); err != nil {
		panic(err)
	}

	json.NewEncoder(response).Encode(trainings)

}
