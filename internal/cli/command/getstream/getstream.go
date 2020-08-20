package getstream

import (
	"fmt"
	"github.com/spf13/cobra"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "getstream",
		Short: "Get job log data",
		Long:  `Get job log by giving arguments`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Incorrect command argument format, the correct format is: \n octavius getstream <job-name>")
				return
			}
			jobName := args[0]
			client := &client.GrpcClient{}
			err := octaviusDaemon.GetStreamLog(jobName, client)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}
