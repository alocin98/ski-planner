package strava

type Athlete struct {
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

type TokenResponse struct {
	TokenType    string  `json:"token_type" bson:"token_type"`
	ExpiresAt    int64   `json:"expires_at" bson:"expires_at"`
	ExpiresIn    int     `json:"expires_in" bson:"expires_in"`
	RefreshToken string  `json:"refresh_token" bson:"refresh_token"`
	AccessToken  string  `json:"access_token" bson:"access_token"`
	Athlete      Athlete `json:"athlete" bson:"athlete"`
}

type TokenDetails struct {
	TokenType    string `json:"token_type" bson:"token_type"`
	ExpiresAt    int64  `json:"expires_at" bson:"expires_at"`
	ExpiresIn    int    `json:"expires_in" bson:"expires_in"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	AccessToken  string `json:"access_token" bson:"access_token"`
}

type MetaAthlete struct {
	ID int `bson:"id"`
}

type LatLng struct {
	Lat float64 `bson:"lat"`
	Lng float64 `bson:"lng"`
}

type PolylineMap struct {
	ID              string `bson:"id"`
	Polyline        string `bson:"polyline"`
	SummaryPolyline string `bson:"summary_polyline"`
}

type SummaryActivity struct {
	ID                   int64       `bson:"_id,omitempty"`
	Long                 int64       `bson:"long"`
	ExternalID           string      `bson:"external_id"`
	UploadID             int64       `bson:"upload_id"`
	Athlete              MetaAthlete `bson:"athlete"`
	Name                 string      `bson:"name"`
	Distance             float64     `bson:"distance"`
	MovingTime           int         `bson:"moving_time"`
	ElapsedTime          int         `bson:"elapsed_time"`
	TotalElevationGain   float64     `bson:"total_elevation_gain"`
	ElevHigh             float64     `bson:"elev_high"`
	ElevLow              float64     `bson:"elev_low"`
	Type                 string      `bson:"type"`
	SportType            string      `bson:"sport_type"`
	StartDate            string      `json:"start_date" bson:"start_date"`
	StartDateLocal       string      `bson:"start_date_local"`
	Timezone             string      `bson:"timezone"`
	StartLatLng          LatLng      `bson:"start_latlng"`
	EndLatLng            LatLng      `bson:"end_latlng"`
	AchievementCount     int         `bson:"achievement_count"`
	KudosCount           int         `bson:"kudos_count"`
	CommentCount         int         `bson:"comment_count"`
	AthleteCount         int         `bson:"athlete_count"`
	PhotoCount           int         `bson:"photo_count"`
	TotalPhotoCount      int         `bson:"total_photo_count"`
	Map                  PolylineMap `bson:"map"`
	Trainer              bool        `bson:"trainer"`
	Commute              bool        `bson:"commute"`
	Manual               bool        `bson:"manual"`
	Private              bool        `bson:"private"`
	Flagged              bool        `bson:"flagged"`
	WorkoutType          int         `bson:"workout_type"`
	UploadIDStr          string      `bson:"upload_id_str"`
	AverageSpeed         float64     `bson:"average_speed"`
	MaxSpeed             float64     `bson:"max_speed"`
	HasKudoed            bool        `bson:"has_kudoed"`
	HideFromHome         bool        `bson:"hide_from_home"`
	GearID               string      `bson:"gear_id"`
	Kilojoules           float64     `bson:"kilojoules"`
	AverageWatts         float64     `bson:"average_watts"`
	DeviceWatts          bool        `bson:"device_watts"`
	MaxWatts             int         `bson:"max_watts"`
	WeightedAverageWatts int         `bson:"weighted_average_watts"`
	AverageHeartrate     int         `bson:"average_heartrate`
	MaxHeartrate         int         `bson:max_heartrate`
}
