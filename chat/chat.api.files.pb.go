// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.api.files.proto

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

// FileType
type FileType int32

const (
	FileType_FileTypeUnknown FileType = 0
	FileType_FileTypePartial FileType = 1
	FileType_FileTypeJpeg    FileType = 2
	FileType_FileTypeGif     FileType = 3
	FileType_FileTypePng     FileType = 4
	FileType_FileTypeWebp    FileType = 5
	FileType_FileTypeMp3     FileType = 6
	FileType_FileTypeMp4     FileType = 7
	FileType_FileTypeMov     FileType = 8
)

var FileType_name = map[int32]string{
	0: "FileTypeUnknown",
	1: "FileTypePartial",
	2: "FileTypeJpeg",
	3: "FileTypeGif",
	4: "FileTypePng",
	5: "FileTypeWebp",
	6: "FileTypeMp3",
	7: "FileTypeMp4",
	8: "FileTypeMov",
}

var FileType_value = map[string]int32{
	"FileTypeUnknown": 0,
	"FileTypePartial": 1,
	"FileTypeJpeg":    2,
	"FileTypeGif":     3,
	"FileTypePng":     4,
	"FileTypeWebp":    5,
	"FileTypeMp3":     6,
	"FileTypeMp4":     7,
	"FileTypeMov":     8,
}

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}

func (x *FileType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FileType_value, data, "FileType")
	if err != nil {
		return err
	}
	*x = FileType(value)
	return nil
}

func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05ced85bdf32a77a, []int{0}
}

// FileSavePart
// @Function
// @Return: Bool
type FileSavePart struct {
	FileID     int64  `protobuf:"varint,1,req,name=FileID" json:"FileID"`
	PartID     int32  `protobuf:"varint,2,req,name=PartID" json:"PartID"`
	TotalParts int32  `protobuf:"varint,3,req,name=TotalParts" json:"TotalParts"`
	Bytes      []byte `protobuf:"bytes,4,req,name=Bytes" json:"Bytes"`
}

func (m *FileSavePart) Reset()         { *m = FileSavePart{} }
func (m *FileSavePart) String() string { return proto.CompactTextString(m) }
func (*FileSavePart) ProtoMessage()    {}
func (*FileSavePart) Descriptor() ([]byte, []int) {
	return fileDescriptor_05ced85bdf32a77a, []int{0}
}
func (m *FileSavePart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileSavePart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileSavePart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileSavePart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSavePart.Merge(m, src)
}
func (m *FileSavePart) XXX_Size() int {
	return m.Size()
}
func (m *FileSavePart) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSavePart.DiscardUnknown(m)
}

var xxx_messageInfo_FileSavePart proto.InternalMessageInfo

func (m *FileSavePart) GetFileID() int64 {
	if m != nil {
		return m.FileID
	}
	return 0
}

func (m *FileSavePart) GetPartID() int32 {
	if m != nil {
		return m.PartID
	}
	return 0
}

func (m *FileSavePart) GetTotalParts() int32 {
	if m != nil {
		return m.TotalParts
	}
	return 0
}

func (m *FileSavePart) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

// FileGetPart
// @Function
// @Return: File
type FileGet struct {
	Location *InputFileLocation `protobuf:"bytes,1,req,name=Location" json:"Location,omitempty"`
	Offset   int32              `protobuf:"varint,2,req,name=Offset" json:"Offset"`
	Limit    int32              `protobuf:"varint,3,req,name=Limit" json:"Limit"`
}

func (m *FileGet) Reset()         { *m = FileGet{} }
func (m *FileGet) String() string { return proto.CompactTextString(m) }
func (*FileGet) ProtoMessage()    {}
func (*FileGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_05ced85bdf32a77a, []int{1}
}
func (m *FileGet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileGet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileGet.Merge(m, src)
}
func (m *FileGet) XXX_Size() int {
	return m.Size()
}
func (m *FileGet) XXX_DiscardUnknown() {
	xxx_messageInfo_FileGet.DiscardUnknown(m)
}

var xxx_messageInfo_FileGet proto.InternalMessageInfo

