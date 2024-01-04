package usecase

import (
	"context"

	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type GameUsecaseInterface interface {
	GetGamesByWeek(ctx context.Context, week int) ([]domain.Game, error)
}

type GameUsecase struct {
	gameRepo adapters.GameRepositoryInterface
}

func NewGameUsecase(gameRepo adapters.GameRepositoryInterface) *GameUsecase {
	return &GameUsecase{
		gameRepo: gameRepo,
	}
}

func (uc *GameUsecase) GetGamesByWeek(ctx context.Context, week int) ([]domain.Game, error) {
	games, err := uc.gameRepo.GetGamesByFilter(ctx, domain.GameFilter{
		Week: week,
	})
	if err != nil {
		return nil, err
	}

	return games, nil
}
