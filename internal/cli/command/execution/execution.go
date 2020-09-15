package execution

import (
	"fmt"
	"github.com/fatih/color"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"
	"strings"

	"github.com/spf13/cobra"
)

// NewCmd create a command for execution
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	var jobName string
	var jobArgs string
	executionCmd := &cobra.Command{
		Use:     "execute",
		Short:   "Execute the existing job",
		Long:    "This command helps to execute the job which is already created in server",
		Example: fmt.Sprintf("octavius execute --job-name <job-name> --args arg1=value1,arg2=value2"),

		Run: func(cmd *cobra.Command, args []string) {

			printer.Println(fmt.Sprintf("Job %s is being added to pending list", jobName), color.FgBlack)
			args = strings.Split(jobArgs, ",")
			jobData := map[string]string{}

			for i := 0; i < len(args); i++ {
				arg := strings.Split(args[i], "=")
				jobData[arg[0]] = arg[1]
			}

			client := &client.GrpcClient{}
			response, err := octaviusDaemon.ExecuteJob(jobName, jobData, client)
			if err != nil {
				log.Error(err, "error in executing job")
				printer.Println("error in executing job", color.FgRed)
				return
			}
			log.Info(response.Status)
			printer.Println(fmt.Sprintf("Job has been added to pending list successfully.\nYou can see the execution logs using getstream %s", response.Status), color.FgGreen)
		},
	}

	executionCmd.Flags().StringVarP(&jobName, "job-name", "", "", "It contains Job Name")
	executionCmd.MarkFlagRequired("job-name")
	executionCmd.Flags().StringVarP(&jobArgs, "args", "", "", "It contains Job arguments")

	return executionCmd
}
