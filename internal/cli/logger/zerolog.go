package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"octavius/internal/cli/printer"
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
	f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("log file opening issue, aborting session")
		os.Exit(1)
	}
	logInit := zerolog.New(f).With().Timestamp().CallerWithSkipFrameCount(constant.LoggerSkipFrameCount).Logger().Level(logLevel)
	zerolog.TimeFieldFormat = time.RFC822
	logInit.Level(logLevel)
	log = Logger{
		logger: &logInit,
	}
	return
}

func Debug(msg string, printer printer.Printer) {
	printer.Println(fmt.Sprint(msg), color.FgHiGreen)
	log.logger.Debug().Msg(msg)
}

func Warn(msg string, printer printer.Printer) {
	printer.Println(fmt.Sprint(msg), color.FgYellow)
	log.logger.Warn().Msg(msg)
}

func Panic(msg string, err error, printer printer.Printer) {
	printer.Println(fmt.Sprint(msg, err), color.FgHiRed)
	log.logger.Panic().Msgf(msg, err)
}

func Info(msg string, printer printer.Printer) {
	printer.Println(fmt.Sprint(msg), color.FgHiGreen)
	log.logger.Info().Msg(msg)
}

func Fatal(msg string, printer printer.Printer) {
	printer.Println(fmt.Sprint(msg, "\n"), color.FgHiRed)
	log.logger.Fatal().Msgf(msg)
}

func Error(err error, msg string, printer printer.Printer) {
	log.logger.Err(err).Msg(msg)
	if err != nil {
		printer.Println(fmt.Sprint(msg, err), color.FgHiRed)
	} else {
		printer.Println(fmt.Sprint(msg), color.FgHiGreen)
	}
}
