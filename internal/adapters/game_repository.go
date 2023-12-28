package adapters

import mongodb "github.com/svbnbyrk/nba/pkg/db"

type GameRepositoryInterface interface {
}

type GameRepository struct {
	*mongodb.MongoDB
}

func NewGameRepository(db *mongodb.MongoDB) *GameRepository {
	return &GameRepository{db}
}
