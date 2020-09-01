package server

import (
	"context"
	"fmt"
	"io"
	"octavius/internal/controller/server/execution"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	executorCPproto "octavius/internal/pkg/protofiles/executor_cp"
	"octavius/internal/pkg/util"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type executorCPServicesServer struct {
	procExec execution.Execution
}

// NewExecutorServiceServer used to create a new execution context
func NewExecutorServiceServer(exec execution.Execution) executorCPproto.ExecutorCPServicesServer {
	return &executorCPServicesServer{
		procExec: exec,
	}
}

func (e *executorCPServicesServer) HealthCheck(ctx context.Context, ping *executorCPproto.Ping) (*executorCPproto.HealthResponse, error) {
	res, err := e.procExec.UpdateExecutorStatus(ctx, ping)
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error in running health check", ping.ID))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	uuid, err := idgen.NextID()
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning id to the request", request.ID))
		return nil, status.Error(codes.Internal, err.Error())
	}

	ctx = context.WithValue(ctx, util.ContextKeyUUID, uuid)
	log.Info(fmt.Sprintf("request id: %v, recieve register request from executor with id %s", uuid, request.ID))

	res, err := e.procExec.RegisterExecutor(ctx, request)
	if err != nil {
		log.Error(err, fmt.Sprintf("request id: %v, error in registering executor with id %s", uuid, request.ID))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) GetJob(ctx context.Context, start *executorCPproto.Start) (*executorCPproto.Job, error) {
	res, err := e.procExec.GetJob(ctx, start)
	//GetJob searches for jobs under executor namespace first and returns from it
	//if there is none, it then picks jobs from the jobs/pending namespace
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error while assigning job to executor", start.Id))
		return nil, err
	}
	return res, err
}

func (e *executorCPServicesServer) StreamLog(stream executorCPproto.ExecutorCPServices_StreamLogServer) error {
	var logCount int32
	logs := []*executorCPproto.JobLog{}
	startTime := time.Now()

	for {
		log, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&executorCPproto.LogSummary{
				Recieved:    true,
				LogCount:    logCount,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		logCount++
		logs = append(logs, log)
	}
}
