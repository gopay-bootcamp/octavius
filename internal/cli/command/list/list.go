package list

import (
	"fmt"
	"github.com/fatih/color"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/printer"

	"github.com/spf13/cobra"
)

// NewCmd create a command for list
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Get job list",
		Long:  `Get job list will give available jobs in octavius`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			client := &client.GrpcClient{}
			jobList, err := octaviusDaemon.GetJobList(client)
			if err != nil {
				printer.Println(fmt.Sprintln("error when getting job list"), color.FgRed)
			}
			if len(jobList.Jobs) == 0 {
				printer.Println(fmt.Sprintln("No jobs available"), color.FgGreen)
			}
			for _, job := range jobList.Jobs {
				printer.Println(fmt.Sprintln(job), color.FgGreen)
			}
		},
	}
}
