package adapters

import mongodb "github.com/svbnbyrk/nba/pkg/db"

type PlayerRepositoryInterface interface {
}

type PlayerRepository struct {
	*mongodb.MongoDB
}

func NewPlayerRepository(db *mongodb.MongoDB) *PlayerRepository {
	return &PlayerRepository{db}
}
