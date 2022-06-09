package logs

import (
	"errors"
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

type Format string

const (
	FormatProductionJSON Format = "productionJSON"
	FormatDevelopment           = "development"
)

func NewLogger(format Format) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	switch format {
	case FormatProductionJSON:
		logger, err = zapdriver.NewProduction()
	case FormatDevelopment:
		logger, err = zap.NewDevelopment()
	default:
		return nil, errors.New("cannot determine logger format")
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func SyncLogger(logger *zap.Logger) error {
	return logger.Sync()
}
