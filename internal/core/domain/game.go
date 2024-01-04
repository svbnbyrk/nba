package domain

type Game struct {
	ID         int  `json:"id" gorm:"id"`
	Week       int  `json:"week" gorm:"week"`
	IsFinished bool `json:"is_finished" gorm:"is_finished"`
	AwayTeamID int  `json:"-" gorm:"away_team_id"`
	AwayTeam   Team `json:"away_team" gorm:"foreignKey:AwayTeamID"`
	HomeTeamID int  `json:"-" gorm:"home_team_id"`
	HomeTeam   Team `json:"home_team" gorm:"foreignKey:HomeTeamID"`
}

type GameFilter struct {
	Week       int
	IsFinished *bool
}
