// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressBook.proto

package adressBook

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_756294b5e39fe676, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_756294b5e39fe676, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=adressBook.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_756294b5e39fe676, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_756294b5e39fe676, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("adressBook.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "adressBook.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "adressBook.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "adressBook.AddressBook")
}

func init() { proto.RegisterFile("addressBook.proto", fileDescriptor_756294b5e39fe676) }

var fileDescriptor_756294b5e39fe676 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x9a, 0x06, 0x3b, 0x91, 0x52, 0x07, 0x91, 0x50, 0x44, 0x43, 0x4f, 0x41, 0x61,
	0x2b, 0x15, 0x04, 0x0f, 0x1e, 0x2c, 0x14, 0x14, 0xad, 0x2d, 0x4b, 0xa5, 0x47, 0xd9, 0x92, 0xb1,
	0x06, 0x93, 0xec, 0x92, 0x6c, 0x0e, 0xfd, 0x75, 0xfe, 0x35, 0xc9, 0x26, 0x6d, 0x05, 0xf1, 0xf6,
	0x66, 0xf6, 0x63, 0xf6, 0xbd, 0x19, 0x38, 0x16, 0x51, 0x94, 0x53, 0x51, 0x8c, 0xa5, 0xfc, 0x62,
	0x2a, 0x97, 0x5a, 0x22, 0x88, 0x5d, 0xa7, 0x7f, 0xb1, 0x96, 0x72, 0x9d, 0xd0, 0xd0, 0xbc, 0xac,
	0xca, 0x8f, 0xa1, 0x8e, 0x53, 0x2a, 0xb4, 0x48, 0x55, 0x0d, 0x0f, 0xbe, 0x6d, 0x70, 0xe7, 0x94,
	0x17, 0x32, 0x43, 0x04, 0x27, 0x13, 0x29, 0xf9, 0x56, 0x60, 0x85, 0x1d, 0x6e, 0x34, 0x76, 0xc1,
	0x8e, 0x23, 0xdf, 0x0e, 0xac, 0xb0, 0xcd, 0xed, 0x38, 0xc2, 0x13, 0x68, 0x53, 0x2a, 0xe2, 0xc4,
	0x6f, 0x19, 0xa8, 0x2e, 0xf0, 0x16, 0x5c, 0xf5, 0x29, 0x33, 0x2a, 0x7c, 0x27, 0x68, 0x85, 0xde,
	0xe8, 0x9c, 0xed, 0x2d, 0xb0, 0x7a, 0x3a, 0x9b, 0x57, 0xc0, 0x6b, 0x99, 0xae, 0x28, 0xe7, 0x0d,
	0x8d, 0xf7, 0x70, 0x94, 0x88, 0x42, 0xbf, 0x97, 0x2a, 0x12, 0x9a, 0x22, 0xbf, 0x1d, 0x58, 0xa1,
	0x37, 0xea, 0xb3, 0xda, 0x34, 0xdb, 0x9a, 0x66, 0x8b, 0xad, 0x69, 0xee, 0x55, 0xfc, 0x5b, 0x8d,
	0xf7, 0x97, 0xe0, 0xfd, 0x9a, 0x8a, 0xa7, 0xe0, 0x66, 0x46, 0x35, 0x09, 0x9a, 0x0a, 0xaf, 0xc1,
	0xd1, 0x1b, 0x45, 0x26, 0x45, 0x77, 0x74, 0xf6, 0x9f, 0xb7, 0xc5, 0x46, 0x11, 0x37, 0xe4, 0xe0,
	0x0a, 0x3a, 0xbb, 0x16, 0x02, 0xb8, 0xd3, 0xd9, 0xf8, 0xe9, 0x65, 0xd2, 0x3b, 0xc0, 0x43, 0x70,
	0x1e, 0x67, 0xd3, 0x49, 0xcf, 0xaa, 0xd4, 0x72, 0xc6, 0x9f, 0x7b, 0xf6, 0xe0, 0x0e, 0xbc, 0x87,
	0xfd, 0x0d, 0xf0, 0x12, 0x5c, 0x45, 0x52, 0x25, 0xd5, 0x1e, 0xab, 0x5d, 0xe0, 0xdf, 0xff, 0x78,
	0x43, 0xac, 0x5c, 0x93, 0xf0, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x23, 0x24, 0xc2, 0xc5,
	0x01, 0x00, 0x00,
}
