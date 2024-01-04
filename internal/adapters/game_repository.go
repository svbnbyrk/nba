package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
)

type GameRepositoryInterface interface {
	GetGamesByFilter(ctx context.Context, filter domain.GameFilter) ([]domain.Game, error)
	UpsertGame(ctx context.Context, game domain.Game) error
}

type GameRepository struct {
	*db.Gorm
}

func NewGameRepository(db *db.Gorm) *GameRepository {
	return &GameRepository{db}
}

func (r *GameRepository) GetGamesByFilter(ctx context.Context, filter domain.GameFilter) ([]domain.Game, error) {
	var games []domain.Game
	query := r.Tx.Where("week = ?", filter.Week)
	if filter.IsFinished != nil {
		query = query.Where("is_finished = ?", filter.IsFinished)
	}

	if err := query.Find(&games).Error; err != nil {
		return nil, err
	}

	for i, game := range games {
		if err := r.Tx.Preload("AwayTeam.TeamStats", "game_id = ?", game.ID).
			Preload("AwayTeam.Players").
			Preload("AwayTeam.Players.PlayerStats", "game_id = ?", game.ID).
			Preload("HomeTeam.TeamStats", "game_id = ?", game.ID).
			Preload("HomeTeam.Players").
			Preload("HomeTeam.Players.PlayerStats", "game_id = ?", game.ID).
			First(&games[i], game.ID).Error; err != nil {
			return nil, err
		}
	}

	return games, nil
}

func (r *GameRepository) UpsertGame(ctx context.Context, game domain.Game) error {
	result := r.Tx.WithContext(ctx).Save(&game)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
