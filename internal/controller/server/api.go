package server

import (
	"fmt"
	"net"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution"
	executorRepo "octavius/internal/controller/server/repository/executor"
	repository "octavius/internal/controller/server/repository/job"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start() error {
	dialTimeout := config.Config().EtcdDialTimeout
	etcdHost := config.Config().EtcdHost + ":" + config.Config().EtcdPort
	appPort := config.Config().AppPort

	etcdClient, err := etcd.NewClient(dialTimeout, etcdHost)
	defer etcdClient.Close()

	metadataRepository := metadataRepo.NewMetadataRepository(etcdClient)

	executorRepository := executorRepo.NewExecutorRepository(etcdClient)

	randomIdGenerator := idgen.NewRandomIdGenerator()
	jobExecutionRepository := repository.NewJobExecutionRepository(etcdClient)

	exec := execution.NewExec(metadataRepository, executorRepository,jobExecutionRepository,randomIdGenerator,scheduler.NewScheduler(etcdClient,randomIdGenerator))
	clientCPGrpcServer := NewProcServiceServer(exec,randomIdGenerator)
	executorCPGrpcServer := NewExecutorServiceServer(exec,randomIdGenerator)

	server := grpc.NewServer()
	clientCPproto.RegisterClientCPServicesServer(server, clientCPGrpcServer)
	executorCPproto.RegisterExecutorCPServicesServer(server, executorCPGrpcServer)

	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintln("Started server at port: ", listener.Addr().String()))
	err = server.Serve(listener)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
