package server

import (
	"context"
	"octavius/pkg/model/proc"
	procProto "octavius/pkg/protobuf"
)

type procServiceServer struct {
	procExec execution.Execution
}

func NewProcServiceServer(exec execution.Execution) procProto.ProcServiceServer {
	return &procServiceServer{
		procExec: exec,
	}
}

func (s *procServiceServer) CreateProc(ctx context.Context, request *procProto.RequestForCreateProc) (*procProto.ProcID, error) {
	var proc model.Proc
	proc.Name = request.Name
	proc.Author = request.Author
	id, err := s.procExec.CreateProc(ctx, &proc)
	if err != nil {
		return nil, err
	}
	resp := &procProto.ProcID{Value: id, Message: "successfully created Proc"}
	return resp, nil
}

func (s *procServiceServer) ReadAllProcs(ctx context.Context, request *procProto.RequestForReadAllProcs) (*procProto.ProcList, error) {
	procList, err := s.procExec.ReadAllProc(ctx)
	if err != nil {
		return nil, err
	}
	protoProcs := []*procProto.Proc{}
	for _, proc := range procList {
		protoProc := procProto.Proc{ID: proc.ID, Name: proc.Name, Author: proc.Author}
		protoProcs = append(protoProcs, &protoProc)
	}
	return &procProto.ProcList{Procs: protoProcs}, nil
}
