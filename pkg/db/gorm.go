package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
	"github.com/svbnbyrk/nba/internal/core/domain"
	"github.com/svbnbyrk/nba/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

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

	sql, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres instance - sql.Open %w", err)
	}

	goose.SetBaseFS(embedMigrations)

	if err = goose.SetDialect("postgres"); err != nil {
		return nil, fmt.Errorf("goose.SetDialect %w", err)
	}

	if err := goose.Up(sql, "migrations"); err != nil {
		return nil, fmt.Errorf("goose.Up %w", err)
	}

	return &Gorm{
		Tx: db,
	}, nil
}
