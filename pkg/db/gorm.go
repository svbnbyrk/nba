package db

import (
	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	Tx *gorm.DB
}

func New(url string) (*Gorm, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Logger().Fatal("connecting database", zap.Error(err))
	}

	// Perform database migration
	err = db.AutoMigrate(domain.Team{}, domain.Game{}, domain.Player{}, domain.PlayerStat{}, domain.TeamStat{})
	if err != nil {
		log.Logger().Fatal("database auto migration", zap.Error(err))
	}

	return &Gorm{
		Tx: db,
	}, nil
}
