package main

import (
	"log"
	"octavius/internal/controller/command"
	"octavius/internal/controller/config"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logLevel := config.Config().LogLevel
	//TODO: add log file path if any.
	if err := octlog.Init(logLevel, "./auditLogs.txt"); err != nil {
		log.Fatal("fail to initialize octavius log")
	}
	command.Execute()
}
