// Code generated by protoc-gen-go.
// source: zvelo.io/msg/stream_result.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StreamResult struct {
	Url     string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Dataset *DataSet `protobuf:"bytes,2,opt,name=dataset" json:"dataset,omitempty"`
}

func (m *StreamResult) Reset()                    { *m = StreamResult{} }
func (m *StreamResult) String() string            { return proto.CompactTextString(m) }
func (*StreamResult) ProtoMessage()               {}
func (*StreamResult) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *StreamResult) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type StreamResults struct {
	Values []*StreamResult `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *StreamResults) Reset()                    { *m = StreamResults{} }
func (m *StreamResults) String() string            { return proto.CompactTextString(m) }
func (*StreamResults) ProtoMessage()               {}
func (*StreamResults) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *StreamResults) GetValues() []*StreamResult {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamResult)(nil), "msg.StreamResult")
	proto.RegisterType((*StreamResults)(nil), "msg.StreamResults")
}

var fileDescriptor5 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xa8, 0x2a, 0x4b, 0xcd,
	0xc9, 0xd7, 0xcb, 0xcc, 0xd7, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x8d,
	0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x4a,
	0x48, 0x49, 0xa1, 0x28, 0x4b, 0x49, 0x2c, 0x49, 0x2c, 0x4e, 0x85, 0x2a, 0x50, 0xf2, 0xe0, 0xe2,
	0x09, 0x06, 0xeb, 0x0b, 0x02, 0x6b, 0x13, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0xd4, 0xb8, 0xd8, 0xa1, 0x5a, 0x24, 0x98, 0x80, 0xa2,
	0xdc, 0x46, 0x3c, 0x7a, 0x40, 0x63, 0xf4, 0x5c, 0x80, 0x62, 0xc1, 0xa9, 0x25, 0x41, 0x30, 0x49,
	0x25, 0x2b, 0x2e, 0x5e, 0x64, 0x93, 0x8a, 0x85, 0x34, 0xb9, 0xd8, 0xca, 0x12, 0x73, 0x4a, 0x53,
	0x8b, 0x81, 0xa6, 0x31, 0x03, 0xf5, 0x09, 0x82, 0xf5, 0x21, 0xab, 0x09, 0x82, 0x2a, 0x70, 0xe2,
	0x8b, 0xe2, 0x41, 0x76, 0x63, 0x12, 0x1b, 0xd8, 0x71, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0xb4, 0x63, 0x1a, 0xe1, 0x00, 0x00, 0x00,
}
