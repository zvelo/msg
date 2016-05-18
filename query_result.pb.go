// Code generated by protoc-gen-go.
// source: zvelo.io/msg/query_result.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QueryResult struct {
	RequestId       string        `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	TrackingId      string        `protobuf:"bytes,2,opt,name=tracking_id,json=trackingId" json:"tracking_id,omitempty"`
	Status          *Status       `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	RequestDataset  []DataSetType `protobuf:"varint,6,rep,name=request_dataset,json=requestDataset,enum=msg.DataSetType" json:"request_dataset,omitempty"`
	ResponseDataset *DataSet      `protobuf:"bytes,7,opt,name=response_dataset,json=responseDataset" json:"response_dataset,omitempty"`
	Url             string        `protobuf:"bytes,8,opt,name=url" json:"url,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *QueryResult) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryResult) GetResponseDataset() *DataSet {
	if m != nil {
		return m.ResponseDataset
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryResult)(nil), "msg.QueryResult")
}

var fileDescriptor3 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0x26, 0x84, 0xf6, 0x52, 0xa5, 0x96, 0xa7, 0x10, 0x09, 0xb5, 0x82, 0xa5, 0x53,
	0x2a, 0x95, 0x01, 0xb1, 0xa2, 0x2e, 0xb0, 0xe1, 0x32, 0xb1, 0x54, 0x81, 0x58, 0x55, 0x44, 0xa8,
	0x83, 0x7d, 0x41, 0x2a, 0x33, 0x3f, 0x9c, 0x8b, 0x63, 0x57, 0xed, 0x12, 0x45, 0xef, 0x7d, 0xf7,
	0xde, 0xf9, 0x60, 0xf6, 0xfb, 0x23, 0x6b, 0x95, 0x57, 0x6a, 0xf9, 0x65, 0x76, 0xcb, 0xef, 0x56,
	0xea, 0xc3, 0x56, 0x4b, 0xd3, 0xd6, 0x98, 0x37, 0x5a, 0xa1, 0xe2, 0x01, 0xe9, 0x59, 0x76, 0x46,
	0x95, 0x05, 0x16, 0x46, 0x3a, 0x20, 0xbb, 0x3a, 0xf3, 0x0c, 0x16, 0xd8, 0x9a, 0xde, 0xba, 0xf9,
	0x1b, 0x42, 0xfc, 0xd2, 0x45, 0x0a, 0x9b, 0xc8, 0xaf, 0x01, 0xb4, 0xa4, 0x0e, 0x83, 0xdb, 0xaa,
	0x4c, 0x07, 0xf3, 0xc1, 0x62, 0x2c, 0xc6, 0x4e, 0x79, 0x2a, 0xf9, 0x0c, 0x62, 0xd4, 0xc5, 0xc7,
	0x67, 0xb5, 0xdf, 0x75, 0xfe, 0xd0, 0xfa, 0xe0, 0x25, 0x02, 0x6e, 0x21, 0xea, 0xf3, 0xd3, 0x90,
	0xbc, 0x78, 0x15, 0xe7, 0x54, 0x99, 0x6f, 0xac, 0x24, 0x9c, 0xc5, 0x1f, 0x60, 0xea, 0x4b, 0xdc,
	0xa2, 0x69, 0x34, 0x0f, 0x16, 0xc9, 0x8a, 0x59, 0x7a, 0x4d, 0xda, 0x46, 0xe2, 0xeb, 0xa1, 0x91,
	0x22, 0x71, 0xe0, 0xba, 0xe7, 0xf8, 0x3d, 0x30, 0x7a, 0x7b, 0xa3, 0xf6, 0x46, 0x1e, 0x67, 0x2f,
	0x6d, 0xd3, 0xe4, 0x74, 0x56, 0x4c, 0x3d, 0xe5, 0x07, 0x19, 0x04, 0xad, 0xae, 0xd3, 0x91, 0xdd,
	0xb8, 0xfb, 0x7d, 0x0e, 0x47, 0x01, 0x0b, 0xe9, 0x7b, 0xc1, 0xa2, 0xc7, 0xe4, 0x6d, 0x72, 0x7a,
	0xa3, 0xf7, 0xc8, 0x5e, 0xe7, 0xee, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x49, 0x81, 0x7e, 0xf1, 0x7c,
	0x01, 0x00, 0x00,
}