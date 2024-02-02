// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.2
// source: chat.messages.actions.proto

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

// MessageActionGroupAddUser
type MessageActionGroupAddUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []int64 `protobuf:"varint,1,rep,packed,name=UserIDs,proto3" json:"UserIDs,omitempty"`
}

func (x *MessageActionGroupAddUser) Reset() {
	*x = MessageActionGroupAddUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionGroupAddUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionGroupAddUser) ProtoMessage() {}

func (x *MessageActionGroupAddUser) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionGroupAddUser.ProtoReflect.Descriptor instead.
func (*MessageActionGroupAddUser) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{0}
}

func (x *MessageActionGroupAddUser) GetUserIDs() []int64 {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

// MessageActionGroupDeleteUser
type MessageActionGroupDeleteUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []int64 `protobuf:"varint,1,rep,packed,name=UserIDs,proto3" json:"UserIDs,omitempty"`
}

func (x *MessageActionGroupDeleteUser) Reset() {
	*x = MessageActionGroupDeleteUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionGroupDeleteUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionGroupDeleteUser) ProtoMessage() {}

func (x *MessageActionGroupDeleteUser) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionGroupDeleteUser.ProtoReflect.Descriptor instead.
func (*MessageActionGroupDeleteUser) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{1}
}

func (x *MessageActionGroupDeleteUser) GetUserIDs() []int64 {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

// MessageActionGroupCreated
type MessageActionGroupCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupTitle string  `protobuf:"bytes,1,opt,name=GroupTitle,proto3" json:"GroupTitle,omitempty"`
	UserIDs    []int64 `protobuf:"varint,2,rep,packed,name=UserIDs,proto3" json:"UserIDs,omitempty"`
}

func (x *MessageActionGroupCreated) Reset() {
	*x = MessageActionGroupCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionGroupCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionGroupCreated) ProtoMessage() {}

func (x *MessageActionGroupCreated) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionGroupCreated.ProtoReflect.Descriptor instead.
func (*MessageActionGroupCreated) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{2}
}

func (x *MessageActionGroupCreated) GetGroupTitle() string {
	if x != nil {
		return x.GroupTitle
	}
	return ""
}

func (x *MessageActionGroupCreated) GetUserIDs() []int64 {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

// MessageActionGroupTitleChanged
type MessageActionGroupTitleChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupTitle string `protobuf:"bytes,1,opt,name=GroupTitle,proto3" json:"GroupTitle,omitempty"`
}

func (x *MessageActionGroupTitleChanged) Reset() {
	*x = MessageActionGroupTitleChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionGroupTitleChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionGroupTitleChanged) ProtoMessage() {}

func (x *MessageActionGroupTitleChanged) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionGroupTitleChanged.ProtoReflect.Descriptor instead.
func (*MessageActionGroupTitleChanged) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{3}
}

func (x *MessageActionGroupTitleChanged) GetGroupTitle() string {
	if x != nil {
		return x.GroupTitle
	}
	return ""
}

// MessageActionGroupPhotoChanged
type MessageActionGroupPhotoChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Photo *GroupPhoto `protobuf:"bytes,1,opt,name=Photo,proto3" json:"Photo,omitempty"`
}

func (x *MessageActionGroupPhotoChanged) Reset() {
	*x = MessageActionGroupPhotoChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionGroupPhotoChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionGroupPhotoChanged) ProtoMessage() {}

func (x *MessageActionGroupPhotoChanged) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionGroupPhotoChanged.ProtoReflect.Descriptor instead.
func (*MessageActionGroupPhotoChanged) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{4}
}

func (x *MessageActionGroupPhotoChanged) GetPhoto() *GroupPhoto {
	if x != nil {
		return x.Photo
	}
	return nil
}

// MessageActionClearHistory
type MessageActionClearHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxID  int64 `protobuf:"varint,1,opt,name=MaxID,proto3" json:"MaxID,omitempty"`
	Delete bool  `protobuf:"varint,2,opt,name=Delete,proto3" json:"Delete,omitempty"`
}

func (x *MessageActionClearHistory) Reset() {
	*x = MessageActionClearHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionClearHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionClearHistory) ProtoMessage() {}

