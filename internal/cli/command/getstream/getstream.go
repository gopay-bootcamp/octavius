package getstream

import (
	"fmt"
	"github.com/spf13/cobra"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/logger"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "getstream",
		Short: "Get job log data",
		Long:  `Get job log by giving arguments`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			jobName := args[0]
			client := &client.GrpcClient{}
			logResponse, err := octaviusDaemon.GetStreamLog(jobName, client)
			logger.Error(err, "Getting Stream")
			for _, log := range *logResponse {
				logger.Info(fmt.Sprintln(log.Log))
			}
		},
	}
}
