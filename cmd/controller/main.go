package main

import (
	"fmt"
	"log"
	"octavius/internal/controller/command"
	"octavius/internal/controller/config"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logfilePath := config.Config().LogFilePath
	logLevel := config.Config().LogLevel
	if err := octlog.Init(logLevel, logfilePath, true); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize config %v", err))
	}

	err := command.Execute()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
