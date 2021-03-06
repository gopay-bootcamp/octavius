// Package command consists of all commands in controller
package command

import (
	"octavius/internal/controller/command/start"

	"github.com/spf13/cobra"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
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
