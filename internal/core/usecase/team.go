package usecase

import (
	"context"

	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type TeamUsecaseInterface interface {
	GetScoreboard(ctx context.Context, week int) ([]domain.Game, error)
}

type TeamUsecase struct {
	teamRepo   adapters.TeamRepositoryInterface
	gameRepo   adapters.GameRepositoryInterface
	playerRepo adapters.PlayerRepositoryInterface
}

func NewTeamUsecase(teamRepo adapters.TeamRepositoryInterface, gameRepo adapters.GameRepositoryInterface, playerRepo adapters.PlayerRepositoryInterface) *TeamUsecase {
	return &TeamUsecase{
		teamRepo:   teamRepo,
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
	}
}
