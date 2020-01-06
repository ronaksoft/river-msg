// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.core.password.algorithms.proto

package msg

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
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

type PasswordAlgorithmKdfSha512Iter100000 struct {
	Salt1 []byte `protobuf:"bytes,1,req,name=Salt1" json:"Salt1"`
	Salt2 []byte `protobuf:"bytes,2,req,name=Salt2" json:"Salt2"`
	G     int32  `protobuf:"varint,3,req,name=G" json:"G"`
	P     []byte `protobuf:"bytes,4,req,name=P" json:"P"`
}

func (m *PasswordAlgorithmKdfSha512Iter100000) Reset()         { *m = PasswordAlgorithmKdfSha512Iter100000{} }
func (m *PasswordAlgorithmKdfSha512Iter100000) String() string { return proto.CompactTextString(m) }
func (*PasswordAlgorithmKdfSha512Iter100000) ProtoMessage()    {}
func (*PasswordAlgorithmKdfSha512Iter100000) Descriptor() ([]byte, []int) {
	return fileDescriptor_aea8cee14a2df555, []int{0}
}
func (m *PasswordAlgorithmKdfSha512Iter100000) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PasswordAlgorithmKdfSha512Iter100000) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PasswordAlgorithmKdfSha512Iter100000.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PasswordAlgorithmKdfSha512Iter100000) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordAlgorithmKdfSha512Iter100000.Merge(m, src)
}
func (m *PasswordAlgorithmKdfSha512Iter100000) XXX_Size() int {
	return m.Size()
}
func (m *PasswordAlgorithmKdfSha512Iter100000) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordAlgorithmKdfSha512Iter100000.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordAlgorithmKdfSha512Iter100000 proto.InternalMessageInfo

func (m *PasswordAlgorithmKdfSha512Iter100000) GetSalt1() []byte {
	if m != nil {
		return m.Salt1
	}
	return nil
}

func (m *PasswordAlgorithmKdfSha512Iter100000) GetSalt2() []byte {
	if m != nil {
		return m.Salt2
	}
	return nil
}

func (m *PasswordAlgorithmKdfSha512Iter100000) GetG() int32 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *PasswordAlgorithmKdfSha512Iter100000) GetP() []byte {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*PasswordAlgorithmKdfSha512Iter100000)(nil), "msg.PasswordAlgorithmKdfSha512Iter100000")
}

func init() {
	proto.RegisterFile("chat.core.password.algorithms.proto", fileDescriptor_aea8cee14a2df555)
}

var fileDescriptor_aea8cee14a2df555 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x48, 0x2c,
	0xd1, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0xd1, 0x4b,
	0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xce, 0x2d, 0x4e, 0x57, 0x6a, 0x63, 0xe4, 0x52, 0x09, 0x80, 0x2a, 0x71, 0x84, 0xa9, 0xf0,
	0x4e, 0x49, 0x0b, 0xce, 0x48, 0x34, 0x35, 0x34, 0xf2, 0x2c, 0x49, 0x2d, 0x32, 0x34, 0x00, 0x01,
	0x21, 0x29, 0x2e, 0xd6, 0xe0, 0xc4, 0x9c, 0x12, 0x43, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x1e, 0x27,
	0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x20, 0x42, 0x30, 0x39, 0x23, 0x09, 0x26, 0x74, 0x39, 0x23,
	0x21, 0x21, 0x2e, 0x46, 0x77, 0x09, 0x66, 0x05, 0x26, 0x0d, 0x56, 0xa8, 0x38, 0xa3, 0x3b, 0x48,
	0x2c, 0x40, 0x82, 0x05, 0x49, 0x2d, 0x63, 0x80, 0x93, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x8e, 0x3b, 0x43, 0xcd, 0x00, 0x00,
	0x00,
}

func (m *PasswordAlgorithmKdfSha512Iter100000) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PasswordAlgorithmKdfSha512Iter100000) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PasswordAlgorithmKdfSha512Iter100000) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.P != nil {
		i -= len(m.P)
		copy(dAtA[i:], m.P)
		i = encodeVarintChatCorePasswordAlgorithms(dAtA, i, uint64(len(m.P)))
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintChatCorePasswordAlgorithms(dAtA, i, uint64(m.G))
	i--
	dAtA[i] = 0x18
	if m.Salt2 != nil {
		i -= len(m.Salt2)
		copy(dAtA[i:], m.Salt2)
		i = encodeVarintChatCorePasswordAlgorithms(dAtA, i, uint64(len(m.Salt2)))
		i--
		dAtA[i] = 0x12
	}
	if m.Salt1 != nil {
		i -= len(m.Salt1)
		copy(dAtA[i:], m.Salt1)
		i = encodeVarintChatCorePasswordAlgorithms(dAtA, i, uint64(len(m.Salt1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChatCorePasswordAlgorithms(dAtA []byte, offset int, v uint64) int {
	offset -= sovChatCorePasswordAlgorithms(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PasswordAlgorithmKdfSha512Iter100000) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Salt1 != nil {
		l = len(m.Salt1)
		n += 1 + l + sovChatCorePasswordAlgorithms(uint64(l))
	}
	if m.Salt2 != nil {
		l = len(m.Salt2)
		n += 1 + l + sovChatCorePasswordAlgorithms(uint64(l))
	}
	n += 1 + sovChatCorePasswordAlgorithms(uint64(m.G))
	if m.P != nil {
		l = len(m.P)
		n += 1 + l + sovChatCorePasswordAlgorithms(uint64(l))
	}
	return n
}

func sovChatCorePasswordAlgorithms(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChatCorePasswordAlgorithms(x uint64) (n int) {
	return sovChatCorePasswordAlgorithms(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PasswordAlgorithmKdfSha512Iter100000) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatCorePasswordAlgorithms
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
			return fmt.Errorf("proto: PasswordAlgorithmKdfSha512Iter100000: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PasswordAlgorithmKdfSha512Iter100000: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt1", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatCorePasswordAlgorithms
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt1 = append(m.Salt1[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt1 == nil {
				m.Salt1 = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt2", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatCorePasswordAlgorithms
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt2 = append(m.Salt2[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt2 == nil {
				m.Salt2 = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field G", wireType)
			}
			m.G = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatCorePasswordAlgorithms
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.G |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatCorePasswordAlgorithms
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.P = append(m.P[:0], dAtA[iNdEx:postIndex]...)
			if m.P == nil {
				m.P = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipChatCorePasswordAlgorithms(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatCorePasswordAlgorithms
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Salt1")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Salt2")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("G")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("P")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChatCorePasswordAlgorithms(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChatCorePasswordAlgorithms
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
					return 0, ErrIntOverflowChatCorePasswordAlgorithms
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
					return 0, ErrIntOverflowChatCorePasswordAlgorithms
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
				return 0, ErrInvalidLengthChatCorePasswordAlgorithms
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChatCorePasswordAlgorithms
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChatCorePasswordAlgorithms
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChatCorePasswordAlgorithms        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChatCorePasswordAlgorithms          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChatCorePasswordAlgorithms = fmt.Errorf("proto: unexpected end of group")
)
