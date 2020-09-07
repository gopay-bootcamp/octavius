package main

import (
	"fmt"
	"log"
	"octavius/internal/controller/command"
	"octavius/internal/controller/config"
	octlog "octavius/internal/pkg/log"
)

func main() {
	controllerConfig, err := config.Loader()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config %v", err))
	}
	logfilePath := controllerConfig.LogFilePath
	logLevel := controllerConfig.LogLevel
	if err = octlog.Init(logLevel, logfilePath, true); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize logger %v", err))
	}

	err = command.Execute()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
