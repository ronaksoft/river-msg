// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bot.api.proto

package msg

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StartBot struct {
	Bot        *InputPeer `protobuf:"bytes,1,req,name=Bot" json:"Bot,omitempty"`
	RandomID   int64      `protobuf:"varint,2,req,name=RandomID" json:"RandomID"`
	StartParam string     `protobuf:"bytes,3,req,name=StartParam" json:"StartParam"`
}

func (m *StartBot) Reset()         { *m = StartBot{} }
func (m *StartBot) String() string { return proto.CompactTextString(m) }
func (*StartBot) ProtoMessage()    {}
func (*StartBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_802c818ed586bbc7, []int{0}
}
func (m *StartBot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartBot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBot.Merge(m, src)
}
func (m *StartBot) XXX_Size() int {
	return m.Size()
}
func (m *StartBot) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBot.DiscardUnknown(m)
}

var xxx_messageInfo_StartBot proto.InternalMessageInfo

func (m *StartBot) GetBot() *InputPeer {
	if m != nil {
		return m.Bot
	}
	return nil
}

func (m *StartBot) GetRandomID() int64 {
	if m != nil {
		return m.RandomID
	}
	return 0
}

func (m *StartBot) GetStartParam() string {
	if m != nil {
		return m.StartParam
	}
	return ""
}

// BotIsStarted
// @Function
// @Return: Bool
type BotIsStarted struct {
	Bot      *InputPeer `protobuf:"bytes,1,req,name=Bot" json:"Bot,omitempty"`
	RandomID int64      `protobuf:"varint,2,req,name=RandomID" json:"RandomID"`
}

func (m *BotIsStarted) Reset()         { *m = BotIsStarted{} }
func (m *BotIsStarted) String() string { return proto.CompactTextString(m) }
func (*BotIsStarted) ProtoMessage()    {}
func (*BotIsStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_802c818ed586bbc7, []int{1}
}
func (m *BotIsStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BotIsStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BotIsStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BotIsStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotIsStarted.Merge(m, src)
}
func (m *BotIsStarted) XXX_Size() int {
	return m.Size()
}
func (m *BotIsStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_BotIsStarted.DiscardUnknown(m)
}

var xxx_messageInfo_BotIsStarted proto.InternalMessageInfo

func (m *BotIsStarted) GetBot() *InputPeer {
	if m != nil {
		return m.Bot
	}
	return nil
}

func (m *BotIsStarted) GetRandomID() int64 {
	if m != nil {
		return m.RandomID
	}
	return 0
}

// BotSetInfo
// @Function
// @Return: Bool
type BotSetInfo struct {
	BotID       int64          `protobuf:"varint,1,req,name=BotID" json:"BotID"`
	RandomID    int64          `protobuf:"varint,2,req,name=RandomID" json:"RandomID"`
	Owner       int64          `protobuf:"varint,3,req,name=Owner" json:"Owner"`
	BotCommands []*BotCommands `protobuf:"bytes,4,rep,name=BotCommands" json:"BotCommands,omitempty"`
	Description string         `protobuf:"bytes,5,req,name=Description" json:"Description"`
}

func (m *BotSetInfo) Reset()         { *m = BotSetInfo{} }
func (m *BotSetInfo) String() string { return proto.CompactTextString(m) }
func (*BotSetInfo) ProtoMessage()    {}
func (*BotSetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_802c818ed586bbc7, []int{2}
}
func (m *BotSetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BotSetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BotSetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BotSetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSetInfo.Merge(m, src)
}
func (m *BotSetInfo) XXX_Size() int {
	return m.Size()
}
func (m *BotSetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BotSetInfo proto.InternalMessageInfo

func (m *BotSetInfo) GetBotID() int64 {
	if m != nil {
		return m.BotID
	}
	return 0
}

func (m *BotSetInfo) GetRandomID() int64 {
	if m != nil {
		return m.RandomID
	}
	return 0
}

func (m *BotSetInfo) GetOwner() int64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *BotSetInfo) GetBotCommands() []*BotCommands {
	if m != nil {
		return m.BotCommands
	}
	return nil
}

func (m *BotSetInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Bots
// @Function
// @Return: BotInfo
type Bots struct {
	UserID int64 `protobuf:"varint,1,req,name=UserID" json:"UserID"`
	Limit  int32 `protobuf:"varint,2,opt,name=Limit" json:"Limit"`
}

func (m *Bots) Reset()         { *m = Bots{} }
func (m *Bots) String() string { return proto.CompactTextString(m) }
func (*Bots) ProtoMessage()    {}
func (*Bots) Descriptor() ([]byte, []int) {
	return fileDescriptor_802c818ed586bbc7, []int{3}
}
func (m *Bots) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bots.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bots.Merge(m, src)
}
func (m *Bots) XXX_Size() int {
	return m.Size()
}
func (m *Bots) XXX_DiscardUnknown() {
	xxx_messageInfo_Bots.DiscardUnknown(m)
}

