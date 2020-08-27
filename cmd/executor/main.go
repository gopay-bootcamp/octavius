package main

import (
	"log"
	"octavius/internal/executor/client"
	"octavius/internal/executor/command"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logLevel := config.Config().LogLevel
	//TODO: add log file path if any.
	if err := octlog.Init(logLevel, ""); err != nil {
		log.Fatal("fail to initialize octavius log")
	}
	executorDaemon := daemon.NewExecutorClient(&client.GrpcClient{})
	command.Execute(executorDaemon)
}
