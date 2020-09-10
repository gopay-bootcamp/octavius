package main

import (
	"fmt"
	"log"
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logfilePath := config.LogFilePath
	logFileSize := config.LogFileSize
	if err := octlog.Init("info", logfilePath, false, logFileSize); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize logger %v", err))
	}

	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)

	err := command.Execute(octaviusDaemon)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
