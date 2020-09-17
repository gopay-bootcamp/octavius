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
	return &cobra.Command{
		Use:     "describe",
		Short:   "Describe the existing job",
		Long:    "This command helps to describe the job which is already created in server",
		Example: fmt.Sprintf("octavius describe <job-name>"),
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			jobName := args[0]
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
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiGreen.Sprintf("YES")})
				} else {
					t.AppendRow([]interface{}{arg.Name, arg.Description, text.FgHiRed.Sprintf("NO")})
				}
			}
			t.Render()
		},
	}
}
