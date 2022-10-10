// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo_service.proto

package my_grpc

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

type GreetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9346b74f56d243b9, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetReply struct {
	Greet                string   `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetReply) Reset()         { *m = GreetReply{} }
func (m *GreetReply) String() string { return proto.CompactTextString(m) }
func (*GreetReply) ProtoMessage()    {}
func (*GreetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9346b74f56d243b9, []int{1}
}

func (m *GreetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetReply.Unmarshal(m, b)
}
func (m *GreetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetReply.Marshal(b, m, deterministic)
}
func (m *GreetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetReply.Merge(m, src)
}
func (m *GreetReply) XXX_Size() int {
	return xxx_messageInfo_GreetReply.Size(m)
}
func (m *GreetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GreetReply proto.InternalMessageInfo

func (m *GreetReply) GetGreet() string {
	if m != nil {
		return m.Greet
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "GreetRequest")
	proto.RegisterType((*GreetReply)(nil), "GreetReply")
}

func init() {
	proto.RegisterFile("demo_service.proto", fileDescriptor_9346b74f56d243b9)
}

var fileDescriptor_9346b74f56d243b9 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x49, 0xcd, 0xcd,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52,
	0xe2, 0xe2, 0x71, 0x2f, 0x4a, 0x4d, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95,
	0x94, 0xb8, 0xb8, 0xa0, 0x6a, 0x0a, 0x72, 0x2a, 0x85, 0x44, 0xb8, 0x58, 0xd3, 0x41, 0x3c, 0xa8,
	0x12, 0x08, 0xc7, 0xc8, 0x92, 0x8b, 0xd3, 0xb7, 0x12, 0xac, 0x2a, 0xb5, 0x48, 0x48, 0x07, 0x6a,
	0xa8, 0x67, 0x5e, 0x72, 0x7e, 0x6e, 0x6a, 0x91, 0x10, 0xaf, 0x1e, 0xb2, 0x1d, 0x52, 0xdc, 0x7a,
	0x08, 0xe3, 0x94, 0x18, 0x9c, 0x64, 0xa3, 0xa4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xd3, 0x53, 0x73, 0x0b, 0x52, 0x73, 0x72, 0x72, 0xf5, 0x73, 0x2b, 0x75, 0xd3,
	0x8b, 0x0a, 0x92, 0x93, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x31,
	0x89, 0x52, 0xbe, 0x00, 0x00, 0x00,
}