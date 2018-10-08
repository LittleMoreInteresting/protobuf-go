// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nopackage/nopackage.proto

package nopackage

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Enum int32

const (
	Enum_ZERO Enum = 0
)

var Enum_name = map[int32]string{
	0: "ZERO",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (x *Enum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum_value, data, "Enum")
	if err != nil {
		return err
	}
	*x = Enum(value)
	return nil
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f33a1d5d178c43c9, []int{0}
}

type Message struct {
	StringField          *string  `protobuf:"bytes,1,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	EnumField            *Enum    `protobuf:"varint,2,opt,name=enum_field,json=enumField,enum=Enum" json:"enum_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f33a1d5d178c43c9, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetStringField() string {
	if m != nil && m.StringField != nil {
		return *m.StringField
	}
	return ""
}

func (m *Message) GetEnumField() Enum {
	if m != nil && m.EnumField != nil {
		return *m.EnumField
	}
	return Enum_ZERO
}

func init() {
	proto.RegisterEnum("Enum", Enum_name, Enum_value)
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("nopackage/nopackage.proto", fileDescriptor_f33a1d5d178c43c9) }

var fileDescriptor_f33a1d5d178c43c9 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcb, 0x2f, 0x48,
	0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0xd5, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x95, 0x82,
	0xb8, 0xd8, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x14, 0xb9, 0x78, 0x8a, 0x4b, 0x8a,
	0x32, 0xf3, 0xd2, 0xe3, 0xd3, 0x32, 0x53, 0x73, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xb8, 0x21, 0x62, 0x6e, 0x20, 0x21, 0x21, 0x15, 0x2e, 0xae, 0xd4, 0xbc, 0xd2, 0x5c, 0xa8, 0x02,
	0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x56, 0x3d, 0xd7, 0xbc, 0xd2, 0xdc, 0x20, 0x4e, 0x90, 0x04,
	0x58, 0x95, 0x96, 0x00, 0x17, 0x0b, 0x48, 0x48, 0x88, 0x83, 0x8b, 0x25, 0xca, 0x35, 0xc8, 0x5f,
	0x80, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xbc, 0x78, 0x1f, 0x81, 0x00, 0x00, 0x00,
}
