package controllers

import (
	"encoding/json"
	"net/http"

	middlewares "github.com/alocin98/ski-planner-api/handlers"
	"github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/services"
	"github.com/alocin98/ski-planner-api/validators"
	"github.com/julienschmidt/httprouter"
)

func Login(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	user.IssuerId = request.Header.Get("uid")

	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(user); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}

	user, err = service.UpsertUser(user)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func GetUser(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	userId := request.Header.Get("uid")
	user, err := service.GetUser(userId)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	json.NewEncoder(response).Encode(user)
}
