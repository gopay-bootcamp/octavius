package main

import (
	"octavius/internal/control_plane/command"
	log "octavius/internal/control_plane/logger"
)

func main() {
	log.Setup()
	command.Execute()
}
