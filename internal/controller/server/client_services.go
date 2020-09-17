package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/util"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	clientCPproto "octavius/internal/pkg/protofiles/client_cp"
)

type clientCPServicesServer struct {
	procExec execution.Execution
	idgen    idgen.RandomIdGenerator
}

type metadataServicesServer struct {
	procExec execution.Execution
	idgen    idgen.RandomIdGenerator
}

type jobServicesServer struct {
	procExec execution.Execution
	idgen    idgen.RandomIdGenerator
}

func NewMetadataServicesServer(exec execution.Execution, idgen idgen.RandomIdGenerator) clientCPproto.MetadataServicesServer {
	return &metadataServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func NewJobServicesServer(exec execution.Execution, idgen idgen.RandomIdGenerator) clientCPproto.JobServicesServer {
	return &jobServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func (m *metadataServicesServer) PostMetadata(ctx context.Context, request *clientCPproto.RequestToPostMetadata) (*clientCPproto.MetadataName, error) {
	uuid, err := m.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, PostMetadata request received with metadata %v", uuid, request.Metadata))

	name, err := m.procExec.SaveMetadata(ctx, request.Metadata)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in saving to etcd", uuid))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return name, nil
}

func (m *metadataServicesServer) GetAllMetadata(ctx context.Context, request *clientCPproto.RequestToGetAllMetadata) (*clientCPproto.MetadataArray, error) {
	uuid, err := m.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, GetAllMetadata request received", uuid))

	dataList, err := m.procExec.ReadAllMetadata(ctx)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in getting all metadata from etcd", uuid))
	}
	return dataList, status.Error(codes.Internal, err.Error())
}

func (j *jobServicesServer) GetLogs(ctx context.Context, request *clientCPproto.RequestForLogs) (*clientCPproto.Log, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	uuid, err := j.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info(fmt.Sprintf("request id: %v, getlogs request received", uuid))
	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	jobLogs, err := j.procExec.GetJobLogs(ctx, request.JobName)
	if err != nil {
		log.Error(fmt.Errorf("request id: %v, error in fetching logs, error details: %v", uuid, err), "")
		return nil, status.Error(codes.Internal, err.Error())
	}
	logString := &clientCPproto.Log{Log: jobLogs}
	return logString, nil
}

// ExecuteJob will call ExecuteJob function of execution and get jobId
func (j *jobServicesServer) ExecuteJob(ctx context.Context, executionData *clientCPproto.RequestForExecute) (*clientCPproto.Response, error) {
	uuid, err := j.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, ExecuteJob request received with executionData %+v", uuid, executionData))

	jobID, err := j.procExec.ExecuteJob(ctx, executionData)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in job execution", uuid))
		return &clientCPproto.Response{Status: "failure"}, err
	}

	jobIDString := strconv.FormatUint(jobID, 10)
	return &clientCPproto.Response{Status: jobIDString}, err
}

// GetJobList will call GetJobList function of execution and return list of available jobs
func (j *jobServicesServer) GetJobList(ctx context.Context, request *clientCPproto.RequestForGetJobList) (*clientCPproto.JobList, error) {
	uuid, err := j.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, GetJobList request received with clientInfo %+v", uuid, request))
	return j.procExec.GetJobList(ctx)
}

func (j *jobServicesServer) DescribeJob(ctx context.Context, descriptionData *clientCPproto.RequestForDescribe) (*clientCPproto.Metadata, error) {
	uuid, err := j.idgen.Generate()
	if err != nil {
		log.Error(err, "error while assigning id to the request")
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request ID: %v, DescribeJob request received with name %+v", uuid, descriptionData))
	metadata, err := j.procExec.GetMetadata(ctx, descriptionData)
	if err != nil {
		log.Error(err, "error in fetching metadata of job")
		return metadata, status.Errorf(codes.NotFound, err.Error())
	}
	return metadata, nil
}
