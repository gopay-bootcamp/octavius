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
	if err := octlog.Init("info", logfilePath, true); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize config %v", err))
	}

	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)

	err := command.Execute(octaviusDaemon)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
