package logger

import (
	_ "github.com/prometheus/common/log"
	"github.com/rs/zerolog"
	"octavius/internal/config"
	"os"
	"time"
)

var Log *zerolog.Logger

func Setup() {
	var logInit zerolog.Logger
	logLevel, err := zerolog.ParseLevel(config.Config().LogLevel)
	if err != nil {
		logLevel = zerolog.Level(1)
	}
	zerolog.TimeFieldFormat = time.RFC822
	logInit = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Level(logLevel)
	Log = &logInit
}
