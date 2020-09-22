package create

import (
	"fmt"
	client "octavius/internal/cli/client/metadata"
	daemon "octavius/internal/cli/daemon/metadata"
	"octavius/internal/pkg/file"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/printer"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// NewCmd Returns an instance of Create command for registering Job Metadata in Octavius
func NewCmd(octaviusDaemon daemon.Client, fileUtil file.File) *cobra.Command {
	var metadataFilePath string

	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create new octavius job metadata",
		Long:    "This command helps create new job metadata to your CP host with proper metadata.json file",
		Example: fmt.Sprintf("octavius create --job-path <filepath>/metadata.json"),
		Args:    cobra.MaximumNArgs(0),

		Run: func(cmd *cobra.Command, args []string) {

			printer.Println("Creating a new job.", color.FgBlack)
			metadataFileIoReader, err := fileUtil.GetIoReader(metadataFilePath)
			if err != nil {
				log.Error(err, "error in reading file")
				printer.Println(fmt.Sprintf("error in reading file %v", err.Error()), color.FgRed)
				return
			}

			client := &client.GrpcClient{}
			res, err := octaviusDaemon.Post(metadataFileIoReader, client)
			if err != nil {
				log.Error(err, "error in creating metadata")
				printer.Println(fmt.Sprintf("error in creating metadata, %v", status.Convert(err).Message()), color.FgRed)
				return
			}

			log.Info(fmt.Sprintf("%s job created", res.Name))
			printer.Println(fmt.Sprintf("%s job created", res.Name), color.FgGreen)
		},
	}
	createCmd.Flags().StringVarP(&metadataFilePath, "job-path", "", "", "path to metadata.json(required)")
	err := createCmd.MarkFlagRequired("job-path")
	if err != nil {
		log.Error(err, "error while setting the flag required")
		printer.Println("error while setting the flag required", color.FgRed)
		return nil
	}
	return createCmd
}
