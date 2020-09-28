// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: internal/pkg/protofiles/metadata_messages.proto

package protofiles

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

// PostMetadata service
type RequestToPostMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata   *Metadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ClientInfo *ClientInfo `protobuf:"bytes,2,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
}

func (x *RequestToPostMetadata) Reset() {
	*x = RequestToPostMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToPostMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToPostMetadata) ProtoMessage() {}

func (x *RequestToPostMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToPostMetadata.ProtoReflect.Descriptor instead.
func (*RequestToPostMetadata) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RequestToPostMetadata) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RequestToPostMetadata) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

type MetadataName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MetadataName) Reset() {
	*x = MetadataName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataName) ProtoMessage() {}

func (x *MetadataName) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataName.ProtoReflect.Descriptor instead.
func (*MetadataName) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GetAllMetadata service
type RequestToGetAllMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestToGetAllMetadata) Reset() {
	*x = RequestToGetAllMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToGetAllMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToGetAllMetadata) ProtoMessage() {}

func (x *RequestToGetAllMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToGetAllMetadata.ProtoReflect.Descriptor instead.
func (*RequestToGetAllMetadata) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP(), []int{2}
}

type MetadataArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Metadata `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MetadataArray) Reset() {
	*x = MetadataArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataArray) ProtoMessage() {}

func (x *MetadataArray) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataArray.ProtoReflect.Descriptor instead.
func (*MetadataArray) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataArray) GetValues() []*Metadata {
	if x != nil {
		return x.Values
	}
	return nil
}

// Describe Service
type RequestToDescribe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientInfo *ClientInfo `protobuf:"bytes,1,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	JobName    string      `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *RequestToDescribe) Reset() {
	*x = RequestToDescribe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToDescribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToDescribe) ProtoMessage() {}

func (x *RequestToDescribe) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToDescribe.ProtoReflect.Descriptor instead.
func (*RequestToDescribe) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP(), []int{4}
}

func (x *RequestToDescribe) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *RequestToDescribe) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

var File_internal_pkg_protofiles_metadata_messages_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_metadata_messages_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x6f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x6f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x32, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x21, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x6f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_metadata_messages_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_metadata_messages_proto_rawDescData = file_internal_pkg_protofiles_metadata_messages_proto_rawDesc
)

func file_internal_pkg_protofiles_metadata_messages_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_metadata_messages_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_metadata_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_metadata_messages_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_metadata_messages_proto_rawDescData
}

var file_internal_pkg_protofiles_metadata_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pkg_protofiles_metadata_messages_proto_goTypes = []interface{}{
	(*RequestToPostMetadata)(nil),   // 0: RequestToPostMetadata
	(*MetadataName)(nil),            // 1: MetadataName
	(*RequestToGetAllMetadata)(nil), // 2: RequestToGetAllMetadata
	(*MetadataArray)(nil),           // 3: MetadataArray
	(*RequestToDescribe)(nil),       // 4: RequestToDescribe
	(*Metadata)(nil),                // 5: Metadata
	(*ClientInfo)(nil),              // 6: ClientInfo
}
var file_internal_pkg_protofiles_metadata_messages_proto_depIdxs = []int32{
	5, // 0: RequestToPostMetadata.metadata:type_name -> Metadata
	6, // 1: RequestToPostMetadata.client_info:type_name -> ClientInfo
	5, // 2: MetadataArray.values:type_name -> Metadata
	6, // 3: RequestToDescribe.client_info:type_name -> ClientInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_metadata_messages_proto_init() }
func file_internal_pkg_protofiles_metadata_messages_proto_init() {
	if File_internal_pkg_protofiles_metadata_messages_proto != nil {
		return
	}
	file_internal_pkg_protofiles_metadata_proto_init()
	file_internal_pkg_protofiles_client_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestToPostMetadata); i {
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
		file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataName); i {
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
		file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestToGetAllMetadata); i {
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
		file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataArray); i {
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
		file_internal_pkg_protofiles_metadata_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestToDescribe); i {
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
			RawDescriptor: file_internal_pkg_protofiles_metadata_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_metadata_messages_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_metadata_messages_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_metadata_messages_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_metadata_messages_proto = out.File
	file_internal_pkg_protofiles_metadata_messages_proto_rawDesc = nil
	file_internal_pkg_protofiles_metadata_messages_proto_goTypes = nil
	file_internal_pkg_protofiles_metadata_messages_proto_depIdxs = nil
}
