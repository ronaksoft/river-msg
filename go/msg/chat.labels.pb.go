// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: chat.labels.proto

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

// LabelsCreate
// @Function
// @Returns Bool
type LabelsCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RandomID int64  `protobuf:"varint,1,opt,name=RandomID,proto3" json:"RandomID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Colour   string `protobuf:"bytes,3,opt,name=Colour,proto3" json:"Colour,omitempty"`
}

func (x *LabelsCreate) Reset() {
	*x = LabelsCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsCreate) ProtoMessage() {}

func (x *LabelsCreate) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsCreate.ProtoReflect.Descriptor instead.
func (*LabelsCreate) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{0}
}

func (x *LabelsCreate) GetRandomID() int64 {
	if x != nil {
		return x.RandomID
	}
	return 0
}

func (x *LabelsCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelsCreate) GetColour() string {
	if x != nil {
		return x.Colour
	}
	return ""
}

// LabelsEdit
// @Function
// @Returns: Bool
type LabelsEdit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelID int32  `protobuf:"varint,1,opt,name=LabelID,proto3" json:"LabelID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Colour  string `protobuf:"bytes,3,opt,name=Colour,proto3" json:"Colour,omitempty"`
}

func (x *LabelsEdit) Reset() {
	*x = LabelsEdit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsEdit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsEdit) ProtoMessage() {}

func (x *LabelsEdit) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsEdit.ProtoReflect.Descriptor instead.
func (*LabelsEdit) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{1}
}

func (x *LabelsEdit) GetLabelID() int32 {
	if x != nil {
		return x.LabelID
	}
	return 0
}

func (x *LabelsEdit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelsEdit) GetColour() string {
	if x != nil {
		return x.Colour
	}
	return ""
}

// LabelsDelete
// @Function
// @Returns: Bool
type LabelsDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelIDs []int32 `protobuf:"varint,1,rep,packed,name=LabelIDs,proto3" json:"LabelIDs,omitempty"`
}

func (x *LabelsDelete) Reset() {
	*x = LabelsDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsDelete) ProtoMessage() {}

func (x *LabelsDelete) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsDelete.ProtoReflect.Descriptor instead.
func (*LabelsDelete) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{2}
}

func (x *LabelsDelete) GetLabelIDs() []int32 {
	if x != nil {
		return x.LabelIDs
	}
	return nil
}

// LabelsGet
// @Function
// @Returns: LabelsMany
type LabelsGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LabelsGet) Reset() {
	*x = LabelsGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsGet) ProtoMessage() {}

func (x *LabelsGet) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsGet.ProtoReflect.Descriptor instead.
func (*LabelsGet) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{3}
}

// LabelsAddToMessage
// @Function
// @Returns: Bool
type LabelsAddToMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer       *InputPeer `protobuf:"bytes,1,opt,name=Peer,proto3" json:"Peer,omitempty"`
	LabelIDs   []int32    `protobuf:"varint,3,rep,packed,name=LabelIDs,proto3" json:"LabelIDs,omitempty"`
	MessageIDs []int64    `protobuf:"varint,4,rep,packed,name=MessageIDs,proto3" json:"MessageIDs,omitempty"`
}

func (x *LabelsAddToMessage) Reset() {
	*x = LabelsAddToMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsAddToMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsAddToMessage) ProtoMessage() {}

func (x *LabelsAddToMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsAddToMessage.ProtoReflect.Descriptor instead.
func (*LabelsAddToMessage) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{4}
}

func (x *LabelsAddToMessage) GetPeer() *InputPeer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *LabelsAddToMessage) GetLabelIDs() []int32 {
	if x != nil {
		return x.LabelIDs
	}
	return nil
}

func (x *LabelsAddToMessage) GetMessageIDs() []int64 {
	if x != nil {
		return x.MessageIDs
	}
	return nil
}

// LabelsRemoveFromMessage
// @Function
// @Returns: Bool
type LabelsRemoveFromMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer       *InputPeer `protobuf:"bytes,1,opt,name=Peer,proto3" json:"Peer,omitempty"`
	LabelIDs   []int32    `protobuf:"varint,3,rep,packed,name=LabelIDs,proto3" json:"LabelIDs,omitempty"`
	MessageIDs []int64    `protobuf:"varint,4,rep,packed,name=MessageIDs,proto3" json:"MessageIDs,omitempty"`
}

func (x *LabelsRemoveFromMessage) Reset() {
	*x = LabelsRemoveFromMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsRemoveFromMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsRemoveFromMessage) ProtoMessage() {}

func (x *LabelsRemoveFromMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsRemoveFromMessage.ProtoReflect.Descriptor instead.
func (*LabelsRemoveFromMessage) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{5}
}

func (x *LabelsRemoveFromMessage) GetPeer() *InputPeer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *LabelsRemoveFromMessage) GetLabelIDs() []int32 {
	if x != nil {
		return x.LabelIDs
	}
	return nil
}

func (x *LabelsRemoveFromMessage) GetMessageIDs() []int64 {
	if x != nil {
		return x.MessageIDs
	}
	return nil
}

