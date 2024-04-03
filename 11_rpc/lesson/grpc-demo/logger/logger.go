package logger

import (
	"github.com/rs/zerolog/log"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

func NewZeroLogLogger() Logger {
	return &ZeroLogLogger{}
}

type ZeroLogLogger struct{}

func (zl *ZeroLogLogger) Info(msg string, args ...any) {
	log.Info().Msgf(msg, args...)
}

func (zl *ZeroLogLogger) Error(msg string, args ...any) {
	log.Error().Msgf(msg, args...)
}
