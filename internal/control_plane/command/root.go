package command

import (
	"github.com/spf13/cobra"
	"octavius/internal/control_plane/command/start"
	"octavius/internal/control_plane/logger"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

// Execute the root command and no error returned
func Execute() {
	rootCmd.AddCommand(start.NewCmd())
	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err, "Root  command execution")
	}
}
