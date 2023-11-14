package routes

import (
	"github.com/alocin98/ski-planner-api/controllers"
	"github.com/julienschmidt/httprouter"
)

// Routes defines endpoints
func Routes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/healthcheck", controllers.HealthCheck)

	return router
}
