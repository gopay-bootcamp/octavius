package start

import (
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  `Start server at AppPort`,
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Start()
			if err != nil {
				logger.Error(err, "Execution error in start server")
			}
		},
	}
}

func init() {
}
