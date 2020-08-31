package log

import (
	"github.com/rs/zerolog"
	"io/ioutil"
	"octavius/internal/pkg/constant"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

type engine struct {
	logger *zerolog.Logger
}

var logEngine engine // contain cli configarution

// Init intializes the logger object
func Init(configLogLevel string, logFile string, logInConsole bool) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	var (
		f       *os.File
		multi   zerolog.LevelWriter
		logPath string
	)
	dirName := filepath.Join(usr.HomeDir, ".octavius")
	logLevel, err := zerolog.ParseLevel(configLogLevel)
	if err != nil {
		return err
	}

	err = createDir(dirName)
	if err != nil {
		return err
	}

	if logFile == "" {
		logPath = filepath.Join(dirName, "test.log")
	} else {
		logPath = filepath.Join(dirName, logFile)
	}

	err = createFile(logPath)
	if err != nil {
		return err
	}

	f, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return err
	}

	if logInConsole {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
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

func createFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		// if it's error other than IsNotExist return it
		if !os.IsNotExist(err) {
			return err
		}
		// if err is file is not exist then we create the file
		if err = ioutil.WriteFile(path, []byte(""), 0644); err != nil {
			return err
		}
	}
	return nil
}

func createDir(path string) error {
	var err = os.Mkdir(path, 0755)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
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
