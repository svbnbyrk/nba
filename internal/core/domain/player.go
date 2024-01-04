package domain

import "gorm.io/gorm"

type Player struct {
	gorm.Model  `json:"-"`
	ID          int          `json:"id" gorm:"id"`
	Fullname    string       `json:"fullname" gorm:"player_name"`
	TeamID      int          `json:"-" gorm:"team_id"`
	PlayerStats []PlayerStat `json:"stats"`
}

type PlayerStat struct {
	gorm.Model        `json:"-"`
	ID                int `json:"id" gorm:"id,primaryKey"`
	GameID            int `json:"-" gorm:"game_id"`
	PlayerID          int `json:"-" gorm:"player_id"`
	TwoPointAttempt   int `json:"two_point_attemp" gorm:"two_point_attemp"`
	TwoPointMade      int `json:"two_point_made" gorm:"two_point_made"`
	ThreePointAttempt int `json:"three_point_attemp" gorm:"three_point_attemp"`
	ThreePointMade    int `json:"three_point_made" gorm:"three_point_made"`
}

type PlayerStatFilter struct {
	GameID   int
	PlayerID int
}
