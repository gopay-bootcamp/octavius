package start

import (
	"github.com/rs/zerolog"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server"

	"github.com/spf13/cobra"
)

var Logger zerolog.Logger
var createCmd = &cobra.Command{
	Use:   "start",
	Short: "Start server",
	Long:  `Start server at AppPort`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start()
		if err != nil {
			Logger.Err(err).Msg("Execution error in start server")
		}
	},
}

//GetCmd returns start command
func GetCmd() *cobra.Command {
	return createCmd
}

func init() {
	Logger = logger.Setup()
}
