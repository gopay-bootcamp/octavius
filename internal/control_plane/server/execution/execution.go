package execution

import (
	"context"
	"octavius/internal/control_plane/db/etcd"
	"octavius/pkg/protobuf"
)

type Execution interface {
	CreateProc(ctx context.Context, proc *protobuf.Proc) (string, error)
	ReadAllProc(ctx context.Context) ([]protobuf.Proc, error)
}

type execution struct {
	client etcd.EtcdClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewExec(dbClient etcd.EtcdClient) Execution {
	return &execution{
		client: dbClient,
	}
}

func (e *execution) CreateProc(ctx context.Context, proc *protobuf.Proc) (string, error) {

	result, err := e.client.PutValue(ctx, proc.Name, proc)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (e *execution) ReadAllProc(ctx context.Context) ([]protobuf.Proc, error) {
	procs, err := e.client.GetAllValues(ctx)
	if err != nil {
		return nil, err
	}
	return procs, nil
}
