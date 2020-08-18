package create

import (
	"fmt"
	"octavius/internal/cli/client"
	"octavius/internal/cli/daemon"
	"os"
	"strings"

	"github.com/spf13/cobra"
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
			metadataFile := arg[1]
			metadataFileHandler, err := os.Open(metadataFile)
			if err != nil {
				fmt.Println("Error opening the file given")
				return
			}
			defer metadataFileHandler.Close()

			client := &client.GrpcClient{}
			res, err := octaviusDaemon.CreateMetadata(metadataFileHandler, client)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(res)
		},
	}
}
