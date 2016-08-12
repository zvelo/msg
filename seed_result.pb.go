// Code generated by protoc-gen-go.
// source: zvelo/msg/seed_result.proto
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

func init() { proto.RegisterFile("zvelo/msg/seed_result.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xae, 0x2a, 0x4b, 0xcd,
	0xc9, 0xd7, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x4e, 0x4d, 0x4d, 0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x06, 0x0a, 0x4b, 0x89, 0x23, 0x54, 0xa4,
	0x24, 0x96, 0x24, 0x16, 0xa7, 0x42, 0x65, 0x95, 0xdc, 0xb8, 0xb8, 0x82, 0x81, 0x5a, 0x82, 0xc0,
	0x3a, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x8b, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40,
	0x4c, 0x21, 0x35, 0x2e, 0x76, 0xa8, 0x06, 0x09, 0x26, 0xa0, 0x28, 0xb7, 0x11, 0x8f, 0x1e, 0xd0,
	0x10, 0x3d, 0x17, 0xa0, 0x58, 0x70, 0x6a, 0x49, 0x10, 0x4c, 0x52, 0xc9, 0x8c, 0x8b, 0x1b, 0x61,
	0x4e, 0xb1, 0x90, 0x3a, 0x17, 0x5b, 0x59, 0x62, 0x4e, 0x69, 0x6a, 0x31, 0xd0, 0x2c, 0x66, 0xa0,
	0x2e, 0x7e, 0xb0, 0x2e, 0x84, 0x8a, 0x20, 0xa8, 0xb4, 0x13, 0x5f, 0x14, 0x0f, 0xd8, 0x69, 0x7a,
	0x99, 0x60, 0xd7, 0x25, 0xb1, 0x81, 0x9d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x46, 0x94,
	0xe2, 0x12, 0xd3, 0x00, 0x00, 0x00,
}
