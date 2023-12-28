package domain

type Player struct {
	ID       int          `json:"id" bson:"_id"`
	Fullname string       `json:"fullname" bson:"name"`
	TeamID   int          `json:"team_id" bson:"team_id"`
	Stats    []PlayerStat `json:"stats,omitempty" bson:"-"`
}

type PlayerStat struct {
	ID                int `json:"id" bson:"_id"`
	GameID            int `json:"game_id" bson:"game_id"`
	PlayerID          int `json:"player_id" bson:"player_id"`
	TwoPointAttempts  int `json:"two_point_attemp" bson:"two_point_attemp"`
	TwoPointMade      int `json:"two_point_made" bson:"two_point_made"`
	ThreePointAttempt int `json:"three_point_attemp" bson:"three_point_attemp"`
	ThreePointMade    int `json:"three_point_made" bson:"three_point_made"`
}
