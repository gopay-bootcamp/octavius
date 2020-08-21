package logger

import (
	"github.com/rs/zerolog"
	"octavius/internal/config"
	"octavius/pkg/constant"
	"os"
	"time"
)

type Logger struct {
	logger *zerolog.Logger
}

var log Logger

func Setup() {
	if (log != Logger{}) {
		return
	}
	logLevel, err := zerolog.ParseLevel(config.Config().LogLevel)
	if err != nil {
		logLevel, _ = zerolog.ParseLevel("info")
	}
	logInit := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(constant.LoggerSkipFrameCount).Logger().Level(logLevel)
	zerolog.TimeFieldFormat = time.RFC822
	logInit.Level(logLevel)
	log = Logger{
		logger: &logInit,
	}
	return
}

func Debug(msg string) {
	log.logger.Debug().Msg(msg)
}

func Warn(msg string) {
	log.logger.Warn().Msg(msg)
}

func Panic(msg string, err error) {
	log.logger.Panic().Msgf(msg, err)
}

func Info(msg string) {
	log.logger.Info().Msg(msg)
}

func Fatal(msg string) {
	log.logger.Fatal().Msgf(msg)
}

func Error(err error, action string) {
	log.logger.Error().Msgf(action, err.Error())
}

func ErrorCheck(err error, msg string) {
	if err != nil {
		log.logger.Error().Msgf(msg, err.Error())
	} else {
		log.logger.Info().Msg(msg)
	}
}
