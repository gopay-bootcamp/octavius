package execution

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"google.golang.org/grpc/status"
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
		Args:    cobra.MaximumNArgs(0),

		Run: func(cmd *cobra.Command, args []string) {

			printer.Println(fmt.Sprintf("Job %s is being added to pending list", jobName), color.FgBlack)
			args = strings.Split(jobArgs, ",")
			if args[0] == "" {
				args = nil
			}
			jobData := map[string]string{}
			for i := range args {
				arg := strings.Split(args[i], "=")
				if len(arg) != 2 {
					log.Error(errors.New("invalid argument format"), "")
					printer.Println("Invalid argument format. Use octavius execute --help for more details.", color.FgRed)
					return
				}
				jobData[arg[0]] = arg[1]
			}

			client := &client.GrpcClient{}
			response, err := octaviusDaemon.ExecuteJob(jobName, jobData, client)
			if err != nil {
				log.Error(err, "error in executing job")
				printer.Println(fmt.Sprintf("error in executing job, %v", status.Convert(err).Message()), color.FgRed)
				return
			}
			log.Info(response.Status)
			printer.Println(fmt.Sprintf("Job has been added to pending list successfully.\nYou can see the execution logs using getlogs --job-id %s", response.Status), color.FgGreen)
		},
	}

	executionCmd.Flags().StringVarP(&jobName, "job-name", "", "", "It contains Job Name")
	err := executionCmd.MarkFlagRequired("job-name")
	if err != nil {
		log.Error(err, "error while setting the flag required")
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	executionCmd.Flags().StringVarP(&jobArgs, "args", "", "", "It contains Job arguments")

	return executionCmd
}
