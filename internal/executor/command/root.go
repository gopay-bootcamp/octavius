package command

import (
	"octavius/internal/executor/command/register"
	"octavius/internal/executor/command/start"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon/health"
	"octavius/internal/executor/daemon/job"
	"octavius/internal/executor/daemon/registration"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "octavius_executor",
	Short: "kubernetes executor of octavius",
	Long:  `kubernetes executor of octavius takes request from cli`,
}

// Execute executes the root command of octavius control plane
func Execute(jobDaemon job.JobServicesClient, registrationDaemon registration.RegistrationServicesClient, healthDaemon health.HeathServicesClient, executorConfig config.OctaviusExecutorConfig) error {

	registerCmd := register.NewCmd(registrationDaemon, executorConfig)
	rootCmd.AddCommand(registerCmd)

	startCmd := start.NewCmd(jobDaemon, healthDaemon, executorConfig)
	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
