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
