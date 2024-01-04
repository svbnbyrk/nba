package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/svbnbyrk/nba/config"
	"github.com/svbnbyrk/nba/internal/adapters"
	"github.com/svbnbyrk/nba/internal/core/usecase"
	"github.com/svbnbyrk/nba/internal/ports"
	"github.com/svbnbyrk/nba/pkg/db"
	"github.com/svbnbyrk/nba/pkg/http"
	"github.com/svbnbyrk/nba/pkg/log"
	"go.uber.org/zap"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Logger().Sugar().Fatalf("Config error: %w", err)
	}

	// connect mongodb
	database, err := db.New(viper.GetString(config.DB_URL))
	if err != nil {
		log.Logger().Fatal("db.New", zap.String("err", err.Error()))
	}

	teamRepo := adapters.NewTeamRepository(database)
	gameRepo := adapters.NewGameRepository(database)
	playerRepo := adapters.NewPlayerRepository(database)
	simulationUsecase := usecase.NewSimulationUsecase(teamRepo, gameRepo, playerRepo)
	gameUsecase := usecase.NewGameUsecase(gameRepo)
	//teamUsecase := usecase.NewTeamUsecase()

	e := echo.New()
	e = ports.SetupMiddleware(e)

	e.GET("/v1/simulate", ports.SimulateHandler(simulationUsecase))
	e.GET("v1/schedule", ports.GetGameScheduleHandler(gameUsecase))
	// setup http server
	httpServer := http.New(e)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Logger().Info("signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Logger().Error("httpServer.Notify", zap.String("err", err.Error()))

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			log.Logger().Error("httpServer.Shutdown", zap.String("err", err.Error()))
		}
	}
}
