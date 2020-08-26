package command

import (
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/command/execution"
	"octavius/internal/cli/command/getstream"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

// Execute Executes the root command of Octavius Cli
func Execute(octaviusDaemon daemon.Client) {

	fileUtil := file.New()

	configCmd := config.NewCmd(fileUtil)
	rootCmd.AddCommand(configCmd)

	createCmd := create.NewCmd(octaviusDaemon, fileUtil)
	rootCmd.AddCommand(createCmd)

	getstreamCmd := getstream.NewCmd(octaviusDaemon)
	rootCmd.AddCommand(getstreamCmd)

	executeCmd := execution.NewCmd(octaviusDaemon)
	rootCmd.AddCommand(executeCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err, "input command execution: ")
	}
}
