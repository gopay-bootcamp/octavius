package server

import (
	"context"
	"octavius/internal/control_plane/server/execution"
	executorCPproto "octavius/internal/pkg/protofiles/executor_CP"
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
	return e.procExec.UpdateExecutorStatus(ctx, ping)
}

func (e *executorCPServicesServer) Register(ctx context.Context, request *executorCPproto.RegisterRequest) (*executorCPproto.RegisterResponse, error) {
	return e.procExec.RegisterExecutor(ctx, request)
}
