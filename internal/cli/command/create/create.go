package create

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"octavius/internal/cli/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewCmd Returns an instance of Create command for registering Job Metadata in Octavius
func NewCmd(octaviusDaemon daemon.Client, fileUtil fileUtil.FileUtil, printer printer.Printer) *cobra.Command {
	var metadataFilePath string

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create new octavius job metadata",
		Long:    "This command helps create new jobmetadata to your CP host with proper metadata.json file",
		Example: fmt.Sprintf("octavius create --job-path <filepath>/metadata.json"),

		Run: func(cmd *cobra.Command, args []string) {
			metadataFileIoReader, err := fileUtil.GetIoReader(metadataFilePath)
			if err != nil {
				printer.Println(fmt.Sprintln(err))
				return
			}

			client := &client.GrpcClient{}
			res, err := octaviusDaemon.CreateMetadata(metadataFileIoReader, client)
			if err != nil {
				printer.Println(fmt.Sprintln(err), color.FgRed)
				return
			}
			printer.Println(fmt.Sprintf("%s job created", res.Name), color.FgGreen)
		},
	}
	createCmd.Flags().StringVarP(&metadataFilePath, "job-path", "", "", "path to metadata.json(required)")
	createCmd.MarkFlagRequired("job-path")

	return createCmd
}
