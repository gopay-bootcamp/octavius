package main

import (
	"octavius/internal/control_plane/command"
	"octavius/internal/control_plane/logger"
)

func main() {
	logger.Setup()
	command.Execute()
}
