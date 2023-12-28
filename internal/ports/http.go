package ports

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/svbnbyrk/nba/internal/core/usecase"
)

func SimulateHandler(uc usecase.SimulationUsecaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _ = uc.StartSimulation(context.Background(), 5)
		return nil
	}
}
