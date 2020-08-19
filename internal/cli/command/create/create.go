package create

import (
	"fmt"
	"github.com/spf13/cobra"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"octavius/internal/cli/fileUtil"
	"strings"
)

func NewCmd(octaviusDaemon daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:     "create",
		Short:   "Create new octavius job metadata",
		Long:    "This command helps create new jobmetadata to your CP host with proper metadata.json file",
		Example: fmt.Sprintf("octavius create PATH=<filepath>/metadata.json"),
		Args:    cobra.MinimumNArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			arg := strings.Split(args[0], "=")
			if len(arg) < 2 || arg[0] != "PATH" {
				fmt.Println("Incorrect command argument format, the correct format is: \n octavius create PATH=<filepath>/metadata.json")
				return
			}
			metadataFilePath := arg[1]

			fileUtil:=fileUtil.NewFileUtil()
			metadataFileIoReader,err:=fileUtil.GetIoReader(metadataFilePath)
			if err!=nil {
				fmt.Println(err)
				return
			}

			client := &client.GrpcClient{}
			res, err := octaviusDaemon.CreateMetadata(metadataFileIoReader, client)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(res)
		},
	}
}
