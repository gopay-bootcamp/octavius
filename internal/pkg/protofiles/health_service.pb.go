// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: internal/pkg/protofiles/health_service.proto

package protofiles

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var File_internal_pkg_protofiles_health_service_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_health_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x31, 0x0a, 0x0e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a,
	0x20, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_pkg_protofiles_health_service_proto_goTypes = []interface{}{
	(*Ping)(nil),           // 0: Ping
	(*HealthResponse)(nil), // 1: HealthResponse
}
var file_internal_pkg_protofiles_health_service_proto_depIdxs = []int32{
	0, // 0: HealthServices.Check:input_type -> Ping
	1, // 1: HealthServices.Check:output_type -> HealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_health_service_proto_init() }
func file_internal_pkg_protofiles_health_service_proto_init() {
	if File_internal_pkg_protofiles_health_service_proto != nil {
		return
	}
	file_internal_pkg_protofiles_ping_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_health_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_protofiles_health_service_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_health_service_proto_depIdxs,
	}.Build()
	File_internal_pkg_protofiles_health_service_proto = out.File
	file_internal_pkg_protofiles_health_service_proto_rawDesc = nil
	file_internal_pkg_protofiles_health_service_proto_goTypes = nil
	file_internal_pkg_protofiles_health_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServicesClient is the client API for HealthServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServicesClient interface {
	Check(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServicesClient(cc grpc.ClientConnInterface) HealthServicesClient {
	return &healthServicesClient{cc}
}

func (c *healthServicesClient) Check(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/HealthServices/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServicesServer is the server API for HealthServices service.
type HealthServicesServer interface {
	Check(context.Context, *Ping) (*HealthResponse, error)
}

// UnimplementedHealthServicesServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServicesServer struct {
}

func (*UnimplementedHealthServicesServer) Check(context.Context, *Ping) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterHealthServicesServer(s *grpc.Server, srv HealthServicesServer) {
	s.RegisterService(&_HealthServices_serviceDesc, srv)
}

func _HealthServices_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServicesServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HealthServices/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServicesServer).Check(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HealthServices",
	HandlerType: (*HealthServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HealthServices_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pkg/protofiles/health_service.proto",
}
