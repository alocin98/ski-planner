package strava

import "time"

func ParseDateToTimestamp(date string) int64 {
	timestamp, err := time.Parse("2006-02-01", date)
	if err != nil {
		panic(err)
	}
	return timestamp.Unix()
}
