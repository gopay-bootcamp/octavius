package server

import (
	"fmt"
	"net"
	"octavius/internal/control_plane/config"
	"octavius/internal/control_plane/db/etcd"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/metadata/repository"
	octerr "octavius/internal/pkg/errors"
	"time"

	"octavius/internal/control_plane/server/execution"
	protobuf "octavius/internal/pkg/protofiles/client_CP"

	"google.golang.org/grpc"
)

// Start the grpc server
func Start() error {

	dialTimeout := 2 * time.Second
	etcdHost := "localhost:" + config.Config().EtcdPort
	appPort := config.Config().AppPort
	listener, err := net.Listen("tcp", "localhost:"+appPort)
	if err != nil {
		return octerr.New(2, err)
	}
	server := grpc.NewServer()
	etcdClient := etcd.NewClient(dialTimeout, etcdHost)
	defer etcdClient.Close()
	metadataRepository := repository.NewMetadataRepository(etcdClient)
	exec := execution.NewExec(metadataRepository)
	clientCPGrpcServer := NewProcServiceServer(exec)
	protobuf.RegisterClientCPServicesServer(server, clientCPGrpcServer)
	logger.Info(fmt.Sprintf("grpc server started on port %v", appPort))
	err = server.Serve(listener)
	return err
}
