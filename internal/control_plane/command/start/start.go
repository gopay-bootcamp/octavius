package start

import (
	"octavius/internal/control_plane/server"
	"octavius/internal/logger"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server at AppPort`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start()
		if err != nil {
			logger.Error("error in starting server", err)
		}
	},
}

//GetCmd returns start command
func GetCmd() *cobra.Command {
	return createCmd
}

func init() {
	logger.Setup()
}
