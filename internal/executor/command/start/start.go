package start

import (
	"octavius/internal/executor/daemon"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(executorDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Executor",
		Long:  `Start Executor for communicating with Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			err := executorDaemon.StartClient()
			log.Error(err, "client started successfully")
			executorDaemon.StartPing()
		},
	}
}
