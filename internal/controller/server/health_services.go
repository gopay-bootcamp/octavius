package server

import (
	"context"
	"fmt"
	"octavius/internal/controller/config"
	"octavius/internal/controller/server/execution/health"
	"octavius/internal/pkg/idgen"
	"octavius/internal/pkg/log"
	"octavius/internal/pkg/protofiles"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type healthServicesServer struct {
	procExec health.HealthExecution
	idgen    idgen.RandomIdGenerator
}

// NewExecutorServiceServer used to create a new execution context
func NewHealthServiceServer(exec health.HealthExecution, idgen idgen.RandomIdGenerator) protofiles.HealthServicesServer {
	return &healthServicesServer{
		procExec: exec,
		idgen:    idgen,
	}
}

func (e *healthServicesServer) Check(ctx context.Context, ping *protofiles.Ping) (*protofiles.HealthResponse, error) {
	controllerConfig, err := config.Loader()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pingTimeOut := controllerConfig.ExecutorPingDeadline
	res, err := e.procExec.UpdatePingStatus(ctx, ping, pingTimeOut)
	if err != nil {
		log.Error(err, fmt.Sprintf("executor id: %s, error in running health check", ping.ID))
		return nil, err
	}
	return res, err
}
