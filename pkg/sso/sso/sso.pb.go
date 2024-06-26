// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.14.0
// source: sso/sso.proto

package sso

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

type DataOwners int32

const (
	DataOwners_TEST DataOwners = 0
	DataOwners_PP   DataOwners = 1
)

// Enum value maps for DataOwners.
var (
	DataOwners_name = map[int32]string{
		0: "TEST",
		1: "PP",
	}
	DataOwners_value = map[string]int32{
		"TEST": 0,
		"PP":   1,
	}
)

func (x DataOwners) Enum() *DataOwners {
	p := new(DataOwners)
	*p = x
	return p
}

func (x DataOwners) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataOwners) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_sso_proto_enumTypes[0].Descriptor()
}

func (DataOwners) Type() protoreflect.EnumType {
	return &file_sso_sso_proto_enumTypes[0]
}

func (x DataOwners) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataOwners.Descriptor instead.
func (DataOwners) EnumDescriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

type GoogleTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GoogleTokenRequest) Reset() {
	*x = GoogleTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleTokenRequest) ProtoMessage() {}

func (x *GoogleTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleTokenRequest.ProtoReflect.Descriptor instead.
func (*GoogleTokenRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GoogleTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	UserInfo  *UserInfo   `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	DataOwner *DataOwners `protobuf:"varint,3,opt,name=dataOwner,proto3,enum=auth.DataOwners,oneof" json:"dataOwner,omitempty"`
}

func (x *GoogleTokenResponse) Reset() {
	*x = GoogleTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleTokenResponse) ProtoMessage() {}

func (x *GoogleTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleTokenResponse.ProtoReflect.Descriptor instead.
func (*GoogleTokenResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleTokenResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GoogleTokenResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *GoogleTokenResponse) GetDataOwner() DataOwners {
	if x != nil && x.DataOwner != nil {
		return *x.DataOwner
	}
	return DataOwners_TEST
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_sso_sso_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

var File_sso_sso_proto protoreflect.FileDescriptor

var file_sso_sso_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x33, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x6a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x1e,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x50, 0x10, 0x01, 0x32, 0x52,
	0x0a, 0x03, 0x53, 0x73, 0x6f, 0x12, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x78, 0x6f, 0x70, 0x65, 0x2d, 0x69, 0x64, 0x2e, 0x73,
	0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_sso_proto_rawDescOnce sync.Once
	file_sso_sso_proto_rawDescData = file_sso_sso_proto_rawDesc
)

func file_sso_sso_proto_rawDescGZIP() []byte {
	file_sso_sso_proto_rawDescOnce.Do(func() {
		file_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_sso_proto_rawDescData)
	})
	return file_sso_sso_proto_rawDescData
}

var file_sso_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sso_sso_proto_goTypes = []interface{}{
	(DataOwners)(0),             // 0: auth.DataOwners
	(*GoogleTokenRequest)(nil),  // 1: auth.GoogleTokenRequest
	(*GoogleTokenResponse)(nil), // 2: auth.GoogleTokenResponse
	(*UserInfo)(nil),            // 3: auth.UserInfo
}
var file_sso_sso_proto_depIdxs = []int32{
	3, // 0: auth.GoogleTokenResponse.userInfo:type_name -> auth.UserInfo
	0, // 1: auth.GoogleTokenResponse.dataOwner:type_name -> auth.DataOwners
	1, // 2: auth.Sso.GetInfoByGoogleToken:input_type -> auth.GoogleTokenRequest
	2, // 3: auth.Sso.GetInfoByGoogleToken:output_type -> auth.GoogleTokenResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sso_sso_proto_init() }
func file_sso_sso_proto_init() {
	if File_sso_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleTokenRequest); i {
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
		file_sso_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleTokenResponse); i {
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
		file_sso_sso_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
	file_sso_sso_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sso_sso_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_sso_proto_goTypes,
		DependencyIndexes: file_sso_sso_proto_depIdxs,
		EnumInfos:         file_sso_sso_proto_enumTypes,
		MessageInfos:      file_sso_sso_proto_msgTypes,
	}.Build()
	File_sso_sso_proto = out.File
	file_sso_sso_proto_rawDesc = nil
	file_sso_sso_proto_goTypes = nil
	file_sso_sso_proto_depIdxs = nil
}
