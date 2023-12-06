package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required,alpha"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,alpha"`
}

type User struct {
	IssuerId     string      `json:"issuerId,omitempty" bson:"issuerId,omitempty" validate: "required"`
	Email        string      `json:"email,omitempty" bson:"email,omitempty" validate: "required,email"`
	IsFirstLogin bool        `json:"isFirstLogin" bson:"isFirstLogin" validate: ""`
	StravaInfo   *StravaInfo `json:"stravaInfo,omitempty" bson:"stravaInfo,omitempty" validate: ""`
}

type StravaInfo struct {
	StravaConnected       bool               `json:"stravaConnected,omitempty" bson:"stravaConnected,omitempty" validate: ""`
	StravaAthlete         StravaAthlete      `json:"stravaAthlete,omitempty" bson:"stravaAthlete,omitempty" validate: ""`
	StravaTokenDetails    StravaTokenDetails `json:"stravaTokenDetails,omitempty" bson:"stravaTokenDetails,omitempty" validate: ""`
	HasImportedFromStrava bool               `bson:"hasImportedFromStrava", json: "hasImportedFromStrava"`
	StravaImportedUntil   time.Time          `bson:"stravaImportedUntil", json: "stravaImportedUntil"`
}

type Training struct {
	Origin               string    `json:"origin,omitempty" bson:"origin,omitempty" validate: "required"`
	TrainingId           string    `json:"trainingId,omitempty" bson:"trainingId,omitempty" validate: "required"`
	AthleteId            string    `json:"athleteId,omitempty" bson:"athleteId,omitempty" validate: "required"`
	Name                 string    `bson:"name"`
	Distance             float64   `bson:"distance"`
	MovingTime           int       `bson:"moving_time"`
	ElapsedTime          int       `bson:"elapsed_time"`
	TotalElevationGain   float64   `bson:"total_elevation_gain"`
	ElevHigh             float64   `bson:"elev_high"`
	ElevLow              float64   `bson:"elev_low"`
	Type                 string    `bson:"type"`
	SportType            string    `bson:"sport_type"`
	StartDate            time.Time `bson:"start_date"`
	StartDateLocal       time.Time `bson:"start_date_local"`
	Week                 int       `bson:"week"`
	Year                 int       `bson:"year"`
	Month                int       `bson:"month"`
	Timezone             string    `bson:"timezone"`
	StartLatLng          []float64 `bson:"start_latlng" json:"start_latlng"`
	EndLatLng            []float64 `bson:"end_latlng" json:"end_latlng"`
	AverageSpeed         float64   `bson:"average_speed"`
	MaxSpeed             float64   `bson:"max_speed"`
	Kilojoules           float64   `bson:"kilojoules"`
	AverageWatts         float64   `bson:"average_watts"`
	DeviceWatts          bool      `bson:"device_watts"`
	MaxWatts             int       `bson:"max_watts"`
	WeightedAverageWatts int       `bson:"weighted_average_watts"`
	AverageHeartrate     float32   `bson:"average_heartrate`
	MaxHeartrate         float32   `bson:max_heartrate`
}

type StravaAthlete struct {
	ID            int     `json:"id" bson:"_id"`
	Username      string  `json:"username" bson:"username"`
	ResourceState int     `json:"resource_state" bson:"resource_state"`
	Firstname     string  `json:"firstname" bson:"firstname"`
	Lastname      string  `json:"lastname" bson:"lastname"`
	Bio           string  `json:"bio" bson:"bio"`
	City          string  `json:"city" bson:"city"`
	State         string  `json:"state" bson:"state"`
	Country       string  `json:"country" bson:"country"`
	Sex           string  `json:"sex" bson:"sex"`
	Premium       bool    `json:"premium" bson:"premium"`
	Summit        bool    `json:"summit" bson:"summit"`
	CreatedAt     string  `json:"created_at" bson:"created_at"`
	UpdatedAt     string  `json:"updated_at" bson:"updated_at"`
	BadgeTypeID   int     `json:"badge_type_id" bson:"badge_type_id"`
	Weight        float64 `json:"weight" bson:"weight"`
	ProfileMedium string  `json:"profile_medium" bson:"profile_medium"`
	Profile       string  `json:"profile" bson:"profile"`
	Friend        string  `json:"friend" bson:"friend"`
	Follower      string  `json:"follower" bson:"follower"`
}

