package command

import (
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/command/describe"
	"octavius/internal/cli/command/execution"
	"octavius/internal/cli/command/getlogs"
	"octavius/internal/cli/command/list"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/file"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

// Execute Executes the root command of Octavius Cli
func Execute(octaviusDaemon daemon.Client) error {

	fileUtil := file.New()

	configCmd := config.NewCmd(fileUtil)
	if configCmd != nil {
		rootCmd.AddCommand(configCmd)
	}

	createCmd := create.NewCmd(octaviusDaemon, fileUtil)
	if createCmd != nil {
		rootCmd.AddCommand(createCmd)
	}

	getLogsCmd := getlogs.NewCmd(octaviusDaemon)
	if getLogsCmd != nil {
		rootCmd.AddCommand(getLogsCmd)
	}

	executeCmd := execution.NewCmd(octaviusDaemon)
	if executeCmd != nil {
		rootCmd.AddCommand(executeCmd)
	}

	listCmd := list.NewCmd(octaviusDaemon)
	rootCmd.AddCommand(listCmd)

	describeCmd := describe.NewCmd(octaviusDaemon)
	if describeCmd != nil {
		rootCmd.AddCommand(describeCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
