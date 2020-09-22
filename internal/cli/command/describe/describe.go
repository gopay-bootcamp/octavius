//Package describe provides cli command to get the job description of a specific job
package describe

import (
	"fmt"

	client "octavius/internal/cli/client/metadata"
	daemon "octavius/internal/cli/daemon/metadata"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// NewCmd create a command for describing job
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	var jobName string
	describeCmd := &cobra.Command{
		Use:     "describe",
		Short:   "Describe the existing job",
		Long:    "This command helps to describe the job which is already created in server",
		Example: "octavius describe --job-name <job-name>",
		Args:    cobra.MaximumNArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			client := &client.GrpcClient{}
			res, err := octaviusDaemon.Describe(jobName, client)
			if err != nil {
				log.Error(err, "error in describing job")
				printer.Println(fmt.Sprintf("error in describing job, %v", status.Convert(err).Message()), color.FgRed)
				return
			}
			log.Info(fmt.Sprintf("describe command for %v executed with metadata response %v", jobName, res))
			printer.Println(fmt.Sprintf("Job name: %v", res.Name), color.FgGreen)
			printer.Println(fmt.Sprintf("Job Description: %v", res.Description), color.FgGreen)
			printer.Println("Usage with arguments : ", color.FgGreen)
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Argument", "Description", "Is Required"})
			jobArgs := res.EnvVars.Args
			for _, arg := range jobArgs {
				if arg.Required {
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiGreen.Sprintf("YES")})
				} else {
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiRed.Sprintf("NO")})
				}
			}
			t.Render()
		},
	}
	describeCmd.Flags().StringVarP(&jobName, "job-name", "", "", "It contains job name")
	err := describeCmd.MarkFlagRequired("job-name")
	if err != nil {
		log.Error(err, "error while setting the flag required")
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	return describeCmd
}
