package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/id_generator"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	octerr "octavius/internal/pkg/errors"
	clientCPproto "octavius/internal/pkg/errors/protofiles/client_CP"
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
	logger.Info(fmt.Sprintf("%v Job Create Request Received - Posting Metadata to etcd", uid))
	name, err := s.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		logger.Error(err, "error in saving to etcd")
	}
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

	// TODO: relay stream logs from executor
	logString := &clientCPproto.Log{RequestId: uid, Log: "lorem ipsum logger logger logger dumb"}
	err = stream.Send(logString)
	logger.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uid))
	errMsg := octerr.New(2, err)
	if err != nil {
		return errMsg
	}
	return nil
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	//will be utilized after implementation
	//uid, err := id_generator.NextID()
	return nil, errors.New("not implemented yet")
}
