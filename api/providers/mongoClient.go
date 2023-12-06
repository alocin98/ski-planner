package providers

import (
	"context"
	"log"

	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDb *mongo.Database

// Dbconnect -> connects mongo
func Dbconnect(db string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(middlewares.DotEnvVariable("MONGO_URL"))
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	MongoDb = MongoClient.Database(db)
	if err != nil {
		log.Fatal(err)
	}
	return MongoClient
}

func MongoUpsertOne(collection string, filter interface{}, update interface{}) {

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	c := MongoDb.Collection(collection)
	update = c.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(update)
}
