package server

import (
	"fmt"
	"net"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution/health"
	"octavius/internal/controller/server/execution/job"
	"octavius/internal/controller/server/execution/metadata"
	"octavius/internal/controller/server/execution/registration"
	executorRepo "octavius/internal/controller/server/repository/executor"
	jobRepo "octavius/internal/controller/server/repository/job"
	metadataRepo "octavius/internal/controller/server/repository/metadata"
	"octavius/internal/controller/server/scheduler"
	"octavius/internal/pkg/db/etcd"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start() error {
	controllerConfig, err := config.Loader()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	dialTimeout := controllerConfig.EtcdDialTimeout
	etcdHost := controllerConfig.EtcdHost + ":" + controllerConfig.EtcdPort
	appPort := controllerConfig.AppPort

	etcdClient, err := etcd.NewClient(dialTimeout, etcdHost)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	defer etcdClient.Close()

	metadataRepository := metadataRepo.NewMetadataRepository(etcdClient)
	executorRepository := executorRepo.NewExecutorRepository(etcdClient)
	jobRepository := jobRepo.NewJobRepository(etcdClient)

	randomIdGenerator := idgen.NewRandomIdGenerator()

	scheduler := scheduler.NewScheduler(randomIdGenerator, jobRepository)

	registrationExec := registration.NewRegistrationExec(executorRepository)
	registrationServicesGrpcServer := NewRegistrationServiceServer(registrationExec, randomIdGenerator)

	jobExec := job.NewJobExec(jobRepository, randomIdGenerator, scheduler)
	jobServicesGrpcServer := NewJobServiceServer(jobExec, randomIdGenerator)

	metadataExec := metadata.NewMetadataExec(metadataRepository, scheduler)
	metadataServicesGrpcServer := NewMetadataServiceServer(metadataExec, randomIdGenerator)

	healthExec := health.NewHealthExec(executorRepository)
	healthServicesGrpcServer := NewHealthServiceServer(healthExec, randomIdGenerator)

	server := grpc.NewServer()
	protofiles.RegisterHealthServicesServer(server, healthServicesGrpcServer)
	protofiles.RegisterJobServiceServer(server, jobServicesGrpcServer)
	protofiles.RegisterMetadataServicesServer(server, metadataServicesGrpcServer)
	protofiles.RegisterRegistrationServiceServer(server, registrationServicesGrpcServer)
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintln("Started controller at port: ", listener.Addr().String()))
	err = server.Serve(listener)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
