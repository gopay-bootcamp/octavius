package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/control_plane/id_generator"
	"octavius/internal/control_plane/logger"
	"octavius/internal/control_plane/server/execution"
	"octavius/internal/control_plane/util"

	octerr "octavius/internal/pkg/errors"
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
	uuid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	logger.Info(fmt.Sprintf("request ID: %v, PostMetadata request received", uuid))
	name, err := s.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		logger.Error(err, fmt.Sprintf("request ID: %v, error in saving to etcd", uuid))
	}
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *clientCPproto.RequestToGetAllMetadata) (*clientCPproto.MetadataArray, error) {
	uuid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	logger.Info(fmt.Sprintf("request ID: %v, GetAllMetadata request received", uuid))
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	return dataList, err
}

func (s *clientCPServicesServer) GetStreamLogs(request *clientCPproto.RequestForStreamLog, stream clientCPproto.ClientCPServices_GetStreamLogsServer) error {
	uuid, err := id_generator.NextID()
	if err != nil {
		logger.Error(err, "Error while assigning is to the request")
	}

	// TODO: relay stream logs from executor
	logString := &clientCPproto.Log{RequestId: uuid, Log: "lorem ipsum logger logger logger dumb"}
	err = stream.Send(logString)
	logger.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uuid))
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
