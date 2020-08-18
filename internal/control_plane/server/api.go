package server

import (
	"fmt"
	"octavius/internal/control_plane/server/metadata/repository"
	"octavius/internal/logger"

	"google.golang.org/grpc"
	"net"
	"octavius/internal/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/server/execution"
	"octavius/pkg/protobuf"
)

func Start() error {
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	server := grpc.NewServer()
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()

	metadataRepository := repository.NewMetadataRepository(etcdClient)
	exec := execution.NewExec(metadataRepository)

	procGrpcServer := NewProcServiceServer(exec)
	protobuf.RegisterOctaviusServicesServer(server, procGrpcServer)
	if err != nil {
		logger.Fatal("grpc server not started", err)
	}
	logger.Info(fmt.Sprintf("grpc server started on port %v", appPort))
	server.Serve(listener)
	return nil
}
