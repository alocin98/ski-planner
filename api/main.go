package main

import (
	"fmt"
	"log"
	"net/http"

	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/alocin98/ski-planner-api/routes"
	"github.com/fatih/color"
	"github.com/rs/cors"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	// Start server
	port := middlewares.DotEnvVariable("PORT")
	color.Cyan("🌏 Server running on localhost:" + port)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Init firebase
	providers.InitFirebase()

	// init routes
	router := routes.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)

	// Start server
	http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
}
