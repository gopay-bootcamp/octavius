package server

import (
	"fmt"
	"net"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution"
	executorRepo "octavius/internal/controller/server/repository/executor"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/pkg/db/etcd"
	octerr "octavius/internal/pkg/errors"
	"octavius/internal/pkg/log"
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

	// TODO: organise code properly as in Ikhsan's PR
	// TODO: change localhost with config's etcd host
	etcdClient, err := etcd.NewClient(dialTimeout, etcdHost)

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
	log.Info(fmt.Sprintln("Started server at port: ", listener.Addr().String()))
	err = server.Serve(listener)
	if err != nil {
		return octerr.New(2, err)
	}
	return nil
}
