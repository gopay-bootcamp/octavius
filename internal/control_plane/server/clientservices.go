package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/id_generator"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
)

type clientCPServicesServer struct {
	procExec execution.Execution
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution) clientCPproto.ClientCPServicesServer {
	return &clientCPServicesServer{
		procExec: exec,
	}
}

func (s *clientCPServicesServer) PostMetadata(ctx context.Context, request *clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error) {
	uid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "Error while assigning is to the request")
	}
	ctx = context.WithValue(ctx, "uid", uid)
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	logger.Error(err, fmt.Sprintf("%v Job Create Request Received - Posting Metadata to etcd", uid))
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *clientCPproto.RequestToGetAllMetadata) (*clientCPproto.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	logger.Error(err, "Getting Metadata")
	return dataList, err
}



func (s *clientCPServicesServer) GetStreamLogs(request *clientCPproto.RequestForStreamLog, stream clientCPproto.ClientCPServices_GetStreamLogsServer) error {
	uid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "Error while assigning is to the request")
	}
	logString := &clientCPproto.Log{RequestId: uid, Log: "lorem ipsum logger logger logger dumb"}
	err = stream.Send(logString)
	logger.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uid))
	return err
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	//will be utilized after implementation
	//uid, err := id_generator.NextID()
	return nil, errors.New("not implemented yet")
}
