package create

import (
	"fmt"
	"octavius/internal/cli/daemon"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"octavius/internal/config"
	"octavius/internal/logger"
	"octavius/pkg/protobuf"
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
			err := octaviusDaemon.CreateMetadata(metadataFile)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
}
