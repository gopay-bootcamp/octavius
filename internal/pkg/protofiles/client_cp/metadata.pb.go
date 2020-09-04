// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: internal/pkg/protofiles/client_cp/metadata.proto

package client_cp

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

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Secret) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Arg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Required    bool   `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *Arg) Reset() {
	*x = Arg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arg) ProtoMessage() {}

func (x *Arg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Arg.ProtoReflect.Descriptor instead.
func (*Arg) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *Arg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Arg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Arg) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

type EnvVars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secrets []*Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	Args    []*Arg    `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *EnvVars) Reset() {
	*x = EnvVars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVars) ProtoMessage() {}

func (x *EnvVars) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVars.ProtoReflect.Descriptor instead.
func (*EnvVars) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *EnvVars) GetSecrets() []*Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *EnvVars) GetArgs() []*Arg {
	if x != nil {
		return x.Args
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description      string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageName        string   `protobuf:"bytes,3,opt,name=image_name,proto3" json:"image_name,omitempty"`
	EnvVars          *EnvVars `protobuf:"bytes,4,opt,name=env_vars,proto3" json:"env_vars,omitempty"`
	AuthorizedGroups []string `protobuf:"bytes,5,rep,name=authorized_groups,proto3" json:"authorized_groups,omitempty"`
	Author           string   `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Contributors     string   `protobuf:"bytes,7,opt,name=contributors,proto3" json:"contributors,omitempty"`
	Organization     string   `protobuf:"bytes,8,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Metadata) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Metadata) GetEnvVars() *EnvVars {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *Metadata) GetAuthorizedGroups() []string {
	if x != nil {
		return x.AuthorizedGroups
	}
	return nil
}

func (x *Metadata) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Metadata) GetContributors() string {
	if x != nil {
		return x.Contributors
	}
	return ""
}

func (x *Metadata) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

var File_internal_pkg_protofiles_client_cp_metadata_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_client_cp_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x70, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x03, 0x41, 0x72, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x07, 0x45,
	0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x41, 0x72, 0x67, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x22, 0x94, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61,
	0x72, 0x73, 0x52, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x6f, 0x63,
	0x74, 0x61, 0x76, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescOnce sync.Once
	file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescData = file_internal_pkg_protofiles_client_cp_metadata_proto_rawDesc
)

func file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescGZIP() []byte {
	file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescOnce.Do(func() {
		file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescData)
	})
	return file_internal_pkg_protofiles_client_cp_metadata_proto_rawDescData
}

var file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_pkg_protofiles_client_cp_metadata_proto_goTypes = []interface{}{
	(*Secret)(nil),   // 0: Secret
	(*Arg)(nil),      // 1: Arg
	(*EnvVars)(nil),  // 2: EnvVars
	(*Metadata)(nil), // 3: Metadata
}
var file_internal_pkg_protofiles_client_cp_metadata_proto_depIdxs = []int32{
	0, // 0: EnvVars.secrets:type_name -> Secret
	1, // 1: EnvVars.args:type_name -> Arg
	2, // 2: Metadata.env_vars:type_name -> EnvVars
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_client_cp_metadata_proto_init() }
func file_internal_pkg_protofiles_client_cp_metadata_proto_init() {
	if File_internal_pkg_protofiles_client_cp_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Arg); i {
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
		file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvVars); i {
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
		file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_internal_pkg_protofiles_client_cp_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_client_cp_metadata_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_client_cp_metadata_proto_depIdxs,
		MessageInfos:      file_internal_pkg_protofiles_client_cp_metadata_proto_msgTypes,
	}.Build()
	File_internal_pkg_protofiles_client_cp_metadata_proto = out.File
	file_internal_pkg_protofiles_client_cp_metadata_proto_rawDesc = nil
	file_internal_pkg_protofiles_client_cp_metadata_proto_goTypes = nil
	file_internal_pkg_protofiles_client_cp_metadata_proto_depIdxs = nil
}
