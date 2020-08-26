package server

import (
	"fmt"
	"net"
	"octavius/internal/control_plane/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/logger"
<<<<<<< HEAD
	executorRepo "octavius/internal/control_plane/server/repository/executor"
	metadataRepo "octavius/internal/control_plane/server/repository/metadata"
	"os"
	"os/signal"
	"syscall"
=======
	"octavius/internal/control_plane/server/metadata/repository"
	octerr "octavius/internal/pkg/errors"
>>>>>>> 441336cc28324122804d568039465967c2c8be26
	"time"

	"octavius/internal/control_plane/server/execution"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"

	"google.golang.org/grpc"
)

func startClientCPServer(listener net.Listener, clientServer clientCPproto.ClientCPServicesServer, errReturn chan error) {
	server := grpc.NewServer()
	clientCPproto.RegisterClientCPServicesServer(server, clientServer)
	logger.Info(fmt.Sprintln("Started client server at port: ", listener.Addr().String()))
	err := server.Serve(listener)
	errReturn <- err
}

func startExecutorCPServer(listener net.Listener, executorServer executorCPproto.ExecutorCPServicesServer, errReturn chan error) {
	server := grpc.NewServer()
	executorCPproto.RegisterExecutorCPServicesServer(server, executorServer)
	logger.Info(fmt.Sprintln("Started executor server at port: ", listener.Addr().String()))
	err := server.Serve(listener)
	errReturn <- err
}

// Start the grpc server
func Start() error {

	dialTimeout := 2 * time.Second
	etcdHost := "localhost:" + config.Config().EtcdPort
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return octerr.New(2, err)
	}
	etcdClient := etcd.NewClient(dialTimeout, etcdHost)
	defer etcdClient.Close()
	metadataRepository := metadataRepo.NewMetadataRepository(etcdClient)
	executorRepository := executorRepo.NewExecutorRepository(etcdClient)
	exec := execution.NewExec(metadataRepository, executorRepository)
	clientCPGrpcServer := NewProcServiceServer(exec)
	executorCPGrpcServer := NewExecutorServiceServer(exec)

	errReturn := make(chan error)
	go startClientCPServer(listener, clientCPGrpcServer, errReturn)
	go startExecutorCPServer(listener, executorCPGrpcServer, errReturn)
	if err := <-errReturn; err != nil {
		logger.Fatal(fmt.Sprintf("error in starting server with value %v", err))
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	return nil
}
