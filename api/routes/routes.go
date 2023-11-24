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
	router.POST("/api/login", WithAuth(controllers.Login))

	return router
}
