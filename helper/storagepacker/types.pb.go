// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/storagepacker/types.proto

package storagepacker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Item struct {
	ID                   string   `sentinel:"" protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              *any.Any `sentinel:"" protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Item) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

type Bucket struct {
	Key                  string   `sentinel:"" protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Items                []*Item  `sentinel:"" protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Bucket) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "storagepacker.Item")
	proto.RegisterType((*Bucket)(nil), "storagepacker.Bucket")
}

func init() { proto.RegisterFile("helper/storagepacker/types.proto", fileDescriptor_c0e98c66c4f51b7f) }

var fileDescriptor_c0e98c66c4f51b7f = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0xaa, 0x15, 0xb7, 0x28, 0xb2, 0x7a, 0x88, 0x9e, 0x42, 0x4f, 0xf1, 0x32, 0x83,
	0xf5, 0x17, 0x58, 0x50, 0xf0, 0x9a, 0xa3, 0xb7, 0x4d, 0x3a, 0x6e, 0x96, 0x64, 0xbb, 0xcb, 0xee,
	0xac, 0xb0, 0xff, 0x5e, 0xda, 0xd8, 0x43, 0xc1, 0xdb, 0xc0, 0xfb, 0xf8, 0xe6, 0x3d, 0x51, 0x0f,
	0x34, 0x79, 0x0a, 0x18, 0xd9, 0x05, 0xa5, 0xc9, 0xab, 0x7e, 0xa4, 0x80, 0x9c, 0x3d, 0x45, 0xf0,
	0xc1, 0xb1, 0x93, 0x37, 0x67, 0xd1, 0xd3, 0xa3, 0x76, 0x4e, 0x4f, 0x84, 0xc7, 0xb0, 0x4b, 0xdf,
	0xa8, 0xf6, 0x79, 0x26, 0xd7, 0x1f, 0xe2, 0xe2, 0x93, 0xc9, 0xca, 0x5b, 0x51, 0x9a, 0x5d, 0x55,
	0xd4, 0x45, 0x73, 0xdd, 0x96, 0x66, 0x27, 0x41, 0x5c, 0x59, 0x8a, 0x51, 0x69, 0xaa, 0xca, 0xba,
	0x68, 0x56, 0x9b, 0x07, 0x98, 0x25, 0x70, 0x92, 0xc0, 0xdb, 0x3e, 0xb7, 0x27, 0x68, 0xfd, 0x2e,
	0x96, 0xdb, 0xd4, 0x8f, 0xc4, 0xf2, 0x4e, 0x2c, 0x46, 0xca, 0x7f, 0xaa, 0xc3, 0x29, 0x9f, 0xc5,
	0xa5, 0x61, 0xb2, 0xb1, 0x2a, 0xeb, 0x45, 0xb3, 0xda, 0xdc, 0xc3, 0x59, 0x3b, 0x38, 0xfc, 0x6f,
	0x67, 0x62, 0xfb, 0xf2, 0x85, 0xda, 0xf0, 0x90, 0x3a, 0xe8, 0x9d, 0xc5, 0x41, 0xc5, 0xc1, 0xf4,
	0x2e, 0x78, 0xfc, 0x51, 0x69, 0x62, 0xfc, 0x6f, 0x77, 0xb7, 0x3c, 0x16, 0x7a, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0x8e, 0xb4, 0xa9, 0x16, 0x01, 0x00, 0x00,
}
