// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: client.updates.proto

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

// ClientPendingMessageDelivery
type ClientUpdatePendingMessageDelivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages       *UserMessage          `protobuf:"bytes,1,opt,name=Messages,proto3" json:"Messages,omitempty"`
	PendingMessage *ClientPendingMessage `protobuf:"bytes,2,opt,name=PendingMessage,proto3" json:"PendingMessage,omitempty"`
	Success        bool                  `protobuf:"varint,3,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *ClientUpdatePendingMessageDelivery) Reset() {
	*x = ClientUpdatePendingMessageDelivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_updates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientUpdatePendingMessageDelivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientUpdatePendingMessageDelivery) ProtoMessage() {}

func (x *ClientUpdatePendingMessageDelivery) ProtoReflect() protoreflect.Message {
	mi := &file_client_updates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientUpdatePendingMessageDelivery.ProtoReflect.Descriptor instead.
func (*ClientUpdatePendingMessageDelivery) Descriptor() ([]byte, []int) {
	return file_client_updates_proto_rawDescGZIP(), []int{0}
}

func (x *ClientUpdatePendingMessageDelivery) GetMessages() *UserMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ClientUpdatePendingMessageDelivery) GetPendingMessage() *ClientPendingMessage {
	if x != nil {
		return x.PendingMessage
	}
	return nil
}

func (x *ClientUpdatePendingMessageDelivery) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ClientUpdateMessageDeleted
type ClientUpdateMessagesDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerID     int64   `protobuf:"varint,1,opt,name=PeerID,proto3" json:"PeerID,omitempty"`
	PeerType   int32   `protobuf:"varint,2,opt,name=PeerType,proto3" json:"PeerType,omitempty"`
	MessageIDs []int64 `protobuf:"varint,3,rep,packed,name=MessageIDs,proto3" json:"MessageIDs,omitempty"`
}

func (x *ClientUpdateMessagesDeleted) Reset() {
	*x = ClientUpdateMessagesDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_updates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientUpdateMessagesDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientUpdateMessagesDeleted) ProtoMessage() {}

func (x *ClientUpdateMessagesDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_client_updates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientUpdateMessagesDeleted.ProtoReflect.Descriptor instead.
func (*ClientUpdateMessagesDeleted) Descriptor() ([]byte, []int) {
	return file_client_updates_proto_rawDescGZIP(), []int{1}
}

func (x *ClientUpdateMessagesDeleted) GetPeerID() int64 {
	if x != nil {
		return x.PeerID
	}
	return 0
}

func (x *ClientUpdateMessagesDeleted) GetPeerType() int32 {
	if x != nil {
		return x.PeerType
	}
	return 0
}

func (x *ClientUpdateMessagesDeleted) GetMessageIDs() []int64 {
	if x != nil {
		return x.MessageIDs
	}
	return nil
}

var File_client_updates_proto protoreflect.FileDescriptor

var file_client_updates_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x10, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x22,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x41, 0x0a, 0x0e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x71, 0x0a,
	0x1b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x73,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_client_updates_proto_rawDescOnce sync.Once
	file_client_updates_proto_rawDescData = file_client_updates_proto_rawDesc
)

func file_client_updates_proto_rawDescGZIP() []byte {
	file_client_updates_proto_rawDescOnce.Do(func() {
		file_client_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_updates_proto_rawDescData)
	})
	return file_client_updates_proto_rawDescData
}

var file_client_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_updates_proto_goTypes = []interface{}{
	(*ClientUpdatePendingMessageDelivery)(nil), // 0: msg.ClientUpdatePendingMessageDelivery
	(*ClientUpdateMessagesDeleted)(nil),        // 1: msg.ClientUpdateMessagesDeleted
	(*UserMessage)(nil),                        // 2: msg.UserMessage
	(*ClientPendingMessage)(nil),               // 3: msg.ClientPendingMessage
}
var file_client_updates_proto_depIdxs = []int32{
	2, // 0: msg.ClientUpdatePendingMessageDelivery.Messages:type_name -> msg.UserMessage
	3, // 1: msg.ClientUpdatePendingMessageDelivery.PendingMessage:type_name -> msg.ClientPendingMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_updates_proto_init() }
func file_client_updates_proto_init() {
	if File_client_updates_proto != nil {
		return
	}
	file_core_types_proto_init()
	file_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_client_updates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientUpdatePendingMessageDelivery); i {
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
		file_client_updates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientUpdateMessagesDeleted); i {
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
			RawDescriptor: file_client_updates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_updates_proto_goTypes,
		DependencyIndexes: file_client_updates_proto_depIdxs,
		MessageInfos:      file_client_updates_proto_msgTypes,
	}.Build()
	File_client_updates_proto = out.File
	file_client_updates_proto_rawDesc = nil
	file_client_updates_proto_goTypes = nil
	file_client_updates_proto_depIdxs = nil
}
