// Code generated by protoc-gen-go.
// source: zvelo.io/msg/stream.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Stream_Status int32

const (
	Stream_CREATED  Stream_Status = 0
	Stream_RUNNING  Stream_Status = 1
	Stream_PAUSED   Stream_Status = 2
	Stream_RETRYING Stream_Status = 3
	Stream_ENDED    Stream_Status = 4
	Stream_FAILED   Stream_Status = 5
)

var Stream_Status_name = map[int32]string{
	0: "CREATED",
	1: "RUNNING",
	2: "PAUSED",
	3: "RETRYING",
	4: "ENDED",
	5: "FAILED",
}
var Stream_Status_value = map[string]int32{
	"CREATED":  0,
	"RUNNING":  1,
	"PAUSED":   2,
	"RETRYING": 3,
	"ENDED":    4,
	"FAILED":   5,
}

func (x Stream_Status) String() string {
	return proto.EnumName(Stream_Status_name, int32(x))
}
func (Stream_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1, 0} }

type StreamRequest struct {
	CustomerTrackingId string        `protobuf:"bytes,1,opt,name=customer_tracking_id,json=customerTrackingId" json:"customer_tracking_id,omitempty"`
	Endpoint           string        `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	Accept             string        `protobuf:"bytes,3,opt,name=accept" json:"accept,omitempty"`
	Dataset            []DataSetType `protobuf:"varint,4,rep,name=dataset,enum=msg.DataSetType" json:"dataset,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Stream struct {
	Uuid          string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	StreamStatus  Stream_Status              `protobuf:"varint,2,opt,name=stream_status,json=streamStatus,enum=msg.Stream_Status" json:"stream_status,omitempty"`
	StreamRequest *StreamRequest             `protobuf:"bytes,3,opt,name=stream_request,json=streamRequest" json:"stream_request,omitempty"`
	Org           string                     `protobuf:"bytes,5,opt,name=org" json:"org,omitempty"`
	User          string                     `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	CreationDate  *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Stream) GetStreamRequest() *StreamRequest {
	if m != nil {
		return m.StreamRequest
	}
	return nil
}

func (m *Stream) GetCreationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

type StreamReply struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Stream *Stream `protobuf:"bytes,2,opt,name=stream" json:"stream,omitempty"`
}

func (m *StreamReply) Reset()                    { *m = StreamReply{} }
func (m *StreamReply) String() string            { return proto.CompactTextString(m) }
func (*StreamReply) ProtoMessage()               {}
func (*StreamReply) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *StreamReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StreamReply) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

type StreamsReply struct {
	Status  *Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Streams []*Stream `protobuf:"bytes,2,rep,name=streams" json:"streams,omitempty"`
}

func (m *StreamsReply) Reset()                    { *m = StreamsReply{} }
func (m *StreamsReply) String() string            { return proto.CompactTextString(m) }
func (*StreamsReply) ProtoMessage()               {}
func (*StreamsReply) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StreamsReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StreamsReply) GetStreams() []*Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "msg.StreamRequest")
	proto.RegisterType((*Stream)(nil), "msg.Stream")
	proto.RegisterType((*StreamReply)(nil), "msg.StreamReply")
	proto.RegisterType((*StreamsReply)(nil), "msg.StreamsReply")
	proto.RegisterEnum("msg.Stream_Status", Stream_Status_name, Stream_Status_value)
}

var fileDescriptor4 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0xe2, 0xb4, 0x13, 0x27, 0x8a, 0x46, 0x08, 0x19, 0x5f, 0xa8, 0x8c, 0x90, 0x2a,
	0x0e, 0x0e, 0x2a, 0x07, 0xc4, 0x09, 0x05, 0x12, 0x50, 0x24, 0x14, 0xa1, 0x8d, 0x2b, 0x44, 0x2f,
	0x91, 0xeb, 0x2c, 0x96, 0x45, 0xec, 0x35, 0xf6, 0x1a, 0xa9, 0xfc, 0x0d, 0xff, 0xc3, 0x47, 0xb1,
	0x9e, 0xdd, 0x2d, 0xb4, 0x37, 0x4e, 0x99, 0xd9, 0xf7, 0xde, 0xf8, 0xcd, 0x9b, 0xc0, 0xe3, 0x9f,
	0x3f, 0xf8, 0x51, 0xc4, 0x85, 0x58, 0x94, 0x6d, 0xbe, 0x68, 0x65, 0xc3, 0xd3, 0x32, 0xae, 0x1b,
	0x21, 0x05, 0xba, 0xea, 0x25, 0x0c, 0xef, 0xe0, 0x87, 0x54, 0xa6, 0x2d, 0x97, 0x9a, 0x10, 0xde,
	0xd7, 0xa6, 0xb2, 0x6b, 0x0d, 0xf4, 0x24, 0x17, 0x22, 0x3f, 0xf2, 0x05, 0x75, 0xd7, 0xdd, 0xd7,
	0x85, 0x2c, 0x4a, 0xae, 0x18, 0x65, 0xad, 0x09, 0xd1, 0x2f, 0x07, 0xa6, 0x3b, 0xfa, 0x1a, 0xe3,
	0xdf, 0x3b, 0x05, 0xe1, 0x0b, 0x78, 0x98, 0x75, 0xad, 0x14, 0x25, 0x6f, 0xf6, 0xb2, 0x49, 0xb3,
	0x6f, 0x45, 0x95, 0xef, 0x8b, 0x43, 0xe0, 0x9c, 0x39, 0xe7, 0xa7, 0x0c, 0x2d, 0x96, 0x18, 0x68,
	0x73, 0xc0, 0x10, 0x4e, 0x78, 0x75, 0xa8, 0x45, 0x51, 0xc9, 0x60, 0x40, 0xac, 0xdb, 0x1e, 0x1f,
	0x81, 0x97, 0x66, 0x19, 0xaf, 0x65, 0xe0, 0x12, 0x62, 0x3a, 0x7c, 0x0e, 0x63, 0xb3, 0x44, 0x30,
	0x3c, 0x73, 0xcf, 0x67, 0x17, 0xf3, 0x58, 0x99, 0x8f, 0x57, 0xea, 0x6d, 0xc7, 0x65, 0x72, 0x53,
	0x73, 0x66, 0x09, 0xd1, 0xef, 0x01, 0x78, 0xda, 0x23, 0x22, 0x0c, 0xbb, 0xee, 0xd6, 0x0c, 0xd5,
	0xf8, 0x0a, 0xa6, 0x3a, 0xaf, 0xbd, 0x5e, 0x9d, 0x3c, 0xcc, 0x2e, 0x90, 0x06, 0x6a, 0x9d, 0xfa,
	0xe9, 0x11, 0xe6, 0x6b, 0xa2, 0xee, 0xf0, 0x35, 0xcc, 0x8c, 0xb0, 0xd1, 0xbb, 0x93, 0xc7, 0xc9,
	0x1d, 0xa5, 0x49, 0x85, 0x99, 0x4f, 0xd8, 0x90, 0xe6, 0xe0, 0x8a, 0x26, 0x0f, 0x46, 0x64, 0xa3,
	0x2f, 0xc9, 0x59, 0xcb, 0x9b, 0xc0, 0x33, 0xce, 0x54, 0x8d, 0x6f, 0x60, 0x9a, 0x29, 0x95, 0x2c,
	0x44, 0xb5, 0x57, 0xcb, 0xf0, 0x60, 0x4c, 0xf3, 0xc3, 0x58, 0x5f, 0x25, 0xb6, 0x57, 0x89, 0x13,
	0x7b, 0x15, 0xe6, 0x5b, 0x81, 0x4a, 0x82, 0x47, 0xbb, 0x7e, 0x71, 0xf2, 0x3a, 0x81, 0xf1, 0x3b,
	0xb6, 0x5e, 0x26, 0xeb, 0xd5, 0xfc, 0x41, 0xdf, 0xb0, 0xcb, 0xed, 0x76, 0xb3, 0xfd, 0x30, 0x77,
	0x10, 0xc0, 0xfb, 0xb4, 0xbc, 0xdc, 0x29, 0x60, 0x80, 0x3e, 0x9c, 0xb0, 0x75, 0xc2, 0xbe, 0xf4,
	0x88, 0x8b, 0xa7, 0x30, 0x5a, 0x6f, 0x57, 0x0a, 0x18, 0xf6, 0xa4, 0xf7, 0xcb, 0xcd, 0x47, 0x55,
	0x8f, 0xa2, 0xcf, 0x30, 0xb1, 0xbb, 0xd5, 0xc7, 0x1b, 0x7c, 0x0a, 0x9e, 0xc9, 0xcd, 0x21, 0x77,
	0x13, 0xb3, 0x3d, 0x05, 0x66, 0x20, 0x4d, 0xea, 0x35, 0x14, 0xee, 0x5f, 0x12, 0x8d, 0x31, 0x50,
	0x74, 0x05, 0xbe, 0x7e, 0x69, 0xff, 0x63, 0xf2, 0x33, 0x18, 0x6b, 0x79, 0x7f, 0x37, 0xf7, 0xfe,
	0x68, 0x8b, 0xbd, 0x9d, 0x5d, 0xf9, 0xff, 0xfe, 0xcb, 0xaf, 0x3d, 0xca, 0xee, 0xe5, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0xe3, 0x8d, 0x57, 0x38, 0x03, 0x00, 0x00,
}
