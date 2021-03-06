// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blank.proto

package blank

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age                  uint64   `protobuf:"fixed64,3,opt,name=age" json:"age,omitempty"`
	Contact              *Contact `protobuf:"bytes,4,opt,name=contact" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_blank_f535a0a5eca36fca, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() uint64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type Contact struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_blank_f535a0a5eca36fca, []int{1}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (dst *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(dst, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "blank.User")
	proto.RegisterType((*Contact)(nil), "blank.Contact")
}

func init() { proto.RegisterFile("blank.proto", fileDescriptor_blank_f535a0a5eca36fca) }

var fileDescriptor_blank_f535a0a5eca36fca = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x41, 0x0e, 0x83, 0x20,
	0x10, 0x45, 0x03, 0x82, 0xc6, 0x31, 0x31, 0xcd, 0xa4, 0x0b, 0x96, 0xc4, 0x15, 0x2b, 0x17, 0x6d,
	0x7a, 0x82, 0xde, 0x80, 0xa4, 0x07, 0x40, 0xa5, 0x2d, 0xa9, 0x82, 0xb1, 0xdc, 0x3f, 0x8d, 0x60,
	0x77, 0xff, 0xbd, 0x59, 0xbc, 0x81, 0x66, 0x98, 0x8d, 0xff, 0xf4, 0xeb, 0x16, 0x62, 0x40, 0x9e,
	0xa0, 0x7b, 0x02, 0x7b, 0x7c, 0xed, 0x86, 0x2d, 0x50, 0x37, 0x09, 0x22, 0x89, 0xe2, 0x9a, 0xba,
	0x09, 0x11, 0x98, 0x37, 0x8b, 0x15, 0x54, 0x12, 0x55, 0xeb, 0xb4, 0xf1, 0x04, 0x85, 0x79, 0x59,
	0x51, 0x48, 0xa2, 0x4a, 0xbd, 0x4f, 0x54, 0x50, 0x8d, 0xc1, 0x47, 0x33, 0x46, 0xc1, 0x24, 0x51,
	0xcd, 0xa5, 0xed, 0x73, 0xe3, 0x9e, 0xad, 0xfe, 0x9f, 0xbb, 0x1b, 0x54, 0x87, 0xc3, 0x33, 0xf0,
	0xf5, 0x1d, 0xbc, 0x4d, 0xb5, 0x5a, 0x67, 0xd8, 0xad, 0x5d, 0x8c, 0x9b, 0x8f, 0x62, 0x86, 0xa1,
	0x4c, 0xcf, 0x5e, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x54, 0x6a, 0x8a, 0xbb, 0x00, 0x00,
	0x00,
}
