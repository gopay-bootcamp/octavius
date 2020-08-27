// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/client_CP/job_messages.proto

package client_CP

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

// GetStreamLogs service
type RequestForStreamLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientInfo *ClientInfo `protobuf:"bytes,1,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	JobName    string      `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *RequestForStreamLog) Reset() {
	*x = RequestForStreamLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestForStreamLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestForStreamLog) ProtoMessage() {}

func (x *RequestForStreamLog) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestForStreamLog.ProtoReflect.Descriptor instead.
func (*RequestForStreamLog) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RequestForStreamLog) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *RequestForStreamLog) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId uint64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Log       string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Log) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

// ExecuteJobs service
type RequestForExecute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientInfo *ClientInfo       `protobuf:"bytes,1,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	JobName    string            `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	JobData    map[string]string `protobuf:"bytes,3,rep,name=job_data,json=jobData,proto3" json:"job_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestForExecute) Reset() {
	*x = RequestForExecute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestForExecute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestForExecute) ProtoMessage() {}

func (x *RequestForExecute) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestForExecute.ProtoReflect.Descriptor instead.
func (*RequestForExecute) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescGZIP(), []int{2}
}

func (x *RequestForExecute) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *RequestForExecute) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *RequestForExecute) GetJobData() map[string]string {
	if x != nil {
		return x.JobData
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_internal_pkg_protofiles_client_CP_job_messages_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x43, 0x50, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x43, 0x50, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x13, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c,
	0x6f, 0x67, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6c, 0x6f, 0x67, 0x22, 0xd4, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x6f, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f,
	0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3a,
	0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x2c,
	0x5a, 0x2a, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x43, 0x50, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescData = file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDesc
)

func file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDescData
}

var file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pkg_protofiles_client_CP_job_messages_proto_goTypes = []interface{}{
	(*RequestForStreamLog)(nil), // 0: RequestForStreamLog
	(*Log)(nil),                 // 1: Log
	(*RequestForExecute)(nil),   // 2: RequestForExecute
	(*Response)(nil),            // 3: Response
	nil,                         // 4: RequestForExecute.JobDataEntry
	(*ClientInfo)(nil),          // 5: ClientInfo
}
var file_internal_pkg_protofiles_client_CP_job_messages_proto_depIdxs = []int32{
	5, // 0: RequestForStreamLog.client_info:type_name -> ClientInfo
	5, // 1: RequestForExecute.client_info:type_name -> ClientInfo
	4, // 2: RequestForExecute.job_data:type_name -> RequestForExecute.JobDataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_client_CP_job_messages_proto_init() }
func file_internal_pkg_protofiles_client_CP_job_messages_proto_init() {
	if File_internal_pkg_protofiles_client_CP_job_messages_proto != nil {
		return
	}
	file_internal_pkg_protofiles_client_CP_client_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestForStreamLog); i {
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
		file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestForExecute); i {
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
		file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_client_CP_job_messages_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_client_CP_job_messages_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_client_CP_job_messages_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_client_CP_job_messages_proto = out.File
	file_internal_pkg_protofiles_client_CP_job_messages_proto_rawDesc = nil
	file_internal_pkg_protofiles_client_CP_job_messages_proto_goTypes = nil
	file_internal_pkg_protofiles_client_CP_job_messages_proto_depIdxs = nil
}
