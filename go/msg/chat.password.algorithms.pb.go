// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: chat.password.algorithms.proto

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

type PasswordAlgorithmVer6A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salt1 []byte `protobuf:"bytes,1,opt,name=Salt1,proto3" json:"Salt1,omitempty"`
	Salt2 []byte `protobuf:"bytes,2,opt,name=Salt2,proto3" json:"Salt2,omitempty"`
	G     int32  `protobuf:"varint,3,opt,name=G,proto3" json:"G,omitempty"`
	P     []byte `protobuf:"bytes,4,opt,name=P,proto3" json:"P,omitempty"`
}

func (x *PasswordAlgorithmVer6A) Reset() {
	*x = PasswordAlgorithmVer6A{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_password_algorithms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordAlgorithmVer6A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordAlgorithmVer6A) ProtoMessage() {}

func (x *PasswordAlgorithmVer6A) ProtoReflect() protoreflect.Message {
	mi := &file_chat_password_algorithms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordAlgorithmVer6A.ProtoReflect.Descriptor instead.
func (*PasswordAlgorithmVer6A) Descriptor() ([]byte, []int) {
	return file_chat_password_algorithms_proto_rawDescGZIP(), []int{0}
}

func (x *PasswordAlgorithmVer6A) GetSalt1() []byte {
	if x != nil {
		return x.Salt1
	}
	return nil
}

func (x *PasswordAlgorithmVer6A) GetSalt2() []byte {
	if x != nil {
		return x.Salt2
	}
	return nil
}

func (x *PasswordAlgorithmVer6A) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *PasswordAlgorithmVer6A) GetP() []byte {
	if x != nil {
		return x.P
	}
	return nil
}

var File_chat_password_algorithms_proto protoreflect.FileDescriptor

var file_chat_password_algorithms_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x60, 0x0a, 0x16, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x56, 0x65, 0x72, 0x36, 0x41, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x61, 0x6c, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x53, 0x61, 0x6c, 0x74, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x61, 0x6c, 0x74, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x53, 0x61, 0x6c, 0x74, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x47,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x47, 0x12, 0x0c, 0x0a, 0x01, 0x50, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x50, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_password_algorithms_proto_rawDescOnce sync.Once
	file_chat_password_algorithms_proto_rawDescData = file_chat_password_algorithms_proto_rawDesc
)

func file_chat_password_algorithms_proto_rawDescGZIP() []byte {
	file_chat_password_algorithms_proto_rawDescOnce.Do(func() {
		file_chat_password_algorithms_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_password_algorithms_proto_rawDescData)
	})
	return file_chat_password_algorithms_proto_rawDescData
}

var file_chat_password_algorithms_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chat_password_algorithms_proto_goTypes = []interface{}{
	(*PasswordAlgorithmVer6A)(nil), // 0: msg.PasswordAlgorithmVer6A
}
var file_chat_password_algorithms_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_password_algorithms_proto_init() }
func file_chat_password_algorithms_proto_init() {
	if File_chat_password_algorithms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_password_algorithms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordAlgorithmVer6A); i {
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
			RawDescriptor: file_chat_password_algorithms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_password_algorithms_proto_goTypes,
		DependencyIndexes: file_chat_password_algorithms_proto_depIdxs,
		MessageInfos:      file_chat_password_algorithms_proto_msgTypes,
	}.Build()
	File_chat_password_algorithms_proto = out.File
	file_chat_password_algorithms_proto_rawDesc = nil
	file_chat_password_algorithms_proto_goTypes = nil
	file_chat_password_algorithms_proto_depIdxs = nil
}
