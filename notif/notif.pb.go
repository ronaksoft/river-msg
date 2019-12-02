// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notif.proto

package msgNotif

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Notification struct {
	Type      int32               `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Title     string              `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body      string              `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	UserID    int64               `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data      []*NotificationData `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
	Tags      []string            `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Timestamp int64               `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return m.Size()
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Notification) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Notification) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Notification) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Notification) GetData() []*NotificationData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Notification) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Notification) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type NotificationData struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NotificationData) Reset()         { *m = NotificationData{} }
func (m *NotificationData) String() string { return proto.CompactTextString(m) }
func (*NotificationData) ProtoMessage()    {}
func (*NotificationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{1}
}
func (m *NotificationData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotificationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotificationData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotificationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationData.Merge(m, src)
}
func (m *NotificationData) XXX_Size() int {
	return m.Size()
}
func (m *NotificationData) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationData.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationData proto.InternalMessageInfo

func (m *NotificationData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NotificationData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SendResponse struct {
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{2}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return m.Size()
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

type CancelRequest struct {
	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (m *CancelRequest) Reset()         { *m = CancelRequest{} }
func (m *CancelRequest) String() string { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()    {}
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{3}
}
func (m *CancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRequest.Merge(m, src)
}
func (m *CancelRequest) XXX_Size() int {
	return m.Size()
}
func (m *CancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRequest proto.InternalMessageInfo

func (m *CancelRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CancelResponse struct {
}

func (m *CancelResponse) Reset()         { *m = CancelResponse{} }
func (m *CancelResponse) String() string { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()    {}
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3a72a17febd07f, []int{4}
}
func (m *CancelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResponse.Merge(m, src)
}
func (m *CancelResponse) XXX_Size() int {
	return m.Size()
}
func (m *CancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Notification)(nil), "msgNotif.Notification")
	proto.RegisterType((*NotificationData)(nil), "msgNotif.NotificationData")
	proto.RegisterType((*SendResponse)(nil), "msgNotif.SendResponse")
	proto.RegisterType((*CancelRequest)(nil), "msgNotif.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "msgNotif.CancelResponse")
}

func init() { proto.RegisterFile("notif.proto", fileDescriptor_bc3a72a17febd07f) }

var fileDescriptor_bc3a72a17febd07f = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0x26, 0xcd, 0xb5, 0xe7, 0x5e, 0x2f, 0x65, 0x14, 0x1d, 0x82, 0xc4, 0x90, 0xbb,
	0xc9, 0xc6, 0x5c, 0xa8, 0x1b, 0x11, 0xdc, 0xd4, 0x6e, 0xba, 0x71, 0x31, 0xc5, 0xb5, 0x4c, 0x92,
	0x69, 0x1c, 0x6c, 0x32, 0x31, 0x33, 0x29, 0xf4, 0x25, 0xc4, 0xc7, 0x72, 0x59, 0x70, 0xe3, 0x4a,
	0x24, 0x7d, 0x11, 0x99, 0x09, 0xfd, 0xa7, 0xee, 0xbe, 0xef, 0xcc, 0x77, 0xce, 0xf9, 0x1d, 0x18,
	0xb8, 0xae, 0xa5, 0x16, 0xab, 0xb4, 0x69, 0xa5, 0x96, 0xf8, 0x61, 0xa5, 0xca, 0xf7, 0xc6, 0x07,
	0x2f, 0x4b, 0xa1, 0x3f, 0x75, 0x59, 0x9a, 0xcb, 0xea, 0xbe, 0x94, 0xa5, 0xbc, 0xb7, 0x81, 0xac,
	0x5b, 0x59, 0x67, 0x8d, 0x55, 0x43, 0x63, 0xfc, 0x03, 0xc1, 0x8d, 0x6d, 0x14, 0x39, 0xd3, 0x42,
	0xd6, 0x18, 0x83, 0xa7, 0xb7, 0x0d, 0x27, 0x28, 0x42, 0xc9, 0x88, 0x5a, 0x8d, 0x9f, 0xc0, 0x48,
	0x0b, 0xbd, 0xe6, 0xe4, 0x41, 0x84, 0x92, 0x31, 0x1d, 0x8c, 0x49, 0x66, 0xb2, 0xd8, 0x12, 0xd7,
	0x16, 0xad, 0xc6, 0x77, 0x70, 0xd5, 0x29, 0xde, 0x7e, 0x14, 0x05, 0xf1, 0x22, 0x94, 0xb8, 0x33,
	0xe8, 0x7f, 0xbd, 0xf0, 0x3f, 0x28, 0xde, 0x2e, 0xe6, 0xd4, 0x37, 0x4f, 0x8b, 0x02, 0xa7, 0xe0,
	0x15, 0x4c, 0x33, 0x32, 0x8a, 0xdc, 0xe4, 0x7a, 0x1a, 0xa4, 0x07, 0xf6, 0xf4, 0x1c, 0x64, 0xce,
	0x34, 0xa3, 0x36, 0x67, 0x91, 0x58, 0xa9, 0x88, 0x1f, 0xb9, 0x66, 0x91, 0xd1, 0xf8, 0x39, 0x8c,
	0xb5, 0xa8, 0xb8, 0xd2, 0xac, 0x6a, 0xc8, 0x95, 0x59, 0x45, 0x4f, 0x85, 0xf8, 0x0d, 0x4c, 0xfe,
	0x9e, 0x85, 0x27, 0xe0, 0x7e, 0xe6, 0x5b, 0x7b, 0xd7, 0x98, 0x1a, 0x69, 0xce, 0xda, 0xb0, 0x75,
	0x77, 0x3c, 0xcb, 0x9a, 0xf8, 0x16, 0x6e, 0x96, 0xbc, 0x2e, 0x28, 0x57, 0x8d, 0xac, 0x15, 0x8f,
	0xef, 0xe0, 0xd1, 0x3b, 0x56, 0xe7, 0x7c, 0x4d, 0xf9, 0x97, 0x8e, 0x2b, 0x7d, 0xc4, 0x41, 0x27,
	0x9c, 0x78, 0x02, 0xb7, 0x87, 0xd0, 0xd0, 0x36, 0xfd, 0x8a, 0xe0, 0xf1, 0x39, 0xc3, 0x92, 0xb7,
	0x1b, 0x91, 0x73, 0xfc, 0x1a, 0x3c, 0x33, 0x1e, 0x3f, 0xfd, 0xff, 0xd9, 0xc1, 0x59, 0xfd, 0x02,
	0xc3, 0xc1, 0x6f, 0xc1, 0x1f, 0x76, 0xe0, 0x67, 0xa7, 0xcc, 0x05, 0x5a, 0x40, 0xfe, 0x7d, 0x38,
	0xb4, 0xcf, 0xc8, 0xf7, 0x3e, 0x44, 0xbb, 0x3e, 0x44, 0xbf, 0xfb, 0x10, 0x7d, 0xdb, 0x87, 0xce,
	0x6e, 0x1f, 0x3a, 0x3f, 0xf7, 0xa1, 0x93, 0xf9, 0xf6, 0x2b, 0xbc, 0xfa, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xcb, 0x4f, 0x05, 0xae, 0x52, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Send(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*SendResponse, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Send(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/msgNotif.NotificationService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/msgNotif.NotificationService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Send(context.Context, *Notification) (*SendResponse, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) Send(ctx context.Context, req *Notification) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedNotificationServiceServer) Cancel(ctx context.Context, req *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgNotif.NotificationService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Send(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgNotif.NotificationService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msgNotif.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _NotificationService_Send_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _NotificationService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notif.proto",
}

func (m *Notification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Notification) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNotif(dAtA, i, uint64(m.Type))
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNotif(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNotif(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if m.UserID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintNotif(dAtA, i, uint64(m.UserID))
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintNotif(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintNotif(dAtA, i, uint64(m.Timestamp))
	}
	return i, nil
}

func (m *NotificationData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotificationData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNotif(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNotif(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *SendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *CancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintNotif(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Notification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovNotif(uint64(m.Type))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovNotif(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovNotif(uint64(l))
	}
	if m.UserID != 0 {
		n += 1 + sovNotif(uint64(m.UserID))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovNotif(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovNotif(uint64(l))
		}
	}
	if m.Timestamp != 0 {
		n += 1 + sovNotif(uint64(m.Timestamp))
	}
	return n
}

func (m *NotificationData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovNotif(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovNotif(uint64(l))
	}
	return n
}

func (m *SendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *CancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovNotif(uint64(l))
		}
	}
	return n
}

func (m *CancelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovNotif(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNotif(x uint64) (n int) {
	return sovNotif(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Notification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotif
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
			return fmt.Errorf("proto: Notification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Notification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &NotificationData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNotif(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotif
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotif
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
func (m *NotificationData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotif
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
			return fmt.Errorf("proto: NotificationData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotificationData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotif(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotif
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotif
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
func (m *SendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotif
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
			return fmt.Errorf("proto: SendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNotif(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotif
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotif
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
func (m *CancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotif
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
			return fmt.Errorf("proto: CancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotif
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
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotif
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotif(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotif
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotif
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
func (m *CancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotif
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
			return fmt.Errorf("proto: CancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNotif(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNotif
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNotif
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
func skipNotif(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotif
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
					return 0, ErrIntOverflowNotif
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
					return 0, ErrIntOverflowNotif
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
				return 0, ErrInvalidLengthNotif
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNotif
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNotif
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
				next, err := skipNotif(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNotif
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
	ErrInvalidLengthNotif = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotif   = fmt.Errorf("proto: integer overflow")
)
