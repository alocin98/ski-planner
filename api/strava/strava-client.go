package strava

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

var clientSecret = "e3cfc4a374bedf06e5e10fe7302c935a3ec28e65"
var clientId = "39719"

func StravaAuthorize(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	url := "https://www.strava.com/oauth/authorize?client_id=39719&response_type=code&redirect_uri=http://localhost:8080/api/strava/exchange-token&approval_prompt=force&scope=activity:read_all"
	http.Redirect(response, request, url, http.StatusSeeOther)
}

func StravaExchangeToken(code string) TokenResponse {

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

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Fatal(err)
	}
	return tokenResponse
}

func StravaRefreshToken(refreshToken string) TokenResponse {

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

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		log.Fatal(err)
	}
	return tokenResponse
}

func StravaGetAthleteActivitiesLastYear(accessToken string) ([]SummaryActivity, time.Time) {
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

	var activities []SummaryActivity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		log.Fatal(err)
	}
	return activities, oneYearAgo
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

	return string(body)

}
