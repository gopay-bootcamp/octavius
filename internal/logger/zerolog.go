package logger

import (
	"github.com/rs/zerolog"
	"octavius/internal/config"
	"os"
)

type Logger struct {
	logger *zerolog.Logger
}

var log Logger

func Setup() {
	if (log != Logger{}) {
		return
	}
	logInit := zerolog.New(os.Stdout).With().Logger().Level(1)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel, err := zerolog.ParseLevel(config.Config().LogLevel)
	if err != nil {
		Panic("Config file load error", err)
	}
	logInit.Output(zerolog.ConsoleWriter{Out: os.Stderr})
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

func Error(action string, err error) {
	log.logger.Print("Error caused ", log.logger.Err(err), " due to ", action)
}

func LogErrors(err error, action string) {
	if err != nil {
		Error(action, err)
	} else {
		Debug(action)
	}
}
