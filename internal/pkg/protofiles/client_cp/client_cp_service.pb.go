// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/client_cp/client_cp_service.proto

package client_cp

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_internal_pkg_protofiles_client_cp_client_cp_service_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_client_cp_client_cp_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x70, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x90, 0x02, 0x0a,
	0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x1a, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x30, 0x01,
	0x12, 0x2b, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x12,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x0e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x2d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x2c, 0x5a, 0x2a, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_pkg_protofiles_client_cp_client_cp_service_proto_goTypes = []interface{}{
	(*RequestForStreamLog)(nil),     // 0: RequestForStreamLog
	(*RequestForExecute)(nil),       // 1: RequestForExecute
	(*RequestToPostMetadata)(nil),   // 2: RequestToPostMetadata
	(*RequestToGetAllMetadata)(nil), // 3: RequestToGetAllMetadata
	(*RequestForGetJobList)(nil),    // 4: RequestForGetJobList
	(*Log)(nil),                     // 5: Log
	(*Response)(nil),                // 6: Response
	(*MetadataName)(nil),            // 7: MetadataName
	(*MetadataArray)(nil),           // 8: MetadataArray
	(*JobList)(nil),                 // 9: JobList
}
var file_internal_pkg_protofiles_client_cp_client_cp_service_proto_depIdxs = []int32{
	0, // 0: ClientCPServices.GetStreamLogs:input_type -> RequestForStreamLog
	1, // 1: ClientCPServices.ExecuteJob:input_type -> RequestForExecute
	2, // 2: ClientCPServices.PostMetadata:input_type -> RequestToPostMetadata
	3, // 3: ClientCPServices.GetAllMetadata:input_type -> RequestToGetAllMetadata
	4, // 4: ClientCPServices.GetJobList:input_type -> RequestForGetJobList
	5, // 5: ClientCPServices.GetStreamLogs:output_type -> Log
	6, // 6: ClientCPServices.ExecuteJob:output_type -> Response
	7, // 7: ClientCPServices.PostMetadata:output_type -> MetadataName
	8, // 8: ClientCPServices.GetAllMetadata:output_type -> MetadataArray
	9, // 9: ClientCPServices.GetJobList:output_type -> JobList
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_client_cp_client_cp_service_proto_init() }
func file_internal_pkg_protofiles_client_cp_client_cp_service_proto_init() {
	if File_internal_pkg_protofiles_client_cp_client_cp_service_proto != nil {
		return
	}
	file_internal_pkg_protofiles_client_cp_metadata_messages_proto_init()
	file_internal_pkg_protofiles_client_cp_job_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_client_cp_client_cp_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_protofiles_client_cp_client_cp_service_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_client_cp_client_cp_service_proto_depIdxs,
	}.Build()
	File_internal_pkg_protofiles_client_cp_client_cp_service_proto = out.File
	file_internal_pkg_protofiles_client_cp_client_cp_service_proto_rawDesc = nil
	file_internal_pkg_protofiles_client_cp_client_cp_service_proto_goTypes = nil
	file_internal_pkg_protofiles_client_cp_client_cp_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientCPServicesClient is the client API for ClientCPServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientCPServicesClient interface {
	GetStreamLogs(ctx context.Context, in *RequestForStreamLog, opts ...grpc.CallOption) (ClientCPServices_GetStreamLogsClient, error)
	ExecuteJob(ctx context.Context, in *RequestForExecute, opts ...grpc.CallOption) (*Response, error)
	PostMetadata(ctx context.Context, in *RequestToPostMetadata, opts ...grpc.CallOption) (*MetadataName, error)
	GetAllMetadata(ctx context.Context, in *RequestToGetAllMetadata, opts ...grpc.CallOption) (*MetadataArray, error)
	GetJobList(ctx context.Context, in *RequestForGetJobList, opts ...grpc.CallOption) (*JobList, error)
}

type clientCPServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewClientCPServicesClient(cc grpc.ClientConnInterface) ClientCPServicesClient {
	return &clientCPServicesClient{cc}
}

func (c *clientCPServicesClient) GetStreamLogs(ctx context.Context, in *RequestForStreamLog, opts ...grpc.CallOption) (ClientCPServices_GetStreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientCPServices_serviceDesc.Streams[0], "/ClientCPServices/GetStreamLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientCPServicesGetStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientCPServices_GetStreamLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type clientCPServicesGetStreamLogsClient struct {
	grpc.ClientStream
}

func (x *clientCPServicesGetStreamLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientCPServicesClient) ExecuteJob(ctx context.Context, in *RequestForExecute, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ClientCPServices/ExecuteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientCPServicesClient) PostMetadata(ctx context.Context, in *RequestToPostMetadata, opts ...grpc.CallOption) (*MetadataName, error) {
	out := new(MetadataName)
	err := c.cc.Invoke(ctx, "/ClientCPServices/PostMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientCPServicesClient) GetAllMetadata(ctx context.Context, in *RequestToGetAllMetadata, opts ...grpc.CallOption) (*MetadataArray, error) {
	out := new(MetadataArray)
	err := c.cc.Invoke(ctx, "/ClientCPServices/GetAllMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientCPServicesClient) GetJobList(ctx context.Context, in *RequestForGetJobList, opts ...grpc.CallOption) (*JobList, error) {
	out := new(JobList)
	err := c.cc.Invoke(ctx, "/ClientCPServices/GetJobList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientCPServicesServer is the server API for ClientCPServices service.
type ClientCPServicesServer interface {
	GetStreamLogs(*RequestForStreamLog, ClientCPServices_GetStreamLogsServer) error
	ExecuteJob(context.Context, *RequestForExecute) (*Response, error)
	PostMetadata(context.Context, *RequestToPostMetadata) (*MetadataName, error)
	GetAllMetadata(context.Context, *RequestToGetAllMetadata) (*MetadataArray, error)
	GetJobList(context.Context, *RequestForGetJobList) (*JobList, error)
}

// UnimplementedClientCPServicesServer can be embedded to have forward compatible implementations.
type UnimplementedClientCPServicesServer struct {
}

func (*UnimplementedClientCPServicesServer) GetStreamLogs(*RequestForStreamLog, ClientCPServices_GetStreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamLogs not implemented")
}
func (*UnimplementedClientCPServicesServer) ExecuteJob(context.Context, *RequestForExecute) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteJob not implemented")
}
func (*UnimplementedClientCPServicesServer) PostMetadata(context.Context, *RequestToPostMetadata) (*MetadataName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMetadata not implemented")
}
func (*UnimplementedClientCPServicesServer) GetAllMetadata(context.Context, *RequestToGetAllMetadata) (*MetadataArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMetadata not implemented")
}
func (*UnimplementedClientCPServicesServer) GetJobList(context.Context, *RequestForGetJobList) (*JobList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobList not implemented")
}

func RegisterClientCPServicesServer(s *grpc.Server, srv ClientCPServicesServer) {
	s.RegisterService(&_ClientCPServices_serviceDesc, srv)
}

func _ClientCPServices_GetStreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestForStreamLog)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientCPServicesServer).GetStreamLogs(m, &clientCPServicesGetStreamLogsServer{stream})
}

type ClientCPServices_GetStreamLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type clientCPServicesGetStreamLogsServer struct {
	grpc.ServerStream
}

func (x *clientCPServicesGetStreamLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientCPServices_ExecuteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientCPServicesServer).ExecuteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientCPServices/ExecuteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientCPServicesServer).ExecuteJob(ctx, req.(*RequestForExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientCPServices_PostMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToPostMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientCPServicesServer).PostMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientCPServices/PostMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientCPServicesServer).PostMetadata(ctx, req.(*RequestToPostMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientCPServices_GetAllMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToGetAllMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientCPServicesServer).GetAllMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientCPServices/GetAllMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientCPServicesServer).GetAllMetadata(ctx, req.(*RequestToGetAllMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientCPServices_GetJobList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForGetJobList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientCPServicesServer).GetJobList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientCPServices/GetJobList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientCPServicesServer).GetJobList(ctx, req.(*RequestForGetJobList))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientCPServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClientCPServices",
	HandlerType: (*ClientCPServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteJob",
			Handler:    _ClientCPServices_ExecuteJob_Handler,
		},
		{
			MethodName: "PostMetadata",
			Handler:    _ClientCPServices_PostMetadata_Handler,
		},
		{
			MethodName: "GetAllMetadata",
			Handler:    _ClientCPServices_GetAllMetadata_Handler,
		},
		{
			MethodName: "GetJobList",
			Handler:    _ClientCPServices_GetJobList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamLogs",
			Handler:       _ClientCPServices_GetStreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/protofiles/client_cp/client_cp_service.proto",
}
