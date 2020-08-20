package command

import (
	"fmt"
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

// Execute Executes the root command of Octavius Cli
func Execute(octaviusDaemon daemon.Client, fileUtil fileUtil.FileUtil, printer printer.Printer) {

	configCmd := config.NewCmd(fileUtil, printer)
	rootCmd.AddCommand(configCmd)

	createCmd := create.NewCmd(octaviusDaemon, fileUtil, printer)
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
