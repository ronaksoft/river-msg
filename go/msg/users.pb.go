// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: users.proto

package msg

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

// UsersGet
// @Function
// @Returns: UsersMany
type UsersGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*InputUser `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *UsersGet) Reset() {
	*x = UsersGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersGet) ProtoMessage() {}

func (x *UsersGet) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersGet.ProtoReflect.Descriptor instead.
func (*UsersGet) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *UsersGet) GetUsers() []*InputUser {
	if x != nil {
		return x.Users
	}
	return nil
}

// UsersGetFull
// @Function
// @Returns: UsersFullMany
type UsersGetFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*InputUser `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *UsersGetFull) Reset() {
	*x = UsersGetFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersGetFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersGetFull) ProtoMessage() {}

func (x *UsersGetFull) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersGetFull.ProtoReflect.Descriptor instead.
func (*UsersGetFull) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *UsersGetFull) GetUsers() []*InputUser {
	if x != nil {
		return x.Users
	}
	return nil
}

// UsersMany
type UsersMany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	Empty bool    `protobuf:"varint,5,opt,name=Empty,proto3" json:"Empty,omitempty"`
}

func (x *UsersMany) Reset() {
	*x = UsersMany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersMany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersMany) ProtoMessage() {}

func (x *UsersMany) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersMany.ProtoReflect.Descriptor instead.
func (*UsersMany) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *UsersMany) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UsersMany) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74,
	0x12, 0x24, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47,
	0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x24, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_users_proto_goTypes = []interface{}{
	(*UsersGet)(nil),     // 0: msg.UsersGet
	(*UsersGetFull)(nil), // 1: msg.UsersGetFull
	(*UsersMany)(nil),    // 2: msg.UsersMany
	(*InputUser)(nil),    // 3: msg.InputUser
	(*User)(nil),         // 4: msg.User
}
var file_users_proto_depIdxs = []int32{
	3, // 0: msg.UsersGet.Users:type_name -> msg.InputUser
	3, // 1: msg.UsersGetFull.Users:type_name -> msg.InputUser
	4, // 2: msg.UsersMany.Users:type_name -> msg.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	file_core_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersGet); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersGetFull); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersMany); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}
