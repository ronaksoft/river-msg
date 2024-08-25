// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: chat.messages.markups.proto

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

// ReplyKeyboardMarkup
type ReplyKeyboardMarkup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual letter-keyboard
	// in the chat – the user can press a special button in the input field
	// to see the custom keyboard again.
	SingleUse bool `protobuf:"varint,1,opt,name=SingleUse,proto3" json:"SingleUse,omitempty"`
	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users
	// that are @mentioned in the text of the Message object; 2) if the bot's message is a reply
	// (has reply_to_message_id), sender of the original message.
	// Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard
	// to select the new language. Other users in the group don’t see the keyboard.
	Selective bool `protobuf:"varint,2,opt,name=Selective,proto3" json:"Selective,omitempty"`
	// Requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons). If not set, the custom
	// keyboard is always of the same height as the app's standard keyboard.
	Resize bool                 `protobuf:"varint,3,opt,name=Resize,proto3" json:"Resize,omitempty"`
	Rows   []*KeyboardButtonRow `protobuf:"bytes,4,rep,name=Rows,proto3" json:"Rows,omitempty"`
}

func (x *ReplyKeyboardMarkup) Reset() {
	*x = ReplyKeyboardMarkup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyKeyboardMarkup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyKeyboardMarkup) ProtoMessage() {}

func (x *ReplyKeyboardMarkup) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyKeyboardMarkup.ProtoReflect.Descriptor instead.
func (*ReplyKeyboardMarkup) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{0}
}

func (x *ReplyKeyboardMarkup) GetSingleUse() bool {
	if x != nil {
		return x.SingleUse
	}
	return false
}

func (x *ReplyKeyboardMarkup) GetSelective() bool {
	if x != nil {
		return x.Selective
	}
	return false
}

func (x *ReplyKeyboardMarkup) GetResize() bool {
	if x != nil {
		return x.Resize
	}
	return false
}

func (x *ReplyKeyboardMarkup) GetRows() []*KeyboardButtonRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

// ReplyInlineMarkup
type ReplyInlineMarkup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*KeyboardButtonRow `protobuf:"bytes,1,rep,name=Rows,proto3" json:"Rows,omitempty"`
}

func (x *ReplyInlineMarkup) Reset() {
	*x = ReplyInlineMarkup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyInlineMarkup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyInlineMarkup) ProtoMessage() {}

func (x *ReplyInlineMarkup) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyInlineMarkup.ProtoReflect.Descriptor instead.
func (*ReplyInlineMarkup) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyInlineMarkup) GetRows() []*KeyboardButtonRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

// ReplyKeyboardHide
type ReplyKeyboardHide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Use this flag if you want to remove the keyboard for specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `protobuf:"varint,1,opt,name=Selective,proto3" json:"Selective,omitempty"`
}

func (x *ReplyKeyboardHide) Reset() {
	*x = ReplyKeyboardHide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyKeyboardHide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyKeyboardHide) ProtoMessage() {}

func (x *ReplyKeyboardHide) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyKeyboardHide.ProtoReflect.Descriptor instead.
func (*ReplyKeyboardHide) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyKeyboardHide) GetSelective() bool {
	if x != nil {
		return x.Selective
	}
	return false
}

// ReplyKeyboardHide
// Force the user to send a reply
type ReplyKeyboardForceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display the usual letter-keyboard
	// in the chat – the user can press a special button in the input field
	// to see the custom keyboard again.
	SingleUse bool `protobuf:"varint,1,opt,name=SingleUse,proto3" json:"SingleUse,omitempty"`
	// Use this flag if you want to remove the keyboard for specific users only. Targets:
	// 1) users that are @mentioned in the text of the Message object;
	// 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	Selective bool `protobuf:"varint,2,opt,name=Selective,proto3" json:"Selective,omitempty"`
}

func (x *ReplyKeyboardForceReply) Reset() {
	*x = ReplyKeyboardForceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyKeyboardForceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyKeyboardForceReply) ProtoMessage() {}

func (x *ReplyKeyboardForceReply) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyKeyboardForceReply.ProtoReflect.Descriptor instead.
func (*ReplyKeyboardForceReply) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyKeyboardForceReply) GetSingleUse() bool {
	if x != nil {
		return x.SingleUse
	}
	return false
}

func (x *ReplyKeyboardForceReply) GetSelective() bool {
	if x != nil {
		return x.Selective
	}
	return false
}