type StravaTokenResponse struct {
	TokenType    string        `json:"token_type" bson:"token_type"`
	ExpiresAt    int64         `json:"expires_at" bson:"expires_at"`
	ExpiresIn    int           `json:"expires_in" bson:"expires_in"`
	RefreshToken string        `json:"refresh_token" bson:"refresh_token"`
	AccessToken  string        `json:"access_token" bson:"access_token"`
	Athlete      StravaAthlete `json:"athlete" bson:"athlete"`
}

type StravaTokenDetails struct {
	TokenType    string `json:"token_type" bson:"token_type"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
	ExpiresIn    int    `json:"expires_in" bson:"expires_in"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	AccessToken  string `json:"access_token" bson:"access_token"`
}

type StravaMetaAthlete struct {
	ID int `bson:"id"`
}

type LatLng struct {
	Lat float64 `bson:"lat"`
	Lng float64 `bson:"lng"`
}

type StravaPolylineMap struct {
	ID              string `bson:"id"`
	Polyline        string `bson:"polyline"`
	SummaryPolyline string `bson:"summary_polyline"`
}

type StravaSummaryActivity struct {
	ID                   int64             `bson:"id,omitempty" json:"id,omitempty"`
	Long                 int64             `bson:"long" json:"long"`
	ExternalID           string            `bson:"external_id" json:"external_id"`
	UploadID             int64             `bson:"upload_id" json:"upload_id"`
	Athlete              StravaMetaAthlete `bson:"athlete" json:"athlete"`
	Name                 string            `bson:"name" json:"name"`
	Distance             float64           `bson:"distance" json:"distance"`
	MovingTime           int               `bson:"moving_time" json:"moving_time"`
	ElapsedTime          int               `bson:"elapsed_time" json:"elapsed_time"`
	TotalElevationGain   float64           `bson:"total_elevation_gain" json:"total_elevation_gain"`
	ElevHigh             float64           `bson:"elev_high" json:"elev_high"`
	ElevLow              float64           `bson:"elev_low" json:"elev_low"`
	Type                 string            `bson:"type" json:"type"`
	SportType            string            `bson:"sport_type" json:"sport_type"`
	StartDate            string            `bson:"start_date" json:"start_date"`
	StartDateLocal       string            `bson:"start_date_local" json:"start_date_local"`
	Timezone             string            `bson:"timezone" json:"timezone"`
	StartLatLng          []float64         `bson:"start_latlng" json:"start_latlng"`
	EndLatLng            []float64         `bson:"end_latlng" json:"end_latlng"`
	AchievementCount     int               `bson:"achievement_count" json:"achievement_count"`
	KudosCount           int               `bson:"kudos_count" json:"kudos_count"`
	CommentCount         int               `bson:"comment_count" json:"comment_count"`
	AthleteCount         int               `bson:"athlete_count" json:"athlete_count"`
	PhotoCount           int               `bson:"photo_count" json:"photo_count"`
	TotalPhotoCount      int               `bson:"total_photo_count" json:"total_photo_count"`
	Map                  StravaPolylineMap `bson:"map" json:"map"`
	Trainer              bool              `bson:"trainer" json:"trainer"`
	Commute              bool              `bson:"commute" json:"commute"`
	Manual               bool              `bson:"manual" json:"manual"`
	Private              bool              `bson:"private" json:"private"`
	Flagged              bool              `bson:"flagged" json:"flagged"`
	WorkoutType          int               `bson:"workout_type" json:"workout_type"`
	UploadIDStr          string            `bson:"upload_id_str" json:"upload_id_str"`
	AverageSpeed         float64           `bson:"average_speed" json:"average_speed"`
	MaxSpeed             float64           `bson:"max_speed" json:"max_speed"`
	HasKudoed            bool              `bson:"has_kudoed" json:"has_kudoed"`
	HideFromHome         bool              `bson:"hide_from_home" json:"hide_from_home"`
	GearID               string            `bson:"gear_id" json:"gear_id"`
	Kilojoules           float64           `bson:"kilojoules" json:"kilojoules"`
	AverageWatts         float64           `bson:"average_watts" json:"average_watts"`
	DeviceWatts          bool              `bson:"device_watts" json:"device_watts"`
	MaxWatts             int               `bson:"max_watts" json:"max_watts"`
	WeightedAverageWatts int               `bson:"weighted_average_watts" json:"weighted_average_watts"`
	AverageHeartrate     float32           `bson:"average_heartrate" json:"average_heartrate"`
	MaxHeartrate         float32           `bson:"max_heartrate" json:"max_heartrate"`
}

