// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: internal/pkg/protofiles/executor_cp/executor_cp_services.proto

package executor_cp

import (
	proto "github.com/golang/protobuf/proto"
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

var File_internal_pkg_protofiles_executor_cp_executor_cp_services_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x36, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8a, 0x01, 0x0a, 0x12,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x05, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x06, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x1a, 0x04, 0x2e, 0x4a, 0x6f, 0x62, 0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x6f, 0x63, 0x74, 0x61,
	0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_goTypes = []interface{}{
	(*Ping)(nil),             // 0: Ping
	(*RegisterRequest)(nil),  // 1: RegisterRequest
	(*Start)(nil),            // 2: Start
	(*HealthResponse)(nil),   // 3: HealthResponse
	(*RegisterResponse)(nil), // 4: RegisterResponse
	(*Job)(nil),              // 5: Job
}
var file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_depIdxs = []int32{
	0, // 0: ExecutorCPServices.HealthCheck:input_type -> Ping
	1, // 1: ExecutorCPServices.Register:input_type -> RegisterRequest
	2, // 2: ExecutorCPServices.StreamJobs:input_type -> Start
	3, // 3: ExecutorCPServices.HealthCheck:output_type -> HealthResponse
	4, // 4: ExecutorCPServices.Register:output_type -> RegisterResponse
	5, // 5: ExecutorCPServices.StreamJobs:output_type -> Job
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_init() }
func file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_init() {
	if File_internal_pkg_protofiles_executor_cp_executor_cp_services_proto != nil {
		return
	}
	file_internal_pkg_protofiles_executor_cp_ping_message_proto_init()
	file_internal_pkg_protofiles_executor_cp_register_message_proto_init()
	file_internal_pkg_protofiles_executor_cp_job_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_depIdxs,
	}.Build()
	File_internal_pkg_protofiles_executor_cp_executor_cp_services_proto = out.File
	file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_rawDesc = nil
	file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_goTypes = nil
	file_internal_pkg_protofiles_executor_cp_executor_cp_services_proto_depIdxs = nil
}
