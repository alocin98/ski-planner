package providers

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/fatih/color"
)

var App *firebase.App
var FbAuthClient *auth.Client

func InitFirebase() {
	color.Red("🔥 Init firebase")
	var err error
	App, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FbAuthClient, err = App.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

}
