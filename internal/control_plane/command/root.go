package command

import (
	"octavius/internal/control_plane/command/start"
	"octavius/internal/logger"
	"os"

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("error in execute", err)
		os.Exit(1)
	}
}

func init() {
	logger.Setup()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.json)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add your client commands here
	rootCmd.AddCommand(start.GetCmd())
}

func initConfig() {
	// Use config file from the flag.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Error("CP Init config issue : ", err)
			os.Exit(1)
		}
		logger.Info("home: " + home)
		viper.AddConfigPath(home)
		logger.Info("Config file utilising from " + home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Info("Using config file:" + viper.ConfigFileUsed())
	}
}
