// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/raw/authorization/authorization.proto

package authorization

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

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raw_authorization_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raw_authorization_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_proto_raw_authorization_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  bool       `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Err *UserError `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raw_authorization_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raw_authorization_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_raw_authorization_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *UserRequest) GetErr() *UserError {
	if x != nil {
		return x.Err
	}
	return nil
}

type UserError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Id  int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserError) Reset() {
	*x = UserError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_raw_authorization_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserError) ProtoMessage() {}

func (x *UserError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_raw_authorization_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserError.ProtoReflect.Descriptor instead.
func (*UserError) Descriptor() ([]byte, []int) {
	return file_proto_raw_authorization_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *UserError) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *UserError) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_raw_authorization_authorization_proto protoreflect.FileDescriptor

var file_proto_raw_authorization_authorization_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x77, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x2a, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x99, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_raw_authorization_authorization_proto_rawDescOnce sync.Once
	file_proto_raw_authorization_authorization_proto_rawDescData = file_proto_raw_authorization_authorization_proto_rawDesc
)

func file_proto_raw_authorization_authorization_proto_rawDescGZIP() []byte {
	file_proto_raw_authorization_authorization_proto_rawDescOnce.Do(func() {
		file_proto_raw_authorization_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_raw_authorization_authorization_proto_rawDescData)
	})
	return file_proto_raw_authorization_authorization_proto_rawDescData
}

var file_proto_raw_authorization_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_raw_authorization_authorization_proto_goTypes = []interface{}{
	(*UserData)(nil),    // 0: authorization.UserData
	(*UserRequest)(nil), // 1: authorization.UserRequest
	(*UserError)(nil),   // 2: authorization.UserError
}
var file_proto_raw_authorization_authorization_proto_depIdxs = []int32{
	2, // 0: authorization.UserRequest.err:type_name -> authorization.UserError
	0, // 1: authorization.UserService.IdentifyUser:input_type -> authorization.UserData
	0, // 2: authorization.UserService.CreateUser:input_type -> authorization.UserData
	1, // 3: authorization.UserService.IdentifyUser:output_type -> authorization.UserRequest
	1, // 4: authorization.UserService.CreateUser:output_type -> authorization.UserRequest
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_raw_authorization_authorization_proto_init() }
func file_proto_raw_authorization_authorization_proto_init() {
	if File_proto_raw_authorization_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_raw_authorization_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_proto_raw_authorization_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_proto_raw_authorization_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserError); i {
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
			RawDescriptor: file_proto_raw_authorization_authorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_raw_authorization_authorization_proto_goTypes,
		DependencyIndexes: file_proto_raw_authorization_authorization_proto_depIdxs,
		MessageInfos:      file_proto_raw_authorization_authorization_proto_msgTypes,
	}.Build()
	File_proto_raw_authorization_authorization_proto = out.File
	file_proto_raw_authorization_authorization_proto_rawDesc = nil
	file_proto_raw_authorization_authorization_proto_goTypes = nil
	file_proto_raw_authorization_authorization_proto_depIdxs = nil
}
