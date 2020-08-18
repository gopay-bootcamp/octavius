package getstream

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getstreamCmd = &cobra.Command{
	Use:   "getstream",
	Short: "Get job log data",
	Long:  `Get job log by giving arguments`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("getstream executed")
	},
}

func GetCmd() *cobra.Command {
	return getstreamCmd
}

func init() {

}