func (x *MessageActionClearHistory) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionClearHistory.ProtoReflect.Descriptor instead.
func (*MessageActionClearHistory) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{5}
}

func (x *MessageActionClearHistory) GetMaxID() int64 {
	if x != nil {
		return x.MaxID
	}
	return 0
}

func (x *MessageActionClearHistory) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

// MessageActionContactRegistered
type MessageActionContactRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessageActionContactRegistered) Reset() {
	*x = MessageActionContactRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionContactRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionContactRegistered) ProtoMessage() {}

func (x *MessageActionContactRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionContactRegistered.ProtoReflect.Descriptor instead.
func (*MessageActionContactRegistered) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{6}
}

// MessageActionScreenShotTaken
type MessageActionScreenShotTaken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinID int64 `protobuf:"varint,1,opt,name=MinID,proto3" json:"MinID,omitempty"`
	MaxID int64 `protobuf:"varint,2,opt,name=MaxID,proto3" json:"MaxID,omitempty"`
}

func (x *MessageActionScreenShotTaken) Reset() {
	*x = MessageActionScreenShotTaken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionScreenShotTaken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionScreenShotTaken) ProtoMessage() {}

func (x *MessageActionScreenShotTaken) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionScreenShotTaken.ProtoReflect.Descriptor instead.
func (*MessageActionScreenShotTaken) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{7}
}

func (x *MessageActionScreenShotTaken) GetMinID() int64 {
	if x != nil {
		return x.MinID
	}
	return 0
}

func (x *MessageActionScreenShotTaken) GetMaxID() int64 {
	if x != nil {
		return x.MaxID
	}
	return 0
}

// MessageActionThreadClosed
type MessageActionThreadClosed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadID int64 `protobuf:"varint,1,opt,name=ThreadID,proto3" json:"ThreadID,omitempty"`
}

func (x *MessageActionThreadClosed) Reset() {
	*x = MessageActionThreadClosed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionThreadClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionThreadClosed) ProtoMessage() {}

func (x *MessageActionThreadClosed) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionThreadClosed.ProtoReflect.Descriptor instead.
func (*MessageActionThreadClosed) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{8}
}

func (x *MessageActionThreadClosed) GetThreadID() int64 {
	if x != nil {
		return x.ThreadID
	}
	return 0
}

// MessageActionCallStarted
type MessageActionCallStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallID int64 `protobuf:"varint,1,opt,name=CallID,proto3" json:"CallID,omitempty"`
}

func (x *MessageActionCallStarted) Reset() {
	*x = MessageActionCallStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionCallStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionCallStarted) ProtoMessage() {}

func (x *MessageActionCallStarted) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionCallStarted.ProtoReflect.Descriptor instead.
func (*MessageActionCallStarted) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{9}
}

func (x *MessageActionCallStarted) GetCallID() int64 {
	if x != nil {
		return x.CallID
	}
	return 0
}

