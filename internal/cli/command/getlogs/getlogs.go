package getlogs

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {

	var jobID string

	getLogsCmd := &cobra.Command{
		Use:   "getlogs",
		Short: "Get job log data",
		Long:  `Get job log by giving arguments`,
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			client := &client.GrpcClient{}
			logResponse, err := octaviusDaemon.GetLogs(jobID, client)
			if err != nil {
				log.Error(err, "error while getting the logs")
				printer.Println(fmt.Sprintf("error while getting the logs, %v", status.Convert(err).Message()), color.FgRed)
				return
			}
			log.Info(fmt.Sprintln(logResponse))
			printer.Println(fmt.Sprintf("%v", logResponse.Log), color.FgMagenta)
		},
	}

	getLogsCmd.Flags().StringVarP(&jobID, "job-id", "", "", "It contains jobID")
	err := getLogsCmd.MarkFlagRequired("job-id")
	if err != nil {
		log.Error(err, "error while setting the flag required")
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	return getLogsCmd
}
