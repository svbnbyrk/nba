package adapters

import (
	"context"

	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/db"
)

type TeamRepositoryInterface interface {
	UpsertTeamStat(ctx context.Context, teamStat domain.TeamStat) error
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
	println(result.RowsAffected)
	return nil
}
