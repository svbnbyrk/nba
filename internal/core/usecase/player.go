package usecase

import (
	"context"

	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type PlayerUsecaseInterface interface {
	GetPlayersStat(ctx context.Context) ([]domain.PlayerAggregatedStat, error)
}

type PlayerUsecase struct {
	playerRepo adapters.PlayerRepositoryInterface
}

func NewPlayerUsecase(PlayerRepo adapters.PlayerRepositoryInterface) *PlayerUsecase {
	return &PlayerUsecase{
		playerRepo: PlayerRepo,
	}
}

func (uc *PlayerUsecase) GetPlayersStat(ctx context.Context) ([]domain.PlayerAggregatedStat, error) {
	players, err := uc.playerRepo.GetPlayerStats(ctx)
	if err != nil {
		return nil, err
	}

	return players, nil
}
