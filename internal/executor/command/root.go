package command

import (
	"octavius/internal/executor/logger"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "octavius_executor",
	Short: "kubernetes executor of octavius",
	Long:  `kubernetes executor of octavius takes request from cli`,
}

// Execute executes the root command of octavius control plane
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err, "Root  command execution")
	}
}
