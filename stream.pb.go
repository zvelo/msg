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
func (Stream_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1, 0} }

type StreamRequest struct {
	CustomerTrackingId string        `protobuf:"bytes,1,opt,name=customer_tracking_id,json=customerTrackingId" json:"customer_tracking_id,omitempty"`
	Endpoint           string        `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	Accept             string        `protobuf:"bytes,3,opt,name=accept" json:"accept,omitempty"`
	Dataset            []DataSetType `protobuf:"varint,4,rep,packed,name=dataset,enum=zvelo.msg.DataSetType" json:"dataset,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Stream struct {
	Uuid          string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	StreamStatus  Stream_Status              `protobuf:"varint,2,opt,name=stream_status,json=streamStatus,enum=zvelo.msg.Stream_Status" json:"stream_status,omitempty"`
	StreamRequest *StreamRequest             `protobuf:"bytes,3,opt,name=stream_request,json=streamRequest" json:"stream_request,omitempty"`
	Org           string                     `protobuf:"bytes,5,opt,name=org" json:"org,omitempty"`
	User          string                     `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	CreationDate  *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

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
func (*StreamReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

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
func (*StreamsReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

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
	proto.RegisterType((*StreamRequest)(nil), "zvelo.msg.StreamRequest")
	proto.RegisterType((*Stream)(nil), "zvelo.msg.Stream")
	proto.RegisterType((*StreamReply)(nil), "zvelo.msg.StreamReply")
	proto.RegisterType((*StreamsReply)(nil), "zvelo.msg.StreamsReply")
	proto.RegisterEnum("zvelo.msg.Stream_Status", Stream_Status_name, Stream_Status_value)
}

func init() { proto.RegisterFile("zvelo/msg/stream.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x69, 0xd3, 0xa6, 0xdb, 0x69, 0x5a, 0x05, 0x0b, 0x95, 0xa8, 0x17, 0x56, 0x39, 0x2d,
	0x42, 0x4a, 0x57, 0xcb, 0x19, 0xad, 0x0a, 0x09, 0xa8, 0x12, 0xaa, 0x90, 0x93, 0x3d, 0xc0, 0xa5,
	0xf2, 0x26, 0xde, 0x28, 0xa2, 0xa9, 0x83, 0xed, 0x20, 0x2d, 0x2f, 0xc4, 0x2b, 0xf1, 0x38, 0x28,
	0x63, 0xa7, 0xb0, 0x68, 0x2f, 0xdc, 0x32, 0x9e, 0xef, 0xf7, 0xfc, 0xfe, 0x47, 0x81, 0xe5, 0x8f,
	0xef, 0xfc, 0x20, 0xd6, 0xb5, 0x2a, 0xd7, 0x4a, 0x4b, 0xce, 0xea, 0xa8, 0x91, 0x42, 0x0b, 0x32,
	0xc5, 0xf3, 0xa8, 0x56, 0xe5, 0xea, 0xf9, 0x1f, 0xa4, 0x60, 0x9a, 0x29, 0xae, 0x0d, 0xb3, 0x7a,
	0xa0, 0x65, 0xba, 0x55, 0xf6, 0xfc, 0x45, 0x29, 0x44, 0x79, 0xe0, 0x6b, 0xac, 0x6e, 0xdb, 0xbb,
	0xb5, 0xae, 0x6a, 0xae, 0x34, 0xab, 0x1b, 0x03, 0x84, 0x3f, 0x07, 0x30, 0x4f, 0x71, 0x1a, 0xe5,
	0xdf, 0x5a, 0xae, 0x34, 0xb9, 0x84, 0x67, 0x79, 0xab, 0xb4, 0xa8, 0xb9, 0xdc, 0x6b, 0xc9, 0xf2,
	0xaf, 0xd5, 0xb1, 0xdc, 0x57, 0x45, 0x30, 0x38, 0x1f, 0x5c, 0x4c, 0x29, 0xe9, 0x7b, 0x99, 0x6d,
	0x6d, 0x0b, 0xb2, 0x82, 0x33, 0x7e, 0x2c, 0x1a, 0x51, 0x1d, 0x75, 0x30, 0x44, 0xea, 0x54, 0x93,
	0x25, 0xb8, 0x2c, 0xcf, 0x79, 0xa3, 0x03, 0x07, 0x3b, 0xb6, 0x22, 0x97, 0x30, 0xb1, 0x2f, 0x08,
	0x46, 0xe7, 0xce, 0xc5, 0xe2, 0x6a, 0x19, 0x9d, 0x9e, 0x19, 0xc5, 0x4c, 0xb3, 0x94, 0xeb, 0xec,
	0xbe, 0xe1, 0xb4, 0xc7, 0xc2, 0x5f, 0x43, 0x70, 0x8d, 0x53, 0x42, 0x60, 0xd4, 0xb6, 0x27, 0x4b,
	0xf8, 0x4d, 0xde, 0xc0, 0xdc, 0xa4, 0xb6, 0x37, 0x01, 0xa0, 0x93, 0xc5, 0x55, 0xf0, 0xd7, 0xb5,
	0x46, 0x1d, 0xa5, 0xd8, 0xa7, 0x9e, 0xc1, 0x4d, 0x45, 0xae, 0x61, 0x61, 0xe5, 0xd2, 0xe4, 0x80,
	0x7e, 0x67, 0x8f, 0xe8, 0x6d, 0x4e, 0xd4, 0x8e, 0xeb, 0x63, 0xf3, 0xc1, 0x11, 0xb2, 0x0c, 0xc6,
	0x68, 0xa9, 0xfb, 0x44, 0x97, 0x8a, 0xcb, 0xc0, 0xb5, 0x2e, 0x15, 0x97, 0xe4, 0x1a, 0xe6, 0xb9,
	0xe4, 0x4c, 0x57, 0xe2, 0xb8, 0x2f, 0x98, 0xe6, 0xc1, 0x04, 0xa7, 0xac, 0x22, 0xb3, 0xa7, 0xa8,
	0xdf, 0x53, 0x94, 0xf5, 0x7b, 0xa2, 0x5e, 0x2f, 0x88, 0x99, 0xe6, 0x61, 0xda, 0x85, 0x80, 0x8e,
	0x67, 0x30, 0x79, 0x47, 0x93, 0x4d, 0x96, 0xc4, 0xfe, 0x93, 0xae, 0xa0, 0x37, 0xbb, 0xdd, 0x76,
	0xf7, 0xc1, 0x1f, 0x10, 0x00, 0xf7, 0xd3, 0xe6, 0x26, 0x4d, 0x62, 0x7f, 0x48, 0x3c, 0x38, 0xa3,
	0x49, 0x46, 0x3f, 0x77, 0x1d, 0x87, 0x4c, 0x61, 0x9c, 0xec, 0xe2, 0x24, 0xf6, 0x47, 0x1d, 0xf4,
	0x7e, 0xb3, 0xfd, 0x98, 0xc4, 0xfe, 0x38, 0xcc, 0x61, 0xd6, 0xbf, 0xad, 0x39, 0xdc, 0x93, 0x97,
	0xe0, 0xda, 0x0c, 0x07, 0xe8, 0xee, 0xe9, 0x83, 0x0c, 0x30, 0x3c, 0x0b, 0x18, 0xb4, 0x53, 0x62,
	0xdc, 0xff, 0xa2, 0x78, 0xa5, 0x05, 0xc2, 0x3b, 0xf0, 0xcc, 0x89, 0xfa, 0xef, 0x29, 0xaf, 0x60,
	0x62, 0x2e, 0xe9, 0xb6, 0xea, 0x3c, 0x3e, 0xa6, 0x27, 0xde, 0x2e, 0xbe, 0x78, 0xa6, 0x59, 0xe1,
	0xff, 0x70, 0xeb, 0x62, 0xa6, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x33, 0x86, 0xc6, 0x68,
	0x5f, 0x03, 0x00, 0x00,
}
