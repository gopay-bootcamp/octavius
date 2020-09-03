package main

import (
	"fmt"
	"log"
	"octavius/internal/executor/client"
	"octavius/internal/executor/command"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logfilePath := config.Config().LogFilePath
	logLevel := config.Config().LogLevel
	if err := octlog.Init(logLevel, logfilePath, true); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize kubeconfig %v", err))
	}

	executorDaemon := daemon.NewExecutorClient(&client.GrpcClient{})
	err := command.Execute(executorDaemon)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execution kubeconfig %v", err))
	}
}
