// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/app/statusbar.proto

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/pb/app/statusbar.proto

It has these top-level messages:
	ActivityIndicator
*/
package app

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

type ActivityIndicator struct {
	Visible bool `protobuf:"varint,1,opt,name=visible" json:"visible,omitempty"`
}

func (m *ActivityIndicator) Reset()                    { *m = ActivityIndicator{} }
func (m *ActivityIndicator) String() string            { return proto.CompactTextString(m) }
func (*ActivityIndicator) ProtoMessage()               {}
func (*ActivityIndicator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivityIndicator) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

func init() {
	proto.RegisterType((*ActivityIndicator)(nil), "app.ActivityIndicator")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/app/statusbar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x92, 0xf4, 0x13, 0x0b,
	0x0a, 0xf4, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x93, 0x12, 0x8b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x98, 0x13, 0x0b, 0x0a, 0x94, 0x74, 0xb9, 0x04, 0x1d, 0x93, 0x4b, 0x32, 0xcb, 0x32,
	0x4b, 0x2a, 0x3d, 0xf3, 0x52, 0x32, 0x93, 0x13, 0x4b, 0xf2, 0x8b, 0x84, 0x24, 0xb8, 0xd8, 0xcb,
	0x32, 0x8b, 0x33, 0x93, 0x72, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0x27,
	0xe1, 0x28, 0x90, 0xae, 0x45, 0x4c, 0xdc, 0xbe, 0x60, 0xa3, 0x1d, 0x0b, 0x0a, 0x02, 0x9c, 0x92,
	0xd8, 0xc0, 0xe6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xe5, 0xfa, 0x6d, 0x7c, 0x00,
	0x00, 0x00,
}
