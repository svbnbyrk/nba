package domain

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model   `json:"-"`
	ID           int        `json:"id" gorm:"id"`
	Name         string     `json:"name" gorm:"name"`
	Abbreviation string     `json:"abbreviation" gorm:"abbreviation"`
	Win          int        `json:"win" gorm:"win"`
	Lose         int        `json:"lose" gorm:"lose"`
	Players      []Player   `json:"players"`
	TeamStats    []TeamStat `json:"stats"`
}

type TeamStat struct {
	gorm.Model   `json:"-"`
	ID           int  `json:"id" gorm:"id"`
	GameID       int  `json:"-" gorm:"game_id"`
	TeamID       int  `json:"-" gorm:"team_id"`
	Team         Team `json:"-" gorm:"foreignKey:TeamID"`
	Score        int  `json:"score" gorm:"score"`
	TotalAttemps int  `json:"total_attemps" gorm:"total_attemps"`
}
