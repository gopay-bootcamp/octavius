package server

import (
	"fmt"
	"github.com/rs/zerolog"
	"octavius/internal/control_plane/server/metadata/repository"

	"net"
	"octavius/internal/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/server/execution"
	"octavius/pkg/protobuf"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start(logger zerolog.Logger) error {
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	server := grpc.NewServer()
	etcdClient := etcd.NewClient()
	defer etcdClient.Close(logger)

	metadataRepository := repository.NewMetadataRepository(etcdClient)
	exec := execution.NewExec(metadataRepository)

	procGrpcServer := NewProcServiceServer(exec, logger)
	protobuf.RegisterOctaviusServicesServer(server, procGrpcServer)
	if err != nil {
		return err
	}
	logger.Info().Msg(fmt.Sprintf("grpc server started on port %v", appPort))
	server.Serve(listener)
	return nil
}
