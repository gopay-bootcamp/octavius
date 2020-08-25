package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"octavius/internal/cli/printer"
	"octavius/internal/pkg/constant"
	"os"
	"time"
)

type Logger struct {
	logger *zerolog.Logger
}

var log Logger
var colorPrinter = printer.GetPrinter()

func Setup() {
	if (log != Logger{}) {
		return
	}
	logLevel, err := zerolog.ParseLevel("info")
	if err != nil {
		fmt.Println("log level parsing problem")
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

func Debug(msg string) {
	colorPrinter.Println(fmt.Sprint(msg), color.FgHiGreen)
	log.logger.Debug().Msg(msg)
}

func Warn(msg string) {
	colorPrinter.Println(fmt.Sprint(msg), color.FgYellow)
	log.logger.Warn().Msg(msg)
}

func Panic(msg string, err error) {
	colorPrinter.Println(fmt.Sprint(msg, err))
	log.logger.Panic().Msgf(msg, err)
}

func Info(msg string) {
	colorPrinter.Println(fmt.Sprint(msg), color.FgHiGreen)
	log.logger.Info().Msg(msg)
}

func Fatal(msg string) {
	colorPrinter.Println(fmt.Sprint(msg, "\n"), color.FgHiRed)
	log.logger.Fatal().Msgf(msg)
}

func Error(err error, msg string) {
	log.logger.Err(err).Msg(msg)
	if err != nil {
		colorPrinter.Println(fmt.Sprint(msg, err), color.FgHiRed)
	} else {
		colorPrinter.Println(fmt.Sprint(msg), color.FgHiGreen)
	}
}
