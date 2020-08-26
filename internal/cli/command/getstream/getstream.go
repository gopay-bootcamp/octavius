package getstream

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
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
			if err != nil {
				log.Error(err, "error when getting stream")
			}
			for _, logResp := range *logResponse {
				log.Info(fmt.Sprintln(logResp.Log))
			}
		},
	}
}
