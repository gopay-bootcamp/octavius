package create

import (
	"context"
	"fmt"
	"log"
	"octavius/internal/config"
	"octavius/pkg/protobuf"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create procs",
	Long:  `Create procs by giving name, author`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create executed")
		fmt.Println("in cli create")
		appPort := config.Config().AppPort
		conn, err := grpc.Dial("localhost:"+appPort, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		client := protobuf.NewProcServiceClient(conn)
		crequest := &protobuf.RequestForCreateProc{
			Name:   "First Job",
			Author: "Author",
		}
		procID, err := client.CreateProc(context.Background(), crequest)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Proc id message:", procID.GetMessage())
			fmt.Println("Proc id value:", procID.GetValue())
		}
	},
}

func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

}
