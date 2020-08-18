package command

import (
	"fmt"
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/daemon"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

func Execute(octaviusDaemon daemon.Client) {

	configCmd := config.NewCmd()
	rootCmd.AddCommand(configCmd)

	createCmd := create.NewCmd(octaviusDaemon)
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