// MessageActionCallEnded
type MessageActionCallEnded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallID   int64         `protobuf:"varint,1,opt,name=CallID,proto3" json:"CallID,omitempty"`
	Reason   DiscardReason `protobuf:"varint,2,opt,name=Reason,proto3,enum=msg.DiscardReason" json:"Reason,omitempty"`
	Duration uint32        `protobuf:"varint,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
}

func (x *MessageActionCallEnded) Reset() {
	*x = MessageActionCallEnded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_actions_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageActionCallEnded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageActionCallEnded) ProtoMessage() {}

func (x *MessageActionCallEnded) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_actions_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageActionCallEnded.ProtoReflect.Descriptor instead.
func (*MessageActionCallEnded) Descriptor() ([]byte, []int) {
	return file_chat_messages_actions_proto_rawDescGZIP(), []int{10}
}

func (x *MessageActionCallEnded) GetCallID() int64 {
	if x != nil {
		return x.CallID
	}
	return 0
}

func (x *MessageActionCallEnded) GetReason() DiscardReason {
	if x != nil {
		return x.Reason
	}
	return DiscardReason_DiscardReasonUnknown
}

func (x *MessageActionCallEnded) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_chat_messages_actions_proto protoreflect.FileDescriptor

var file_chat_messages_actions_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x1a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x22, 0x3c, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22,
	0x59, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x40, 0x0a, 0x1e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x47, 0x0a, 0x1e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x25,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x05,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x4d, 0x61, 0x78, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x22, 0x20, 0x0a, 0x1e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x22, 0x4a, 0x0a, 0x1c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x4d, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x78, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x61, 0x78, 0x49, 0x44, 0x22, 0x37,
	0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x22, 0x32, 0x0a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x44, 0x22, 0x78, 0x0a, 0x16, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c,
	0x45, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x44, 0x12, 0x2a, 0x0a,
	0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_messages_actions_proto_rawDescOnce sync.Once
	file_chat_messages_actions_proto_rawDescData = file_chat_messages_actions_proto_rawDesc
)

func file_chat_messages_actions_proto_rawDescGZIP() []byte {
	file_chat_messages_actions_proto_rawDescOnce.Do(func() {
		file_chat_messages_actions_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_messages_actions_proto_rawDescData)
	})
	return file_chat_messages_actions_proto_rawDescData
}

var file_chat_messages_actions_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chat_messages_actions_proto_goTypes = []interface{}{
	(*MessageActionGroupAddUser)(nil),      // 0: msg.MessageActionGroupAddUser
	(*MessageActionGroupDeleteUser)(nil),   // 1: msg.MessageActionGroupDeleteUser
	(*MessageActionGroupCreated)(nil),      // 2: msg.MessageActionGroupCreated
	(*MessageActionGroupTitleChanged)(nil), // 3: msg.MessageActionGroupTitleChanged
	(*MessageActionGroupPhotoChanged)(nil), // 4: msg.MessageActionGroupPhotoChanged
	(*MessageActionClearHistory)(nil),      // 5: msg.MessageActionClearHistory
	(*MessageActionContactRegistered)(nil), // 6: msg.MessageActionContactRegistered
	(*MessageActionScreenShotTaken)(nil),   // 7: msg.MessageActionScreenShotTaken
	(*MessageActionThreadClosed)(nil),      // 8: msg.MessageActionThreadClosed
	(*MessageActionCallStarted)(nil),       // 9: msg.MessageActionCallStarted
	(*MessageActionCallEnded)(nil),         // 10: msg.MessageActionCallEnded
	(*GroupPhoto)(nil),                     // 11: msg.GroupPhoto
	(DiscardReason)(0),                     // 12: msg.DiscardReason
}
var file_chat_messages_actions_proto_depIdxs = []int32{
	11, // 0: msg.MessageActionGroupPhotoChanged.Photo:type_name -> msg.GroupPhoto
	12, // 1: msg.MessageActionCallEnded.Reason:type_name -> msg.DiscardReason
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_chat_messages_actions_proto_init() }
func file_chat_messages_actions_proto_init() {
	if File_chat_messages_actions_proto != nil {
		return
	}
	file_core_types_proto_init()
	file_chat_phone_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_messages_actions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionGroupAddUser); i {
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
		file_chat_messages_actions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionGroupDeleteUser); i {
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
		file_chat_messages_actions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionGroupCreated); i {
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
		file_chat_messages_actions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionGroupTitleChanged); i {
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
		file_chat_messages_actions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionGroupPhotoChanged); i {
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
		file_chat_messages_actions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionClearHistory); i {
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
		file_chat_messages_actions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionContactRegistered); i {
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
		file_chat_messages_actions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionScreenShotTaken); i {
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
		file_chat_messages_actions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionThreadClosed); i {
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
		file_chat_messages_actions_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionCallStarted); i {
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
		file_chat_messages_actions_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageActionCallEnded); i {
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
			RawDescriptor: file_chat_messages_actions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_messages_actions_proto_goTypes,
		DependencyIndexes: file_chat_messages_actions_proto_depIdxs,
		MessageInfos:      file_chat_messages_actions_proto_msgTypes,
	}.Build()
	File_chat_messages_actions_proto = out.File
	file_chat_messages_actions_proto_rawDesc = nil
	file_chat_messages_actions_proto_goTypes = nil
	file_chat_messages_actions_proto_depIdxs = nil
}
