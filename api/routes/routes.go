package routes

import (
	"net/http"

	"github.com/alocin98/ski-planner-api/controllers"
	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/julienschmidt/httprouter"
)

// Routes defines endpoints
func Routes() *httprouter.Router {
	router := httprouter.New()

	router.GET("/healthcheck", controllers.HealthCheck)
	router.POST("/person", controllers.CreatePersonEndpoint)
	router.GET("/auth", controllers.Auths)
	router.GET("/people", middlewares.IsAuthorized(controllers.GetPeopleEndpoint))
	router.GET("/person/:id", controllers.GetPersonEndpoint)
	router.DELETE("/person/:id", controllers.DeletePersonEndpoint)
	router.PUT("/person/:id", controllers.UpdatePersonEndpoint)
	router.POST("/upload", controllers.UploadFileEndpoint)
	router.ServeFiles("/static/*filepath", http.Dir("./uploaded/"))

	return router
}