var xxx_messageInfo_Bots proto.InternalMessageInfo

func (m *Bots) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Bots) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// BotsMany
type BotsMany struct {
	Bots []*BotInfo `protobuf:"bytes,1,rep,name=Bots" json:"Bots,omitempty"`
}

func (m *BotsMany) Reset()         { *m = BotsMany{} }
func (m *BotsMany) String() string { return proto.CompactTextString(m) }
func (*BotsMany) ProtoMessage()    {}
func (*BotsMany) Descriptor() ([]byte, []int) {
	return fileDescriptor_802c818ed586bbc7, []int{4}
}
func (m *BotsMany) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BotsMany) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BotsMany.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BotsMany) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsMany.Merge(m, src)
}
func (m *BotsMany) XXX_Size() int {
	return m.Size()
}
func (m *BotsMany) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsMany.DiscardUnknown(m)
}

var xxx_messageInfo_BotsMany proto.InternalMessageInfo

func (m *BotsMany) GetBots() []*BotInfo {
	if m != nil {
		return m.Bots
	}
	return nil
}

func init() {
	proto.RegisterType((*StartBot)(nil), "msg.StartBot")
	proto.RegisterType((*BotIsStarted)(nil), "msg.BotIsStarted")
	proto.RegisterType((*BotSetInfo)(nil), "msg.BotSetInfo")
	proto.RegisterType((*Bots)(nil), "msg.Bots")
	proto.RegisterType((*BotsMany)(nil), "msg.BotsMany")
}

func init() { proto.RegisterFile("bot.api.proto", fileDescriptor_802c818ed586bbc7) }

var fileDescriptor_802c818ed586bbc7 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbd, 0x76, 0x52, 0x85, 0x4d, 0xf9, 0xa3, 0x15, 0x48, 0x96, 0x0f, 0xc6, 0xb2, 0x10,
	0xf2, 0x01, 0xdc, 0x2a, 0x8f, 0xb0, 0x0d, 0x87, 0x48, 0x54, 0x8d, 0x1c, 0xf1, 0x00, 0x9b, 0x64,
	0xea, 0x5a, 0xc2, 0x3b, 0x66, 0x77, 0x52, 0xc8, 0x5b, 0xf0, 0x4c, 0x9c, 0x7a, 0xec, 0x91, 0x13,
	0x42, 0xc9, 0x8b, 0x20, 0xaf, 0x93, 0xc8, 0x07, 0x0e, 0x1c, 0xb8, 0xed, 0xfe, 0xe6, 0x9b, 0x6f,
	0xbe, 0x19, 0xfe, 0x74, 0x89, 0x94, 0xab, 0xa6, 0xca, 0x1b, 0x83, 0x84, 0x22, 0xa8, 0x6d, 0x19,
	0xbd, 0x2f, 0x2b, 0xba, 0xdb, 0x2c, 0xf3, 0x15, 0xd6, 0x17, 0x25, 0x96, 0x78, 0xe1, 0x6a, 0xcb,
	0xcd, 0xad, 0xfb, 0xb9, 0x8f, 0x7b, 0x75, 0x3d, 0xd1, 0xab, 0xd5, 0x9d, 0xa2, 0x7c, 0x85, 0x06,
	0x72, 0xda, 0x36, 0x60, 0x3b, 0x9c, 0x12, 0x1f, 0x2d, 0x48, 0x19, 0x92, 0x48, 0x22, 0xe1, 0x81,
	0x44, 0x0a, 0x59, 0xe2, 0x67, 0xe3, 0xc9, 0xb3, 0xbc, 0xb6, 0x65, 0x3e, 0xd3, 0xcd, 0x86, 0xe6,
	0x00, 0xa6, 0x08, 0x3a, 0xc5, 0xa8, 0x50, 0x7a, 0x8d, 0xf5, 0x6c, 0x1a, 0xfa, 0x89, 0x9f, 0x05,
	0x72, 0xf0, 0xf0, 0xeb, 0xb5, 0x57, 0x9c, 0xa8, 0x78, 0xc3, 0xb9, 0xf3, 0x9b, 0x2b, 0xa3, 0xea,
	0x30, 0x48, 0xfc, 0xec, 0xc9, 0x41, 0xd3, 0xe3, 0x69, 0xc1, 0xcf, 0x25, 0xd2, 0xcc, 0x3a, 0x04,
	0xeb, 0xff, 0x31, 0x39, 0xfd, 0xc1, 0x38, 0x97, 0x48, 0x0b, 0xa0, 0x99, 0xbe, 0x45, 0x11, 0xf1,
	0x61, 0x3b, 0x62, 0xea, 0x4c, 0x8f, 0xea, 0x0e, 0xfd, 0xc3, 0x1a, 0x11, 0x1f, 0xde, 0x7c, 0xd5,
	0x60, 0xdc, 0x06, 0xa7, 0x6e, 0x87, 0xc4, 0x84, 0x8f, 0x25, 0xd2, 0x15, 0xd6, 0xb5, 0xd2, 0x6b,
	0x1b, 0x0e, 0x92, 0x20, 0x1b, 0x4f, 0x5e, 0xb8, 0xd0, 0x3d, 0x5e, 0xf4, 0x45, 0xe2, 0x2d, 0x1f,
	0x4f, 0xc1, 0xae, 0x4c, 0xd5, 0x50, 0x85, 0x3a, 0x1c, 0xf6, 0xee, 0xd2, 0x2f, 0xa4, 0x92, 0x0f,
	0x24, 0x92, 0x15, 0x31, 0x3f, 0xfb, 0x64, 0xc1, 0x9c, 0xe2, 0x9f, 0xb5, 0xd2, 0x4b, 0x56, 0x1c,
	0x68, 0x9b, 0xef, 0x63, 0x55, 0x57, 0x14, 0xfa, 0x09, 0xcb, 0x86, 0xc7, 0x7c, 0x0e, 0xa5, 0xef,
	0xf8, 0xa8, 0xf5, 0xb8, 0x56, 0x7a, 0x2b, 0x92, 0xce, 0x2f, 0x64, 0x2e, 0xe4, 0xf9, 0x31, 0x64,
	0x7b, 0xa1, 0xc2, 0x55, 0x26, 0x37, 0x4e, 0x3d, 0x37, 0xf8, 0x6d, 0x2b, 0xae, 0xf8, 0xf3, 0x05,
	0x98, 0x7b, 0x90, 0x48, 0x05, 0x7c, 0xd9, 0x80, 0x25, 0xf1, 0xd2, 0xb5, 0x5c, 0x83, 0xb5, 0xaa,
	0x84, 0x0f, 0xfa, 0x1e, 0x3e, 0x63, 0x03, 0xd1, 0x5f, 0x69, 0xea, 0x65, 0xec, 0x92, 0xc9, 0xf0,
	0x61, 0x17, 0xb3, 0xc7, 0x5d, 0xcc, 0x7e, 0xef, 0x62, 0xf6, 0x7d, 0x1f, 0x7b, 0x8f, 0xfb, 0xd8,
	0xfb, 0xb9, 0x8f, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xa8, 0x18, 0x0e, 0xc6, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BotProxyClient is the client API for BotProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BotProxyClient interface {
	ServeBotRequest(ctx context.Context, opts ...grpc.CallOption) (BotProxy_ServeBotRequestClient, error)
}

