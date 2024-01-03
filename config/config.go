package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DB_URL                       = "DB_URL"
	HTTP_SERVER_PORT             = "HTTP_SERVER_PORT"
	HTTP_SERVER_TIMEOUT_READ     = "HTTP_SERVER_TIMEOUT_READ"
	HTTP_SERVER_TIMEOUT_WRITE    = "HTTP_SERVER_TIMEOUT_WRITE"
	HTTP_SERVER_TIMEOUT_SHUTDOWN = "HTTP_SERVER_TIMEOUT_SHUTDOWN"
	TIME_FACTOR                  = "TIME_FACTOR"
)

// InitConfig init app config.
func InitConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	return nil
}
