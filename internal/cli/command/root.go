//Package command provides execution for all the cli commands
package command

import (
	"octavius/internal/cli/command/config"
	"octavius/internal/cli/command/create"
	"octavius/internal/cli/command/describe"
	"octavius/internal/cli/command/execution"
	"octavius/internal/cli/command/getlogs"
	"octavius/internal/cli/command/list"
	jobDaemon "octavius/internal/cli/daemon/job"
	metadataDaemon "octavius/internal/cli/daemon/metadata"
	"octavius/internal/pkg/file"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "OCTAVIUS",
	Short: "Welcome to the octavius cli",
	Long:  `Easily automate your work using ocatvius' multi-processing capabilities`,
}

// Execute function executes the root command of Octavius Cli
func Execute(jobDaemon jobDaemon.Client, metadataDaemon metadataDaemon.Client) error {

	fileUtil := file.New()

	configCmd := config.NewCmd(fileUtil)
	if configCmd != nil {
		rootCmd.AddCommand(configCmd)
	}

	createCmd := create.NewCmd(metadataDaemon, fileUtil)
	if createCmd != nil {
		rootCmd.AddCommand(createCmd)
	}

	getLogsCmd := getlogs.NewCmd(jobDaemon)
	if getLogsCmd != nil {
		rootCmd.AddCommand(getLogsCmd)
	}

	executeCmd := execution.NewCmd(jobDaemon)
	if executeCmd != nil {
		rootCmd.AddCommand(executeCmd)
	}

	listCmd := list.NewCmd(metadataDaemon)
	rootCmd.AddCommand(listCmd)

	describeCmd := describe.NewCmd(metadataDaemon)
	if describeCmd != nil {
		rootCmd.AddCommand(describeCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
