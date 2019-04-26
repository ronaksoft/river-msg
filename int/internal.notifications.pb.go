// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal.notifications.proto

package internalMsg

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
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
	Type   int32               `protobuf:"varint,1,req,name=Type" json:"Type"`
	Body   string              `protobuf:"bytes,2,req,name=Body" json:"Body"`
	UserID int64               `protobuf:"varint,3,req,name=UserID" json:"UserID"`
	Data   []*NotificationData `protobuf:"bytes,4,rep,name=Data" json:"Data,omitempty"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_63a8736448a9f241, []int{0}
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

type NotificationData struct {
	Key   string `protobuf:"bytes,1,req,name=Key" json:"Key"`
	Value string `protobuf:"bytes,2,req,name=Value" json:"Value"`
}

func (m *NotificationData) Reset()         { *m = NotificationData{} }
func (m *NotificationData) String() string { return proto.CompactTextString(m) }
func (*NotificationData) ProtoMessage()    {}
func (*NotificationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63a8736448a9f241, []int{1}
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

func init() {
	proto.RegisterType((*Notification)(nil), "internalMsg.Notification")
	proto.RegisterType((*NotificationData)(nil), "internalMsg.NotificationData")
}

func init() { proto.RegisterFile("internal.notifications.proto", fileDescriptor_63a8736448a9f241) }

var fileDescriptor_63a8736448a9f241 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcb, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc9, 0xfa, 0x16, 0xa7, 0x2b,
	0x4d, 0x66, 0xe4, 0xe2, 0xf1, 0x43, 0x52, 0x24, 0x24, 0xc1, 0xc5, 0x12, 0x52, 0x59, 0x90, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0xea, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x58, 0x04, 0x24,
	0xe3, 0x94, 0x9f, 0x52, 0x29, 0xc1, 0xa4, 0xc0, 0xa4, 0xc1, 0x09, 0x93, 0x01, 0x89, 0x08, 0xc9,
	0x70, 0xb1, 0x85, 0x16, 0xa7, 0x16, 0x79, 0xba, 0x48, 0x30, 0x2b, 0x30, 0x69, 0x30, 0x43, 0xe5,
	0xa0, 0x62, 0x42, 0x86, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0x89, 0x12, 0x2c, 0x0a, 0xcc, 0x1a, 0xdc,
	0x46, 0xb2, 0x7a, 0x48, 0xd6, 0xeb, 0x21, 0x5b, 0x0d, 0x52, 0x14, 0x04, 0x56, 0xaa, 0xe4, 0xc6,
	0x25, 0x80, 0x2e, 0x23, 0x24, 0xc6, 0xc5, 0xec, 0x9d, 0x5a, 0x09, 0x76, 0x17, 0xcc, 0x76, 0x90,
	0x80, 0x90, 0x14, 0x17, 0x6b, 0x58, 0x62, 0x4e, 0x69, 0x2a, 0x8a, 0xbb, 0x20, 0x42, 0x4e, 0x12,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x0a, 0x32, 0x8d, 0x23, 0x01, 0x00, 0x00,
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
	dAtA[i] = 0x8
	i++
	i = encodeVarintInternalNotifications(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x12
	i++
	i = encodeVarintInternalNotifications(dAtA, i, uint64(len(m.Body)))
	i += copy(dAtA[i:], m.Body)
	dAtA[i] = 0x18
	i++
	i = encodeVarintInternalNotifications(dAtA, i, uint64(m.UserID))
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x22
			i++
			i = encodeVarintInternalNotifications(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
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
	dAtA[i] = 0xa
	i++
	i = encodeVarintInternalNotifications(dAtA, i, uint64(len(m.Key)))
	i += copy(dAtA[i:], m.Key)
	dAtA[i] = 0x12
	i++
	i = encodeVarintInternalNotifications(dAtA, i, uint64(len(m.Value)))
	i += copy(dAtA[i:], m.Value)
	return i, nil
}

func encodeVarintInternalNotifications(dAtA []byte, offset int, v uint64) int {
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
	n += 1 + sovInternalNotifications(uint64(m.Type))
	l = len(m.Body)
	n += 1 + l + sovInternalNotifications(uint64(l))
	n += 1 + sovInternalNotifications(uint64(m.UserID))
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovInternalNotifications(uint64(l))
		}
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
	n += 1 + l + sovInternalNotifications(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovInternalNotifications(uint64(l))
	return n
}

func sovInternalNotifications(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternalNotifications(x uint64) (n int) {
	return sovInternalNotifications(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Notification) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalNotifications
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
					return ErrIntOverflowInternalNotifications
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
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalNotifications
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
				return ErrInvalidLengthInternalNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalNotifications
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
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalNotifications
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
				return ErrInvalidLengthInternalNotifications
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
		default:
			iNdEx = preIndex
			skippy, err := skipInternalNotifications(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalNotifications
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Type")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Body")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("UserID")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotificationData) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternalNotifications
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
					return ErrIntOverflowInternalNotifications
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
				return ErrInvalidLengthInternalNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternalNotifications
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
				return ErrInvalidLengthInternalNotifications
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipInternalNotifications(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternalNotifications
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Key")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Value")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInternalNotifications(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternalNotifications
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
					return 0, ErrIntOverflowInternalNotifications
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
					return 0, ErrIntOverflowInternalNotifications
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
				return 0, ErrInvalidLengthInternalNotifications
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternalNotifications
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
				next, err := skipInternalNotifications(dAtA[start:])
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
	ErrInvalidLengthInternalNotifications = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternalNotifications   = fmt.Errorf("proto: integer overflow")
)
