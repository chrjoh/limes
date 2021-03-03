// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ims.proto

package ims

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{0}
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error           string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Role            string `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
	AccessKeyId     string `protobuf:"bytes,3,opt,name=AccessKeyId,proto3" json:"AccessKeyId,omitempty"`
	SecretAccessKey string `protobuf:"bytes,4,opt,name=SecretAccessKey,proto3" json:"SecretAccessKey,omitempty"`
	SessionToken    string `protobuf:"bytes,5,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	Expiration      string `protobuf:"bytes,6,opt,name=Expiration,proto3" json:"Expiration,omitempty"`
	Region          string `protobuf:"bytes,7,opt,name=Region,proto3" json:"Region,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{1}
}

func (x *StatusReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *StatusReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *StatusReply) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *StatusReply) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

func (x *StatusReply) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *StatusReply) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *StatusReply) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type StopReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *StopReply) Reset() {
	*x = StopReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReply) ProtoMessage() {}

func (x *StopReply) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReply.ProtoReflect.Descriptor instead.
func (*StopReply) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{2}
}

func (x *StopReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AssumeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Mfa  string `protobuf:"bytes,2,opt,name=Mfa,proto3" json:"Mfa,omitempty"`
}

func (x *AssumeRoleRequest) Reset() {
	*x = AssumeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssumeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssumeRoleRequest) ProtoMessage() {}

func (x *AssumeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssumeRoleRequest.ProtoReflect.Descriptor instead.
func (*AssumeRoleRequest) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{3}
}

func (x *AssumeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssumeRoleRequest) GetMfa() string {
	if x != nil {
		return x.Mfa
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwsAccessKeyID     string `protobuf:"bytes,1,opt,name=AwsAccessKeyID,proto3" json:"AwsAccessKeyID,omitempty"`
	AwsSecretAccessKey string `protobuf:"bytes,2,opt,name=AwsSecretAccessKey,proto3" json:"AwsSecretAccessKey,omitempty"`
	AwsSessionToken    string `protobuf:"bytes,3,opt,name=AwsSessionToken,proto3" json:"AwsSessionToken,omitempty"`
	Region             string `protobuf:"bytes,4,opt,name=Region,proto3" json:"Region,omitempty"`
	MFASerial          string `protobuf:"bytes,5,opt,name=MFASerial,proto3" json:"MFASerial,omitempty"`
	RoleARN            string `protobuf:"bytes,6,opt,name=RoleARN,proto3" json:"RoleARN,omitempty"`
	SourceProfile      string `protobuf:"bytes,7,opt,name=SourceProfile,proto3" json:"SourceProfile,omitempty"`
	RoleSessionName    string `protobuf:"bytes,8,opt,name=RoleSessionName,proto3" json:"RoleSessionName,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{4}
}

func (x *Profile) GetAwsAccessKeyID() string {
	if x != nil {
		return x.AwsAccessKeyID
	}
	return ""
}

func (x *Profile) GetAwsSecretAccessKey() string {
	if x != nil {
		return x.AwsSecretAccessKey
	}
	return ""
}

func (x *Profile) GetAwsSessionToken() string {
	if x != nil {
		return x.AwsSessionToken
	}
	return ""
}

func (x *Profile) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Profile) GetMFASerial() string {
	if x != nil {
		return x.MFASerial
	}
	return ""
}

func (x *Profile) GetRoleARN() string {
	if x != nil {
		return x.RoleARN
	}
	return ""
}

func (x *Profile) GetSourceProfile() string {
	if x != nil {
		return x.SourceProfile
	}
	return ""
}

func (x *Profile) GetRoleSessionName() string {
	if x != nil {
		return x.RoleSessionName
	}
	return ""
}

type ConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles map[string]*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfigReply) Reset() {
	*x = ConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ims_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigReply) ProtoMessage() {}

func (x *ConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_ims_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigReply.ProtoReflect.Descriptor instead.
func (*ConfigReply) Descriptor() ([]byte, []int) {
	return file_ims_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigReply) GetProfiles() map[string]*Profile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

var File_ims_proto protoreflect.FileDescriptor

var file_ims_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x6d, 0x73,
	0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x09, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x39, 0x0a,
	0x11, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x66, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x66, 0x61, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x77,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12,
	0x41, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x41, 0x77, 0x73, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x41, 0x77, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x77, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x4d, 0x46, 0x41, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4d, 0x46, 0x41, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x6f, 0x6c, 0x65, 0x41, 0x52, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x6f, 0x6c, 0x65, 0x41, 0x52, 0x4e, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x49, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x82, 0x02,
	0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x09, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x10, 0x2e, 0x69, 0x6d, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x23,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x09, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x0e, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x6d, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x69, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x09, 0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x10,
	0x2e, 0x69, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x06, 0xa2, 0x02, 0x03, 0x49, 0x4d, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ims_proto_rawDescOnce sync.Once
	file_ims_proto_rawDescData = file_ims_proto_rawDesc
)

func file_ims_proto_rawDescGZIP() []byte {
	file_ims_proto_rawDescOnce.Do(func() {
		file_ims_proto_rawDescData = protoimpl.X.CompressGZIP(file_ims_proto_rawDescData)
	})
	return file_ims_proto_rawDescData
}

