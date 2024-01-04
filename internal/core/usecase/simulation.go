package usecase

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/spf13/viper"
	"github.com/svbnbyrk/nba/config"
	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/domain"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_simulation.go -source=./simulation.go -package=usecase
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
	finished := false
	games, err := uc.gameRepo.GetGamesByFilter(ctx, domain.GameFilter{
		Week:       week,
		IsFinished: &finished,
	})
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		go uc.simulateGame(ctx, game)
	}

	return games, nil
}

func (uc *SimulationUsecase) simulateGame(ctx context.Context, game domain.Game) {
	for minute := 0; minute < 48; minute++ {
		uc.simulateMinute(ctx, game, minute)
		time.Sleep(viper.GetDuration(config.TIME_FACTOR) * time.Second)
	}
	game.IsFinished = true
	if game.AwayTeam.TeamStats[0].Score > game.HomeTeam.TeamStats[0].Score {
		game.AwayTeam.Win++
		game.HomeTeam.Lose++
	} else if game.AwayTeam.TeamStats[0].Score < game.HomeTeam.TeamStats[0].Score {
		game.HomeTeam.Win++
		game.AwayTeam.Lose++
	}

	_ = uc.gameRepo.UpsertGame(ctx, game)
	_ = uc.teamRepo.UpsertTeam(ctx, game.AwayTeam)
	_ = uc.teamRepo.UpsertTeam(ctx, game.HomeTeam)
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

	teamStat, err := uc.teamRepo.GetTeamStat(ctx, domain.TeamStatFilter{GameID: gameID, TeamID: team.ID})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	for i := 0; i < attackCount; i++ {
		playerIndex := rand.Intn(len(team.Players)) // Rastgele oyuncu seç
		attackResult, err := uc.simulateAttack(ctx, team.Players[playerIndex], gameID)
		if err != nil {
			return err
		}
		score += attackResult
	}

	teamStat.TotalAttemps += attackCount
	teamStat.Score += score
	teamStat.GameID = gameID
	teamStat.TeamID = team.ID
	err = uc.teamRepo.UpsertTeamStat(ctx, teamStat)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SimulationUsecase) simulateAttack(ctx context.Context, player domain.Player, gameID int) (attackResult int, err error) {
	attackType := rand.Intn(2) // 0: 2 puan, 1: 3 puan
	point := 0
	playerStat, err := uc.playerRepo.GetPlayerStat(ctx, domain.PlayerStatFilter{GameID: gameID, PlayerID: player.ID})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		playerStat.PlayerID = player.ID
		playerStat.GameID = gameID
	}

	if attackType == 0 {
		playerStat.TwoPointAttempt++
		if rand.Float32() < 0.65 { // Basit bir başarı olasılığı
			playerStat.TwoPointMade++
			point = 2
		}
	} else if attackType == 1 {
		playerStat.ThreePointAttempt++
		if rand.Float32() < 0.40 { // Basit bir başarı olasılığı
			playerStat.ThreePointMade++
			point = 3
		}
	}

	err = uc.playerRepo.UpsertPlayerStats(ctx, playerStat)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	return point, nil
}
