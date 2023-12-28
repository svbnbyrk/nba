package domain

import "time"

type Game struct {
	ID         int  `json:"id" bson:"_id"`
	AwayTeamID Team `json:"away" bson:"away"`
	HomeTeamID Team `json:"home" bson:"home"`
	Week       int
	StartTime  time.Time `json:"start_time" bson:"start_time"`
	IsFinished bool      `json:"is_finished" bson:"is_finished"`
}
