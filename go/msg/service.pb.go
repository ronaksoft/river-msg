// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: service.proto

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

type ServiceSendMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnBehalf   int64            `protobuf:"varint,100,opt,name=OnBehalf,proto3" json:"OnBehalf,omitempty"`
	RandomID   int64            `protobuf:"varint,1,opt,name=RandomID,proto3" json:"RandomID,omitempty"`
	Peer       *InputPeer       `protobuf:"bytes,2,opt,name=Peer,proto3" json:"Peer,omitempty"`
	Body       string           `protobuf:"bytes,5,opt,name=Body,proto3" json:"Body,omitempty"`
	ReplyTo    int64            `protobuf:"varint,6,opt,name=ReplyTo,proto3" json:"ReplyTo,omitempty"`
	ClearDraft bool             `protobuf:"varint,7,opt,name=ClearDraft,proto3" json:"ClearDraft,omitempty"`
	Entities   []*MessageEntity `protobuf:"bytes,8,rep,name=Entities,proto3" json:"Entities,omitempty"`
}

func (x *ServiceSendMessage) Reset() {
	*x = ServiceSendMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSendMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSendMessage) ProtoMessage() {}

func (x *ServiceSendMessage) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSendMessage.ProtoReflect.Descriptor instead.
func (*ServiceSendMessage) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceSendMessage) GetOnBehalf() int64 {
	if x != nil {
		return x.OnBehalf
	}
	return 0
}

func (x *ServiceSendMessage) GetRandomID() int64 {
	if x != nil {
		return x.RandomID
	}
	return 0
}

func (x *ServiceSendMessage) GetPeer() *InputPeer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *ServiceSendMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ServiceSendMessage) GetReplyTo() int64 {
	if x != nil {
		return x.ReplyTo
	}
	return 0
}

func (x *ServiceSendMessage) GetClearDraft() bool {
	if x != nil {
		return x.ClearDraft
	}
	return false
}

func (x *ServiceSendMessage) GetEntities() []*MessageEntity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4f, 0x6e, 0x42, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x4f, 0x6e, 0x42, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x52, 0x61, 0x6e,
	0x64, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x44, 0x72, 0x61, 0x66, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x43, 0x6c, 0x65,
	0x61, 0x72, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_service_proto_goTypes = []interface{}{
	(*ServiceSendMessage)(nil), // 0: msg.ServiceSendMessage
	(*InputPeer)(nil),          // 1: msg.InputPeer
	(*MessageEntity)(nil),      // 2: msg.MessageEntity
}
var file_service_proto_depIdxs = []int32{
	1, // 0: msg.ServiceSendMessage.Peer:type_name -> msg.InputPeer
	2, // 1: msg.ServiceSendMessage.Entities:type_name -> msg.MessageEntity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_core_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSendMessage); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
