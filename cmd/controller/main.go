package main

import (
	"octavius/internal/controller/command"
	"octavius/internal/controller/config"
	"octavius/internal/pkg/log"
)

func main() {
	logLevel := config.Config().LogLevel
	//TODO: add log file path if any.
	log.Init(logLevel, "")
	command.Execute()
}
