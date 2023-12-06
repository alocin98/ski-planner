package routes

import (
	"github.com/alocin98/ski-planner-api/controllers"
	. "github.com/alocin98/ski-planner-api/middlewares"
	"github.com/julienschmidt/httprouter"
)

// Routes defines endpoints
func Routes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/healthcheck", controllers.HealthCheck)
	router.POST("/api/login", Cors(WithAuth(controllers.Login)))
	router.GET("/api/user", Cors(WithAuth(controllers.GetUser)))
	router.GET("/api/user/finish-tutorial", Cors(WithAuth(controllers.FinishTutorial)))

	// Strava
	router.GET("/api/strava/authorize", Cors(WithAuth(controllers.StravaAuthorize)))
	router.GET("/api/strava/exchange-token", Cors(WithAuth(controllers.StravaExchangeToken)))
	router.GET("/api/strava/load-training-data", Cors(WithAuth(controllers.StravaLoadTrainingData)))

	router.GET("/api/strava/webhook", Cors(controllers.StravaWebhookVerifier))
	router.POST("/api/strava/webhook", Cors(controllers.StravaWebhook))
	router.GET("/api/strava/webhook/list", controllers.StravaWebhookList)

	// trainings
	router.GET("/api/trainings", Cors(WithAuth(controllers.GetTrainings)))

	return router
}
