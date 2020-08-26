package command

import (
	"octavius/internal/control_plane/command/start"
	"octavius/internal/control_plane/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

// Execute executes the root command of octavius control plane
func Execute() {
	rootCmd.AddCommand(start.NewCmd())
	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err, "Root  command execution")
	}
}
