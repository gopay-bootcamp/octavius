package start

import (
	"github.com/rs/zerolog"
	"octavius/internal/control_plane/server"

	"github.com/spf13/cobra"
)

//NewCmd returns start command
func NewCmd(logger zerolog.Logger) *cobra.Command{
	return &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  `Start server at AppPort`,
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Start(logger)
			if err != nil {
				logger.Err(err).Msg("Execution error in start server")
			}
		},
	}
}

func init() {
}
