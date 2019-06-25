// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: notif.proto

package notif

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
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
	Body      string              `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	UserID    int64               `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Data      []*NotificationData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	Tags      []string            `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Timestamp int64               `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
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
	proto.RegisterType((*Notification)(nil), "notif.Notification")
	proto.RegisterType((*NotificationData)(nil), "notif.NotificationData")
	proto.RegisterType((*SendResponse)(nil), "notif.SendResponse")
	proto.RegisterType((*CancelRequest)(nil), "notif.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "notif.CancelResponse")
}

func init() { proto.RegisterFile("notif.proto", fileDescriptor_bc3a72a17febd07f) }

var fileDescriptor_bc3a72a17febd07f = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x6e, 0xda, 0x40,
	0x14, 0x85, 0x3d, 0xf5, 0x4f, 0xc5, 0x85, 0x22, 0x34, 0x50, 0xd5, 0x42, 0x95, 0x6b, 0x99, 0x8d,
	0xa5, 0xaa, 0x20, 0xd1, 0x45, 0xa5, 0x2e, 0x29, 0x1b, 0x36, 0x5d, 0x0c, 0xca, 0x3a, 0x1a, 0xdb,
	0x83, 0x63, 0x05, 0x3c, 0x8e, 0x3d, 0x46, 0x62, 0x9b, 0x27, 0xc8, 0xf3, 0xe4, 0x09, 0xb2, 0x64,
	0x99, 0x55, 0x14, 0x99, 0x17, 0x89, 0x66, 0x0c, 0xe1, 0x47, 0xd9, 0x9d, 0x7b, 0xe6, 0x9e, 0x7b,
	0x3e, 0xcb, 0xd0, 0x4c, 0xb9, 0x48, 0x16, 0xc3, 0x2c, 0xe7, 0x82, 0x63, 0x53, 0x0d, 0xfd, 0x5f,
	0x71, 0x22, 0x6e, 0xca, 0x60, 0x18, 0xf2, 0xd5, 0x28, 0xe6, 0x31, 0x1f, 0xa9, 0xd7, 0xa0, 0x5c,
	0xa8, 0x49, 0x0d, 0x4a, 0xd5, 0x29, 0xef, 0x11, 0x41, 0xeb, 0xbf, 0x0c, 0x26, 0x21, 0x15, 0x09,
	0x4f, 0x31, 0x06, 0x43, 0x6c, 0x32, 0x66, 0x23, 0x17, 0xf9, 0x26, 0x51, 0x5a, 0x7a, 0x01, 0x8f,
	0x36, 0xf6, 0x27, 0x17, 0xf9, 0x0d, 0xa2, 0x34, 0x1e, 0xc0, 0xe7, 0xb2, 0x60, 0xf9, 0x75, 0x12,
	0xd9, 0xba, 0x8b, 0x7c, 0x7d, 0x02, 0xd5, 0xcb, 0x0f, 0xeb, 0xaa, 0x60, 0xf9, 0x6c, 0x4a, 0x2c,
	0xf9, 0x34, 0x8b, 0xf0, 0x4f, 0x30, 0x22, 0x2a, 0xa8, 0x6d, 0xb8, 0xba, 0xdf, 0x1c, 0x7f, 0x1b,
	0xd6, 0xbc, 0xa7, 0x7d, 0x53, 0x2a, 0x28, 0x51, 0x4b, 0xaa, 0x99, 0xc6, 0x85, 0x6d, 0xba, 0xba,
	0x6c, 0x91, 0x1a, 0x7f, 0x87, 0x86, 0x48, 0x56, 0xac, 0x10, 0x74, 0x95, 0xd9, 0x96, 0xec, 0x21,
	0x47, 0xc3, 0xfb, 0x0b, 0x9d, 0xcb, 0x5b, 0xb8, 0x03, 0xfa, 0x2d, 0xdb, 0x28, 0xfc, 0x06, 0x91,
	0x12, 0xf7, 0xc0, 0x5c, 0xd3, 0x65, 0xc9, 0xf6, 0xf8, 0xf5, 0xe0, 0xb5, 0xa1, 0x35, 0x67, 0x69,
	0x44, 0x58, 0x91, 0xf1, 0xb4, 0x60, 0xde, 0x00, 0xbe, 0xfc, 0xa3, 0x69, 0xc8, 0x96, 0x84, 0xdd,
	0x95, 0xac, 0x10, 0xef, 0x38, 0xe8, 0x88, 0xe3, 0x75, 0xa0, 0x7d, 0x58, 0xaa, 0x63, 0xe3, 0x7b,
	0x04, 0xdd, 0x53, 0x86, 0x39, 0xcb, 0xd7, 0x49, 0xc8, 0xf0, 0x18, 0x0c, 0x79, 0x1e, 0x77, 0x3f,
	0xf8, 0xe6, 0xfe, 0xc1, 0x3c, 0x03, 0xd0, 0xf0, 0x1f, 0xb0, 0xea, 0xeb, 0xb8, 0xb7, 0x5f, 0x38,
	0x23, 0xea, 0x7f, 0xbd, 0x70, 0x0f, 0xc1, 0x89, 0xfd, 0x54, 0x39, 0x68, 0x5b, 0x39, 0xe8, 0xb5,
	0x72, 0xd0, 0xc3, 0xce, 0xd1, 0xb6, 0x3b, 0x47, 0x7b, 0xde, 0x39, 0x5a, 0x60, 0xa9, 0xbf, 0xfc,
	0xfb, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x17, 0x84, 0xa0, 0x0d, 0x2a, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/notif.NotificationService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/notif.NotificationService/Cancel", in, out, opts...)
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
		FullMethod: "/notif.NotificationService/Send",
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
		FullMethod: "/notif.NotificationService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notif.NotificationService",
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
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNotif(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if m.UserID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNotif(dAtA, i, uint64(m.UserID))
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x22
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
			dAtA[i] = 0x2a
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
		dAtA[i] = 0x30
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
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
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
			wire |= (uint64(b) & 0x7F) << shift
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
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
				m.UserID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
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
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &NotificationData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
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
				m.Timestamp |= (int64(b) & 0x7F) << shift
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
			wire |= (uint64(b) & 0x7F) << shift
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
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
			wire |= (uint64(b) & 0x7F) << shift
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
			wire |= (uint64(b) & 0x7F) << shift
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
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNotif
			}
			postIndex := iNdEx + intStringLen
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
			wire |= (uint64(b) & 0x7F) << shift
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
			iNdEx += length
			if length < 0 {
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
