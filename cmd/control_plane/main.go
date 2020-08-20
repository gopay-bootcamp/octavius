package main

import (
	"octavius/internal/control_plane/command"
	"octavius/internal/control_plane/logger"
)

func main() {
	Logger := logger.Setup()
	command.Execute(Logger)
}
