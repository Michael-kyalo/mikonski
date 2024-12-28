package logging

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

// InitLogger initializes the global logger.
func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

// GetLogger returns the global logger instance.
func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

// Sync flushes any buffered log entries.
func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}
