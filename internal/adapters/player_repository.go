package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
)

type PlayerRepositoryInterface interface {
	UpsertPlayerStats(ctx context.Context, playerStat domain.PlayerStat) error
	GetPlayerStat(ctx context.Context, filter domain.PlayerStatFilter) (domain.PlayerStat, error)
}

type PlayerRepository struct {
	*db.Gorm
}

func NewPlayerRepository(db *db.Gorm) *PlayerRepository {
	return &PlayerRepository{db}
}

func (r *PlayerRepository) UpsertPlayerStats(ctx context.Context, playerStat domain.PlayerStat) error {
	result := r.Tx.WithContext(ctx).Save(&playerStat)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PlayerRepository) GetPlayerStat(ctx context.Context, filter domain.PlayerStatFilter) (domain.PlayerStat, error) {
	var stat domain.PlayerStat

	result := r.Tx.WithContext(ctx).Where("game_id = ? AND player_id = ?", filter.GameID, filter.PlayerID).First(&stat)

	if result.Error != nil {
		return stat, result.Error
	}

	return stat, nil
}
