// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bot.proto

package bot

import (
	context "context"
	fmt "fmt"
	ext "git.ronaksoftware.com/river/messages/ext"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StartBotRequest struct {
	User       *ext.InputUser `protobuf:"bytes,1,req,name=User" json:"User,omitempty"`
	Bot        *ext.InputUser `protobuf:"bytes,2,req,name=Bot" json:"Bot,omitempty"`
	Peer       *ext.InputPeer `protobuf:"bytes,3,req,name=Peer" json:"Peer,omitempty"`
	StartParam string         `protobuf:"bytes,4,req,name=StartParam" json:"StartParam"`
}

func (m *StartBotRequest) Reset()         { *m = StartBotRequest{} }
func (m *StartBotRequest) String() string { return proto.CompactTextString(m) }
func (*StartBotRequest) ProtoMessage()    {}
func (*StartBotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{0}
}
func (m *StartBotRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartBotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartBotRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartBotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBotRequest.Merge(m, src)
}
func (m *StartBotRequest) XXX_Size() int {
	return m.Size()
}
func (m *StartBotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartBotRequest proto.InternalMessageInfo

func (m *StartBotRequest) GetUser() *ext.InputUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StartBotRequest) GetBot() *ext.InputUser {
	if m != nil {
		return m.Bot
	}
	return nil
}

func (m *StartBotRequest) GetPeer() *ext.InputPeer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *StartBotRequest) GetStartParam() string {
	if m != nil {
		return m.StartParam
	}
	return ""
}

type StartBotResponse struct {
}

func (m *StartBotResponse) Reset()         { *m = StartBotResponse{} }
func (m *StartBotResponse) String() string { return proto.CompactTextString(m) }
func (*StartBotResponse) ProtoMessage()    {}
func (*StartBotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{1}
}
func (m *StartBotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartBotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartBotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartBotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBotResponse.Merge(m, src)
}
func (m *StartBotResponse) XXX_Size() int {
	return m.Size()
}
func (m *StartBotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartBotResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartBotRequest)(nil), "bot.StartBotRequest")
	proto.RegisterType((*StartBotResponse)(nil), "bot.StartBotResponse")
}

func init() { proto.RegisterFile("bot.proto", fileDescriptor_51d7d70385167023) }

var fileDescriptor_51d7d70385167023 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4a, 0xc4, 0x30,
	0x18, 0x86, 0xfb, 0xb7, 0x98, 0x89, 0xa0, 0x12, 0x14, 0x42, 0x17, 0x99, 0x52, 0x5c, 0xcc, 0xc6,
	0x14, 0x66, 0xe7, 0x36, 0xe0, 0xc2, 0x5d, 0xa9, 0x78, 0x80, 0xb6, 0xc4, 0x8e, 0x8b, 0xce, 0x57,
	0x93, 0xaf, 0x30, 0x73, 0x0b, 0xaf, 0xe0, 0x6d, 0x66, 0x39, 0x4b, 0x57, 0x22, 0xed, 0x45, 0x24,
	0x29, 0xa2, 0x14, 0x77, 0xf9, 0xde, 0xe7, 0xc9, 0x9b, 0x1f, 0xb2, 0xac, 0x00, 0x45, 0xa7, 0x01,
	0x81, 0x86, 0x15, 0x60, 0x7c, 0xdb, 0xbc, 0xe0, 0xb6, 0xaf, 0x44, 0x0d, 0x6d, 0xd6, 0x40, 0x03,
	0x99, 0x63, 0x55, 0xff, 0xec, 0x26, 0x37, 0xb8, 0xd5, 0xb4, 0x27, 0x5e, 0xa9, 0x3d, 0x4e, 0x4a,
	0x56, 0x6f, 0x4b, 0x14, 0x35, 0x68, 0x25, 0xf0, 0xd0, 0x29, 0x33, 0x09, 0xe9, 0xbb, 0x4f, 0x2e,
	0x1e, 0xb1, 0xd4, 0x28, 0x01, 0x0b, 0xf5, 0xda, 0x2b, 0x83, 0x34, 0x25, 0xd1, 0x93, 0x51, 0x9a,
	0xf9, 0x49, 0xb0, 0x3e, 0xdb, 0x9c, 0x8b, 0xd6, 0x34, 0xe2, 0x61, 0xd7, 0xf5, 0x68, 0xd3, 0xc2,
	0x31, 0x9a, 0x90, 0x50, 0x02, 0xb2, 0xe0, 0x5f, 0xc5, 0x22, 0xdb, 0x92, 0x2b, 0xa5, 0x59, 0x38,
	0x57, 0x6c, 0x5a, 0x38, 0x46, 0x6f, 0x08, 0x71, 0x87, 0xe7, 0xa5, 0x2e, 0x5b, 0x16, 0x25, 0xc1,
	0x7a, 0x29, 0xa3, 0xe3, 0xe7, 0xca, 0x2b, 0xfe, 0xe4, 0x29, 0x25, 0x97, 0xbf, 0x57, 0x34, 0x1d,
	0xec, 0x8c, 0xda, 0xdc, 0x93, 0x85, 0x04, 0xcc, 0x35, 0xec, 0x0f, 0xf4, 0x8e, 0x2c, 0x7e, 0x38,
	0xbd, 0x12, 0xf6, 0xc3, 0x66, 0x2f, 0x8a, 0xaf, 0x67, 0xe9, 0x54, 0x92, 0x7a, 0x92, 0x1d, 0x07,
	0xee, 0x9f, 0x06, 0xee, 0x7f, 0x0d, 0xdc, 0x7f, 0x1b, 0xb9, 0x77, 0x1a, 0xb9, 0xf7, 0x31, 0x72,
	0xef, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xdf, 0xc4, 0xe9, 0x79, 0x01, 0x00, 0x00,
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
	StartBot(ctx context.Context, in *StartBotRequest, opts ...grpc.CallOption) (*StartBotResponse, error)
}

