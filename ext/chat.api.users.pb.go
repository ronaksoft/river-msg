// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.users.proto

package msg

import (
	fmt "fmt"
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

// UsersGet
// @Function
// @Returns: UsersMany
type UsersGet struct {
	Users []*InputUser `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UsersGet) Reset()         { *m = UsersGet{} }
func (m *UsersGet) String() string { return proto.CompactTextString(m) }
func (*UsersGet) ProtoMessage()    {}
func (*UsersGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d96293bfdd29d581, []int{0}
}
func (m *UsersGet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UsersGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UsersGet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UsersGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersGet.Merge(m, src)
}
func (m *UsersGet) XXX_Size() int {
	return m.Size()
}
func (m *UsersGet) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersGet.DiscardUnknown(m)
}

var xxx_messageInfo_UsersGet proto.InternalMessageInfo

func (m *UsersGet) GetUsers() []*InputUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// UsersGetFull
// @Function
// @Returns: UsersFullMany
type UsersGetFull struct {
	Users []*InputUser `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UsersGetFull) Reset()         { *m = UsersGetFull{} }
func (m *UsersGetFull) String() string { return proto.CompactTextString(m) }
func (*UsersGetFull) ProtoMessage()    {}
func (*UsersGetFull) Descriptor() ([]byte, []int) {
	return fileDescriptor_d96293bfdd29d581, []int{1}
}
func (m *UsersGetFull) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UsersGetFull) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UsersGetFull.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UsersGetFull) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersGetFull.Merge(m, src)
}
func (m *UsersGetFull) XXX_Size() int {
	return m.Size()
}
func (m *UsersGetFull) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersGetFull.DiscardUnknown(m)
}

var xxx_messageInfo_UsersGetFull proto.InternalMessageInfo

func (m *UsersGetFull) GetUsers() []*InputUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// UsersMany
type UsersMany struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users" json:"Users,omitempty"`
}

func (m *UsersMany) Reset()         { *m = UsersMany{} }
func (m *UsersMany) String() string { return proto.CompactTextString(m) }
func (*UsersMany) ProtoMessage()    {}
func (*UsersMany) Descriptor() ([]byte, []int) {
	return fileDescriptor_d96293bfdd29d581, []int{2}
}
func (m *UsersMany) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UsersMany) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UsersMany.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UsersMany) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersMany.Merge(m, src)
}
func (m *UsersMany) XXX_Size() int {
	return m.Size()
}
func (m *UsersMany) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersMany.DiscardUnknown(m)
}

var xxx_messageInfo_UsersMany proto.InternalMessageInfo

func (m *UsersMany) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UsersGet)(nil), "msg.UsersGet")
	proto.RegisterType((*UsersGetFull)(nil), "msg.UsersGetFull")
	proto.RegisterType((*UsersMany)(nil), "msg.UsersMany")
}

func init() { proto.RegisterFile("chat.api.users.proto", fileDescriptor_d96293bfdd29d581) }

var fileDescriptor_d96293bfdd29d581 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xce, 0x48, 0x2c,
	0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x2d, 0x4e, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x12, 0x05, 0x4b, 0x25, 0xe7, 0x17, 0xa5, 0xea, 0x95, 0x54,
	0x16, 0xa4, 0x42, 0xe5, 0x94, 0x0c, 0xb8, 0x38, 0x42, 0x41, 0x4a, 0xdd, 0x53, 0x4b, 0x84, 0x54,
	0xb8, 0x58, 0xc1, 0x6c, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0xbd, 0xdc, 0xe2, 0x74,
	0x3d, 0xcf, 0xbc, 0x82, 0xd2, 0x12, 0x90, 0x70, 0x10, 0x44, 0x52, 0xc9, 0x84, 0x8b, 0x07, 0xa6,
	0xc3, 0xad, 0x34, 0x27, 0x87, 0x48, 0x5d, 0x3a, 0x5c, 0x9c, 0x60, 0x86, 0x6f, 0x62, 0x5e, 0xa5,
	0x90, 0x3c, 0xaa, 0x16, 0x4e, 0xb0, 0x16, 0x24, 0xd5, 0x4e, 0x12, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x9a, 0x88, 0xe5, 0xe2, 0x00,
	0x00, 0x00,
}

func (m *UsersGet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsersGet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			dAtA[i] = 0xa
			i++
			i = encodeVarintChatApiUsers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UsersGetFull) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsersGetFull) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			dAtA[i] = 0xa
			i++
			i = encodeVarintChatApiUsers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *UsersMany) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UsersMany) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, msg := range m.Users {
			dAtA[i] = 0xa
			i++
			i = encodeVarintChatApiUsers(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintChatApiUsers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UsersGet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovChatApiUsers(uint64(l))
		}
	}
	return n
}

func (m *UsersGetFull) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovChatApiUsers(uint64(l))
		}
	}
	return n
}

func (m *UsersMany) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovChatApiUsers(uint64(l))
		}
	}
	return n
}

func sovChatApiUsers(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChatApiUsers(x uint64) (n int) {
	return sovChatApiUsers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UsersGet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiUsers
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
			return fmt.Errorf("proto: UsersGet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsersGet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiUsers
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
				return ErrInvalidLengthChatApiUsers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &InputUser{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiUsers
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
func (m *UsersGetFull) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiUsers
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
			return fmt.Errorf("proto: UsersGetFull: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsersGetFull: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiUsers
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
				return ErrInvalidLengthChatApiUsers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &InputUser{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiUsers
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
func (m *UsersMany) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiUsers
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
			return fmt.Errorf("proto: UsersMany: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UsersMany: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiUsers
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
				return ErrInvalidLengthChatApiUsers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiUsers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiUsers
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiUsers
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
func skipChatApiUsers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChatApiUsers
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
					return 0, ErrIntOverflowChatApiUsers
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
					return 0, ErrIntOverflowChatApiUsers
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
				return 0, ErrInvalidLengthChatApiUsers
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthChatApiUsers
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChatApiUsers
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
				next, err := skipChatApiUsers(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthChatApiUsers
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
	ErrInvalidLengthChatApiUsers = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChatApiUsers   = fmt.Errorf("proto: integer overflow")
)
