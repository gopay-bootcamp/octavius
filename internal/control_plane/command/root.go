package command

import (
	"github.com/rs/zerolog"
	"octavius/internal/control_plane/command/start"
	"octavius/internal/control_plane/logger"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Logger zerolog.Logger
var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

// Execute the root command and no error returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Logger.Err(err).Msg("Root command error")
	} else {
		Logger.Info().Msg("Root Command Executed")
	}
}

func init() {
	Logger = logger.Setup()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.json)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(start.GetCmd())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			Logger.Err(err).Msg("CP Init config issue ")
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Err(err).Msg("problem in reading config file.")
	} else{
		Logger.Debug().Msg("Data being read from config file.")
	}
}
