package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alocin98/ski-planner-api/services"
	"github.com/alocin98/ski-planner-api/strava"
	"github.com/julienschmidt/httprouter"
)

var clientSecret = "e3cfc4a374bedf06e5e10fe7302c935a3ec28e65"
var clientId = "39719"

func StravaAuthorize(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	url := "https://www.strava.com/oauth/authorize?client_id=39719&response_type=code&redirect_uri=http://localhost:8080/api/strava/exchange-token&approval_prompt=force&scope=activity:read_all"
	http.Redirect(response, request, url, http.StatusSeeOther)
}

func StravaExchangeToken(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	// populate values
	issuerId := request.Header.Get("uid")
	queryValues := request.URL.Query()
	code := queryValues.Get("code")

	// actions
	stravaTokenResponse := strava.StravaExchangeToken(code)
	service.SaveStravaTokenDetails(issuerId, stravaTokenResponse)

	// response
	redirectUrl, err := request.Cookie("strava-authorization-redirect-url")
	if err != nil {
		log.Fatal("!!No strava-authorization-redirect-url set. Please fill a cookie with this value")
		panic(err)
	}
	fmt.Println(redirectUrl.Value)

	http.Redirect(response, request, redirectUrl.Value, http.StatusSeeOther)

}

func StravaLoadTrainingData(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	issuerId := request.Header.Get("uid")
	// populate values
	trainings, loadedUntil := service.LoadTrainingData(issuerId)
	for _, training := range trainings {
		service.SaveStravaTrainingToDb(training, issuerId)
	}
	service.SetHasImportedFromStrava(issuerId, loadedUntil)

	// response
	json.NewEncoder(response).Encode(trainings)

}

func StravaWebhookVerifier(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("Webhook called")
	// populate values
	queryValues := request.URL.Query()
	challenge := queryValues.Get("hub.challenge")
	verifyToken := queryValues.Get("hub.verify_token")
	if verifyToken != "STRAVA" {
		fmt.Println("Verify token does not match")
		return
	}
	// Construct JSON response
	jsonResponse := map[string]string{"hub.challenge": challenge}
	jsonBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write JSON response
	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonBytes)

}

func StravaWebhook(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Println("Webhook called")
	// populate values
	var webhookEvent strava.StravaWebhookEvent
	err := json.NewDecoder(request.Body).Decode(&webhookEvent)
	if err != nil {
		panic(err)
	}
	fmt.Println(webhookEvent)
	// actions
	//service.SaveStravaTrainingToDb(webhookEvent)

	// response
	response.Write([]byte("ok"))

}

func StravaWebhookList(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// actions
	webhookList := strava.ListWebhooks()
	response.Write([]byte(webhookList))
}