type botProxyClient struct {
	cc *grpc.ClientConn
}

func NewBotProxyClient(cc *grpc.ClientConn) BotProxyClient {
	return &botProxyClient{cc}
}

func (c *botProxyClient) StartBot(ctx context.Context, in *StartBotRequest, opts ...grpc.CallOption) (*StartBotResponse, error) {
	out := new(StartBotResponse)
	err := c.cc.Invoke(ctx, "/bot.BotProxy/StartBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotProxyServer is the server API for BotProxy service.
type BotProxyServer interface {
	StartBot(context.Context, *StartBotRequest) (*StartBotResponse, error)
}

// UnimplementedBotProxyServer can be embedded to have forward compatible implementations.
type UnimplementedBotProxyServer struct {
}

func (*UnimplementedBotProxyServer) StartBot(ctx context.Context, req *StartBotRequest) (*StartBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBot not implemented")
}

func RegisterBotProxyServer(s *grpc.Server, srv BotProxyServer) {
	s.RegisterService(&_BotProxy_serviceDesc, srv)
}

func _BotProxy_StartBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotProxyServer).StartBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bot.BotProxy/StartBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotProxyServer).StartBot(ctx, req.(*StartBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bot.BotProxy",
	HandlerType: (*BotProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBot",
			Handler:    _BotProxy_StartBot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot.proto",
}

func (m *StartBotRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartBotRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.User == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("User")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBot(dAtA, i, uint64(m.User.Size()))
		n1, err1 := m.User.MarshalTo(dAtA[i:])
		if err1 != nil {
			return 0, err1
		}
		i += n1
	}
	if m.Bot == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBot(dAtA, i, uint64(m.Bot.Size()))
		n2, err2 := m.Bot.MarshalTo(dAtA[i:])
		if err2 != nil {
			return 0, err2
		}
		i += n2
	}
	if m.Peer == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Peer")
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBot(dAtA, i, uint64(m.Peer.Size()))
		n3, err3 := m.Peer.MarshalTo(dAtA[i:])
		if err3 != nil {
			return 0, err3
		}
		i += n3
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintBot(dAtA, i, uint64(len(m.StartParam)))
	i += copy(dAtA[i:], m.StartParam)
	return i, nil
}

func (m *StartBotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartBotResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintBot(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StartBotRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovBot(uint64(l))
	}
	if m.Bot != nil {
		l = m.Bot.Size()
		n += 1 + l + sovBot(uint64(l))
	}
	if m.Peer != nil {
		l = m.Peer.Size()
		n += 1 + l + sovBot(uint64(l))
	}
	l = len(m.StartParam)
	n += 1 + l + sovBot(uint64(l))
	return n
}

func (m *StartBotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovBot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBot(x uint64) (n int) {
	return sovBot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StartBotRequest) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBot
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
			return fmt.Errorf("proto: StartBotRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartBotRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBot
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
				return ErrInvalidLengthBot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &ext.InputUser{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBot
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
				return ErrInvalidLengthBot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Bot == nil {
				m.Bot = &ext.InputUser{}
			}
			if err := m.Bot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBot
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
				return ErrInvalidLengthBot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Peer == nil {
				m.Peer = &ext.InputPeer{}
			}
			if err := m.Peer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartParam", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBot
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
				return ErrInvalidLengthBot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartParam = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipBot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBot
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("User")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bot")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Peer")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("StartParam")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StartBotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBot
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
			return fmt.Errorf("proto: StartBotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartBotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipBot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBot
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
func skipBot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBot
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
					return 0, ErrIntOverflowBot
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBot
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
				return 0, ErrInvalidLengthBot
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBot
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBot(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBot
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBot   = fmt.Errorf("proto: integer overflow")
)