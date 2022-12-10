package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = initLogger()

func initLogger() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()

	level := zapcore.Level(-1)
	cfg.Level = zap.NewAtomicLevelAt(level)
	logger, _ := cfg.Build()
	defer logger.Sync()
	return logger.Sugar()
}
