package start

import (

	"github.com/spf13/cobra"
	"octavius/internal/control_plane/server"
	"octavius/internal/logger"
)

var createCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server at AppPort`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Start Executed")
		err := server.Start()
		if err != nil {
			logger.Panic("Start command err: ", err)
		}
	},
}

func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

	logger.Setup()

}
