package providers

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func getFirebase() *firebase.App {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}
	return app
}
