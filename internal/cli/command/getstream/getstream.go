package getstream

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/logger"
	"octavius/internal/cli/printer"
)

func NewCmd(octaviusDaemon daemon.Client, printer printer.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "getstream",
		Short: "Get job log data",
		Long:  `Get job log by giving arguments`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			jobName := args[0]
			client := &client.GrpcClient{}
			logResponse, err := octaviusDaemon.GetStreamLog(jobName, client)
			logger.Error(err, "Getting Stream", printer)
			for _, log := range *logResponse {
				printer.Println(fmt.Sprintln(log.Log), color.FgYellow)
			}
		},
	}
}
