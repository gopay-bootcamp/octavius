package main

import (
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
)

func main() {
	// TODO: get log level from config and define log file path if any
	log.Init("info", "")

	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)

	command.Execute(octaviusDaemon)
}
