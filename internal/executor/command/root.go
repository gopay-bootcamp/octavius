package command

import (
	"octavius/internal/executor/command/start"
	"octavius/internal/executor/daemon"
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
func Execute(executorDaemon daemon.Client) {

	startCmd := start.NewCmd(executorDaemon)
	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		logger.Error(err, "root command execution")
	}
}
