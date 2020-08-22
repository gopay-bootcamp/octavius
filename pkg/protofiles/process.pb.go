// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: pkg/protobuf/process.proto

package protobuf

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int64  `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetErrorCode() int64 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

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
		mi := &file_pkg_protobuf_process_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[1]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{1}
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
}

func (x *Arg) Reset() {
	*x = Arg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Arg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Arg) ProtoMessage() {}

func (x *Arg) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[2]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{2}
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
		mi := &file_pkg_protobuf_process_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVars) ProtoMessage() {}

func (x *EnvVars) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[3]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{3}
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
		mi := &file_pkg_protobuf_process_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[4]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{4}
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

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientEmail string `protobuf:"bytes,1,opt,name=client_email,json=clientEmail,proto3" json:"client_email,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{5}
}

func (x *ClientInfo) GetClientEmail() string {
	if x != nil {
		return x.ClientEmail
	}
	return ""
}

func (x *ClientInfo) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

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
		mi := &file_pkg_protobuf_process_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToPostMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToPostMetadata) ProtoMessage() {}

func (x *RequestToPostMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[6]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{6}
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

	Err  *Error `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MetadataName) Reset() {
	*x = MetadataName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataName) ProtoMessage() {}

func (x *MetadataName) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[7]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{7}
}

func (x *MetadataName) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *MetadataName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestToGetAllMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestToGetAllMetadata) Reset() {
	*x = RequestToGetAllMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestToGetAllMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToGetAllMetadata) ProtoMessage() {}

func (x *RequestToGetAllMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[8]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{8}
}

type MetadataArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err   *Error      `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Value []*Metadata `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *MetadataArray) Reset() {
	*x = MetadataArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_process_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataArray) ProtoMessage() {}

func (x *MetadataArray) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_process_proto_msgTypes[9]
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
	return file_pkg_protobuf_process_proto_rawDescGZIP(), []int{9}
}

func (x *MetadataArray) GetErr() *Error {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *MetadataArray) GetValue() []*Metadata {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_pkg_protobuf_process_proto protoreflect.FileDescriptor

var file_pkg_protobuf_process_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x03, 0x61, 0x72, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73,
	0x12, 0x21, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x04, 0x2e, 0x61, 0x72, 0x67, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x95, 0x02,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x52, 0x08, 0x65,
	0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x18, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a,
	0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3c, 0x0a, 0x0c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x8f, 0x01, 0x0a, 0x10, 0x4f, 0x63, 0x74, 0x61, 0x76, 0x69,
	0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x50, 0x6f,
	0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x41, 0x72, 0x72, 0x61, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x6f, 0x63, 0x74, 0x61, 0x76,
	0x69, 0x75, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protobuf_process_proto_rawDescOnce sync.Once
	file_pkg_protobuf_process_proto_rawDescData = file_pkg_protobuf_process_proto_rawDesc
)

func file_pkg_protobuf_process_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_process_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_process_proto_rawDescData)
	})
	return file_pkg_protobuf_process_proto_rawDescData
}

var file_pkg_protobuf_process_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_protobuf_process_proto_goTypes = []interface{}{
	(*Error)(nil),                   // 0: Error
	(*Secret)(nil),                  // 1: secret
	(*Arg)(nil),                     // 2: arg
	(*EnvVars)(nil),                 // 3: env_vars
	(*Metadata)(nil),                // 4: metadata
	(*ClientInfo)(nil),              // 5: client_info
	(*RequestToPostMetadata)(nil),   // 6: request_to_post_metadata
	(*MetadataName)(nil),            // 7: MetadataName
	(*RequestToGetAllMetadata)(nil), // 8: request_to_get_all_metadata
	(*MetadataArray)(nil),           // 9: metadataArray
}
var file_pkg_protobuf_process_proto_depIdxs = []int32{
	1,  // 0: env_vars.secrets:type_name -> secret
	2,  // 1: env_vars.args:type_name -> arg
	3,  // 2: metadata.env_vars:type_name -> env_vars
	4,  // 3: request_to_post_metadata.metadata:type_name -> metadata
	5,  // 4: request_to_post_metadata.client_info:type_name -> client_info
	0,  // 5: MetadataName.err:type_name -> Error
	0,  // 6: metadataArray.err:type_name -> Error
	4,  // 7: metadataArray.value:type_name -> metadata
	6,  // 8: OctaviusServices.Post_metadata:input_type -> request_to_post_metadata
	8,  // 9: OctaviusServices.Get_all_metadata:input_type -> request_to_get_all_metadata
	7,  // 10: OctaviusServices.Post_metadata:output_type -> MetadataName
	9,  // 11: OctaviusServices.Get_all_metadata:output_type -> metadataArray
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_process_proto_init() }
func file_pkg_protobuf_process_proto_init() {
	if File_pkg_protobuf_process_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_process_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_pkg_protobuf_process_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_pkg_protobuf_process_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_process_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_protobuf_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protobuf_process_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_process_proto_depIdxs,
		MessageInfos:      file_pkg_protobuf_process_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_process_proto = out.File
	file_pkg_protobuf_process_proto_rawDesc = nil
	file_pkg_protobuf_process_proto_goTypes = nil
	file_pkg_protobuf_process_proto_depIdxs = nil
}
