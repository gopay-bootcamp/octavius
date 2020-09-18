package command

import (
	"octavius/internal/controller/command/start"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "controller",
	Short: "Controller of octavius",
	Long:  `Controller of octavius takes request from cli`,
}

// Execute executes the root command of octavius control plane
func Execute() error {
	rootCmd.AddCommand(start.NewCmd())
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
