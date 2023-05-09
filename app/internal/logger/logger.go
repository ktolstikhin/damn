package logger

import (
	"io"
	"time"

	"github.com/rs/zerolog"
)

func New(out io.Writer, level zerolog.Level, json bool) zerolog.Logger {
	if !json {
		out = zerolog.ConsoleWriter{
			Out:        out,
			TimeFormat: time.RFC3339,
		}
	}

	log := zerolog.
		New(out).
		With().
		Timestamp().
		Logger().
		Level(level)

	return log
}
