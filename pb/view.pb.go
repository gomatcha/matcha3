// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/mochi/pb/view.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Node struct {
	Id          int64                           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BuildId     int64                           `protobuf:"varint,2,opt,name=buildId" json:"buildId,omitempty"`
	LayoutId    int64                           `protobuf:"varint,3,opt,name=layoutId" json:"layoutId,omitempty"`
	PaintId     int64                           `protobuf:"varint,4,opt,name=paintId" json:"paintId,omitempty"`
	BridgeName  string                          `protobuf:"bytes,7,opt,name=bridgeName" json:"bridgeName,omitempty"`
	BridgeValue *google_protobuf.Any            `protobuf:"bytes,8,opt,name=bridgeValue" json:"bridgeValue,omitempty"`
	Values      map[string]*google_protobuf.Any `protobuf:"bytes,10,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Children    []*Node                         `protobuf:"bytes,5,rep,name=children" json:"children,omitempty"`
	LayoutGuide *Guide                          `protobuf:"bytes,6,opt,name=layoutGuide" json:"layoutGuide,omitempty"`
	PaintStyle  *PaintStyle                     `protobuf:"bytes,9,opt,name=paintStyle" json:"paintStyle,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *Node) GetLayoutId() int64 {
	if m != nil {
		return m.LayoutId
	}
	return 0
}

func (m *Node) GetPaintId() int64 {
	if m != nil {
		return m.PaintId
	}
	return 0
}

func (m *Node) GetBridgeName() string {
	if m != nil {
		return m.BridgeName
	}
	return ""
}

func (m *Node) GetBridgeValue() *google_protobuf.Any {
	if m != nil {
		return m.BridgeValue
	}
	return nil
}

func (m *Node) GetValues() map[string]*google_protobuf.Any {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Node) GetChildren() []*Node {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Node) GetLayoutGuide() *Guide {
	if m != nil {
		return m.LayoutGuide
	}
	return nil
}

func (m *Node) GetPaintStyle() *PaintStyle {
	if m != nil {
		return m.PaintStyle
	}
	return nil
}

type Root struct {
	Node *Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Root) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "view.Node")
	proto.RegisterType((*Root)(nil), "view.Root")
}

func init() { proto.RegisterFile("github.com/overcyn/mochi/pb/view.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xd1, 0x87, 0xbf, 0x46, 0xd4, 0xb4, 0x4b, 0x29, 0x5b, 0x1d, 0x8c, 0xf0, 0xc1, 0x15,
	0x3d, 0xac, 0xa8, 0x0b, 0xa5, 0xf4, 0x56, 0x43, 0x29, 0x3e, 0xd4, 0x35, 0x5b, 0xe8, 0x21, 0x37,
	0x49, 0xbb, 0x91, 0x97, 0xc8, 0x5a, 0x21, 0x4b, 0x0e, 0x7a, 0x9d, 0xbc, 0x4e, 0x5e, 0x2a, 0x68,
	0xd6, 0x76, 0x94, 0x43, 0x7c, 0x9b, 0xff, 0xcc, 0x6f, 0x56, 0xff, 0x99, 0x11, 0x2c, 0x32, 0x55,
	0xef, 0x9a, 0x84, 0xa5, 0x7a, 0x1f, 0xe9, 0xa3, 0xac, 0xd2, 0xb6, 0x88, 0xf6, 0x3a, 0xdd, 0xa9,
	0xa8, 0x4c, 0xa2, 0xa3, 0x92, 0xf7, 0xac, 0xac, 0x74, 0xad, 0x89, 0xdb, 0xc5, 0x7e, 0x78, 0x8d,
	0xce, 0xe3, 0x56, 0x37, 0xb5, 0xe1, 0xfd, 0x4f, 0xd7, 0xc8, 0x32, 0x56, 0xc5, 0x19, 0xfc, 0x98,
	0x69, 0x9d, 0xe5, 0x32, 0x42, 0x95, 0x34, 0xb7, 0x51, 0x5c, 0xb4, 0xa6, 0x34, 0x7f, 0x74, 0xc0,
	0xdd, 0x68, 0x21, 0xc9, 0x14, 0x6c, 0x25, 0xa8, 0x15, 0x58, 0xa1, 0xc3, 0x6d, 0x25, 0x08, 0x85,
	0x51, 0xd2, 0xa8, 0x5c, 0xac, 0x05, 0xb5, 0x31, 0x79, 0x96, 0xc4, 0x87, 0xb1, 0xb1, 0xb1, 0x16,
	0xd4, 0xc1, 0xd2, 0x45, 0x77, 0x5d, 0xf8, 0xe1, 0xb5, 0xa0, 0xae, 0xe9, 0x3a, 0x49, 0x32, 0x03,
	0x48, 0x2a, 0x25, 0x32, 0xb9, 0x89, 0xf7, 0x92, 0x8e, 0x02, 0x2b, 0x9c, 0xf0, 0x5e, 0x86, 0x7c,
	0x03, 0xcf, 0xa8, 0xff, 0x71, 0xde, 0x48, 0x3a, 0x0e, 0xac, 0xd0, 0x5b, 0xbe, 0x67, 0xc6, 0x39,
	0x3b, 0x3b, 0x67, 0x3f, 0x8b, 0x96, 0xf7, 0x41, 0xc2, 0x60, 0x78, 0xec, 0x82, 0x03, 0x85, 0xc0,
	0x09, 0xbd, 0xe5, 0x07, 0x86, 0x1b, 0xed, 0x66, 0x62, 0x48, 0x1c, 0x7e, 0x15, 0x75, 0xd5, 0xf2,
	0x13, 0x45, 0x16, 0x30, 0x4e, 0x77, 0x2a, 0x17, 0x95, 0x2c, 0xe8, 0x00, 0x3b, 0xe0, 0xb9, 0x83,
	0x5f, 0x6a, 0x24, 0x02, 0xcf, 0x4c, 0xf5, 0xbb, 0x51, 0x42, 0xd2, 0x21, 0xfa, 0x79, 0xc3, 0x4e,
	0x07, 0xc0, 0x24, 0xef, 0x13, 0xe4, 0x0b, 0x00, 0xce, 0xfa, 0xaf, 0x6e, 0x73, 0x49, 0x27, 0xc8,
	0xbf, 0x63, 0xe6, 0x0c, 0xdb, 0x4b, 0x81, 0xf7, 0x20, 0xff, 0x2f, 0x78, 0x3d, 0x8b, 0xe4, 0x2d,
	0x38, 0x77, 0xb2, 0xc5, 0x1b, 0x4c, 0x78, 0x17, 0x92, 0xcf, 0x30, 0x40, 0xdb, 0x78, 0x82, 0xd7,
	0xd6, 0x61, 0x90, 0x1f, 0xf6, 0x77, 0x6b, 0xbe, 0x00, 0x97, 0x6b, 0x5d, 0x93, 0x19, 0xb8, 0x85,
	0x16, 0x12, 0x9f, 0x7a, 0x39, 0x20, 0xe6, 0x57, 0xd3, 0x1b, 0xbb, 0x4c, 0x1e, 0xec, 0xd1, 0x9f,
	0xee, 0x6f, 0xd9, 0xae, 0x92, 0x21, 0x3e, 0xf8, 0xf5, 0x29, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x19,
	0x46, 0xd3, 0xaa, 0x02, 0x00, 0x00,
}
