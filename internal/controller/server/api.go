package server

import (
	"fmt"
	"net"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution"
	"octavius/internal/controller/server/metadata/repository"
	"octavius/internal/pkg/db/etcd"
	octerr "octavius/internal/pkg/errors"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles/client_CP"
	"time"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start() error {

	dialTimeout := 2 * time.Second
	// TODO: change localhost with config's etcd host
	etcdHost := "localhost:" + config.Config().EtcdPort

	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return octerr.New(2, err)
	}

	etcdClient, err := etcd.NewClient(dialTimeout, etcdHost)
	if err != nil {
		return octerr.New(2, err)
	}
	defer etcdClient.Close()

	metadataRepository := repository.NewMetadataRepository(etcdClient)

	exec := execution.NewExec(metadataRepository)
	clientCPGrpcServer := NewProcServiceServer(exec)

	server := grpc.NewServer()
	protobuf.RegisterClientCPServicesServer(server, clientCPGrpcServer)
	log.Info(fmt.Sprintf("grpc server started on port %v", appPort))

	err = server.Serve(listener)
	return err
}
