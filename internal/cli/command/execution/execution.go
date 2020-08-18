package execution

import (
	"fmt"

	"github.com/spf13/cobra"
)

var executionCmd = &cobra.Command{
	Use:   "execution",
	Short: "Execution of Job",
	Long:  `Execution of Job by giving arguments`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("create executed")
	},
}

func GetCmd() *cobra.Command {
	return executionCmd
}

func init() {

}
