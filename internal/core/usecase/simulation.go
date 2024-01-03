package usecase

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/svbnbyrk/nba/config"
	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
)

type SimulationUsecaseInterface interface {
	StartSimulation(ctx context.Context, week int) ([]domain.Game, error)
}

type SimulationUsecase struct {
	teamRepo   adapters.TeamRepositoryInterface
	gameRepo   adapters.GameRepositoryInterface
	playerRepo adapters.PlayerRepositoryInterface
}

func NewSimulationUsecase(teamRepo adapters.TeamRepositoryInterface, gameRepo adapters.GameRepositoryInterface, playerRepo adapters.PlayerRepositoryInterface) *SimulationUsecase {
	return &SimulationUsecase{
		teamRepo:   teamRepo,
		gameRepo:   gameRepo,
		playerRepo: playerRepo,
	}
}

func (uc *SimulationUsecase) StartSimulation(ctx context.Context, week int) ([]domain.Game, error) {
	games, err := uc.gameRepo.GetGamesByFilter(ctx, domain.GameFilter{
		Week: week,
	})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	gameUpdate := make(chan string, len(games))

	for _, game := range games {
		wg.Add(1)
		go uc.simulateGame(ctx, game, gameUpdate, &wg)
	}

	wg.Wait()
	close(gameUpdate)

	return games, nil
}

func (uc *SimulationUsecase) simulateGame(ctx context.Context, game domain.Game, updates chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for minute := 0; minute < 48; minute++ {
		uc.simulateMinute(ctx, game, minute)
		time.Sleep(viper.GetDuration(config.TIME_FACTOR) * time.Second)
	}
}

// simulateMinute simulates a single minute of game time.
func (uc *SimulationUsecase) simulateMinute(ctx context.Context, game domain.Game, minute int) {
	// Assume GetAttackCount is a function that calculates attack count for a team.
	awayTeamAttackCount := uc.getRandomAttackCount()
	homeTeamAttackCount := uc.getRandomAttackCount()

	err := uc.simulateTeamAttack(ctx, game.AwayTeam, game.ID, awayTeamAttackCount)
	if err != nil {
		println(err.Error())
	}
	err = uc.simulateTeamAttack(ctx, game.HomeTeam, game.ID, homeTeamAttackCount)
	if err != nil {
		println(err.Error())
	}
}

func (uc *SimulationUsecase) getRandomAttackCount() int {
	return rand.Intn(2) + 1
}

// simulateTeamAttack simulates an attack for a team.
// Additional logic and error handling might be necessary depending on the implementation details.
func (uc *SimulationUsecase) simulateTeamAttack(ctx context.Context, team domain.Team, gameID int, attackCount int) error {
	score := 0
	for i := 0; i < attackCount; i++ {
		playerIndex := rand.Intn(len(team.Players)) // Rastgele oyuncu seç
		attackResult, err := uc.simulateAttack(ctx, team.Players[playerIndex], gameID)
		if err != nil {
			return err
		}
		score += attackResult
	}

	return nil
}

func (uc *SimulationUsecase) ensurePlayerStats(ctx context.Context, player domain.Player, gameID int) (domain.PlayerStat, error) {
	var playerStats domain.PlayerStat
	// Assuming that you have a method to get the current game's player stat
	// Check if there's existing stats, if not, create a new one
	if len(player.PlayerStats) == 0 {
		// Handle creation and possible errors properly
		playerStats.PlayerID = player.ID
		playerStats.GameID = gameID
		err := uc.playerRepo.UpsertPlayerStats(ctx, playerStats)
		if err != nil {
			return playerStats, err
		}
		return playerStats, nil
	}
	return player.PlayerStats[0], nil // Assuming that player always has at least one stat
}

func (uc *SimulationUsecase) simulateAttack(ctx context.Context, player domain.Player, gameID int) (attackResult int, err error) {
	attackType := rand.Intn(2) // 0: 2 puan, 1: 3 puan
	playerStats, err := uc.ensurePlayerStats(ctx, player, gameID)
	if err != nil {
		return 0, err
	}

	if attackType == 0 {
		playerStats.TwoPointAttempt++
		if rand.Float32() < 0.65 { // Basit bir başarı olasılığı
			playerStats.TwoPointMade++
			return 2, err
		} else {
			return 0, err
		}
	} else if attackType == 1 {
		playerStats.ThreePointAttempt++
		if rand.Float32() < 0.40 { // Basit bir başarı olasılığı
			playerStats.ThreePointMade++
			return 3, err
		} else {
			return 0, err
		}
	}

	err = uc.playerRepo.UpsertPlayerStats(ctx, playerStats)
	if err != nil {
		return 0, err
	}

	return 0, err
}
