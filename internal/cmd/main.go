package main

import (
	"context"

	"github.com/spf13/viper"
	"github.com/svbnbyrk/nba/config"
	mongodb "github.com/svbnbyrk/nba/pkg/db"
	"github.com/svbnbyrk/nba/pkg/log"
	"go.uber.org/zap"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Logger().Sugar().Fatalf("Config error: %w", err)
	}

	// connect mongodb
	mdb, err := mongodb.New(context.Background(), viper.GetString(config.DB_NAME))
	if err != nil {
		log.Logger().Fatal("mongodb.New", zap.String("err", err.Error()))
	}
	defer mdb.Close()

}