func (m *FileGet) GetLocation() *InputFileLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *FileGet) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FileGet) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// File
type File struct {
	Type         FileType `protobuf:"varint,1,req,name=Type,enum=msg.FileType" json:"Type"`
	ModifiedTime int64    `protobuf:"varint,2,req,name=ModifiedTime" json:"ModifiedTime"`
	Bytes        []byte   `protobuf:"bytes,4,req,name=Bytes" json:"Bytes"`
	MD5Hash      string   `protobuf:"bytes,5,opt,name=MD5Hash" json:"MD5Hash"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_05ced85bdf32a77a, []int{2}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_File.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return m.Size()
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetType() FileType {
	if m != nil {
		return m.Type
	}
	return FileType_FileTypeUnknown
}

func (m *File) GetModifiedTime() int64 {
	if m != nil {
		return m.ModifiedTime
	}
	return 0
}

func (m *File) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *File) GetMD5Hash() string {
	if m != nil {
		return m.MD5Hash
	}
	return ""
}

func init() {
	proto.RegisterEnum("msg.FileType", FileType_name, FileType_value)
	proto.RegisterType((*FileSavePart)(nil), "msg.FileSavePart")
	proto.RegisterType((*FileGet)(nil), "msg.FileGet")
	proto.RegisterType((*File)(nil), "msg.File")
}

func init() { proto.RegisterFile("chat.api.files.proto", fileDescriptor_05ced85bdf32a77a) }

var fileDescriptor_05ced85bdf32a77a = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xe3, 0xdc, 0x5a, 0xce, 0x14, 0xc6, 0x32, 0x17, 0x45, 0x23, 0x64, 0xa2, 0x0a, 0x89,
	0x88, 0x45, 0x84, 0x0a, 0xbc, 0x40, 0x54, 0x31, 0x14, 0x4d, 0xc5, 0xa8, 0x14, 0xb1, 0x36, 0x1d,
	0x27, 0x63, 0x91, 0xc4, 0x56, 0x63, 0x66, 0xd4, 0x97, 0x40, 0x6c, 0x78, 0x0c, 0xde, 0x63, 0x96,
	0xb3, 0x64, 0x85, 0x50, 0xfb, 0x22, 0xc8, 0xb9, 0xa0, 0x64, 0xc3, 0x2e, 0xfe, 0xfe, 0xef, 0xe4,
	0x3f, 0x89, 0xe1, 0xc1, 0xe6, 0x92, 0xe9, 0x98, 0x29, 0x11, 0xa7, 0x22, 0xe7, 0x55, 0xac, 0xb6,
	0x52, 0x4b, 0xe2, 0x14, 0x55, 0x76, 0xf2, 0xb0, 0x8e, 0x36, 0x72, 0xcb, 0x63, 0xbd, 0x53, 0x5d,
	0x36, 0xfd, 0x86, 0x60, 0xf2, 0x46, 0xe4, 0xfc, 0x03, 0xbb, 0xe2, 0xe7, 0x6c, 0xab, 0x09, 0x05,
	0xdf, 0x9c, 0x17, 0xf3, 0x00, 0x85, 0x76, 0xe4, 0x24, 0xfe, 0xcd, 0xef, 0x27, 0xd6, 0x0b, 0xb4,
	0x6a, 0x29, 0x79, 0x0c, 0xbe, 0xf1, 0x16, 0xf3, 0xc0, 0x0e, 0xed, 0xc8, 0x4b, 0x5c, 0x93, 0xaf,
	0x5a, 0x46, 0x9e, 0x02, 0xac, 0xa5, 0x66, 0xb9, 0x39, 0x56, 0x81, 0xd3, 0x33, 0x7a, 0x9c, 0x9c,
	0x80, 0x97, 0xec, 0x34, 0xaf, 0x02, 0x37, 0xb4, 0xa3, 0x49, 0x2b, 0x34, 0x68, 0x7a, 0x0d, 0x23,
	0xd3, 0x74, 0xca, 0x35, 0x99, 0xc1, 0xf8, 0x4c, 0x6e, 0x98, 0x16, 0xb2, 0xac, 0x97, 0x39, 0x9a,
	0x3d, 0x8a, 0x8b, 0x2a, 0x8b, 0x17, 0xa5, 0xfa, 0xaa, 0x8d, 0xd4, 0xa5, 0xab, 0x7f, 0x9e, 0x59,
	0xef, 0x7d, 0x9a, 0x56, 0x5c, 0x0f, 0xd7, 0x6b, 0x98, 0x29, 0x3e, 0x13, 0x85, 0xd0, 0x83, 0xcd,
	0x1a, 0x34, 0xfd, 0x81, 0xc0, 0x35, 0x2f, 0x25, 0xcf, 0xc0, 0x5d, 0xef, 0x14, 0xaf, 0x2b, 0xef,
	0xcd, 0xee, 0xd6, 0x95, 0x26, 0x30, 0xb0, 0x1d, 0xa9, 0x05, 0x12, 0xc1, 0x64, 0x29, 0x2f, 0x44,
	0x2a, 0xf8, 0xc5, 0x5a, 0x14, 0xbc, 0x6e, 0x74, 0x5a, 0x63, 0x90, 0xfc, 0xef, 0x83, 0x09, 0x85,
	0xd1, 0x72, 0xfe, 0xfa, 0x2d, 0xab, 0x2e, 0x03, 0x2f, 0x44, 0xd1, 0x9d, 0x36, 0xed, 0xe0, 0xf3,
	0x9f, 0x08, 0xc6, 0x5d, 0x3d, 0xb9, 0x0f, 0xc7, 0xdd, 0xf3, 0xc7, 0xf2, 0x4b, 0x29, 0xaf, 0x4b,
	0x6c, 0xf5, 0xa1, 0xf9, 0xbf, 0x82, 0xe5, 0x18, 0x11, 0xdc, 0xdc, 0xab, 0x81, 0xef, 0x14, 0xcf,
	0xb0, 0x4d, 0x8e, 0xe1, 0xa8, 0x23, 0xa7, 0x22, 0xc5, 0x4e, 0x1f, 0x9c, 0x97, 0x19, 0x76, 0xfb,
	0x33, 0x9f, 0xf8, 0x67, 0x85, 0xbd, 0xbe, 0xb2, 0x54, 0x2f, 0xb1, 0x3f, 0x04, 0xaf, 0xf0, 0x68,
	0x00, 0xe4, 0x15, 0x1e, 0x27, 0xc1, 0xcd, 0x9e, 0xa2, 0xdb, 0x3d, 0x45, 0x7f, 0xf6, 0x14, 0x7d,
	0x3f, 0x50, 0xeb, 0xf6, 0x40, 0xad, 0x5f, 0x07, 0x6a, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2b,
	0x67, 0x15, 0x85, 0x9e, 0x02, 0x00, 0x00,
}

func (m *FileSavePart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileSavePart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileSavePart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bytes != nil {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintChatApiFiles(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.TotalParts))
	i--
	dAtA[i] = 0x18
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.PartID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.FileID))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *FileGet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileGet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileGet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.Limit))
	i--
	dAtA[i] = 0x18
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.Offset))
	i--
	dAtA[i] = 0x10
	if m.Location == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("Location")
	} else {
		{
			size, err := m.Location.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChatApiFiles(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *File) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *File) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *File) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.MD5Hash)
	copy(dAtA[i:], m.MD5Hash)
	i = encodeVarintChatApiFiles(dAtA, i, uint64(len(m.MD5Hash)))
	i--
	dAtA[i] = 0x2a
	if m.Bytes != nil {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintChatApiFiles(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.ModifiedTime))
	i--
	dAtA[i] = 0x10
	i = encodeVarintChatApiFiles(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintChatApiFiles(dAtA []byte, offset int, v uint64) int {
	offset -= sovChatApiFiles(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileSavePart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovChatApiFiles(uint64(m.FileID))
	n += 1 + sovChatApiFiles(uint64(m.PartID))
	n += 1 + sovChatApiFiles(uint64(m.TotalParts))
	if m.Bytes != nil {
		l = len(m.Bytes)
		n += 1 + l + sovChatApiFiles(uint64(l))
	}
	return n
}

func (m *FileGet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Location != nil {
		l = m.Location.Size()
		n += 1 + l + sovChatApiFiles(uint64(l))
	}
	n += 1 + sovChatApiFiles(uint64(m.Offset))
	n += 1 + sovChatApiFiles(uint64(m.Limit))
	return n
}

func (m *File) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovChatApiFiles(uint64(m.Type))
	n += 1 + sovChatApiFiles(uint64(m.ModifiedTime))
	if m.Bytes != nil {
		l = len(m.Bytes)
		n += 1 + l + sovChatApiFiles(uint64(l))
	}
	l = len(m.MD5Hash)
	n += 1 + l + sovChatApiFiles(uint64(l))
	return n
}

func sovChatApiFiles(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChatApiFiles(x uint64) (n int) {
	return sovChatApiFiles(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileSavePart) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiFiles
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
			return fmt.Errorf("proto: FileSavePart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileSavePart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileID", wireType)
			}
			m.FileID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartID", wireType)
			}
			m.PartID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PartID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalParts", wireType)
			}
			m.TotalParts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalParts |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
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
				return ErrInvalidLengthChatApiFiles
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiFiles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("FileID")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("PartID")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("TotalParts")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bytes")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileGet) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiFiles
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
			return fmt.Errorf("proto: FileGet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileGet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
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
				return ErrInvalidLengthChatApiFiles
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Location == nil {
				m.Location = &InputFileLocation{}
			}
			if err := m.Location.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
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
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiFiles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Location")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Offset")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Limit")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *File) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChatApiFiles
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
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= FileType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedTime", wireType)
			}
			m.ModifiedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
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
				return ErrInvalidLengthChatApiFiles
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MD5Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChatApiFiles
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
				return ErrInvalidLengthChatApiFiles
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MD5Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChatApiFiles(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChatApiFiles
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChatApiFiles
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
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("ModifiedTime")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Bytes")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChatApiFiles(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChatApiFiles
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
					return 0, ErrIntOverflowChatApiFiles
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
					return 0, ErrIntOverflowChatApiFiles
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
				return 0, ErrInvalidLengthChatApiFiles
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthChatApiFiles
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChatApiFiles
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
				next, err := skipChatApiFiles(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthChatApiFiles
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
	ErrInvalidLengthChatApiFiles = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChatApiFiles   = fmt.Errorf("proto: integer overflow")
)
