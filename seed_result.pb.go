// Code generated by protoc-gen-go.
// source: zvelo.io/msg/seed_result.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SeedResult struct {
	Url     string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Dataset *DataSet `protobuf:"bytes,2,opt,name=dataset" json:"dataset,omitempty"`
}

func (m *SeedResult) Reset()                    { *m = SeedResult{} }
func (m *SeedResult) String() string            { return proto.CompactTextString(m) }
func (*SeedResult) ProtoMessage()               {}
func (*SeedResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SeedResult) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type SeedResults struct {
	Values []*SeedResult `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *SeedResults) Reset()                    { *m = SeedResults{} }
func (m *SeedResults) String() string            { return proto.CompactTextString(m) }
func (*SeedResults) ProtoMessage()               {}
func (*SeedResults) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SeedResults) GetValues() []*SeedResult {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*SeedResult)(nil), "msg.SeedResult")
	proto.RegisterType((*SeedResults)(nil), "msg.SeedResults")
}

var fileDescriptor3 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xab, 0x2a, 0x4b, 0xcd,
	0xc9, 0xd7, 0xcb, 0xcc, 0xd7, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x4e, 0x4d, 0x4d, 0x89, 0x2f, 0x4a,
	0x2d, 0x2e, 0xcd, 0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x0a, 0x4b, 0x49,
	0xa1, 0x28, 0x4a, 0x49, 0x2c, 0x49, 0x2c, 0x4e, 0x85, 0x2a, 0x50, 0x72, 0xe3, 0xe2, 0x0a, 0x06,
	0xea, 0x0a, 0x02, 0x6b, 0x12, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0x31, 0x85, 0xd4, 0xb8, 0xd8, 0xa1, 0x1a, 0x24, 0x98, 0x80, 0xa2, 0xdc, 0x46,
	0x3c, 0x7a, 0x40, 0x43, 0xf4, 0x5c, 0x80, 0x62, 0xc1, 0xa9, 0x25, 0x41, 0x30, 0x49, 0x25, 0x33,
	0x2e, 0x6e, 0x84, 0x39, 0xc5, 0x42, 0xea, 0x5c, 0x6c, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0xc5, 0x40,
	0xb3, 0x98, 0x81, 0xba, 0xf8, 0xc1, 0xba, 0x10, 0x2a, 0x82, 0xa0, 0xd2, 0x4e, 0x7c, 0x51, 0x3c,
	0xc8, 0xae, 0x4b, 0x62, 0x03, 0x3b, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x57, 0xae,
	0x85, 0xd9, 0x00, 0x00, 0x00,
}