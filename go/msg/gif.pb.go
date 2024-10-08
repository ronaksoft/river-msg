// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: gif.proto

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

// GifGetSaved
// @Function
// @Return: SavedGifs
type GifGetSaved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash uint32 `protobuf:"fixed32,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *GifGetSaved) Reset() {
	*x = GifGetSaved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GifGetSaved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GifGetSaved) ProtoMessage() {}

func (x *GifGetSaved) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GifGetSaved.ProtoReflect.Descriptor instead.
func (*GifGetSaved) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{0}
}

func (x *GifGetSaved) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

// GifSave
// @Function
// @Return: Bool
type GifSave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc        *InputDocument       `protobuf:"bytes,1,opt,name=Doc,proto3" json:"Doc,omitempty"`
	Attributes []*DocumentAttribute `protobuf:"bytes,2,rep,name=Attributes,proto3" json:"Attributes,omitempty"`
}

func (x *GifSave) Reset() {
	*x = GifSave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GifSave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GifSave) ProtoMessage() {}

func (x *GifSave) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GifSave.ProtoReflect.Descriptor instead.
func (*GifSave) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{1}
}

func (x *GifSave) GetDoc() *InputDocument {
	if x != nil {
		return x.Doc
	}
	return nil
}

func (x *GifSave) GetAttributes() []*DocumentAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// GifDelete
type GifDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Doc *InputDocument `protobuf:"bytes,1,opt,name=Doc,proto3" json:"Doc,omitempty"`
}

func (x *GifDelete) Reset() {
	*x = GifDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GifDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GifDelete) ProtoMessage() {}

func (x *GifDelete) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GifDelete.ProtoReflect.Descriptor instead.
func (*GifDelete) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{2}
}

func (x *GifDelete) GetDoc() *InputDocument {
	if x != nil {
		return x.Doc
	}
	return nil
}

// GifSearch
// @Function
// @Return: FoundGifs
type GifSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	Hash  int64  `protobuf:"varint,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *GifSearch) Reset() {
	*x = GifSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GifSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GifSearch) ProtoMessage() {}

func (x *GifSearch) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GifSearch.ProtoReflect.Descriptor instead.
func (*GifSearch) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{3}
}

func (x *GifSearch) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GifSearch) GetHash() int64 {
	if x != nil {
		return x.Hash
	}
	return 0
}

// FoundGifs
type FoundGifs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextOffset int32       `protobuf:"varint,1,opt,name=NextOffset,proto3" json:"NextOffset,omitempty"`
	Gifs       []*FoundGif `protobuf:"bytes,2,rep,name=Gifs,proto3" json:"Gifs,omitempty"`
}

func (x *FoundGifs) Reset() {
	*x = FoundGifs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoundGifs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoundGifs) ProtoMessage() {}

func (x *FoundGifs) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoundGifs.ProtoReflect.Descriptor instead.
func (*FoundGifs) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{4}
}

func (x *FoundGifs) GetNextOffset() int32 {
	if x != nil {
		return x.NextOffset
	}
	return 0
}

func (x *FoundGifs) GetGifs() []*FoundGif {
	if x != nil {
		return x.Gifs
	}
	return nil
}

// FoundGif
type FoundGif struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string    `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Doc   *Document `protobuf:"bytes,2,opt,name=Doc,proto3" json:"Doc,omitempty"`
	Thumb *Document `protobuf:"bytes,3,opt,name=Thumb,proto3" json:"Thumb,omitempty"`
}

func (x *FoundGif) Reset() {
	*x = FoundGif{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoundGif) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoundGif) ProtoMessage() {}

func (x *FoundGif) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoundGif.ProtoReflect.Descriptor instead.
func (*FoundGif) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{5}
}

func (x *FoundGif) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FoundGif) GetDoc() *Document {
	if x != nil {
		return x.Doc
	}
	return nil
}

func (x *FoundGif) GetThumb() *Document {
	if x != nil {
		return x.Thumb
	}
	return nil
}

// SavedGifs
type SavedGifs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash        uint32           `protobuf:"fixed32,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Docs        []*MediaDocument `protobuf:"bytes,2,rep,name=Docs,proto3" json:"Docs,omitempty"`
	NotModified bool             `protobuf:"varint,3,opt,name=NotModified,proto3" json:"NotModified,omitempty"`
}

func (x *SavedGifs) Reset() {
	*x = SavedGifs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gif_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedGifs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedGifs) ProtoMessage() {}

