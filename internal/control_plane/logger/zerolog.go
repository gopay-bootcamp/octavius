package logger

import (
	"octavius/internal/config"
	"octavius/pkg/constant"
	"os"
	"time"

	"github.com/rs/zerolog"
)

//Logger holds the pointer to zerolog.Logger object
type Logger struct {
	logger *zerolog.Logger
}

var log Logger

//Setup intializes the logger object
func Setup() {
	if (log != Logger{}) {
		return
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	logLevel, err := zerolog.ParseLevel(config.Config().LogLevel)
	if err != nil {
		logLevel, _ = zerolog.ParseLevel("info")
	}
	logInit := zerolog.New(consoleWriter).With().Timestamp().CallerWithSkipFrameCount(constant.LoggerSkipFrameCount).Logger().Level(logLevel)
	zerolog.TimeFieldFormat = time.RFC822
	logInit.Level(logLevel)
	log = Logger{
		logger: &logInit,
	}
	return
}

//Debug logs the message at debug level
func Debug(msg string) {
	log.logger.Debug().Msg(msg)
}

//Info logs the message at info level
func Info(msg string) {
	log.logger.Info().Msg(msg)
}

//Warn logs the message at warn level
func Warn(msg string) {
	log.logger.Warn().Msg(msg)
}

//Error logs the message and error at error level if err is not nil else it logs message at info level
func Error(err error, msg string) {
	log.logger.Err(err).Msg(msg)
}

//Fatal logs the message at fatal level followed by a os.Exit(1) call
func Fatal(msg string) {
	log.logger.Fatal().Msgf(msg)
}

//Panic logs the message and error at panic level and stops the ordinary flow of goroutine
func Panic(msg string, err error) {
	log.logger.Panic().Err(err).Msg(msg)
}
