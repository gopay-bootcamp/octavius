package server

import (
	"octavius/internal/config"
	"octavius/internal/controlPlane/db/etcd"
	"octavius/internal/logger"

	// "crud-toy/internal/model"
	"octavius/pkg/protobuf"
	"net"

	"google.golang.org/grpc"
)

func Start() error {
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", appPort)
	server := grpc.NewServer()
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	
	procGrpcServer := NewProcServiceServer(exec)
	procProto.RegisterProcServiceServer(server, procGrpcServer)
	if err != nil {
		logger.Fatal("grpc server not started")
		return err
	}
	logger.Info("grpc server started on port 8000")
	server.Serve(listener)
	return nil
}