type StravaWebhookEvent struct {
	AspectType     string                 `json:"aspect_type"`
	EventTime      int64                  `json:"event_time"`
	ObjectID       int64                  `json:"object_id"`
	ObjectType     string                 `json:"object_type"`
	OwnerID        int64                  `json:"owner_id"`
	SubscriptionID int64                  `json:"subscription_id"`
	Updates        map[string]interface{} `json:"updates"`
}

type StravaDetailedActivity struct {
	ID                   int64             `bson:"id,omitempty" json:"id,omitempty"`
	Long                 int64             `bson:"long" json:"long"`
	ExternalID           string            `bson:"external_id" json:"external_id"`
	UploadID             int64             `bson:"upload_id" json:"upload_id"`
	Athlete              StravaMetaAthlete `bson:"athlete" json:"athlete"`
	Name                 string            `bson:"name" json:"name"`
	Distance             float64           `bson:"distance" json:"distance"`
	MovingTime           int               `bson:"moving_time" json:"moving_time"`
	ElapsedTime          int               `bson:"elapsed_time" json:"elapsed_time"`
	TotalElevationGain   float64           `bson:"total_elevation_gain" json:"total_elevation_gain"`
	ElevHigh             float64           `bson:"elev_high" json:"elev_high"`
	ElevLow              float64           `bson:"elev_low" json:"elev_low"`
	Type                 string            `bson:"type" json:"type"`
	StartDate            string            `bson:"start_date" json:"start_date"`
	StartDateLocal       string            `bson:"start_date_local" json:"start_date_local"`
	Timezone             string            `bson:"timezone" json:"timezone"`
	StartLatLng          []float64         `bson:"start_latlng" json:"start_latlng"`
	EndLatLng            []float64         `bson:"end_latlng" json:"end_latlng"`
	AchievementCount     int               `bson:"achievement_count" json:"achievement_count"`
	KudosCount           int               `bson:"kudos_count" json:"kudos_count"`
	CommentCount         int               `bson:"comment_count" json:"comment_count"`
	AthleteCount         int               `bson:"athlete_count" json:"athlete_count"`
	PhotoCount           int               `bson:"photo_count" json:"photo_count"`
	TotalPhotoCount      int               `bson:"total_photo_count" json:"total_photo_count"`
	Map                  StravaPolylineMap `bson:"map" json:"map"`
	Trainer              bool              `bson:"trainer" json:"trainer"`
	Commute              bool              `bson:"commute" json:"commute"`
	Manual               bool              `bson:"manual" json:"manual"`
	Private              bool              `bson:"private" json:"private"`
	Flagged              bool              `bson:"flagged" json:"flagged"`
	WorkoutType          int               `bson:"workout_type" json:"workout_type"`
	UploadIDStr          string            `bson:"upload_id_str" json:"upload_id_str"`
	AverageSpeed         float64           `bson:"average_speed" json:"average_speed"`
	MaxSpeed             float64           `bson:"max_speed" json:"max_speed"`
	HasKudoed            bool              `bson:"has_kudoed" json:"has_kudoed"`
	HideFromHome         bool              `bson:"hide_from_home" json:"hide_from_home"`
	GearID               string            `bson:"gear_id" json:"gear_id"`
	Kilojoules           float64           `bson:"kilojoules" json:"kilojoules"`
	AverageWatts         float64           `bson:"average_watts" json:"average_watts"`
	DeviceWatts          bool              `bson:"device_watts" json:"device_watts"`
	MaxWatts             int               `bson:"max_watts" json:"max_watts"`
	WeightedAverageWatts int               `bson:"weighted_average_watts" json:"weighted_average_watts"`
	AverageHeartrate     float32           `bson:"average_heartrate" json:"average_heartrate"`
	MaxHeartrate         float32           `bson:"max_heartrate" json:"max_heartrate"`
	SufferScore          int               `bson:"suffer_score" json:"suffer_score"`
	Description          string            `bson:"description" json:"description"`
	Calories             float64           `bson:"calories" json:"calories"`
	DeviceName           string            `bson:"device_name" json:"device_name"`
	EmbedToken           string            `bson:"embed_token" json:"embed_token"`
}
