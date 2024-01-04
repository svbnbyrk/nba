package usecase

import (
	"context"

	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type TeamUsecaseInterface interface {
	GetTeams(ctx context.Context, sortParam string) ([]domain.Team, error)
}

type TeamUsecase struct {
	teamRepo adapters.TeamRepositoryInterface
}

func NewTeamUsecase(teamRepo adapters.TeamRepositoryInterface) *TeamUsecase {
	return &TeamUsecase{
		teamRepo: teamRepo,
	}
}

func (uc *TeamUsecase) GetTeams(ctx context.Context, sortParam string) ([]domain.Team, error) {
	teams, err := uc.teamRepo.GetTeams(ctx, domain.TeamFilter{Sort: sortParam})
	if err != nil {
		return nil, err
	}
	return teams, nil
}
