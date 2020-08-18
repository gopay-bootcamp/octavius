package execution

import (
	"fmt"
	"octavius/internal/cli/daemon"
	"github.com/spf13/cobra"
)

var executionCmd = &cobra.Command{
	Use:   "execution",
	Short: "Execution of Job",
	Long:  `Execution of Job by giving arguments`,
	Run: func(cmd *cobra.Command, args []string, octaviusDaemon daemon.Client) {
		if len(args) != 1 {
			fmt.Println("Incorrect command argument format, the correct format is: \n octavius getstream <job-name>")
			return
		}
		jobName := args[0]
		err := octaviusDaemon.GetStreamLog(jobName)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func GetCmd() *cobra.Command {
	return executionCmd
}

func init() {

}