type botProxyClient struct {
	cc *grpc.ClientConn
}

func NewBotProxyClient(cc *grpc.ClientConn) BotProxyClient {
	return &botProxyClient{cc}
}

func (c *botProxyClient) ServeBotRequest(ctx context.Context, opts ...grpc.CallOption) (BotProxy_ServeBotRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BotProxy_serviceDesc.Streams[0], "/msg.BotProxy/ServeBotRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &botProxyServeBotRequestClient{stream}
	return x, nil
}

type BotProxy_ServeBotRequestClient interface {
	Send(*MessageEnvelope) error
	Recv() (*MessageEnvelope, error)
	grpc.ClientStream
}

type botProxyServeBotRequestClient struct {
	grpc.ClientStream
}

func (x *botProxyServeBotRequestClient) Send(m *MessageEnvelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botProxyServeBotRequestClient) Recv() (*MessageEnvelope, error) {
	m := new(MessageEnvelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BotProxyServer is the server API for BotProxy service.
type BotProxyServer interface {
	ServeBotRequest(BotProxy_ServeBotRequestServer) error
}

// UnimplementedBotProxyServer can be embedded to have forward compatible implementations.
type UnimplementedBotProxyServer struct {
}

func (*UnimplementedBotProxyServer) ServeBotRequest(srv BotProxy_ServeBotRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method ServeBotRequest not implemented")
}

func RegisterBotProxyServer(s *grpc.Server, srv BotProxyServer) {
	s.RegisterService(&_BotProxy_serviceDesc, srv)
}

func _BotProxy_ServeBotRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotProxyServer).ServeBotRequest(&botProxyServeBotRequestServer{stream})
}

type BotProxy_ServeBotRequestServer interface {
	Send(*MessageEnvelope) error
	Recv() (*MessageEnvelope, error)
	grpc.ServerStream
}

type botProxyServeBotRequestServer struct {
	grpc.ServerStream
}

