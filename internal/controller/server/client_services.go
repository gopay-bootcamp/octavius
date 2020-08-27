package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"

	octerr "octavius/internal/pkg/errors"
	clientCPproto "octavius/internal/pkg/protofiles/client_CP"
)

type customCTXKey string

var uidKey customCTXKey = "uid"

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
	uid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, uidKey, uid)
	log.Info(fmt.Sprintf("request ID: %v, PostMetadata request received", uid))
	name, err := s.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		log.Error(err, fmt.Sprintf("request ID: %v, error in saving to etcd", uid))
	}
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *clientCPproto.RequestToGetAllMetadata) (*clientCPproto.MetadataArray, error) {
	uid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
	}
	ctx = context.WithValue(ctx, uidKey, uid)
	log.Info(fmt.Sprintf("request ID: %v, GetAllMetadata request received", uid))
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	return dataList, err
}

func (s *clientCPServicesServer) GetStreamLogs(request *clientCPproto.RequestForStreamLog, stream clientCPproto.ClientCPServices_GetStreamLogsServer) error {
	uid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "Error while assigning is to the request")
	}

	// TODO: relay stream logs from executor
	logString := &clientCPproto.Log{RequestId: uid, Log: "lorem ipsum logger logger logger dumb"}
	err = stream.Send(logString)
	log.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uid))
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
