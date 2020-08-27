package start

import (
	"fmt"
	"octavius/internal/executor/daemon"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(executorDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Executor",
		Long:  `Start Executor for communicating with Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			executorDaemon.StartClient()
			executorDaemon.StartPing()

			fmt.Println("Starting Executor")
		},
	}
}

func init() {
}
