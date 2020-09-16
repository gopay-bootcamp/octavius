package describe

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
)

// NewCmd create a command for describing job
func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	var jobName string
	describeCmd := &cobra.Command{
		Use:     "describe",
		Short:   "Describe the existing job",
		Long:    "This command helps to describe the job which is already created in server",
		Example: fmt.Sprintf("octavius describe --job-name <job-name>"),
		Args:    cobra.MaximumNArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			client := &client.GrpcClient{}
			res, err := octaviusDaemon.DescribeJob(jobName, client)
			if err != nil {
				log.Error(err, "error in describing job")
				printer.Println("error in describing job", color.FgRed)
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
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiRed.Sprintf("YES")})
				} else {
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiGreen.Sprintf("NO")})
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