func (x *botProxyServeBotRequestServer) Send(m *MessageEnvelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botProxyServeBotRequestServer) Recv() (*MessageEnvelope, error) {
	m := new(MessageEnvelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _BotProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.BotProxy",
	HandlerType: (*BotProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServeBotRequest",
			Handler:       _BotProxy_ServeBotRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bot.api.proto",
}

func (m *StartBot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartBot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StartBot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.StartParam)
	copy(dAtA[i:], m.StartParam)
	i = encodeVarintBotApi(dAtA, i, uint64(len(m.StartParam)))
	i--
	dAtA[i] = 0x1a
	i = encodeVarintBotApi(dAtA, i, uint64(m.RandomID))
	i--
	dAtA[i] = 0x10
	if m.Bot == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	} else {
		{
			size, err := m.Bot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBotApi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BotIsStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BotIsStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BotIsStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintBotApi(dAtA, i, uint64(m.RandomID))
	i--
	dAtA[i] = 0x10
	if m.Bot == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	} else {
		{
			size, err := m.Bot.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBotApi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BotSetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BotSetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BotSetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Description)
	copy(dAtA[i:], m.Description)
	i = encodeVarintBotApi(dAtA, i, uint64(len(m.Description)))
	i--
	dAtA[i] = 0x2a
	if len(m.BotCommands) > 0 {
		for iNdEx := len(m.BotCommands) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BotCommands[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBotApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	i = encodeVarintBotApi(dAtA, i, uint64(m.Owner))
	i--
	dAtA[i] = 0x18
	i = encodeVarintBotApi(dAtA, i, uint64(m.RandomID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintBotApi(dAtA, i, uint64(m.BotID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *Bots) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bots) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bots) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintBotApi(dAtA, i, uint64(m.Limit))
	i--
	dAtA[i] = 0x10
	i = encodeVarintBotApi(dAtA, i, uint64(m.UserID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *BotsMany) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BotsMany) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BotsMany) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bots) > 0 {
		for iNdEx := len(m.Bots) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bots[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBotApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBotApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovBotApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StartBot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bot != nil {
		l = m.Bot.Size()
		n += 1 + l + sovBotApi(uint64(l))
	}
	n += 1 + sovBotApi(uint64(m.RandomID))
	l = len(m.StartParam)
	n += 1 + l + sovBotApi(uint64(l))
	return n
}

func (m *BotIsStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Bot != nil {
		l = m.Bot.Size()
		n += 1 + l + sovBotApi(uint64(l))
	}
	n += 1 + sovBotApi(uint64(m.RandomID))
	return n
}

func (m *BotSetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovBotApi(uint64(m.BotID))
	n += 1 + sovBotApi(uint64(m.RandomID))
	n += 1 + sovBotApi(uint64(m.Owner))
	if len(m.BotCommands) > 0 {
		for _, e := range m.BotCommands {
			l = e.Size()
			n += 1 + l + sovBotApi(uint64(l))
		}
	}
	l = len(m.Description)
	n += 1 + l + sovBotApi(uint64(l))
	return n
}

func (m *Bots) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovBotApi(uint64(m.UserID))
	n += 1 + sovBotApi(uint64(m.Limit))
	return n
}

func (m *BotsMany) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Bots) > 0 {
		for _, e := range m.Bots {
			l = e.Size()
			n += 1 + l + sovBotApi(uint64(l))
		}
	}
	return n
}

func sovBotApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBotApi(x uint64) (n int) {
	return sovBotApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StartBot) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StartBot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartBot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bot == nil {
				m.Bot = &InputPeer{}
			}
			if err := m.Bot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomID", wireType)
			}
			m.RandomID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartParam", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartParam = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipBotApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("RandomID")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("StartParam")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BotIsStarted) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BotIsStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BotIsStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bot == nil {
				m.Bot = &InputPeer{}
			}
			if err := m.Bot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomID", wireType)
			}
			m.RandomID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipBotApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("RandomID")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BotSetInfo) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BotSetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BotSetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BotID", wireType)
			}
			m.BotID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BotID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RandomID", wireType)
			}
			m.RandomID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RandomID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			m.Owner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Owner |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BotCommands", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BotCommands = append(m.BotCommands, &BotCommands{})
			if err := m.BotCommands[len(m.BotCommands)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipBotApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("BotID")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("RandomID")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Owner")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Description")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Bots) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bots: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bots: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBotApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("UserID")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BotsMany) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BotsMany: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BotsMany: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bots", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBotApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBotApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bots = append(m.Bots, &BotInfo{})
			if err := m.Bots[len(m.Bots)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBotApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBotApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBotApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBotApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBotApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBotApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBotApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBotApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBotApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBotApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBotApi = fmt.Errorf("proto: unexpected end of group")
)
