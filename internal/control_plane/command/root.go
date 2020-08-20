package command

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"octavius/internal/control_plane/command/start"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

// Execute the root command and no error returned
func Execute(logger zerolog.Logger) {
	rootCmd.AddCommand(start.NewCmd(logger))
	if err := rootCmd.Execute(); err != nil {
		logger.Err(err).Msg("Root command error")
	} else {
		logger.Info().Msg("Root Command Executed")
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.json)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
