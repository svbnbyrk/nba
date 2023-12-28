package log

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zapcore.InfoLevel)
	if os.Getenv("env") == "local" {
		cfg.Level.SetLevel(zapcore.DebugLevel)
	}
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	zapLogger, err := cfg.Build()
	if err != nil {
		log.Fatalf("Fail to build zap logger. err: %s", err)
	}

	logger = zapLogger.With(zap.String("app", "nba"))
}

func Logger() *zap.Logger {
	return logger
}
