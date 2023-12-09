package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/alocin98/ski-planner-api/models"
	. "github.com/alocin98/ski-planner-api/models"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

var clientSecret = "e3cfc4a374bedf06e5e10fe7302c935a3ec28e65"
var clientId = "39719"

func StravaAuthorize(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	Host := request.URL.Host
	url := "https://www.strava.com/oauth/authorize?client_id=39719&response_type=code&redirect_uri=http://" + Host + "/api/strava/exchange-token&approval_prompt=force&scope=activity:read_all"
	http.Redirect(response, request, url, http.StatusSeeOther)
}

func StravaExchangeToken(code string) StravaTokenResponse {

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", "https://www.strava.com/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	var tokenResponse StravaTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Fatal(err)
	}
	return tokenResponse
}

func StravaRefreshToken(refreshToken string) StravaTokenResponse {

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("refresh_token", refreshToken)
	data.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", "https://www.strava.com/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	var tokenResponse StravaTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Fatal(err)
	}
	return tokenResponse
}

func StravaGetAthleteActivitiesLastYear(accessToken string) ([]StravaSummaryActivity, time.Time) {
	// Get today's date and the date from one year ago
	now := time.Now().UTC()
	oneYearAgo := now.AddDate(-1, 0, 0)

	after := oneYearAgo.Unix()
	before := now.Unix()

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities?after="+strconv.FormatInt(after, 10)+"&before="+strconv.FormatInt(before, 10)+"&per_page=100", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	var activities []StravaSummaryActivity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		log.Fatal(err)
	}
	return activities, oneYearAgo
}

func StravaGetActivityDetails(accessToken string, activityId string) StravaSummaryActivity {
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/activities/"+activityId, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}

	var activity StravaSummaryActivity
	if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
		log.Fatal(err)
	}
	return activity
}

func ListWebhooks() string {
	apiURL := "https://www.strava.com/api/v3/push_subscriptions"

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("GET", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Panic("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error reading response body:", err)
		return ""
	}
	fmt.Println(string(body))

	return string(body)

}

func StravaGetAccessToken(issuerId string) string {
	fmt.Println(issuerId)
	// actions
	doc := providers.MongoDb.Collection("users").FindOne(context.Background(), bson.D{
		{Key: "issuerId", Value: issuerId},
	})
	var user models.User
	doc.Decode(&user)
	fmt.Println(user)

	// check if token is still valid
	if user.StravaInfo.StravaTokenDetails.ExpiresAt > time.Now().Unix() {
		return user.StravaInfo.StravaTokenDetails.AccessToken
	}

	// refresh token
	tokenResponse := StravaRefreshToken(user.StravaInfo.StravaTokenDetails.RefreshToken)

	// update user
	user.StravaInfo.StravaTokenDetails.AccessToken = tokenResponse.AccessToken
	user.StravaInfo.StravaTokenDetails.RefreshToken = tokenResponse.RefreshToken
	user.StravaInfo.StravaTokenDetails.ExpiresAt = tokenResponse.ExpiresAt
	providers.MongoDb.Collection("users").FindOneAndUpdate(context.Background(), bson.D{
		{Key: "issuerId", Value: issuerId},
	}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "stravaInfo.stravaTokenDetails", Value: user.StravaInfo.StravaTokenDetails},
		}},
	})

	return tokenResponse.AccessToken
}

func SendRequestAndDecodeAnswer(method string, apiURL string, data url.Values, accessToken string) string {

	req, err := http.NewRequest(method, apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Panic("Error creating request:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Error reading response body:", err)
		return ""
	}
	fmt.Println(string(body))

	return string(body)
}

func StravaGetUserByStravaId(stravaId int64) models.User {
	// actions
	doc := providers.MongoDb.Collection("users").FindOne(context.Background(), bson.D{
		{Key: "stravaInfo.stravaAthlete._id", Value: stravaId},
	})
	var user models.User
	doc.Decode(&user)

	return user
}

func StravaHandleWebhookEvent(webhookEvent StravaWebhookEvent) (models.StravaSummaryActivity, models.User) {
	// actions
	user := StravaGetUserByStravaId(webhookEvent.OwnerID)
	var accessToken = StravaGetAccessToken(user.IssuerId)
	activity := StravaGetActivityDetails(accessToken, strconv.FormatInt(webhookEvent.ObjectID, 10))
	return activity, user

}
