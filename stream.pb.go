// Code generated by protoc-gen-go.
// source: zvelo/msg/stream.proto
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
func (Stream_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{1, 0} }

type StreamRequest struct {
	CustomerTrackingId string        `protobuf:"bytes,1,opt,name=customer_tracking_id,json=customerTrackingId" json:"customer_tracking_id,omitempty"`
	Endpoint           string        `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	Accept             string        `protobuf:"bytes,3,opt,name=accept" json:"accept,omitempty"`
	Dataset            []DataSetType `protobuf:"varint,4,rep,packed,name=dataset,enum=msg.DataSetType" json:"dataset,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

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
func (*StreamReply) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

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
func (*StreamsReply) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

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

func init() { proto.RegisterFile("zvelo/msg/stream.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x35, 0x4d, 0x9b, 0xee, 0xde, 0xa6, 0x25, 0x0c, 0xb2, 0x86, 0xbe, 0xb8, 0x44, 0x84, 0xc5,
	0x87, 0x54, 0xd6, 0x07, 0xf1, 0x49, 0xaa, 0x89, 0x52, 0x90, 0x22, 0x93, 0x2c, 0xe2, 0xbe, 0x94,
	0xd9, 0x64, 0x0c, 0xc1, 0x26, 0x13, 0x67, 0x26, 0xc2, 0xfa, 0x37, 0xfe, 0x8f, 0x1f, 0x25, 0xb9,
	0x33, 0xd9, 0x65, 0x7d, 0xf3, 0x6d, 0xee, 0x9c, 0x73, 0xee, 0x9c, 0x7b, 0xee, 0xc0, 0xd9, 0xaf,
	0x9f, 0xfc, 0x28, 0x36, 0x8d, 0xaa, 0x36, 0x4a, 0x4b, 0xce, 0x9a, 0xb8, 0x93, 0x42, 0x0b, 0xe2,
	0x36, 0xaa, 0x5a, 0x3f, 0xb9, 0x07, 0x4b, 0xa6, 0x99, 0xe2, 0xda, 0xa0, 0xeb, 0x07, 0x2a, 0xa6,
	0x7b, 0x65, 0xef, 0x9f, 0x56, 0x42, 0x54, 0x47, 0xbe, 0xc1, 0xea, 0xa6, 0xff, 0xb6, 0xd1, 0x75,
	0xc3, 0x95, 0x66, 0x4d, 0x67, 0x08, 0xd1, 0x6f, 0x07, 0x96, 0x19, 0xbe, 0x43, 0xf9, 0x8f, 0x9e,
	0x2b, 0x4d, 0x5e, 0xc2, 0xe3, 0xa2, 0x57, 0x5a, 0x34, 0x5c, 0x1e, 0xb4, 0x64, 0xc5, 0xf7, 0xba,
	0xad, 0x0e, 0x75, 0x19, 0x3a, 0xe7, 0xce, 0xc5, 0x29, 0x25, 0x23, 0x96, 0x5b, 0x68, 0x57, 0x92,
	0x35, 0x9c, 0xf0, 0xb6, 0xec, 0x44, 0xdd, 0xea, 0x70, 0x82, 0xac, 0xbb, 0x9a, 0x9c, 0x81, 0xc7,
	0x8a, 0x82, 0x77, 0x3a, 0x74, 0x11, 0xb1, 0x15, 0x79, 0x01, 0x73, 0x3b, 0x41, 0x38, 0x3d, 0x77,
	0x2f, 0x56, 0x97, 0x41, 0xdc, 0xa8, 0x2a, 0x4e, 0x98, 0x66, 0x19, 0xd7, 0xf9, 0x6d, 0xc7, 0xe9,
	0x48, 0x88, 0xfe, 0x4c, 0xc0, 0x33, 0x1e, 0x09, 0x81, 0x69, 0xdf, 0xdf, 0x99, 0xc1, 0x33, 0x79,
	0x0d, 0x4b, 0x93, 0xd4, 0xc1, 0x8c, 0x8e, 0x1e, 0x56, 0x97, 0x04, 0x1b, 0x1a, 0x5d, 0x9c, 0x21,
	0x42, 0x7d, 0x43, 0x34, 0x15, 0x79, 0x03, 0x2b, 0x2b, 0x94, 0x66, 0x76, 0xf4, 0xb8, 0x78, 0xa0,
	0xb4, 0xa9, 0x50, 0xfb, 0xc4, 0x18, 0x52, 0x00, 0xae, 0x90, 0x55, 0x38, 0x43, 0x1b, 0xc3, 0x11,
	0x9d, 0x29, 0x2e, 0x43, 0xcf, 0x3a, 0x53, 0x5c, 0x92, 0xb7, 0xb0, 0x2c, 0x24, 0x67, 0xba, 0x16,
	0xed, 0xa1, 0x64, 0x9a, 0x87, 0x73, 0xec, 0xbf, 0x8e, 0xcd, 0x56, 0xe2, 0x71, 0x2b, 0x71, 0x3e,
	0x6e, 0x85, 0xfa, 0xa3, 0x20, 0x61, 0x9a, 0x47, 0xd9, 0x30, 0x38, 0x7a, 0x5d, 0xc0, 0xfc, 0x3d,
	0x4d, 0xb7, 0x79, 0x9a, 0x04, 0x8f, 0x86, 0x82, 0x5e, 0xed, 0xf7, 0xbb, 0xfd, 0xc7, 0xc0, 0x21,
	0x00, 0xde, 0xe7, 0xed, 0x55, 0x96, 0x26, 0xc1, 0x84, 0xf8, 0x70, 0x42, 0xd3, 0x9c, 0x7e, 0x1d,
	0x10, 0x97, 0x9c, 0xc2, 0x2c, 0xdd, 0x27, 0x69, 0x12, 0x4c, 0x07, 0xd2, 0x87, 0xed, 0xee, 0x53,
	0x9a, 0x04, 0xb3, 0xe8, 0x0b, 0x2c, 0xc6, 0xd9, 0xba, 0xe3, 0x2d, 0x79, 0x06, 0x9e, 0xcd, 0xcd,
	0x41, 0x77, 0x0b, 0x3b, 0x3d, 0x06, 0x66, 0x21, 0x43, 0x1a, 0x34, 0x18, 0xee, 0x3d, 0x09, 0xdb,
	0x58, 0x28, 0xba, 0x06, 0xdf, 0xdc, 0xa8, 0xff, 0xe8, 0xfc, 0x1c, 0xe6, 0x46, 0x3e, 0xec, 0xcd,
	0xfd, 0xb7, 0xf5, 0x88, 0xbd, 0x5b, 0x5d, 0xfb, 0xf8, 0xc5, 0xe3, 0x1a, 0x7f, 0xf9, 0x8d, 0x87,
	0xd9, 0xbd, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc6, 0xd4, 0xa5, 0x2f, 0x03, 0x00, 0x00,
}