var file_ims_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ims_proto_goTypes = []interface{}{
	(*Void)(nil),              // 0: ims.Void
	(*StatusReply)(nil),       // 1: ims.StatusReply
	(*StopReply)(nil),         // 2: ims.StopReply
	(*AssumeRoleRequest)(nil), // 3: ims.AssumeRoleRequest
	(*Profile)(nil),           // 4: ims.Profile
	(*ConfigReply)(nil),       // 5: ims.ConfigReply
	nil,                       // 6: ims.ConfigReply.ProfilesEntry
}
var file_ims_proto_depIdxs = []int32{
	6, // 0: ims.ConfigReply.profiles:type_name -> ims.ConfigReply.ProfilesEntry
	4, // 1: ims.ConfigReply.ProfilesEntry.value:type_name -> ims.Profile
	0, // 2: ims.InstanceMetaService.Status:input_type -> ims.Void
	0, // 3: ims.InstanceMetaService.Stop:input_type -> ims.Void
	3, // 4: ims.InstanceMetaService.AssumeRole:input_type -> ims.AssumeRoleRequest
	3, // 5: ims.InstanceMetaService.RetrieveRole:input_type -> ims.AssumeRoleRequest
	0, // 6: ims.InstanceMetaService.Config:input_type -> ims.Void
	1, // 7: ims.InstanceMetaService.Status:output_type -> ims.StatusReply
	2, // 8: ims.InstanceMetaService.Stop:output_type -> ims.StopReply
	1, // 9: ims.InstanceMetaService.AssumeRole:output_type -> ims.StatusReply
	1, // 10: ims.InstanceMetaService.RetrieveRole:output_type -> ims.StatusReply
	5, // 11: ims.InstanceMetaService.Config:output_type -> ims.ConfigReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ims_proto_init() }
func file_ims_proto_init() {
	if File_ims_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ims_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_ims_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_ims_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReply); i {
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
		file_ims_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssumeRoleRequest); i {
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
		file_ims_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_ims_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigReply); i {
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
			RawDescriptor: file_ims_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ims_proto_goTypes,
		DependencyIndexes: file_ims_proto_depIdxs,
		MessageInfos:      file_ims_proto_msgTypes,
	}.Build()
	File_ims_proto = out.File
	file_ims_proto_rawDesc = nil
	file_ims_proto_goTypes = nil
	file_ims_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InstanceMetaServiceClient is the client API for InstanceMetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstanceMetaServiceClient interface {
	Status(ctx context.Context, in *Void, opts ...grpc.CallOption) (*StatusReply, error)
	Stop(ctx context.Context, in *Void, opts ...grpc.CallOption) (*StopReply, error)
	AssumeRole(ctx context.Context, in *AssumeRoleRequest, opts ...grpc.CallOption) (*StatusReply, error)
	RetrieveRole(ctx context.Context, in *AssumeRoleRequest, opts ...grpc.CallOption) (*StatusReply, error)
	Config(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ConfigReply, error)
}

type instanceMetaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceMetaServiceClient(cc grpc.ClientConnInterface) InstanceMetaServiceClient {
	return &instanceMetaServiceClient{cc}
}

func (c *instanceMetaServiceClient) Status(ctx context.Context, in *Void, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/ims.InstanceMetaService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceMetaServiceClient) Stop(ctx context.Context, in *Void, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := c.cc.Invoke(ctx, "/ims.InstanceMetaService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceMetaServiceClient) AssumeRole(ctx context.Context, in *AssumeRoleRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/ims.InstanceMetaService/AssumeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceMetaServiceClient) RetrieveRole(ctx context.Context, in *AssumeRoleRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/ims.InstanceMetaService/RetrieveRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceMetaServiceClient) Config(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ConfigReply, error) {
	out := new(ConfigReply)
	err := c.cc.Invoke(ctx, "/ims.InstanceMetaService/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceMetaServiceServer is the server API for InstanceMetaService service.
type InstanceMetaServiceServer interface {
	Status(context.Context, *Void) (*StatusReply, error)
	Stop(context.Context, *Void) (*StopReply, error)
	AssumeRole(context.Context, *AssumeRoleRequest) (*StatusReply, error)
	RetrieveRole(context.Context, *AssumeRoleRequest) (*StatusReply, error)
	Config(context.Context, *Void) (*ConfigReply, error)
}

// UnimplementedInstanceMetaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInstanceMetaServiceServer struct {
}

func (*UnimplementedInstanceMetaServiceServer) Status(context.Context, *Void) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedInstanceMetaServiceServer) Stop(context.Context, *Void) (*StopReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedInstanceMetaServiceServer) AssumeRole(context.Context, *AssumeRoleRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssumeRole not implemented")
}
func (*UnimplementedInstanceMetaServiceServer) RetrieveRole(context.Context, *AssumeRoleRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveRole not implemented")
}
func (*UnimplementedInstanceMetaServiceServer) Config(context.Context, *Void) (*ConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}

func RegisterInstanceMetaServiceServer(s *grpc.Server, srv InstanceMetaServiceServer) {
	s.RegisterService(&_InstanceMetaService_serviceDesc, srv)
}

func _InstanceMetaService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceMetaServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ims.InstanceMetaService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceMetaServiceServer).Status(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceMetaService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceMetaServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ims.InstanceMetaService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceMetaServiceServer).Stop(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceMetaService_AssumeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssumeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceMetaServiceServer).AssumeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ims.InstanceMetaService/AssumeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceMetaServiceServer).AssumeRole(ctx, req.(*AssumeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceMetaService_RetrieveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssumeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceMetaServiceServer).RetrieveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ims.InstanceMetaService/RetrieveRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceMetaServiceServer).RetrieveRole(ctx, req.(*AssumeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceMetaService_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceMetaServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ims.InstanceMetaService/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceMetaServiceServer).Config(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _InstanceMetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ims.InstanceMetaService",
	HandlerType: (*InstanceMetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _InstanceMetaService_Status_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _InstanceMetaService_Stop_Handler,
		},
		{
			MethodName: "AssumeRole",
			Handler:    _InstanceMetaService_AssumeRole_Handler,
		},
		{
			MethodName: "RetrieveRole",
			Handler:    _InstanceMetaService_RetrieveRole_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _InstanceMetaService_Config_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ims.proto",
}
