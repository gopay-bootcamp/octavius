package command

import (
	"octavius/internal/cli/command/create"
	"os"
	"octavius/internal/logger"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using octavius' multi-processing capabilities`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("Execution error ", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// add your client commands here
	rootCmd.AddCommand(create.GetCmd())

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.json)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Error("Config Setup Issue : ", err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Info("Using Config file: " + viper.ConfigFileUsed())
	}
}
