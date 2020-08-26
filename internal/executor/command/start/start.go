package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Executor",
		Long:  `Start Executor for communicating with Control Plane`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Executor")
		},
	}
}

func init() {
}
