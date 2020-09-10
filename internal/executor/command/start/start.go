package start

import (
	"octavius/internal/executor/config"
	"octavius/internal/executor/daemon"
	"octavius/internal/pkg/log"
	"fmt"
	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(executorDaemon daemon.Client, executorConfig config.OctaviusExecutorConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Executor",
		Long:  `Start Executor for communicating with Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			err := executorDaemon.StartClient(executorConfig)
			if err != nil {
				log.Error(err, "failed to configure client, see config")
			}
			log.Info(fmt.Sprintf("executor with id: %s started",executorConfig.ID))
			go executorDaemon.StartPing()
			executorDaemon.StartKubernetesService()
		},
	}
}
