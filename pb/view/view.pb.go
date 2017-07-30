// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/view.proto

/*
Package view is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/pb/view/view.proto

It has these top-level messages:
	BuildNode
	LayoutPaintNode
	Root
*/
package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha_paint "gomatcha.io/matcha/pb/paint"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildNode struct {
	Id          int64                           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BuildId     int64                           `protobuf:"varint,2,opt,name=buildId" json:"buildId,omitempty"`
	BridgeName  string                          `protobuf:"bytes,3,opt,name=bridgeName" json:"bridgeName,omitempty"`
	BridgeValue *google_protobuf.Any            `protobuf:"bytes,4,opt,name=bridgeValue" json:"bridgeValue,omitempty"`
	Values      map[string]*google_protobuf.Any `protobuf:"bytes,5,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Children    []int64                         `protobuf:"varint,6,rep,packed,name=children" json:"children,omitempty"`
}

func (m *BuildNode) Reset()                    { *m = BuildNode{} }
func (m *BuildNode) String() string            { return proto.CompactTextString(m) }
func (*BuildNode) ProtoMessage()               {}
func (*BuildNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildNode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BuildNode) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *BuildNode) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *BuildNode) GetBridgeValue() *google_protobuf.Any {
	if m != nil {
		return m.BridgeValue
	}
	return nil
}

func (m *BuildNode) GetValues() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *BuildNode) GetChildren() []int64 {
	if m != nil {
		return m.Children
	}
	return nil
}

type LayoutPaintNode struct {
	Id       int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LayoutId int64 `protobuf:"varint,2,opt,name=layoutId" json:"layoutId,omitempty"`
	PaintId  int64 `protobuf:"varint,3,opt,name=paintId" json:"paintId,omitempty"`
	// matcha.layout.Guide layoutGuide = 4;
	// Guide
	Minx       float64             `protobuf:"fixed64,4,opt,name=minx" json:"minx,omitempty"`
	Miny       float64             `protobuf:"fixed64,5,opt,name=miny" json:"miny,omitempty"`
	Maxx       float64             `protobuf:"fixed64,6,opt,name=maxx" json:"maxx,omitempty"`
	Maxy       float64             `protobuf:"fixed64,7,opt,name=maxy" json:"maxy,omitempty"`
	ZIndex     int64               `protobuf:"varint,8,opt,name=zIndex" json:"zIndex,omitempty"`
	ChildOrder []int64             `protobuf:"varint,9,rep,packed,name=childOrder" json:"childOrder,omitempty"`
	PaintStyle *matcha_paint.Style `protobuf:"bytes,10,opt,name=paintStyle" json:"paintStyle,omitempty"`
}

func (m *LayoutPaintNode) Reset()                    { *m = LayoutPaintNode{} }
func (m *LayoutPaintNode) String() string            { return proto.CompactTextString(m) }
func (*LayoutPaintNode) ProtoMessage()               {}
func (*LayoutPaintNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LayoutPaintNode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LayoutPaintNode) GetLayoutId() int64 {
	if m != nil {
		return m.LayoutId
	}
	return 0
}

func (m *LayoutPaintNode) GetPaintId() int64 {
	if m != nil {
		return m.PaintId
	}
	return 0
}

func (m *LayoutPaintNode) GetMinx() float64 {
	if m != nil {
		return m.Minx
	}
	return 0
}

func (m *LayoutPaintNode) GetMiny() float64 {
	if m != nil {
		return m.Miny
	}
	return 0
}

func (m *LayoutPaintNode) GetMaxx() float64 {
	if m != nil {
		return m.Maxx
	}
	return 0
}

func (m *LayoutPaintNode) GetMaxy() float64 {
	if m != nil {
		return m.Maxy
	}
	return 0
}

func (m *LayoutPaintNode) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func (m *LayoutPaintNode) GetChildOrder() []int64 {
	if m != nil {
		return m.ChildOrder
	}
	return nil
}

func (m *LayoutPaintNode) GetPaintStyle() *matcha_paint.Style {
	if m != nil {
		return m.PaintStyle
	}
	return nil
}

type Root struct {
	LayoutPaintNodes map[int64]*LayoutPaintNode      `protobuf:"bytes,2,rep,name=layoutPaintNodes" json:"layoutPaintNodes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	BuildNodes       map[int64]*BuildNode            `protobuf:"bytes,3,rep,name=buildNodes" json:"buildNodes,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Middleware       map[string]*google_protobuf.Any `protobuf:"bytes,4,rep,name=middleware" json:"middleware,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Root) GetLayoutPaintNodes() map[int64]*LayoutPaintNode {
	if m != nil {
		return m.LayoutPaintNodes
	}
	return nil
}

func (m *Root) GetBuildNodes() map[int64]*BuildNode {
	if m != nil {
		return m.BuildNodes
	}
	return nil
}

func (m *Root) GetMiddleware() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Middleware
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildNode)(nil), "matcha.view.BuildNode")
	proto.RegisterType((*LayoutPaintNode)(nil), "matcha.view.LayoutPaintNode")
	proto.RegisterType((*Root)(nil), "matcha.view.Root")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/view.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0x92, 0xed, 0xd8, 0xe3, 0x52, 0x87, 0x6d, 0x13, 0xb6, 0xa2, 0x14, 0xd7, 0x50, 0x62,
	0x4a, 0x91, 0xc1, 0x81, 0x52, 0x72, 0x8b, 0xa1, 0x07, 0x43, 0xf3, 0x83, 0x4c, 0x73, 0xe8, 0x6d,
	0x95, 0xdd, 0x3a, 0x4b, 0x65, 0xad, 0x91, 0xe5, 0x58, 0xdb, 0xe7, 0xe8, 0x13, 0xf4, 0x21, 0xfa,
	0x5c, 0x7d, 0x84, 0xb2, 0x23, 0x59, 0x5d, 0x3b, 0xca, 0x29, 0x17, 0xb1, 0xf3, 0xed, 0x37, 0xdf,
	0xce, 0x7c, 0x33, 0x82, 0x77, 0x73, 0xb5, 0x60, 0xd9, 0xed, 0x1d, 0x0b, 0xa4, 0x1a, 0x15, 0xa7,
	0xd1, 0x32, 0x1a, 0xdd, 0x4b, 0xb1, 0xc1, 0x4f, 0xb0, 0x4c, 0x55, 0xa6, 0x48, 0xb7, 0x24, 0x19,
	0xc8, 0x3f, 0xa9, 0xcf, 0x59, 0x32, 0x99, 0x64, 0xc5, 0xb7, 0xc8, 0xf2, 0x5f, 0xcd, 0x95, 0x9a,
	0xc7, 0x62, 0x84, 0x51, 0xb4, 0xfe, 0x3e, 0x62, 0x89, 0x2e, 0xae, 0x06, 0x7f, 0x5c, 0xe8, 0x4c,
	0xd6, 0x32, 0xe6, 0x97, 0x8a, 0x0b, 0xf2, 0x1c, 0x5c, 0xc9, 0xa9, 0xd3, 0x77, 0x86, 0x5e, 0xe8,
	0x4a, 0x4e, 0x28, 0x1c, 0x44, 0xe6, 0x72, 0xca, 0xa9, 0x8b, 0xe0, 0x36, 0x24, 0x6f, 0x00, 0xa2,
	0x54, 0xf2, 0xb9, 0xb8, 0x64, 0x0b, 0x41, 0xbd, 0xbe, 0x33, 0xec, 0x84, 0x16, 0x42, 0x3e, 0x42,
	0xb7, 0x88, 0x6e, 0x58, 0xbc, 0x16, 0xb4, 0xd1, 0x77, 0x86, 0xdd, 0xf1, 0xcb, 0xa0, 0x28, 0x24,
	0xd8, 0x16, 0x12, 0x9c, 0x27, 0x3a, 0xb4, 0x89, 0xe4, 0x0c, 0x5a, 0xf7, 0xe6, 0xb0, 0xa2, 0xcd,
	0xbe, 0x37, 0xec, 0x8e, 0x07, 0x81, 0xd5, 0x71, 0x50, 0x55, 0x1a, 0x20, 0x7b, 0xf5, 0x39, 0xc9,
	0x52, 0x1d, 0x96, 0x19, 0xc4, 0x87, 0xf6, 0xed, 0x9d, 0x8c, 0x79, 0x2a, 0x12, 0xda, 0xea, 0x7b,
	0x43, 0x2f, 0xac, 0x62, 0xff, 0x0a, 0xba, 0x56, 0x0a, 0x39, 0x04, 0xef, 0x87, 0xd0, 0xd8, 0x69,
	0x27, 0x34, 0x47, 0xf2, 0x1e, 0x9a, 0x28, 0x83, 0x8d, 0x3e, 0x56, 0x6a, 0x41, 0x39, 0x73, 0x3f,
	0x39, 0x83, 0x5f, 0x2e, 0xf4, 0xbe, 0x30, 0xad, 0xd6, 0xd9, 0xb5, 0x71, 0xba, 0xd6, 0x3e, 0x1f,
	0xda, 0x31, 0x52, 0x2a, 0xff, 0xaa, 0xd8, 0x58, 0x8b, 0x23, 0x9a, 0x72, 0x74, 0xcf, 0x0b, 0xb7,
	0x21, 0x21, 0xd0, 0x58, 0xc8, 0x24, 0x47, 0xcf, 0x9c, 0x10, 0xcf, 0x25, 0xa6, 0x69, 0xb3, 0xc2,
	0x34, 0x62, 0x2c, 0xcf, 0x69, 0xab, 0xc4, 0x58, 0x9e, 0x97, 0x98, 0xa6, 0x07, 0x15, 0xa6, 0xc9,
	0x31, 0xb4, 0x7e, 0x4e, 0x13, 0x2e, 0x72, 0xda, 0xc6, 0x87, 0xca, 0xc8, 0x8c, 0x10, 0xed, 0xb9,
	0x4a, 0xb9, 0x48, 0x69, 0x07, 0x0d, 0xb3, 0x10, 0x72, 0x0a, 0x80, 0x25, 0xcd, 0x32, 0x1d, 0x0b,
	0x0a, 0x68, 0xcb, 0x8b, 0xed, 0x38, 0x8a, 0xf5, 0xc2, 0xab, 0xd0, 0xa2, 0x0d, 0xfe, 0x7a, 0xd0,
	0x08, 0x95, 0xca, 0xc8, 0x0c, 0x0e, 0xe3, 0x5d, 0x7b, 0x56, 0xd4, 0xc5, 0x91, 0x9e, 0xec, 0x8c,
	0xd4, 0x90, 0x83, 0x3d, 0x23, 0xcb, 0xb9, 0x3e, 0x10, 0x20, 0xe7, 0x00, 0xd1, 0x76, 0x05, 0x56,
	0xd4, 0x43, 0xb9, 0xb7, 0x0f, 0xe5, 0xaa, 0x35, 0x29, 0x85, 0xac, 0x24, 0x23, 0xb1, 0x90, 0x9c,
	0xc7, 0x62, 0xc3, 0x52, 0xb3, 0x97, 0x8f, 0x48, 0x5c, 0x54, 0x9c, 0x52, 0xe2, 0x7f, 0x92, 0xcf,
	0xe0, 0xa8, 0xb6, 0x60, 0x7b, 0xab, 0xbc, 0x62, 0xab, 0xc6, 0xbb, 0x5b, 0xf5, 0x7a, 0xe7, 0xa1,
	0x3d, 0x11, 0x6b, 0xbb, 0xfc, 0xaf, 0xd0, 0xdb, 0x6b, 0xa2, 0x46, 0xfc, 0xc3, 0xae, 0xf8, 0x71,
	0xfd, 0xaf, 0x62, 0xcb, 0xce, 0xa0, 0xb7, 0xd7, 0xd8, 0xd3, 0xff, 0x84, 0xc9, 0xd1, 0xb7, 0x86,
	0x79, 0xf1, 0xb7, 0xfb, 0xec, 0x02, 0xdf, 0xbf, 0x91, 0x62, 0x73, 0x3d, 0x89, 0x5a, 0xc8, 0x3f,
	0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x85, 0xe1, 0x67, 0xda, 0x04, 0x00, 0x00,
}
