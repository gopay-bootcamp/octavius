package server

import (
	"fmt"
	"net"
	"octavius/internal/control_plane/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/logger"
	repository "octavius/internal/control_plane/server/repository/metadata"
	"os"
	"os/signal"
	"syscall"
	"time"

	"octavius/internal/control_plane/server/execution"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"

	"google.golang.org/grpc"
)

func startClientCPServer(listener net.Listener, clientServer clientCPproto.ClientCPServicesServer) {
	server := grpc.NewServer()
	clientCPproto.RegisterClientCPServicesServer(server, clientServer)
	logger.Info(fmt.Sprintln("Started client server at port: ", listener.Addr().String()))
	err := server.Serve(listener)
	if err != nil {
		logger.Fatal("Failed to start Client Server")
	}
}

func startExecutorCPServer(listener net.Listener, executorServer executorCPproto.ExecutorCPServicesServer) {
	server := grpc.NewServer()
	executorCPproto.RegisterExecutorCPServicesServer(server, executorServer)
	logger.Info(fmt.Sprintln("Started executor server at port: ", listener.Addr().String()))
	err := server.Serve(listener)
	if err != nil {
		logger.Fatal("Failed to start executor server")
	}
}

// Start the grpc server
func Start() error {

	dialTimeout := 2 * time.Second
	etcdHost := "localhost:" + config.Config().EtcdPort

	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return err
	}

	etcdClient := etcd.NewClient(dialTimeout, etcdHost)
	defer etcdClient.Close()
	metadataRepository := repository.NewMetadataRepository(etcdClient)
	exec := execution.NewExec(metadataRepository)
	clientCPGrpcServer := NewProcServiceServer(exec)
	executorCPGrpcServer := NewExecutorServiceServer(exec)

	go startClientCPServer(listener, clientCPGrpcServer)
	go startExecutorCPServer(listener, executorCPGrpcServer)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	return nil
}
