package logger

import (
	_ "github.com/prometheus/common/log"
	"github.com/rs/zerolog"
	"octavius/internal/config"
	"os"
)

var Log zerolog.Logger

func Setup() zerolog.Logger {
	Log = zerolog.New(os.Stdout).With().Caller().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel, err := zerolog.ParseLevel(config.Config().LogLevel)
	if err != nil {
		Log.Panic().Err(err).Msg("Config file load error")
	}
	Log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	Log.Level(logLevel)
	return Log
}

