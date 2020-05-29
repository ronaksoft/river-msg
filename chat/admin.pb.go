// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin.proto

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

type PushProviderType int32

const (
	PushProviderType_PushProviderUnknown  PushProviderType = 0
	PushProviderType_PushProviderFirebase PushProviderType = 1
	PushProviderType_PushProviderApple    PushProviderType = 2
	PushProviderType_PushProviderPushKit  PushProviderType = 3
)

var PushProviderType_name = map[int32]string{
	0: "PushProviderUnknown",
	1: "PushProviderFirebase",
	2: "PushProviderApple",
	3: "PushProviderPushKit",
}

var PushProviderType_value = map[string]int32{
	"PushProviderUnknown":  0,
	"PushProviderFirebase": 1,
	"PushProviderApple":    2,
	"PushProviderPushKit":  3,
}

func (x PushProviderType) Enum() *PushProviderType {
	p := new(PushProviderType)
	*p = x
	return p
}

func (x PushProviderType) String() string {
	return proto.EnumName(PushProviderType_name, int32(x))
}

func (x *PushProviderType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PushProviderType_value, data, "PushProviderType")
	if err != nil {
		return err
	}
	*x = PushProviderType(value)
	return nil
}

func (PushProviderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}

// WelcomeMessage
type WelcomeMessage struct {
	Lang     string `protobuf:"bytes,2,req,name=Lang" json:"Lang"`
	Template string `protobuf:"bytes,3,req,name=Template" json:"Template"`
}

func (m *WelcomeMessage) Reset()         { *m = WelcomeMessage{} }
func (m *WelcomeMessage) String() string { return proto.CompactTextString(m) }
func (*WelcomeMessage) ProtoMessage()    {}
func (*WelcomeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}
func (m *WelcomeMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WelcomeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WelcomeMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WelcomeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeMessage.Merge(m, src)
}
func (m *WelcomeMessage) XXX_Size() int {
	return m.Size()
}
func (m *WelcomeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeMessage proto.InternalMessageInfo

func (m *WelcomeMessage) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *WelcomeMessage) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

// PushProvider
type PushProvider struct {
	Name        string           `protobuf:"bytes,1,req,name=Name" json:"Name"`
	Type        PushProviderType `protobuf:"varint,2,req,name=Type,enum=msg.PushProviderType" json:"Type"`
	TestMode    bool             `protobuf:"varint,3,req,name=TestMode" json:"TestMode"`
	Credentials []byte           `protobuf:"bytes,4,opt,name=Credentials" json:"Credentials"`
	KeyID       string           `protobuf:"bytes,5,opt,name=KeyID" json:"KeyID"`
	TeamID      string           `protobuf:"bytes,6,opt,name=TeamID" json:"TeamID"`
	Topic       string           `protobuf:"bytes,7,opt,name=Topic" json:"Topic"`
}

func (m *PushProvider) Reset()         { *m = PushProvider{} }
func (m *PushProvider) String() string { return proto.CompactTextString(m) }
func (*PushProvider) ProtoMessage()    {}
func (*PushProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{1}
}
func (m *PushProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PushProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PushProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PushProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushProvider.Merge(m, src)
}
func (m *PushProvider) XXX_Size() int {
	return m.Size()
}
func (m *PushProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_PushProvider.DiscardUnknown(m)
}

var xxx_messageInfo_PushProvider proto.InternalMessageInfo

func (m *PushProvider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PushProvider) GetType() PushProviderType {
	if m != nil {
		return m.Type
	}
	return PushProviderType_PushProviderUnknown
}

func (m *PushProvider) GetTestMode() bool {
	if m != nil {
		return m.TestMode
	}
	return false
}

func (m *PushProvider) GetCredentials() []byte {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *PushProvider) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *PushProvider) GetTeamID() string {
	if m != nil {
		return m.TeamID
	}
	return ""
}

