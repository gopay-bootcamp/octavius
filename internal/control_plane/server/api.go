package server

import (
	"log"
	"octavius/internal/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/server/execution"
	"octavius/pkg/protobuf"
	"net"
	"google.golang.org/grpc"
)

func Start() error {
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	server := grpc.NewServer()
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	exec := execution.NewExec(etcdClient)

	procGrpcServer := NewProcServiceServer(exec)
	protobuf.RegisterProcServiceServer(server, procGrpcServer)
	if err != nil {
		log.Fatal("grpc server not started")
		return err
	}
	log.Printf("grpc server started on port %v", appPort)
	server.Serve(listener)
	return nil
}
