// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api/protos/v1/user/user.proto

package user

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt string `protobuf:"bytes,4,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	UserId    string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	UserType  string `protobuf:"bytes,6,opt,name=userType,proto3" json:"userType,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *User) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt  string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  string `protobuf:"bytes,3,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt  string `protobuf:"bytes,4,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	User       *User  `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Name       string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Birth      string `protobuf:"bytes,7,opt,name=birth,proto3" json:"birth,omitempty"`
	Gender     string `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	ProfileImg string `protobuf:"bytes,9,opt,name=profileImg,proto3" json:"profileImg,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_user_user_proto_msgTypes[1]
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
	return file_api_protos_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Profile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Profile) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Profile) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *Profile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Profile) GetProfileImg() string {
	if x != nil {
		return x.ProfileImg
	}
	return ""
}

type GetMyProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetMyProfileReq) Reset() {
	*x = GetMyProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyProfileReq) ProtoMessage() {}

func (x *GetMyProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyProfileReq.ProtoReflect.Descriptor instead.
func (*GetMyProfileReq) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyProfileReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetMyProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RsltCd    string   `protobuf:"bytes,1,opt,name=rsltCd,proto3" json:"rsltCd,omitempty"`
	RsltMsg   string   `protobuf:"bytes,2,opt,name=rsltMsg,proto3" json:"rsltMsg,omitempty"`
	MyProfile *Profile `protobuf:"bytes,3,opt,name=myProfile,proto3" json:"myProfile,omitempty"`
}

func (x *GetMyProfileRes) Reset() {
	*x = GetMyProfileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protos_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyProfileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyProfileRes) ProtoMessage() {}

func (x *GetMyProfileRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_protos_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyProfileRes.ProtoReflect.Descriptor instead.
func (*GetMyProfileRes) Descriptor() ([]byte, []int) {
	return file_api_protos_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetMyProfileRes) GetRsltCd() string {
	if x != nil {
		return x.RsltCd
	}
	return ""
}

func (x *GetMyProfileRes) GetRsltMsg() string {
	if x != nil {
		return x.RsltMsg
	}
	return ""
}

func (x *GetMyProfileRes) GetMyProfile() *Profile {
	if x != nil {
		return x.MyProfile
	}
	return nil
}

var File_api_protos_v1_user_user_proto protoreflect.FileDescriptor

var file_api_protos_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x6d, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x6d, 0x67, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x79, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x6c, 0x74, 0x43, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x6c, 0x74, 0x43, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x73,
	0x6c, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x73, 0x6c,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x09, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x09, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x5f, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x6b, 0x2d, 0x4b,
	0x77, 0x6f, 0x6e, 0x73, 0x6f, 0x6f, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protos_v1_user_user_proto_rawDescOnce sync.Once
	file_api_protos_v1_user_user_proto_rawDescData = file_api_protos_v1_user_user_proto_rawDesc
)

func file_api_protos_v1_user_user_proto_rawDescGZIP() []byte {
	file_api_protos_v1_user_user_proto_rawDescOnce.Do(func() {
		file_api_protos_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protos_v1_user_user_proto_rawDescData)
	})
	return file_api_protos_v1_user_user_proto_rawDescData
}

var file_api_protos_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_protos_v1_user_user_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: v1.user_proto.User
	(*Profile)(nil),         // 1: v1.user_proto.Profile
	(*GetMyProfileReq)(nil), // 2: v1.user_proto.GetMyProfileReq
	(*GetMyProfileRes)(nil), // 3: v1.user_proto.GetMyProfileRes
}
var file_api_protos_v1_user_user_proto_depIdxs = []int32{
	0, // 0: v1.user_proto.Profile.user:type_name -> v1.user_proto.User
	1, // 1: v1.user_proto.GetMyProfileRes.myProfile:type_name -> v1.user_proto.Profile
	2, // 2: v1.user_proto.UserService.GetMyProfile:input_type -> v1.user_proto.GetMyProfileReq
	3, // 3: v1.user_proto.UserService.GetMyProfile:output_type -> v1.user_proto.GetMyProfileRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_protos_v1_user_user_proto_init() }
func file_api_protos_v1_user_user_proto_init() {
	if File_api_protos_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protos_v1_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_protos_v1_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_protos_v1_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyProfileReq); i {
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
		file_api_protos_v1_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyProfileRes); i {
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
			RawDescriptor: file_api_protos_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protos_v1_user_user_proto_goTypes,
		DependencyIndexes: file_api_protos_v1_user_user_proto_depIdxs,
		MessageInfos:      file_api_protos_v1_user_user_proto_msgTypes,
	}.Build()
	File_api_protos_v1_user_user_proto = out.File
	file_api_protos_v1_user_user_proto_rawDesc = nil
	file_api_protos_v1_user_user_proto_goTypes = nil
	file_api_protos_v1_user_user_proto_depIdxs = nil
}