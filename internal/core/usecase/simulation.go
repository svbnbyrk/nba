package usecase

import (
	"context"

	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type SimulationUsecaseInterface interface {
	StartSimulation(ctx context.Context, week int) (*domain.Simulation, error)
}

type SimulationUsecase struct {
	teamRepo adapters.TeamRepositoryInterface
}

func NewSimulationUsecase(teamRepo adapters.TeamRepositoryInterface) *SimulationUsecase {
	return &SimulationUsecase{
		teamRepo: teamRepo,
	}
}
