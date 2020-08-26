package main

import (
	"octavius/internal/executor/client"
	"octavius/internal/executor/command"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon"
	"octavius/internal/executor/logger"
)

func main() {
	logLevel := config.Config().LogLevel
	logger.Setup(logLevel)
	executorDaemon := daemon.NewExecutorClient(&client.GrpcClient{})
	command.Execute(executorDaemon)
}
