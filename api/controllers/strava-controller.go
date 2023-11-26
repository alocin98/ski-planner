package controllers

import (
	"fmt"
	"log"
	"net/http"

	userservice "github.com/alocin98/ski-planner-api/services"
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

	issuerId := request.Header.Get("uid")
	queryValues := request.URL.Query()

	code := queryValues.Get("code")
	grant_type := "authorization_code"

	stravaTokenResponse := strava.StravaExchangeToken(code, grant_type)

	user, err := userservice.ConnectToStrava(issuerId, stravaTokenResponse)
	redirectUrl, err := request.Cookie("strava-authorization-redirect-url")
	if err != nil {
		log.Fatal("!!No strava-authorization-redirect-url set. Please fill a cookie with this value")
		panic(err)
	}
	fmt.Println(user)

	redirectUrl.Value = ""

	http.Redirect(response, request, redirectUrl.Value, http.StatusSeeOther)

}
