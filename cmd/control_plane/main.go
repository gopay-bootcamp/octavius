package main

import (
	"octavius/internal/control_plane/command"
	"octavius/internal/control_plane/config"
	"octavius/internal/control_plane/logger"
)

func main() {
	logLevel := config.Config().LogLevel
	logger.Setup(logLevel)
	command.Execute()
}
