package ports

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/svbnbyrk/nba/internal/core/usecase"
	"github.com/svbnbyrk/nba/pkg/log"
	"go.uber.org/zap"
)

func SimulateHandler(uc usecase.SimulationUsecaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		week, err := strconv.Atoi(c.QueryParam("week"))
		if err != nil {
			week = 1
		}

		games, err := uc.StartSimulation(context.Background(), week)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, games)
	}
}

func GetGameScheduleHandler(uc usecase.GameUsecaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		week, err := strconv.Atoi(c.QueryParam("week"))
		if err != nil {
			week = 1
		}

		games, err := uc.GetGamesByWeek(context.Background(), week)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, games)
	}
}

func SetupMiddleware(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var errStr string
			if v.Error != nil {
				errStr = v.Error.Error()
			}
			log.Logger().Info("HTTP server", zap.String(echo.HeaderXRequestID, c.Response().Header().Get(echo.HeaderXRequestID)), zap.String("uri", v.URI), zap.String("method", c.Request().Method), zap.Int("status", v.Status), zap.String("remoteIP", c.Request().RemoteAddr), zap.String("err", errStr))

			return nil
		},
	}))

	return e
}
