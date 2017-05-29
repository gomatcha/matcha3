// Code generated by protoc-gen-go. DO NOT EDIT.
// source: layout.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	layout.proto
	text.proto
	paint.proto
	view.proto
	scrollview.proto
	imageview.proto
	button.proto
	touch.proto

It has these top-level messages:
	Point
	Rect
	Insets
	Guide
	Text
	Font
	Color
	TextStyle
	PaintStyle
	Node
	Root
	ScrollView
	Image
	ImageView
	Button
	Recognizer
	RecognizerList
	TapRecognizer
	TapEvent
	PressRecognizer
	PanRecognizer
*/
package pb

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

type Point struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rect struct {
	Min *Point `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max *Point `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
}

func (m *Rect) Reset()                    { *m = Rect{} }
func (m *Rect) String() string            { return proto.CompactTextString(m) }
func (*Rect) ProtoMessage()               {}
func (*Rect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rect) GetMin() *Point {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Rect) GetMax() *Point {
	if m != nil {
		return m.Max
	}
	return nil
}

type Insets struct {
	Top    float64 `protobuf:"fixed64,1,opt,name=top" json:"top,omitempty"`
	Left   float64 `protobuf:"fixed64,2,opt,name=left" json:"left,omitempty"`
	Bottom float64 `protobuf:"fixed64,3,opt,name=bottom" json:"bottom,omitempty"`
	Right  float64 `protobuf:"fixed64,4,opt,name=right" json:"right,omitempty"`
}

func (m *Insets) Reset()                    { *m = Insets{} }
func (m *Insets) String() string            { return proto.CompactTextString(m) }
func (*Insets) ProtoMessage()               {}
func (*Insets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Insets) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Insets) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Insets) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Insets) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

type Guide struct {
	Frame  *Rect   `protobuf:"bytes,1,opt,name=frame" json:"frame,omitempty"`
	Insets *Insets `protobuf:"bytes,2,opt,name=insets" json:"insets,omitempty"`
	ZIndex int64   `protobuf:"varint,3,opt,name=zIndex" json:"zIndex,omitempty"`
}

func (m *Guide) Reset()                    { *m = Guide{} }
func (m *Guide) String() string            { return proto.CompactTextString(m) }
func (*Guide) ProtoMessage()               {}
func (*Guide) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Guide) GetFrame() *Rect {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *Guide) GetInsets() *Insets {
	if m != nil {
		return m.Insets
	}
	return nil
}

func (m *Guide) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "layout.Point")
	proto.RegisterType((*Rect)(nil), "layout.Rect")
	proto.RegisterType((*Insets)(nil), "layout.Insets")
	proto.RegisterType((*Guide)(nil), "layout.Guide")
}

func init() { proto.RegisterFile("layout.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xd9, 0xfc, 0x59, 0x61, 0x8c, 0x45, 0x06, 0x91, 0xdc, 0x94, 0x15, 0xc4, 0x53, 0x0f,
	0xfa, 0x06, 0xbd, 0x68, 0x0f, 0x42, 0xd9, 0xa3, 0x78, 0x49, 0xda, 0xad, 0x5d, 0x6c, 0xb2, 0x21,
	0x9d, 0xc2, 0xc6, 0xc7, 0xf1, 0x49, 0x65, 0x67, 0x37, 0x37, 0x6f, 0xf3, 0xcd, 0xfc, 0x26, 0xf3,
	0x65, 0xa1, 0x3a, 0x36, 0x93, 0x3b, 0xd3, 0x72, 0x18, 0x1d, 0x39, 0x94, 0x91, 0xd4, 0x03, 0x94,
	0x1b, 0x67, 0x7b, 0xc2, 0x0a, 0x84, 0xaf, 0xc5, 0xbd, 0x78, 0x12, 0x5a, 0xf8, 0x40, 0x53, 0x9d,
	0x45, 0x9a, 0xd4, 0x1b, 0x14, 0xda, 0x6c, 0x09, 0xef, 0x20, 0xef, 0x6c, 0xcf, 0xa9, 0xcb, 0xe7,
	0xab, 0x65, 0xfa, 0x20, 0xef, 0xeb, 0x30, 0xe1, 0x40, 0xe3, 0x79, 0xf1, 0x9f, 0x40, 0xe3, 0xd5,
	0x27, 0xc8, 0x75, 0x7f, 0x32, 0x74, 0xc2, 0x6b, 0xc8, 0xc9, 0x0d, 0xe9, 0x62, 0x28, 0x11, 0xa1,
	0x38, 0x9a, 0x3d, 0xa5, 0xb3, 0x5c, 0xe3, 0x2d, 0xc8, 0xd6, 0x11, 0xb9, 0xae, 0xce, 0xb9, 0x9b,
	0x08, 0x6f, 0xa0, 0x1c, 0xed, 0xd7, 0x81, 0xea, 0x82, 0xdb, 0x11, 0xd4, 0x37, 0x94, 0xaf, 0x67,
	0xbb, 0x33, 0xa8, 0xa0, 0xdc, 0x8f, 0x4d, 0x67, 0x92, 0x6a, 0x35, 0x9b, 0x84, 0xbf, 0xd0, 0x71,
	0x84, 0x8f, 0x20, 0x2d, 0xab, 0x24, 0xdd, 0xc5, 0x1c, 0x8a, 0x82, 0x3a, 0x4d, 0x83, 0xc2, 0xcf,
	0xba, 0xdf, 0x19, 0xcf, 0x0a, 0xb9, 0x4e, 0xb4, 0x5a, 0x7c, 0x64, 0x43, 0xfb, 0x9b, 0x5d, 0xbc,
	0xbb, 0xed, 0xc1, 0x6e, 0x56, 0xad, 0xe4, 0x87, 0x7d, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x05,
	0x9f, 0x16, 0xd0, 0x68, 0x01, 0x00, 0x00,
}
