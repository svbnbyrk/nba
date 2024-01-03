package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
	"gorm.io/gorm"
)

type GameRepositoryInterface interface {
	GetGamesByFilter(ctx context.Context, filter domain.GameFilter) ([]domain.Game, error)
}

type GameRepository struct {
	*db.Gorm
}

func NewGameRepository(db *db.Gorm) *GameRepository {
	return &GameRepository{db}
}

func (r *GameRepository) GetGamesByFilter(ctx context.Context, filter domain.GameFilter) ([]domain.Game, error) {
	var games []domain.Game
	result := r.Tx.Preload("AwayTeam").Preload("AwayTeam.TeamStats").Preload("AwayTeam.Players").Preload("AwayTeam.Players.PlayerStats").Preload("HomeTeam").Preload("HomeTeam.TeamStats").Preload("HomeTeam.Players").Preload("HomeTeam.Players.PlayerStats").Where("week = ? AND is_finished = ?", filter.Week, filter.IsFinished).Find(&games)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return games, nil
}
