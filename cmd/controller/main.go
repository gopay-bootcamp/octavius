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
		log.Fatal(fmt.Sprintf("failed to initialize kubeconfig %v", err))
	}

	command.Execute()
}