// LabelsListItems
// @Function
// @Returns: LabelItems
type LabelsListItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelID int32 `protobuf:"varint,1,opt,name=LabelID,proto3" json:"LabelID,omitempty"`
	MinID   int64 `protobuf:"varint,2,opt,name=MinID,proto3" json:"MinID,omitempty"`
	MaxID   int64 `protobuf:"varint,3,opt,name=MaxID,proto3" json:"MaxID,omitempty"`
	Limit   int32 `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *LabelsListItems) Reset() {
	*x = LabelsListItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelsListItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelsListItems) ProtoMessage() {}

func (x *LabelsListItems) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelsListItems.ProtoReflect.Descriptor instead.
func (*LabelsListItems) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{6}
}

func (x *LabelsListItems) GetLabelID() int32 {
	if x != nil {
		return x.LabelID
	}
	return 0
}

func (x *LabelsListItems) GetMinID() int64 {
	if x != nil {
		return x.MinID
	}
	return 0
}

func (x *LabelsListItems) GetMaxID() int64 {
	if x != nil {
		return x.MaxID
	}
	return 0
}

func (x *LabelsListItems) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// LabelItems
type LabelItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LabelID  int32          `protobuf:"varint,1,opt,name=LabelID,proto3" json:"LabelID,omitempty"`
	Messages []*UserMessage `protobuf:"bytes,2,rep,name=Messages,proto3" json:"Messages,omitempty"`
	Dialogs  []*Dialog      `protobuf:"bytes,3,rep,name=Dialogs,proto3" json:"Dialogs,omitempty"`
	Users    []*User        `protobuf:"bytes,4,rep,name=Users,proto3" json:"Users,omitempty"`
	Groups   []*Group       `protobuf:"bytes,5,rep,name=Groups,proto3" json:"Groups,omitempty"`
}

func (x *LabelItems) Reset() {
	*x = LabelItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_labels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelItems) ProtoMessage() {}

func (x *LabelItems) ProtoReflect() protoreflect.Message {
	mi := &file_chat_labels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelItems.ProtoReflect.Descriptor instead.
func (*LabelItems) Descriptor() ([]byte, []int) {
	return file_chat_labels_proto_rawDescGZIP(), []int{7}
}

func (x *LabelItems) GetLabelID() int32 {
	if x != nil {
		return x.LabelID
	}
	return 0
}

func (x *LabelItems) GetMessages() []*UserMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *LabelItems) GetDialogs() []*Dialog {
	if x != nil {
		return x.Dialogs
	}
	return nil
}

func (x *LabelItems) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *LabelItems) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_chat_labels_proto protoreflect.FileDescriptor

var file_chat_labels_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x52, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f,
	0x6c, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6c, 0x6f,
	0x75, 0x72, 0x22, 0x52, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x64, 0x69, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x22, 0x2a, 0x0a, 0x0c, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49,
	0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49,
	0x44, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x47, 0x65, 0x74, 0x22,
	0x74, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x49, 0x44, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x73, 0x22, 0x79, 0x0a, 0x17, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04,
	0x50, 0x65, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x73,
	0x22, 0x6d, 0x0a, 0x0f, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x4d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x69,
	0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x78, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4d, 0x61, 0x78, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0xc0, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x07, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x22,
	0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_labels_proto_rawDescOnce sync.Once
	file_chat_labels_proto_rawDescData = file_chat_labels_proto_rawDesc
)

func file_chat_labels_proto_rawDescGZIP() []byte {
	file_chat_labels_proto_rawDescOnce.Do(func() {
		file_chat_labels_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_labels_proto_rawDescData)
	})
	return file_chat_labels_proto_rawDescData
}

var file_chat_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chat_labels_proto_goTypes = []interface{}{
	(*LabelsCreate)(nil),            // 0: msg.LabelsCreate
	(*LabelsEdit)(nil),              // 1: msg.LabelsEdit
	(*LabelsDelete)(nil),            // 2: msg.LabelsDelete
	(*LabelsGet)(nil),               // 3: msg.LabelsGet
	(*LabelsAddToMessage)(nil),      // 4: msg.LabelsAddToMessage
	(*LabelsRemoveFromMessage)(nil), // 5: msg.LabelsRemoveFromMessage
	(*LabelsListItems)(nil),         // 6: msg.LabelsListItems
	(*LabelItems)(nil),              // 7: msg.LabelItems
	(*InputPeer)(nil),               // 8: msg.InputPeer
	(*UserMessage)(nil),             // 9: msg.UserMessage
	(*Dialog)(nil),                  // 10: msg.Dialog
	(*User)(nil),                    // 11: msg.User
	(*Group)(nil),                   // 12: msg.Group
}
var file_chat_labels_proto_depIdxs = []int32{
	8,  // 0: msg.LabelsAddToMessage.Peer:type_name -> msg.InputPeer
	8,  // 1: msg.LabelsRemoveFromMessage.Peer:type_name -> msg.InputPeer
	9,  // 2: msg.LabelItems.Messages:type_name -> msg.UserMessage
	10, // 3: msg.LabelItems.Dialogs:type_name -> msg.Dialog
	11, // 4: msg.LabelItems.Users:type_name -> msg.User
	12, // 5: msg.LabelItems.Groups:type_name -> msg.Group
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chat_labels_proto_init() }
func file_chat_labels_proto_init() {
	if File_chat_labels_proto != nil {
		return
	}
	file_core_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_labels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsCreate); i {
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
		file_chat_labels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsEdit); i {
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
		file_chat_labels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsDelete); i {
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
		file_chat_labels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsGet); i {
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
		file_chat_labels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsAddToMessage); i {
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
		file_chat_labels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsRemoveFromMessage); i {
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
		file_chat_labels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelsListItems); i {
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
		file_chat_labels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelItems); i {
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
			RawDescriptor: file_chat_labels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_labels_proto_goTypes,
		DependencyIndexes: file_chat_labels_proto_depIdxs,
		MessageInfos:      file_chat_labels_proto_msgTypes,
	}.Build()
	File_chat_labels_proto = out.File
	file_chat_labels_proto_rawDesc = nil
	file_chat_labels_proto_goTypes = nil
	file_chat_labels_proto_depIdxs = nil
}
