// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/protos/v1/auth/login.proto

package auth

import (
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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginType string `protobuf:"bytes,1,opt,name=loginType,proto3" json:"loginType,omitempty"`
	MemId     string `protobuf:"bytes,2,opt,name=memId,proto3" json:"memId,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetLoginType() string {
	if x != nil {
		return x.LoginType
	}
	return ""
}

func (x *LoginReq) GetMemId() string {
	if x != nil {
		return x.MemId
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RsltCd  string `protobuf:"bytes,1,opt,name=rsltCd,proto3" json:"rsltCd,omitempty"`
	RsltMsg string `protobuf:"bytes,2,opt,name=rsltMsg,proto3" json:"rsltMsg,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetRsltCd() string {
	if x != nil {
		return x.RsltCd
	}
	return ""
}

func (x *LoginRes) GetRsltMsg() string {
	if x != nil {
		return x.RsltMsg
	}
	return ""
}

func (x *LoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PasswordCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword string `protobuf:"bytes,1,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
}

func (x *PasswordCheckReq) Reset() {
	*x = PasswordCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCheckReq) ProtoMessage() {}

func (x *PasswordCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordCheckReq.ProtoReflect.Descriptor instead.
func (*PasswordCheckReq) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{2}
}

func (x *PasswordCheckReq) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

type PasswordCheckRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RsltCd    string `protobuf:"bytes,1,opt,name=rsltCd,proto3" json:"rsltCd,omitempty"`
	RsltMsg   string `protobuf:"bytes,2,opt,name=rsltMsg,proto3" json:"rsltMsg,omitempty"`
	IsChecked bool   `protobuf:"varint,3,opt,name=isChecked,proto3" json:"isChecked,omitempty"`
}

func (x *PasswordCheckRes) Reset() {
	*x = PasswordCheckRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordCheckRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCheckRes) ProtoMessage() {}

func (x *PasswordCheckRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordCheckRes.ProtoReflect.Descriptor instead.
func (*PasswordCheckRes) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{3}
}

func (x *PasswordCheckRes) GetRsltCd() string {
	if x != nil {
		return x.RsltCd
	}
	return ""
}

func (x *PasswordCheckRes) GetRsltMsg() string {
	if x != nil {
		return x.RsltMsg
	}
	return ""
}

func (x *PasswordCheckRes) GetIsChecked() bool {
	if x != nil {
		return x.IsChecked
	}
	return false
}

type PasswordChangeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassword string `protobuf:"bytes,1,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *PasswordChangeReq) Reset() {
	*x = PasswordChangeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordChangeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordChangeReq) ProtoMessage() {}

func (x *PasswordChangeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordChangeReq.ProtoReflect.Descriptor instead.
func (*PasswordChangeReq) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{4}
}

func (x *PasswordChangeReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type PasswordChangeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RsltCd  string `protobuf:"bytes,1,opt,name=rsltCd,proto3" json:"rsltCd,omitempty"`
	RsltMsg string `protobuf:"bytes,2,opt,name=rsltMsg,proto3" json:"rsltMsg,omitempty"`
}

func (x *PasswordChangeRes) Reset() {
	*x = PasswordChangeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_auth_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordChangeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordChangeRes) ProtoMessage() {}

func (x *PasswordChangeRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_auth_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordChangeRes.ProtoReflect.Descriptor instead.
func (*PasswordChangeRes) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_auth_login_proto_rawDescGZIP(), []int{5}
}

func (x *PasswordChangeRes) GetRsltCd() string {
	if x != nil {
		return x.RsltCd
	}
	return ""
}

func (x *PasswordChangeRes) GetRsltMsg() string {
	if x != nil {
		return x.RsltMsg
	}
	return ""
}

var File_api_protos_v1_auth_login_proto protoreflect.FileDescriptor

var file_api_protos_v1_auth_login_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65,
	0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x52, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x6c, 0x74,
	0x43, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x6c, 0x74, 0x43, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x34, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x62, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73,
	0x6c, 0x74, 0x43, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x6c, 0x74,
	0x43, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x35, 0x0a, 0x11, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x45, 0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x6c, 0x74, 0x43, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x6c, 0x74, 0x43, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x73, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x32, 0xfe, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x76, 0x31,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x0e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x6b, 0x2d, 0x4b, 0x77, 0x6f,
	0x6e, 0x73, 0x6f, 0x6f, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_v1_auth_login_proto_rawDescOnce sync.Once
	file_api_protos_v1_auth_login_proto_rawDescData = file_api_protos_v1_auth_login_proto_rawDesc
)

func file_api_protos_v1_auth_login_proto_rawDescGZIP() []byte {
	file_api_protos_v1_auth_login_proto_rawDescOnce.Do(func() {
		file_api_protos_v1_auth_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_v1_auth_login_proto_rawDescData)
	})
	return file_api_protos_v1_auth_login_proto_rawDescData
}

var file_api_protos_v1_auth_login_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_protos_v1_auth_login_proto_goTypes = []interface{}{
	(*LoginReq)(nil),          // 0: v1.login_proto.LoginReq
	(*LoginRes)(nil),          // 1: v1.login_proto.LoginRes
	(*PasswordCheckReq)(nil),  // 2: v1.login_proto.PasswordCheckReq
	(*PasswordCheckRes)(nil),  // 3: v1.login_proto.PasswordCheckRes
	(*PasswordChangeReq)(nil), // 4: v1.login_proto.PasswordChangeReq
	(*PasswordChangeRes)(nil), // 5: v1.login_proto.PasswordChangeRes
}
var file_api_protos_v1_auth_login_proto_depIdxs = []int32{
	0, // 0: v1.login_proto.LoginService.Login:input_type -> v1.login_proto.LoginReq
	2, // 1: v1.login_proto.LoginService.PasswordCheck:input_type -> v1.login_proto.PasswordCheckReq
	4, // 2: v1.login_proto.LoginService.PasswordChange:input_type -> v1.login_proto.PasswordChangeReq
	1, // 3: v1.login_proto.LoginService.Login:output_type -> v1.login_proto.LoginRes
	3, // 4: v1.login_proto.LoginService.PasswordCheck:output_type -> v1.login_proto.PasswordCheckRes
	5, // 5: v1.login_proto.LoginService.PasswordChange:output_type -> v1.login_proto.PasswordChangeRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protos_v1_auth_login_proto_init() }
func file_api_protos_v1_auth_login_proto_init() {
	if File_api_protos_v1_auth_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_v1_auth_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_protos_v1_auth_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_api_protos_v1_auth_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordCheckReq); i {
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
		file_api_protos_v1_auth_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordCheckRes); i {
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
		file_api_protos_v1_auth_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordChangeReq); i {
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
		file_api_protos_v1_auth_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordChangeRes); i {
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
			RawDescriptor: file_api_protos_v1_auth_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protos_v1_auth_login_proto_goTypes,
		DependencyIndexes: file_api_protos_v1_auth_login_proto_depIdxs,
		MessageInfos:      file_api_protos_v1_auth_login_proto_msgTypes,
	}.Build()
	File_api_protos_v1_auth_login_proto = out.File
	file_api_protos_v1_auth_login_proto_rawDesc = nil
	file_api_protos_v1_auth_login_proto_goTypes = nil
	file_api_protos_v1_auth_login_proto_depIdxs = nil
}
