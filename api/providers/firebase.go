package providers

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/alocin98/ski-planner-api/handlers"
	"github.com/fatih/color"
	"google.golang.org/api/option"
)

var App *firebase.App
var FbAuthClient *auth.Client

func InitFirebase() {
	color.Red("🔥 Init firebase")
	var err error
	credentials := handlers.DotEnvVariable("GOOGLE_APPLICATION_CREDENTIALS")
	opt := option.WithCredentialsFile(credentials)
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	FbAuthClient, err = App.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

}
