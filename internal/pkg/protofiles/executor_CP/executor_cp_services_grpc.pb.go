// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package executor_cp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExecutorCPServicesClient is the client API for ExecutorCPServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorCPServicesClient interface {
	HealthCheck(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*HealthResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	StreamJobs(ctx context.Context, in *Start, opts ...grpc.CallOption) (ExecutorCPServices_StreamJobsClient, error)
}

type executorCPServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorCPServicesClient(cc grpc.ClientConnInterface) ExecutorCPServicesClient {
	return &executorCPServicesClient{cc}
}

func (c *executorCPServicesClient) HealthCheck(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/ExecutorCPServices/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorCPServicesClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/ExecutorCPServices/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorCPServicesClient) StreamJobs(ctx context.Context, in *Start, opts ...grpc.CallOption) (ExecutorCPServices_StreamJobsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ExecutorCPServices_serviceDesc.Streams[0], "/ExecutorCPServices/StreamJobs", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorCPServicesStreamJobsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExecutorCPServices_StreamJobsClient interface {
	Recv() (*Job, error)
	grpc.ClientStream
}

type executorCPServicesStreamJobsClient struct {
	grpc.ClientStream
}

func (x *executorCPServicesStreamJobsClient) Recv() (*Job, error) {
	m := new(Job)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExecutorCPServicesServer is the server API for ExecutorCPServices service.
// All implementations should embed UnimplementedExecutorCPServicesServer
// for forward compatibility
type ExecutorCPServicesServer interface {
	HealthCheck(context.Context, *Ping) (*HealthResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	StreamJobs(*Start, ExecutorCPServices_StreamJobsServer) error
}

// UnimplementedExecutorCPServicesServer should be embedded to have forward compatible implementations.
type UnimplementedExecutorCPServicesServer struct {
}

func (*UnimplementedExecutorCPServicesServer) HealthCheck(context.Context, *Ping) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (*UnimplementedExecutorCPServicesServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedExecutorCPServicesServer) StreamJobs(*Start, ExecutorCPServices_StreamJobsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamJobs not implemented")
}

func RegisterExecutorCPServicesServer(s *grpc.Server, srv ExecutorCPServicesServer) {
	s.RegisterService(&_ExecutorCPServices_serviceDesc, srv)
}

func _ExecutorCPServices_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorCPServicesServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExecutorCPServices/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorCPServicesServer).HealthCheck(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutorCPServices_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorCPServicesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExecutorCPServices/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorCPServicesServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutorCPServices_StreamJobs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Start)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorCPServicesServer).StreamJobs(m, &executorCPServicesStreamJobsServer{stream})
}

type ExecutorCPServices_StreamJobsServer interface {
	Send(*Job) error
	grpc.ServerStream
}

type executorCPServicesStreamJobsServer struct {
	grpc.ServerStream
}

func (x *executorCPServicesStreamJobsServer) Send(m *Job) error {
	return x.ServerStream.SendMsg(m)
}

var _ExecutorCPServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExecutorCPServices",
	HandlerType: (*ExecutorCPServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ExecutorCPServices_HealthCheck_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ExecutorCPServices_Register_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamJobs",
			Handler:       _ExecutorCPServices_StreamJobs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/protofiles/executor_cp/executor_cp_services.proto",
}
