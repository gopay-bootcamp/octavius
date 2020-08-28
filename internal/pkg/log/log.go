package log

import (
	"octavius/internal/pkg/constant"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type engine struct {
	logger *zerolog.Logger
	// put configurabe field here
}

var logEngine engine // contain cli configarution

// Init intializes the logger object
func Init(configLogLevel string, logFile string, logInConsole bool) error {

	var (
		f     *os.File
		err   error
		multi zerolog.LevelWriter
	)

	logLevel, err := zerolog.ParseLevel(configLogLevel)
	if err != nil {
		return err
	}

	f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	if logInConsole {
		consoleWriter := zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
		multi = zerolog.MultiLevelWriter(f, consoleWriter)
	} else {
		multi = zerolog.MultiLevelWriter(f)
	}

	zerolog.TimeFieldFormat = time.RFC850

	zerologInstance := zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(constant.LoggerSkipFrameCount).Logger().Level(logLevel)
	logEngine = engine{
		logger: &zerologInstance,
	}

	return nil
}

//Debug logs the message at debug level
func Debug(msg string) {
	logEngine.logger.Debug().Msg(msg)
}

//Info logs the message at info level
func Info(msg string) {
	logEngine.logger.Info().Msg(msg)
}

//Warn logs the message at warn level
func Warn(msg string) {
	logEngine.logger.Warn().Msg(msg)
}

//Error logs the message and error at error level if err is not nil else it logs message at info level
func Error(err error, msg string) {
	logEngine.logger.Err(err).Msg(msg)
}

//Fatal logs the message at fatal level followed by a os.Exit(1) call
func Fatal(msg string) {
	logEngine.logger.Fatal().Msgf(msg)
}

//Panic logs the message and error at panic level and stops the ordinary flow of goroutine
func Panic(msg string, err error) {
	logEngine.logger.Panic().Err(err).Msg(msg)
}
