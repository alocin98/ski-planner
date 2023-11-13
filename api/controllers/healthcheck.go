package controllers

import (
	"net/http"

	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/julienschmidt/httprouter"
)

func HealthCheck(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	middlewares.SuccessResponse("OK", response)
}
