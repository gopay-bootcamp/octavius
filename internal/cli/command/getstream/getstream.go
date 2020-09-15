package getstream

import (
	"fmt"
	"github.com/fatih/color"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"

	"github.com/spf13/cobra"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {

	var jobID string

	getStreamCmd := &cobra.Command{
		Use:   "getstream",
		Short: "Get job log data",
		Long:  `Get job log by giving arguments`,
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			client := &client.GrpcClient{}
			logResponse, err := octaviusDaemon.GetStreamLog(jobID, client)
			if err != nil {
				log.Error(err, "error while getting the stream")
				printer.Println("error while getting the stream", color.FgRed)
				return
			}
			log.Info(fmt.Sprintln(logResponse))
			for _, logs := range *logResponse {
				log.Info(fmt.Sprintf("%v", logs.Log))
				printer.Println(fmt.Sprintf("%v", logs.Log), color.FgMagenta)
			}
		},
	}

	getStreamCmd.Flags().StringVarP(&jobID, "job-id", "", "", "It contains jobID")
	getStreamCmd.MarkFlagRequired("job-id")
	return getStreamCmd
}