// KeyboardButtonRow
// This is a list of buttons
type KeyboardButtonRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buttons []*KeyboardButtonEnvelope `protobuf:"bytes,1,rep,name=Buttons,proto3" json:"Buttons,omitempty"`
}

func (x *KeyboardButtonRow) Reset() {
	*x = KeyboardButtonRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardButtonRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardButtonRow) ProtoMessage() {}

func (x *KeyboardButtonRow) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardButtonRow.ProtoReflect.Descriptor instead.
func (*KeyboardButtonRow) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{4}
}

func (x *KeyboardButtonRow) GetButtons() []*KeyboardButtonEnvelope {
	if x != nil {
		return x.Buttons
	}
	return nil
}

// KeyboardButtonEnvelope
// This is an envelope for buttons
type KeyboardButtonEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Constructor int64  `protobuf:"varint,1,opt,name=Constructor,proto3" json:"Constructor,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *KeyboardButtonEnvelope) Reset() {
	*x = KeyboardButtonEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyboardButtonEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyboardButtonEnvelope) ProtoMessage() {}

func (x *KeyboardButtonEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyboardButtonEnvelope.ProtoReflect.Descriptor instead.
func (*KeyboardButtonEnvelope) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{5}
}

func (x *KeyboardButtonEnvelope) GetConstructor() int64 {
	if x != nil {
		return x.Constructor
	}
	return 0
}

func (x *KeyboardButtonEnvelope) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Button
// This button do nothing and is just showing a text
type Button struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *Button) Reset() {
	*x = Button{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Button) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Button) ProtoMessage() {}

func (x *Button) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Button.ProtoReflect.Descriptor instead.
func (*Button) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{6}
}

func (x *Button) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// ButtonUrl
// Client must open the url
type ButtonUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *ButtonUrl) Reset() {
	*x = ButtonUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonUrl) ProtoMessage() {}

func (x *ButtonUrl) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonUrl.ProtoReflect.Descriptor instead.
func (*ButtonUrl) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{7}
}

func (x *ButtonUrl) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ButtonUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// ButtonCallback
// Client must use BotGetCallbackQuery to get a CallbackAnswer from the bot
type ButtonCallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ButtonCallback) Reset() {
	*x = ButtonCallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonCallback) ProtoMessage() {}

func (x *ButtonCallback) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonCallback.ProtoReflect.Descriptor instead.
func (*ButtonCallback) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{8}
}

func (x *ButtonCallback) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ButtonCallback) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ButtonRequestPhone
// Button to request a user's phone number
type ButtonRequestPhone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *ButtonRequestPhone) Reset() {
	*x = ButtonRequestPhone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonRequestPhone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonRequestPhone) ProtoMessage() {}

func (x *ButtonRequestPhone) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonRequestPhone.ProtoReflect.Descriptor instead.
func (*ButtonRequestPhone) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{9}
}

func (x *ButtonRequestPhone) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// ButtonRequestGeoLocation
// Button to request a user's geolocation
type ButtonRequestGeoLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *ButtonRequestGeoLocation) Reset() {
	*x = ButtonRequestGeoLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonRequestGeoLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonRequestGeoLocation) ProtoMessage() {}

func (x *ButtonRequestGeoLocation) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonRequestGeoLocation.ProtoReflect.Descriptor instead.
func (*ButtonRequestGeoLocation) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{10}
}

func (x *ButtonRequestGeoLocation) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// ButtonSwitchInline
// Button to force a user to switch to inline mode Pressing the button will prompt the user
// to select one of their chats, open that chat and insert the bot‘s username and the
// specified inline query in the input field.
type ButtonSwitchInline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Query    string `protobuf:"bytes,2,opt,name=Query,proto3" json:"Query,omitempty"`
	SamePeer bool   `protobuf:"varint,3,opt,name=SamePeer,proto3" json:"SamePeer,omitempty"`
}

func (x *ButtonSwitchInline) Reset() {
	*x = ButtonSwitchInline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonSwitchInline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonSwitchInline) ProtoMessage() {}

func (x *ButtonSwitchInline) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonSwitchInline.ProtoReflect.Descriptor instead.
func (*ButtonSwitchInline) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{11}
}

func (x *ButtonSwitchInline) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ButtonSwitchInline) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ButtonSwitchInline) GetSamePeer() bool {
	if x != nil {
		return x.SamePeer
	}
	return false
}

// ButtonBuy
// Button to buy a product
type ButtonBuy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *ButtonBuy) Reset() {
	*x = ButtonBuy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_markups_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonBuy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonBuy) ProtoMessage() {}

func (x *ButtonBuy) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_markups_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonBuy.ProtoReflect.Descriptor instead.
func (*ButtonBuy) Descriptor() ([]byte, []int) {
	return file_chat_messages_markups_proto_rawDescGZIP(), []int{12}
}

func (x *ButtonBuy) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_chat_messages_markups_proto protoreflect.FileDescriptor

var file_chat_messages_markups_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x12,
	0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x22, 0x31, 0x0a, 0x11, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x48, 0x69, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x55,
	0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x4a, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x6f, 0x77, 0x12, 0x35, 0x0a, 0x07, 0x42, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x07, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x73, 0x22, 0x4e, 0x0a, 0x16, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x1c, 0x0a, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22,
	0x31, 0x0a, 0x09, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x72, 0x6c, 0x22, 0x38, 0x0a, 0x0e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x12,
	0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x61, 0x6d, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x53, 0x61, 0x6d, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x42, 0x75, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_messages_markups_proto_rawDescOnce sync.Once
	file_chat_messages_markups_proto_rawDescData = file_chat_messages_markups_proto_rawDesc
)

func file_chat_messages_markups_proto_rawDescGZIP() []byte {
	file_chat_messages_markups_proto_rawDescOnce.Do(func() {
		file_chat_messages_markups_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_messages_markups_proto_rawDescData)
	})
	return file_chat_messages_markups_proto_rawDescData
}

var file_chat_messages_markups_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chat_messages_markups_proto_goTypes = []interface{}{
	(*ReplyKeyboardMarkup)(nil),      // 0: msg.ReplyKeyboardMarkup
	(*ReplyInlineMarkup)(nil),        // 1: msg.ReplyInlineMarkup
	(*ReplyKeyboardHide)(nil),        // 2: msg.ReplyKeyboardHide
	(*ReplyKeyboardForceReply)(nil),  // 3: msg.ReplyKeyboardForceReply
	(*KeyboardButtonRow)(nil),        // 4: msg.KeyboardButtonRow
	(*KeyboardButtonEnvelope)(nil),   // 5: msg.KeyboardButtonEnvelope
	(*Button)(nil),                   // 6: msg.Button
	(*ButtonUrl)(nil),                // 7: msg.ButtonUrl
	(*ButtonCallback)(nil),           // 8: msg.ButtonCallback
	(*ButtonRequestPhone)(nil),       // 9: msg.ButtonRequestPhone
	(*ButtonRequestGeoLocation)(nil), // 10: msg.ButtonRequestGeoLocation
	(*ButtonSwitchInline)(nil),       // 11: msg.ButtonSwitchInline
	(*ButtonBuy)(nil),                // 12: msg.ButtonBuy
}
var file_chat_messages_markups_proto_depIdxs = []int32{
	4, // 0: msg.ReplyKeyboardMarkup.Rows:type_name -> msg.KeyboardButtonRow
	4, // 1: msg.ReplyInlineMarkup.Rows:type_name -> msg.KeyboardButtonRow
	5, // 2: msg.KeyboardButtonRow.Buttons:type_name -> msg.KeyboardButtonEnvelope
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chat_messages_markups_proto_init() }
func file_chat_messages_markups_proto_init() {
	if File_chat_messages_markups_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_messages_markups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyKeyboardMarkup); i {
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
		file_chat_messages_markups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyInlineMarkup); i {
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
		file_chat_messages_markups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyKeyboardHide); i {
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
		file_chat_messages_markups_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyKeyboardForceReply); i {
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
		file_chat_messages_markups_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardButtonRow); i {
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
		file_chat_messages_markups_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyboardButtonEnvelope); i {
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
		file_chat_messages_markups_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Button); i {
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
		file_chat_messages_markups_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonUrl); i {
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
		file_chat_messages_markups_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonCallback); i {
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
		file_chat_messages_markups_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonRequestPhone); i {
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
		file_chat_messages_markups_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonRequestGeoLocation); i {
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
		file_chat_messages_markups_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonSwitchInline); i {
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
		file_chat_messages_markups_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonBuy); i {
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
			RawDescriptor: file_chat_messages_markups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_messages_markups_proto_goTypes,
		DependencyIndexes: file_chat_messages_markups_proto_depIdxs,
		MessageInfos:      file_chat_messages_markups_proto_msgTypes,
	}.Build()
	File_chat_messages_markups_proto = out.File
	file_chat_messages_markups_proto_rawDesc = nil
	file_chat_messages_markups_proto_goTypes = nil
	file_chat_messages_markups_proto_depIdxs = nil
}
