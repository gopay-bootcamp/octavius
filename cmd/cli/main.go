package main

import (
	"fmt"
	"log"
	"octavius/internal/cli/command"
	"octavius/internal/cli/config"
	"octavius/internal/cli/daemon/job"
	"octavius/internal/cli/daemon/metadata"
	octlog "octavius/internal/pkg/log"
)

func main() {
	logfilePath := config.LogFilePath
	logFileSize := config.LogFileSize
	if err := octlog.Init("info", logfilePath, false, logFileSize); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize logger %v", err))
	}

	clientConfigLoader := config.NewLoader()
	metadataDaemon := metadata.NewClient(clientConfigLoader)
	jobDaemon := job.NewClient(clientConfigLoader)

	err := command.Execute(jobDaemon, metadataDaemon)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
