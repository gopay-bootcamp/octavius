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

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("error in execute", err)
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
			logger.Fatal("CP Init config issue : ", err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("problem in reading config file", err)
	}
}
