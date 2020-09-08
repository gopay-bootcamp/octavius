package list

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

// NewCmd create a command for list
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Get job list",
		Long:  `Get job list will give available jobs in octavius`,
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			client := &client.GrpcClient{}
			jobList, err := octaviusDaemon.GetJobList(client)
			if err != nil {
				log.Error(err, "error when getting job list")
			}
			if len(jobList.Jobs) == 0 {
				log.Info(fmt.Sprintln("No jobs available"))
			}
			for _, job := range jobList.Jobs {
				log.Info(fmt.Sprintln(job))
			}
		},
	}
}
