package start

import (
	"fmt"
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon/health"
	"octavius/internal/executor/daemon/job"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(jobDaemon job.JobServicesClient, healthDaemon health.HealthServicesClient, executorConfig config.OctaviusExecutorConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Executor",
		Long:  `Start Executor for communicating with Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info(fmt.Sprintf("executor with id: %s started", executorConfig.ID))
			healthDaemon.StartPing(executorConfig)
			jobDaemon.ConfigureKubernetesClient(executorConfig)
			jobDaemon.StartKubernetesService(executorConfig)
		},
	}
}
