// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/executor_CP/register_message.proto

package executor_CP

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type ExecutorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ExecutorInfo) Reset() {
	*x = ExecutorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorInfo) ProtoMessage() {}

func (x *ExecutorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorInfo.ProtoReflect.Descriptor instead.
func (*ExecutorInfo) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutorInfo) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ExecutorInfo *ExecutorInfo `protobuf:"bytes,2,opt,name=executor_info,json=executorInfo,proto3" json:"executor_info,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RegisterRequest) GetExecutorInfo() *ExecutorInfo {
	if x != nil {
		return x.ExecutorInfo
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registered bool `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

var File_internal_pkg_protofiles_executor_CP_register_message_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x43, 0x50, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x42, 0x2e, 0x5a, 0x2c, 0x6f,
	0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x43, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescData = file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDesc
)

func file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDescData
}

var file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_pkg_protofiles_executor_CP_register_message_proto_goTypes = []interface{}{
	(*ExecutorInfo)(nil),     // 0: ExecutorInfo
	(*RegisterRequest)(nil),  // 1: RegisterRequest
	(*RegisterResponse)(nil), // 2: RegisterResponse
}
var file_internal_pkg_protofiles_executor_CP_register_message_proto_depIdxs = []int32{
	0, // 0: RegisterRequest.executor_info:type_name -> ExecutorInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_executor_CP_register_message_proto_init() }
func file_internal_pkg_protofiles_executor_CP_register_message_proto_init() {
	if File_internal_pkg_protofiles_executor_CP_register_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutorInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_executor_CP_register_message_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_executor_CP_register_message_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_executor_CP_register_message_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_executor_CP_register_message_proto = out.File
	file_internal_pkg_protofiles_executor_CP_register_message_proto_rawDesc = nil
	file_internal_pkg_protofiles_executor_CP_register_message_proto_goTypes = nil
	file_internal_pkg_protofiles_executor_CP_register_message_proto_depIdxs = nil
}
