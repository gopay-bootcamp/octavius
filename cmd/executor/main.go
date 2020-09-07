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
	executorConfig, err := config.Loader()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config %v", err))
	}
	logfilePath := executorConfig.LogFilePath
	logLevel := executorConfig.LogLevel
	if err = octlog.Init(logLevel, logfilePath, true); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize logger %v", err))
	}

	executorDaemon := daemon.NewExecutorClient(&client.GrpcClient{})
	err = command.Execute(executorDaemon, executorConfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
