package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create procs",
	Long:  `Create procs by giving name, author`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create executed")
	},
}

func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

}
