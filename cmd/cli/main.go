package main

import (
	"log"
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon"
	octlog "octavius/internal/pkg/log"
)

func main() {
	// TODO: get log level from config and define log file path if any
	if err := octlog.Init("info", "./logs.txt"); err != nil {
		log.Fatal("fail to initialize octavius log")
	}

	clientConfigLoader := config.NewLoader()
	octaviusDaemon := daemon.NewClient(clientConfigLoader)

	command.Execute(octaviusDaemon)
}
