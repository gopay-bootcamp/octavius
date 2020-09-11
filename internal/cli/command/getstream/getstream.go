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
				log.Error(err, "error while getting the stream")
				printer.Println("error while getting the stream", color.FgRed)
			}
			log.Info(fmt.Sprintln(logResponse))
			for _,logs := range *logResponse{
				printer.Println(fmt.Sprintf("%v",logs.Log), color.FgMagenta)
			}
		},
	}
}
