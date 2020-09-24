package main

import (
	"fmt"
	"log"
	healthclient "octavius/internal/executor/client/health"
	jobclient "octavius/internal/executor/client/job"
	registrationclient "octavius/internal/executor/client/registration"
	"octavius/internal/executor/command"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon/health"
	"octavius/internal/executor/daemon/job"
	"octavius/internal/executor/daemon/registration"
	octlog "octavius/internal/pkg/log"
)

func main() {
	executorConfig, err := config.Loader()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init config %v", err))
	}
	logfilePath := executorConfig.LogFilePath
	logLevel := executorConfig.LogLevel
	logFileSize := executorConfig.LogFileSize
	if err = octlog.Init(logLevel, logfilePath, true, logFileSize); err != nil {
		log.Fatal(fmt.Sprintf("failed to initialize logger %v", err))
	}

	jobDaemon := job.NewJobServicesClient(&jobclient.GrpcClient{})
	registrationDaemon := registration.NewRegistrationServicesClient(&registrationclient.GrpcClient{})
	healthDaemon := health.NewHealthServicesClient(&healthclient.GrpcClient{})

	err = command.Execute(jobDaemon, registrationDaemon, healthDaemon, executorConfig)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to execute command %v", err))
	}
}
