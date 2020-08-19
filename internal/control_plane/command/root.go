package command

import (
	"octavius/internal/control_plane/command/start"
	"octavius/internal/logger"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "control_plane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

// Execute the root command and no error returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("error in root command execution", err)
	}
}

func init() {
	logger.Setup()
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
			logger.Error("CP Init config issue : ", err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("problem in reading config file", err)
	}
}
