package command

import (
	"fmt"
	"octavius/internal/cli/command/config"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

func Execute() {

	configCmd := config.NewCmd()
	rootCmd.AddCommand(configCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
