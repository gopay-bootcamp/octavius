package server

import (
	"context"
	"octavius/internal/control_plane/server/execution"
	procProto "octavius/pkg/protobuf"
)

type procServiceServer struct {
	procExec execution.Execution
}

func NewProcServiceServer(exec execution.Execution) procProto.OctaviusServicesServer {
	return &procServiceServer{
		procExec: exec,
	}
}

func (s *procServiceServer) PostMetadata(ctx context.Context, request *procProto.RequestToPostMetadata) (*procProto.MetadataID, error) {
	id := s.procExec.SaveMetadataToDb(ctx, request.Metadata)
	return id, nil
}

// func (s *procServiceServer) ReadAllProcs(ctx context.Context, request *procProto.RequestForReadAllProcs) (*procProto.ProcList, error) {
// 	procList, err := s.procExec.ReadAllProc(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	protoProcs := []*procProto.Proc{}
// 	for _, proc := range procList {
// 		protoProc := procProto.Proc{ID: proc.ID, Name: proc.Name, Author: proc.Author}
// 		protoProcs = append(protoProcs, &protoProc)
// 	}
// 	return &procProto.ProcList{Procs: protoProcs}, nil
// }
