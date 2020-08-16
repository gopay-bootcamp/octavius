package command

import (
	"fmt"
	"octavius/internal/controlPlane/command/start"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ControlPlane",
	Short: "Control plane of octavius",
	Long:  `Control plane of octavius takes request from cli`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("in controlPlane root init")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.json)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add your client commands here
	rootCmd.AddCommand(start.GetCmd())
}

func initConfig() {
	fmt.Println("in controlPlane root initconfig")
	// Use config file from the flag.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("home: ", home)
		viper.AddConfigPath(home)
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
