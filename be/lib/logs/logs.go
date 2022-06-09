package logs

import (
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

type Format string

const (
	FormatProductionJSON Format = "productionJSON"
	FormatDevelopment           = "development"
)

func NewLogger(format Format) *zap.Logger {
	var logger *zap.Logger
	var err error

	switch format {
	case FormatProductionJSON:
		logger, err = zapdriver.NewProduction()
	case FormatDevelopment:
		logger, err = zap.NewDevelopment()
	default:
		panic("cannot determine logger format")
	}

	if err != nil {
		panic(err)
	}

	return logger
}

func SyncLogger(logger *zap.Logger) {
	_ = logger.Sync()
}
