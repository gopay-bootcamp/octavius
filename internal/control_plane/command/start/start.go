package start

import (
	"fmt"
	"log"
	"octavius/internal/control_plane/server"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server at AppPort`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start executed")
		err := server.Start()
		if err == nil {
			log.Panic(err)
		}
	},
}

func GetCmd() *cobra.Command {
	return createCmd
}

func init() {
	fmt.Println("in start init")

}
