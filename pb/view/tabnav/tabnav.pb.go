// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/view/tabnav/tabnav.proto

/*
Package tabnav is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/mochi/pb/view/tabnav/tabnav.proto

It has these top-level messages:
	Screen
	TabNav
	Event
*/
package tabnav

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import imageview "github.com/overcyn/mochi/pb"

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
	Id           int64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title        string           `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Icon         *imageview.Image `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	SelectedIcon *imageview.Image `protobuf:"bytes,4,opt,name=selectedIcon" json:"selectedIcon,omitempty"`
	Badge        string           `protobuf:"bytes,5,opt,name=badge" json:"badge,omitempty"`
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

func (m *Screen) GetIcon() *imageview.Image {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *Screen) GetSelectedIcon() *imageview.Image {
	if m != nil {
		return m.SelectedIcon
	}
	return nil
}

func (m *Screen) GetBadge() string {
	if m != nil {
		return m.Badge
	}
	return ""
}

type TabNav struct {
	Screens       []*Screen `protobuf:"bytes,1,rep,name=screens" json:"screens,omitempty"`
	SelectedIndex int64     `protobuf:"varint,2,opt,name=selectedIndex" json:"selectedIndex,omitempty"`
	EventFunc     int64     `protobuf:"varint,3,opt,name=eventFunc" json:"eventFunc,omitempty"`
}

func (m *TabNav) Reset()                    { *m = TabNav{} }
func (m *TabNav) String() string            { return proto.CompactTextString(m) }
func (*TabNav) ProtoMessage()               {}
func (*TabNav) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TabNav) GetScreens() []*Screen {
	if m != nil {
		return m.Screens
	}
	return nil
}

func (m *TabNav) GetSelectedIndex() int64 {
	if m != nil {
		return m.SelectedIndex
	}
	return 0
}

func (m *TabNav) GetEventFunc() int64 {
	if m != nil {
		return m.EventFunc
	}
	return 0
}

type Event struct {
	SelectedIndex int64 `protobuf:"varint,1,opt,name=selectedIndex" json:"selectedIndex,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Event) GetSelectedIndex() int64 {
	if m != nil {
		return m.SelectedIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Screen)(nil), "tabnav.Screen")
	proto.RegisterType((*TabNav)(nil), "tabnav.TabNav")
	proto.RegisterType((*Event)(nil), "tabnav.Event")
}

func init() {
	proto.RegisterFile("github.com/overcyn/mochi/pb/view/tabnav/tabnav.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x5f, 0x4b, 0xfb, 0x30,
	0x14, 0x25, 0xed, 0xd6, 0x1f, 0xcb, 0x7e, 0x1b, 0x12, 0x7c, 0x08, 0xe2, 0x43, 0x19, 0x7b, 0x28,
	0x88, 0x2d, 0xcc, 0x7d, 0x82, 0x81, 0xc2, 0x1e, 0x14, 0x89, 0x3e, 0xf9, 0x96, 0xa4, 0x97, 0x2d,
	0xb0, 0x25, 0xa3, 0xcd, 0xa2, 0x7e, 0x1b, 0xf1, 0x93, 0x4a, 0x92, 0xea, 0x18, 0xfe, 0x79, 0x6a,
	0xcf, 0xbd, 0xe7, 0x9e, 0x7b, 0xcf, 0x09, 0x9e, 0xaf, 0x94, 0x5d, 0xef, 0x45, 0x29, 0xcd, 0xb6,
	0x32, 0x0e, 0x1a, 0xf9, 0xaa, 0xab, 0xad, 0x91, 0x6b, 0x55, 0xed, 0x44, 0xe5, 0x14, 0x3c, 0x57,
	0x96, 0x0b, 0xcd, 0x5d, 0xf7, 0x29, 0x77, 0x8d, 0xb1, 0x86, 0x64, 0x11, 0x9d, 0x5d, 0xfc, 0x35,
	0xad, 0xb6, 0x7c, 0x05, 0x5e, 0x22, 0x0e, 0x4d, 0xde, 0x10, 0xce, 0x1e, 0x64, 0x03, 0xa0, 0xc9,
	0x18, 0x27, 0xaa, 0xa6, 0x28, 0x47, 0x45, 0xca, 0x12, 0x55, 0x93, 0x53, 0xdc, 0xb7, 0xca, 0x6e,
	0x80, 0x26, 0x39, 0x2a, 0x06, 0x2c, 0x02, 0x32, 0xc5, 0x3d, 0x25, 0x8d, 0xa6, 0x69, 0x8e, 0x8a,
	0xe1, 0xec, 0xa4, 0x3c, 0x08, 0x2e, 0xfd, 0x1f, 0x0b, 0x5d, 0x32, 0xc7, 0xff, 0x5b, 0xd8, 0x80,
	0xb4, 0x50, 0x2f, 0x3d, 0xbb, 0xf7, 0x0b, 0xfb, 0x88, 0xe5, 0x37, 0x0a, 0x5e, 0xaf, 0x80, 0xf6,
	0xe3, 0xc6, 0x00, 0x26, 0x0e, 0x67, 0x8f, 0x5c, 0xdc, 0x71, 0x47, 0x0a, 0xfc, 0xaf, 0x0d, 0xb7,
	0xb6, 0x14, 0xe5, 0x69, 0x31, 0x9c, 0x8d, 0xcb, 0x2e, 0x81, 0x68, 0x81, 0x7d, 0xb6, 0xc9, 0x14,
	0x8f, 0xbe, 0x94, 0x75, 0x0d, 0x2f, 0xc1, 0x43, 0xca, 0x8e, 0x8b, 0xe4, 0x1c, 0x0f, 0xc0, 0x81,
	0xb6, 0x37, 0x7b, 0x2d, 0x83, 0xa1, 0x94, 0x1d, 0x0a, 0x93, 0x4b, 0xdc, 0xbf, 0xf6, 0xe0, 0xbb,
	0x18, 0xfa, 0x41, 0x6c, 0x41, 0x9f, 0xba, 0x07, 0x78, 0x4f, 0x46, 0xb7, 0x3e, 0xee, 0xfb, 0x45,
	0x3c, 0x5b, 0x64, 0x21, 0xea, 0xab, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x44, 0x09, 0x09,
	0xd7, 0x01, 0x00, 0x00,
}