func (x *SavedGifs) ProtoReflect() protoreflect.Message {
	mi := &file_gif_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedGifs.ProtoReflect.Descriptor instead.
func (*SavedGifs) Descriptor() ([]byte, []int) {
	return file_gif_proto_rawDescGZIP(), []int{6}
}

func (x *SavedGifs) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}

func (x *SavedGifs) GetDocs() []*MediaDocument {
	if x != nil {
		return x.Docs
	}
	return nil
}

func (x *SavedGifs) GetNotModified() bool {
	if x != nil {
		return x.NotModified
	}
	return false
}

var File_gif_proto protoreflect.FileDescriptor

var file_gif_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67,
	0x1a, 0x1a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21,
	0x0a, 0x0b, 0x47, 0x69, 0x66, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x22, 0x67, 0x0a, 0x07, 0x47, 0x69, 0x66, 0x53, 0x61, 0x76, 0x65, 0x12, 0x24, 0x0a, 0x03,
	0x44, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03, 0x44,
	0x6f, 0x63, 0x12, 0x36, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x09, 0x47, 0x69,
	0x66, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x44, 0x6f, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03, 0x44, 0x6f, 0x63, 0x22, 0x35, 0x0a,
	0x09, 0x47, 0x69, 0x66, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x4e, 0x0a, 0x09, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x69, 0x66,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x04, 0x47, 0x69, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x69, 0x66, 0x52, 0x04,
	0x47, 0x69, 0x66, 0x73, 0x22, 0x62, 0x0a, 0x08, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x69, 0x66,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x03, 0x44, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03,
	0x44, 0x6f, 0x63, 0x12, 0x23, 0x0a, 0x05, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x22, 0x69, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x47, 0x69, 0x66, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x6f, 0x63,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x44, 0x6f, 0x63,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gif_proto_rawDescOnce sync.Once
	file_gif_proto_rawDescData = file_gif_proto_rawDesc
)

func file_gif_proto_rawDescGZIP() []byte {
	file_gif_proto_rawDescOnce.Do(func() {
		file_gif_proto_rawDescData = protoimpl.X.CompressGZIP(file_gif_proto_rawDescData)
	})
	return file_gif_proto_rawDescData
}

var file_gif_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gif_proto_goTypes = []interface{}{
	(*GifGetSaved)(nil),       // 0: msg.GifGetSaved
	(*GifSave)(nil),           // 1: msg.GifSave
	(*GifDelete)(nil),         // 2: msg.GifDelete
	(*GifSearch)(nil),         // 3: msg.GifSearch
	(*FoundGifs)(nil),         // 4: msg.FoundGifs
	(*FoundGif)(nil),          // 5: msg.FoundGif
	(*SavedGifs)(nil),         // 6: msg.SavedGifs
	(*InputDocument)(nil),     // 7: msg.InputDocument
	(*DocumentAttribute)(nil), // 8: msg.DocumentAttribute
	(*Document)(nil),          // 9: msg.Document
	(*MediaDocument)(nil),     // 10: msg.MediaDocument
}
var file_gif_proto_depIdxs = []int32{
	7,  // 0: msg.GifSave.Doc:type_name -> msg.InputDocument
	8,  // 1: msg.GifSave.Attributes:type_name -> msg.DocumentAttribute
	7,  // 2: msg.GifDelete.Doc:type_name -> msg.InputDocument
	5,  // 3: msg.FoundGifs.Gifs:type_name -> msg.FoundGif
	9,  // 4: msg.FoundGif.Doc:type_name -> msg.Document
	9,  // 5: msg.FoundGif.Thumb:type_name -> msg.Document
	10, // 6: msg.SavedGifs.Docs:type_name -> msg.MediaDocument
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_gif_proto_init() }
func file_gif_proto_init() {
	if File_gif_proto != nil {
		return
	}
	file_chat_messages_medias_proto_init()
	file_core_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GifGetSaved); i {
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
		file_gif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GifSave); i {
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
		file_gif_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GifDelete); i {
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
		file_gif_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GifSearch); i {
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
		file_gif_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoundGifs); i {
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
		file_gif_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoundGif); i {
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
		file_gif_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedGifs); i {
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
			RawDescriptor: file_gif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gif_proto_goTypes,
		DependencyIndexes: file_gif_proto_depIdxs,
		MessageInfos:      file_gif_proto_msgTypes,
	}.Build()
	File_gif_proto = out.File
	file_gif_proto_rawDesc = nil
	file_gif_proto_goTypes = nil
	file_gif_proto_depIdxs = nil
}
