package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
)

type TeamRepositoryInterface interface {
	UpsertTeamStat(ctx context.Context, teamStat domain.TeamStat) error
	UpsertTeam(ctx context.Context, team domain.Team) error
	GetTeamStat(ctx context.Context, filter domain.TeamStatFilter) (domain.TeamStat, error)
}

type TeamRepository struct {
	*db.Gorm
}

func NewTeamRepository(db *db.Gorm) *TeamRepository {
	return &TeamRepository{db}
}

func (r *TeamRepository) UpsertTeamStat(ctx context.Context, teamStat domain.TeamStat) error {
	result := r.Tx.WithContext(ctx).Save(&teamStat)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TeamRepository) UpsertTeam(ctx context.Context, team domain.Team) error {
	result := r.Tx.WithContext(ctx).Save(&team)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TeamRepository) GetTeamStat(ctx context.Context, filter domain.TeamStatFilter) (domain.TeamStat, error) {
	var stat domain.TeamStat

	result := r.Tx.WithContext(ctx).Where("game_id = ? AND team_id = ?", filter.GameID, filter.TeamID).First(&stat)

	if result.Error != nil {
		return stat, result.Error
	}

	return stat, nil
}
