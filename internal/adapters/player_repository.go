package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
)

type PlayerRepositoryInterface interface {
	UpsertPlayerStats(ctx context.Context, playerStat domain.PlayerStat) error
	GetPlayerStat(ctx context.Context, filter domain.PlayerStatFilter) (domain.PlayerStat, error)
	GetPlayerStats(ctx context.Context) ([]domain.PlayerAggregatedStat, error)
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

func (r *PlayerRepository) GetPlayerStats(ctx context.Context) ([]domain.PlayerAggregatedStat, error) {

	var aggregatedStats []domain.PlayerAggregatedStat
	result := r.Tx.Model(&domain.PlayerStat{}).
		Select([]string{
			"player_id",
			"SUM(two_point_attempt) as two_point_attempt_total",
			"SUM(two_point_made) as two_point_made_total",
			"SUM(three_point_attempt) as three_point_attempt_total",
			"SUM(three_point_made) as three_point_made_total",
			"AVG(two_point_made::decimal / NULLIF(two_point_attempt, 0)) as two_point_percentage",
			"AVG(three_point_made::decimal / NULLIF(three_point_attempt, 0)) as three_point_percentage",
		}).
		Group("player_id").
		Find(&aggregatedStats)

	if result.Error != nil {
		return nil, result.Error
	}

	return aggregatedStats, nil
}
