package execution

import (
	"fmt"
	"github.com/spf13/cobra"
	"octavius/internal/cli/daemon"
	"strings"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:     "execute",
		Short:   "Execute the existing job",
		Long:    "This command helps to execute the job which is already created in server",
		Example: fmt.Sprintf("octavius execute <job-name> arg1=argvalue1 arg2=argvalue2"),
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Incorrect command argument format, the correct format is: \n octavius execute <job-name> arg1=argvalue1 arg2=argvalue2 ...")
				return
			}
			jobName := args[0]
			jobData := map[string]string{}

			for i := 1; i < len(args); i++ {
				arg := strings.Split(args[i], "=")
				jobData[arg[0]] = arg[1]
			}

			err := octaviusDaemon.Execute(jobName, jobData)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}
