package server

import (
	"fmt"
	"net"
	"octavius/internal/control_plane/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	executorRepo "octavius/internal/control_plane/server/repository/executor"
	metadataRepo "octavius/internal/control_plane/server/repository/metadata"
	octerr "octavius/internal/pkg/errors"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
	"time"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start() error {
	dialTimeout := 2 * time.Second
	etcdHost := "localhost:" + config.Config().EtcdPort
	appPort := config.Config().AppPort

	etcdClient := etcd.NewClient(dialTimeout, etcdHost)
	defer etcdClient.Close()
	metadataRepository := metadataRepo.NewMetadataRepository(etcdClient)
	executorRepository := executorRepo.NewExecutorRepository(etcdClient)
	exec := execution.NewExec(metadataRepository, executorRepository)
	clientCPGrpcServer := NewProcServiceServer(exec)
	executorCPGrpcServer := NewExecutorServiceServer(exec)

	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return octerr.New(2, err)
	}

	server := grpc.NewServer()
	clientCPproto.RegisterClientCPServicesServer(server, clientCPGrpcServer)
	executorCPproto.RegisterExecutorCPServicesServer(server, executorCPGrpcServer)
	logger.Info(fmt.Sprintln("Started server at port: ", listener.Addr().String()))
	err = server.Serve(listener)
	if err != nil {
		return octerr.New(2, err)
	}
	return nil
}
