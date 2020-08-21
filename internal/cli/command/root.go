package command

import (
	"fmt"
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/command/execution"
	"octavius/internal/cli/command/getstream"
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

	getstreamCmd := getstream.NewCmd(octaviusDaemon, printer)
	rootCmd.AddCommand(getstreamCmd)

	executeCmd := execution.NewCmd(octaviusDaemon, printer)
	rootCmd.AddCommand(executeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
