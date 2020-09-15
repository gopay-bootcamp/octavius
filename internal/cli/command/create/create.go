package create

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"
)

// NewCmd Returns an instance of Create command for registering Job Metadata in Octavius
func NewCmd(octaviusDaemon daemon.Client, fileUtil file.File) *cobra.Command {
	var metadataFilePath string

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create new octavius job metadata",
		Long:    "This command helps create new job metadata to your CP host with proper metadata.json file",
		Example: fmt.Sprintf("octavius create --job-path <filepath>/metadata.json"),

		Run: func(cmd *cobra.Command, args []string) {

			printer.Println("Creating a new job.", color.FgBlack)
			metadataFileIoReader, err := fileUtil.GetIoReader(metadataFilePath)
			if err != nil {
				log.Error(err, "error in reading file")
				printer.Println("error in reading file", color.FgRed)
				return
			}

			client := &client.GrpcClient{}
			res, err := octaviusDaemon.CreateMetadata(metadataFileIoReader, client)
			if err != nil {
				log.Error(err, "error in creating metadata")
				printer.Println("error in creating metadata", color.FgRed)
				return
			}

			log.Info(fmt.Sprintf("%s job created", res.Name))
			printer.Println(fmt.Sprintf("%s job created", res.Name), color.FgGreen)
		},
	}
	createCmd.Flags().StringVarP(&metadataFilePath, "job-path", "", "", "path to metadata.json(required)")
	_ = createCmd.MarkFlagRequired("job-path")

	return createCmd
}
