package server

import (
	"context"
	"errors"
	"fmt"
	"octavius/internal/controller/server/execution"
	octerr "octavius/internal/pkg/errors"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	protobuf "octavius/internal/pkg/protofiles/client_CP"
)

type customCTXKey string

var uidKey customCTXKey = "uid"

type clientCPServicesServer struct {
	procExec execution.Execution
}

// NewProcServiceServer used to create a new execution context
func NewProcServiceServer(exec execution.Execution) protobuf.ClientCPServicesServer {
	return &clientCPServicesServer{
		procExec: exec,
	}
}

func (s *clientCPServicesServer) PostMetadata(ctx context.Context, request *protobuf.RequestToPostMetadata) (*protobuf.MetadataName, error) {
	uid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "error while assigning is to the request")
	}
	// should not use basic type string as key in context.WithValue
	ctx = context.WithValue(ctx, uidKey, uid)
	name, err := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	log.Error(err, fmt.Sprintf("%v Job Create Request Received - Posting Metadata to etcd", uid))
	return name, err
}

func (s *clientCPServicesServer) GetAllMetadata(ctx context.Context, request *protobuf.RequestToGetAllMetadata) (*protobuf.MetadataArray, error) {
	dataList, err := s.procExec.ReadAllMetadata(ctx)
	log.Error(err, "Getting Metadata")
	return dataList, err
}

func (s *clientCPServicesServer) GetStreamLogs(request *protobuf.RequestForStreamLog, stream protobuf.ClientCPServices_GetStreamLogsServer) error {

	uid, err := idgen.NextID()
	if err != nil {
		log.Error(err, "Error while assigning is to the request")
	}
	// RequestId is missing.
	// logString := &protobuf.Log{RequestId: uid, Log: "lorem ipsum logger logger logger dumb"}
	logString := &protobuf.Log{}
	err = stream.Send(logString)
	//handle the error first
	if err != nil {
		log.Error(err, fmt.Sprintf("%v GetStream Request Received - Sending stream to client", uid))
		errMsg := octerr.New(2, err)
		return errMsg
	}
	return nil
}

func (s *clientCPServicesServer) ExecuteJob(ctx context.Context, execute *protobuf.RequestForExecute) (*protobuf.Response, error) {
	//will be utilized after implementation
	//uid, err := idgen.NextID()
	return nil, errors.New("not implemented yet")
}
