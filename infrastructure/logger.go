package infrastructure

import (
	"go.uber.org/zap"
)

//Logger struct
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger sets up new logger
func NewLogger() Logger {
	config := zap.NewDevelopmentConfig()
	logger, _ := config.Build()
	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}
}