func (m *PushProvider) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// AppUpdate
type Version struct {
	Vendor         string   `protobuf:"bytes,1,req,name=Vendor" json:"Vendor"`
	Stage          string   `protobuf:"bytes,2,req,name=Stage" json:"Stage"`
	OS             string   `protobuf:"bytes,3,req,name=OS" json:"OS"`
	MinVersion     string   `protobuf:"bytes,4,req,name=MinVersion" json:"MinVersion"`
	CurrentVersion string   `protobuf:"bytes,5,req,name=CurrentVersion" json:"CurrentVersion"`
	ForcedVersions []string `protobuf:"bytes,6,rep,name=ForcedVersions" json:"ForcedVersions,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{2}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Version) GetStage() string {
	if m != nil {
		return m.Stage
	}
	return ""
}

func (m *Version) GetOS() string {
	if m != nil {
		return m.OS
	}
	return ""
}

func (m *Version) GetMinVersion() string {
	if m != nil {
		return m.MinVersion
	}
	return ""
}

func (m *Version) GetCurrentVersion() string {
	if m != nil {
		return m.CurrentVersion
	}
	return ""
}

func (m *Version) GetForcedVersions() []string {
	if m != nil {
		return m.ForcedVersions
	}
	return nil
}

func init() {
	proto.RegisterEnum("msg.PushProviderType", PushProviderType_name, PushProviderType_value)
	proto.RegisterType((*WelcomeMessage)(nil), "msg.WelcomeMessage")
	proto.RegisterType((*PushProvider)(nil), "msg.PushProvider")
	proto.RegisterType((*Version)(nil), "msg.Version")
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_73a7fc70dcc2027c) }

var fileDescriptor_73a7fc70dcc2027c = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x33, 0x49, 0x76, 0xdb, 0x9e, 0x96, 0x10, 0xc7, 0x16, 0x07, 0x91, 0x18, 0x16, 0x29,
	0x8b, 0xc8, 0x0a, 0xbe, 0x81, 0x6d, 0x29, 0x94, 0x76, 0x6d, 0xe9, 0xc6, 0x7a, 0x3d, 0x6e, 0x0e,
	0x71, 0x30, 0x99, 0x09, 0x33, 0x59, 0xa5, 0x6f, 0xe1, 0x63, 0xf5, 0xb2, 0x37, 0x82, 0x57, 0x22,
	0xbb, 0x4f, 0xe1, 0x9d, 0xcc, 0x26, 0xad, 0xb3, 0xb9, 0x1b, 0xbe, 0xef, 0x9c, 0xff, 0x4f, 0x66,
	0x60, 0x97, 0xe7, 0x95, 0x90, 0x93, 0x5a, 0xab, 0x46, 0xd1, 0xa0, 0x32, 0xc5, 0xe8, 0x02, 0xa2,
	0x4f, 0x58, 0xce, 0x55, 0x85, 0x53, 0x34, 0x86, 0x17, 0x48, 0x19, 0x84, 0x17, 0x5c, 0x16, 0xcc,
	0x4f, 0xfd, 0xf1, 0xce, 0x51, 0x78, 0xf7, 0xfb, 0xa5, 0x77, 0xbd, 0x26, 0x34, 0x85, 0xed, 0x0c,
	0xab, 0xba, 0xe4, 0x0d, 0xb2, 0xc0, 0xb1, 0x8f, 0x74, 0xf4, 0x97, 0xc0, 0xde, 0xd5, 0xc2, 0x7c,
	0xb9, 0xd2, 0xea, 0x9b, 0xc8, 0x51, 0xdb, 0xb0, 0x0f, 0xbc, 0x42, 0x46, 0xdc, 0x30, 0x4b, 0xe8,
	0x5b, 0x08, 0xb3, 0xdb, 0x1a, 0xd7, 0x35, 0xd1, 0xbb, 0x83, 0x49, 0x65, 0x8a, 0x89, 0xbb, 0x6a,
	0xe5, 0xc3, 0x82, 0x3d, 0xb7, 0xed, 0xa6, 0x99, 0xaa, 0xbc, 0x6d, 0xdf, 0xfe, 0xdf, 0xde, 0x52,
	0x7a, 0x08, 0xbb, 0xc7, 0x1a, 0x73, 0x94, 0x8d, 0xe0, 0xa5, 0x61, 0x61, 0x4a, 0xc6, 0x7b, 0xdd,
	0x90, 0x2b, 0xe8, 0x73, 0x18, 0x9c, 0xe3, 0xed, 0xd9, 0x09, 0x1b, 0xa4, 0xe4, 0xf1, 0xab, 0x5a,
	0x44, 0x5f, 0xc0, 0x30, 0x43, 0x5e, 0x9d, 0x9d, 0xb0, 0xa1, 0x23, 0x3b, 0x66, 0x37, 0x33, 0x55,
	0x8b, 0x39, 0xdb, 0x72, 0x37, 0xd7, 0x68, 0xf4, 0x93, 0xc0, 0xd6, 0x0d, 0x6a, 0x23, 0x94, 0xb4,
	0x29, 0x37, 0x28, 0x73, 0xa5, 0x37, 0x7e, 0xbc, 0x63, 0x36, 0x65, 0xd6, 0xf0, 0x02, 0x37, 0xae,
	0xb8, 0x45, 0x74, 0x1f, 0xfc, 0xcb, 0xd9, 0xc6, 0xed, 0xfa, 0x97, 0x33, 0xfa, 0x0a, 0x60, 0x2a,
	0x64, 0x97, 0xce, 0x42, 0xc7, 0x3a, 0x9c, 0xbe, 0x81, 0xe8, 0x78, 0xa1, 0x35, 0xca, 0xe6, 0x61,
	0x72, 0xe0, 0x4c, 0xf6, 0x1c, 0x3d, 0x84, 0xe8, 0x54, 0xe9, 0x39, 0xe6, 0x1d, 0x30, 0x6c, 0x98,
	0x06, 0xe3, 0x9d, 0xeb, 0x1e, 0x7d, 0xbd, 0x80, 0xb8, 0xff, 0x2e, 0xf4, 0x19, 0x3c, 0x75, 0xd9,
	0x47, 0xf9, 0x55, 0xaa, 0xef, 0x32, 0xf6, 0x28, 0x83, 0x7d, 0x57, 0x9c, 0x0a, 0x8d, 0x9f, 0xb9,
	0xc1, 0x98, 0xd0, 0x03, 0x78, 0xe2, 0x9a, 0xf7, 0x75, 0x5d, 0x62, 0xec, 0xf7, 0x93, 0xec, 0xf9,
	0x5c, 0x34, 0x71, 0x70, 0xc4, 0xee, 0x96, 0x09, 0xb9, 0x5f, 0x26, 0xe4, 0xcf, 0x32, 0x21, 0x3f,
	0x56, 0x89, 0x77, 0xbf, 0x4a, 0xbc, 0x5f, 0xab, 0xc4, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x72,
	0xa5, 0xcc, 0x31, 0xc5, 0x02, 0x00, 0x00,
}

func (m *WelcomeMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WelcomeMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WelcomeMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Template)
	copy(dAtA[i:], m.Template)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Template)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Lang)
	copy(dAtA[i:], m.Lang)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Lang)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *PushProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PushProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PushProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Topic)
	copy(dAtA[i:], m.Topic)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Topic)))
	i--
	dAtA[i] = 0x3a
	i -= len(m.TeamID)
	copy(dAtA[i:], m.TeamID)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.TeamID)))
	i--
	dAtA[i] = 0x32
	i -= len(m.KeyID)
	copy(dAtA[i:], m.KeyID)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.KeyID)))
	i--
	dAtA[i] = 0x2a
	if m.Credentials != nil {
		i -= len(m.Credentials)
		copy(dAtA[i:], m.Credentials)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.Credentials)))
		i--
		dAtA[i] = 0x22
	}
	i--
	if m.TestMode {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i = encodeVarintAdmin(dAtA, i, uint64(m.Type))
	i--
	dAtA[i] = 0x10
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Version) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ForcedVersions) > 0 {
		for iNdEx := len(m.ForcedVersions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ForcedVersions[iNdEx])
			copy(dAtA[i:], m.ForcedVersions[iNdEx])
			i = encodeVarintAdmin(dAtA, i, uint64(len(m.ForcedVersions[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	i -= len(m.CurrentVersion)
	copy(dAtA[i:], m.CurrentVersion)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.CurrentVersion)))
	i--
	dAtA[i] = 0x2a
	i -= len(m.MinVersion)
	copy(dAtA[i:], m.MinVersion)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.MinVersion)))
	i--
	dAtA[i] = 0x22
	i -= len(m.OS)
	copy(dAtA[i:], m.OS)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.OS)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Stage)
	copy(dAtA[i:], m.Stage)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Stage)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Vendor)
	copy(dAtA[i:], m.Vendor)
	i = encodeVarintAdmin(dAtA, i, uint64(len(m.Vendor)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WelcomeMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Lang)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.Template)
	n += 1 + l + sovAdmin(uint64(l))
	return n
}

func (m *PushProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovAdmin(uint64(l))
	n += 1 + sovAdmin(uint64(m.Type))
	n += 2
	if m.Credentials != nil {
		l = len(m.Credentials)
		n += 1 + l + sovAdmin(uint64(l))
	}
	l = len(m.KeyID)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.TeamID)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.Topic)
	n += 1 + l + sovAdmin(uint64(l))
	return n
}

func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Vendor)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.Stage)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.OS)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.MinVersion)
	n += 1 + l + sovAdmin(uint64(l))
	l = len(m.CurrentVersion)
	n += 1 + l + sovAdmin(uint64(l))
	if len(m.ForcedVersions) > 0 {
		for _, s := range m.ForcedVersions {
			l = len(s)
			n += 1 + l + sovAdmin(uint64(l))
		}
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WelcomeMessage) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: WelcomeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WelcomeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lang", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lang = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Template = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Lang")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Template")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PushProvider) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: PushProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PushProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PushProviderType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestMode", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TestMode = bool(v != 0)
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = append(m.Credentials[:0], dAtA[iNdEx:postIndex]...)
			if m.Credentials == nil {
				m.Credentials = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Type")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("TestMode")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Version) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForcedVersions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForcedVersions = append(m.ForcedVersions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdmin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Vendor")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("Stage")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("OS")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("MinVersion")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("CurrentVersion")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)