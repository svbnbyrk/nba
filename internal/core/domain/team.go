package domain

type Team struct {
	ID           int        `json:"id" bson:"_id"`
	Name         string     `json:"name" bson:"name"`
	Abbreviation string     `json:"abbreviation" bson:"abbreviation"`
	Players      []Player   `json:"players,omitempty" bson:"-"`
	Win          int        `json:"win" bson:"win"`
	Lose         int        `json:"lose" bson:"lose"`
	Stats        []TeamStat `json:"stats,omitempty" bson:"-"`
}

type TeamStat struct {
	ID           int `json:"id" bson:"_id"`
	GameID       int `json:"game_id" bson:"game_id"`
	TeamID       int `json:"team_id" bson:"team_id"`
	Score        int `json:"score" bson:"team_id"`
	TotalAttemps int `json:"total_attemps" bson:"total_attemps"`
}
