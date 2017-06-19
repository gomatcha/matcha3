// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/view/stacknav/stacknavigator.proto

/*
Package stacknav is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/view/stacknav/stacknavigator.proto

It has these top-level messages:
	Screen
	StackNav
	StackEvent
*/
package stacknav

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

type Screen struct {
	Id                    int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title                 string  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	CustomBackButtonTitle bool    `protobuf:"varint,8,opt,name=customBackButtonTitle" json:"customBackButtonTitle,omitempty"`
	BackButtonTitle       string  `protobuf:"bytes,3,opt,name=backButtonTitle" json:"backButtonTitle,omitempty"`
	BackButtonHidden      bool    `protobuf:"varint,4,opt,name=backButtonHidden" json:"backButtonHidden,omitempty"`
	TitleViewId           int64   `protobuf:"varint,5,opt,name=titleViewId" json:"titleViewId,omitempty"`
	RightViewIds          []int64 `protobuf:"varint,6,rep,packed,name=rightViewIds" json:"rightViewIds,omitempty"`
	LeftViewIds           []int64 `protobuf:"varint,7,rep,packed,name=leftViewIds" json:"leftViewIds,omitempty"`
}

func (m *Screen) Reset()                    { *m = Screen{} }
func (m *Screen) String() string            { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()               {}
func (*Screen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Screen) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Screen) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Screen) GetCustomBackButtonTitle() bool {
	if m != nil {
		return m.CustomBackButtonTitle
	}
	return false
}

func (m *Screen) GetBackButtonTitle() string {
	if m != nil {
		return m.BackButtonTitle
	}
	return ""
}

func (m *Screen) GetBackButtonHidden() bool {
	if m != nil {
		return m.BackButtonHidden
	}
	return false
}

func (m *Screen) GetTitleViewId() int64 {
	if m != nil {
		return m.TitleViewId
	}
	return 0
}

func (m *Screen) GetRightViewIds() []int64 {
	if m != nil {
		return m.RightViewIds
	}
	return nil
}

func (m *Screen) GetLeftViewIds() []int64 {
	if m != nil {
		return m.LeftViewIds
	}
	return nil
}

type StackNav struct {
	Screens []*Screen `protobuf:"bytes,1,rep,name=screens" json:"screens,omitempty"`
}

func (m *StackNav) Reset()                    { *m = StackNav{} }
func (m *StackNav) String() string            { return proto.CompactTextString(m) }
func (*StackNav) ProtoMessage()               {}
func (*StackNav) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StackNav) GetScreens() []*Screen {
	if m != nil {
		return m.Screens
	}
	return nil
}

type StackEvent struct {
	Id []int64 `protobuf:"varint,1,rep,packed,name=id" json:"id,omitempty"`
}

func (m *StackEvent) Reset()                    { *m = StackEvent{} }
func (m *StackEvent) String() string            { return proto.CompactTextString(m) }
func (*StackEvent) ProtoMessage()               {}
func (*StackEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StackEvent) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*Screen)(nil), "stacknav.Screen")
	proto.RegisterType((*StackNav)(nil), "stacknav.StackNav")
	proto.RegisterType((*StackEvent)(nil), "stacknav.StackEvent")
}

func init() {
	proto.RegisterFile("github.com/overcyn/matcha/pb/view/stacknav/stacknavigator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xeb, 0xb6, 0xfa, 0x26, 0x3a, 0x82, 0x42, 0x0e, 0x0a, 0xa1, 0xa7, 0xb0, 0x43,
	0x0b, 0x2a, 0x5e, 0x85, 0x82, 0xa0, 0x07, 0x45, 0x3a, 0xf1, 0xe0, 0x2d, 0x4d, 0xe3, 0x16, 0xb6,
	0x25, 0xa3, 0x7d, 0xeb, 0xf0, 0xeb, 0xec, 0x93, 0xca, 0x52, 0x5a, 0xe7, 0xf4, 0x96, 0xf7, 0xcb,
	0x3f, 0x2f, 0x8f, 0xdf, 0x83, 0xfb, 0xa9, 0xc6, 0xd9, 0x3a, 0x8f, 0xa5, 0x5d, 0x26, 0xb6, 0x56,
	0xa5, 0xfc, 0x32, 0xc9, 0x52, 0xa0, 0x9c, 0x89, 0x64, 0x95, 0x27, 0xb5, 0x56, 0x9b, 0xa4, 0x42,
	0x21, 0xe7, 0x46, 0xd4, 0xdd, 0x41, 0x4f, 0x05, 0xda, 0x32, 0x5e, 0x95, 0x16, 0x2d, 0x09, 0x5b,
	0x1a, 0x6d, 0x7d, 0xe8, 0x4f, 0x64, 0xa9, 0x94, 0x21, 0xa7, 0xe0, 0xeb, 0x82, 0x7a, 0xcc, 0xe3,
	0x41, 0xe6, 0xeb, 0x82, 0x9c, 0x43, 0x0f, 0x35, 0x2e, 0x14, 0xf5, 0x99, 0xc7, 0x8f, 0xb3, 0xa6,
	0x20, 0xb7, 0x70, 0x21, 0xd7, 0x15, 0xda, 0x65, 0x2a, 0xe4, 0x3c, 0x5d, 0x23, 0x5a, 0xf3, 0xe6,
	0x52, 0x21, 0xf3, 0x78, 0x98, 0xfd, 0x7f, 0x49, 0x38, 0x9c, 0xe5, 0x07, 0xf9, 0xc0, 0x75, 0x3d,
	0xc4, 0x64, 0x0c, 0xa3, 0x1f, 0xf4, 0xa8, 0x8b, 0x42, 0x19, 0x7a, 0xe4, 0x5a, 0xff, 0xe1, 0x84,
	0xc1, 0xd0, 0x0d, 0xf5, 0xae, 0xd5, 0xe6, 0xa9, 0xa0, 0x3d, 0x37, 0xfa, 0x3e, 0x22, 0x11, 0x9c,
	0x94, 0x7a, 0x3a, 0xc3, 0xa6, 0xac, 0x68, 0x9f, 0x05, 0x3c, 0xc8, 0x7e, 0xb1, 0x5d, 0x97, 0x85,
	0xfa, 0xec, 0x22, 0x03, 0x17, 0xd9, 0x47, 0xd1, 0x1d, 0x84, 0x93, 0x9d, 0xb0, 0x17, 0x51, 0x93,
	0x31, 0x0c, 0x2a, 0xe7, 0xab, 0xa2, 0x1e, 0x0b, 0xf8, 0xf0, 0x7a, 0x14, 0xb7, 0x32, 0xe3, 0x46,
	0x64, 0xd6, 0x06, 0xa2, 0x4b, 0x00, 0xf7, 0xee, 0xa1, 0x56, 0x06, 0x3b, 0xbf, 0x41, 0xe3, 0x37,
	0xbd, 0xfa, 0xe8, 0xd6, 0xb0, 0xf5, 0x47, 0xcf, 0x6e, 0x7f, 0xaf, 0x69, 0xfb, 0x51, 0xde, 0x77,
	0xab, 0xba, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xde, 0x9f, 0x4d, 0x4e, 0xed, 0x01, 0x00, 0x00,
}
