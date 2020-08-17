package create

import (
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"octavius/internal/config"
	"octavius/internal/logger"
	"octavius/pkg/protobuf"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create procs",
	Long:  `Create procs by giving name, author`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("create executed")
		logger.Info("in cli create")
		appPort := config.Config().AppPort
		conn, err := grpc.Dial("localhost:"+appPort, grpc.WithInsecure())
		if err != nil {
			logger.Fatal("", err)
		}
		client := protobuf.NewProcServiceClient(conn)
		crequest := &protobuf.RequestForCreateProc{
			Name:   "First Job",
			Author: "Author",
		}
		procID, err := client.CreateProc(context.Background(), crequest)
		if err != nil {
			logger.Error("",err)
		} else {
			logger.Info("Proc id message:" + procID.GetMessage())
			logger.Info("Proc id value:" + procID.GetValue())
		}
	},
}

func GetCmd() *cobra.Command {
	return createCmd
}

func init() {

}
