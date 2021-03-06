// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyboard.proto

package laptop

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_AZWETY  Keyboard_Layout = 2
)

var Keyboard_Layout_name = map[int32]string{
	0: "UNKNOWN",
	1: "QWERTY",
	2: "AZWETY",
}

var Keyboard_Layout_value = map[string]int32{
	"UNKNOWN": 0,
	"QWERTY":  1,
	"AZWETY":  2,
}

func (x Keyboard_Layout) String() string {
	return proto.EnumName(Keyboard_Layout_name, int32(x))
}

func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c6f4f419d62a664, []int{0, 0}
}

type Keyboard struct {
	Layout               Keyboard_Layout `protobuf:"varint,1,opt,name=layout,proto3,enum=laptop.Keyboard_Layout" json:"layout,omitempty"`
	Backlit              bool            `protobuf:"varint,2,opt,name=backlit,proto3" json:"backlit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Keyboard) Reset()         { *m = Keyboard{} }
func (m *Keyboard) String() string { return proto.CompactTextString(m) }
func (*Keyboard) ProtoMessage()    {}
func (*Keyboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6f4f419d62a664, []int{0}
}

func (m *Keyboard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyboard.Unmarshal(m, b)
}
func (m *Keyboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyboard.Marshal(b, m, deterministic)
}
func (m *Keyboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyboard.Merge(m, src)
}
func (m *Keyboard) XXX_Size() int {
	return xxx_messageInfo_Keyboard.Size(m)
}
func (m *Keyboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyboard.DiscardUnknown(m)
}

var xxx_messageInfo_Keyboard proto.InternalMessageInfo

func (m *Keyboard) GetLayout() Keyboard_Layout {
	if m != nil {
		return m.Layout
	}
	return Keyboard_UNKNOWN
}

func (m *Keyboard) GetBacklit() bool {
	if m != nil {
		return m.Backlit
	}
	return false
}

func init() {
	proto.RegisterEnum("laptop.Keyboard_Layout", Keyboard_Layout_name, Keyboard_Layout_value)
	proto.RegisterType((*Keyboard)(nil), "laptop.Keyboard")
}

func init() {
	proto.RegisterFile("keyboard.proto", fileDescriptor_4c6f4f419d62a664)
}

var fileDescriptor_4c6f4f419d62a664 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4e, 0xad, 0x4c,
	0xca, 0x4f, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x49, 0x2c, 0x28,
	0xc9, 0x2f, 0x50, 0x6a, 0x61, 0xe4, 0xe2, 0xf0, 0x86, 0x4a, 0x09, 0xe9, 0x73, 0xb1, 0xe5, 0x24,
	0x56, 0xe6, 0x97, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x89, 0xeb, 0x41, 0x54, 0xe9,
	0xc1, 0x54, 0xe8, 0xf9, 0x80, 0xa5, 0x83, 0xa0, 0xca, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x12, 0x93,
	0xb3, 0x73, 0x32, 0x4b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0x25, 0x5d, 0x2e,
	0x36, 0x88, 0x5a, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x06,
	0x21, 0x2e, 0x2e, 0xb6, 0xc0, 0x70, 0xd7, 0xa0, 0x90, 0x48, 0x01, 0x46, 0x10, 0xdb, 0x31, 0x2a,
	0xdc, 0x35, 0x24, 0x52, 0x80, 0xc9, 0x89, 0x23, 0x0a, 0xea, 0xa0, 0x24, 0x36, 0xb0, 0xfb, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x61, 0x64, 0x43, 0xb1, 0x00, 0x00, 0x00,
}
