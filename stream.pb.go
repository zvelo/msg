// Code generated by protoc-gen-go.
// source: zvelo/msg/stream.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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
func (*StreamResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *StreamResult) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type StreamParams struct {
}

func (m *StreamParams) Reset()                    { *m = StreamParams{} }
func (m *StreamParams) String() string            { return proto.CompactTextString(m) }
func (*StreamParams) ProtoMessage()               {}
func (*StreamParams) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*StreamResult)(nil), "zvelo.msg.StreamResult")
	proto.RegisterType((*StreamParams)(nil), "zvelo.msg.StreamParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Stream service

type StreamClient interface {
	Events(ctx context.Context, in *StreamParams, opts ...grpc.CallOption) (Stream_EventsClient, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Events(ctx context.Context, in *StreamParams, opts ...grpc.CallOption) (Stream_EventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Stream_serviceDesc.Streams[0], c.cc, "/zvelo.msg.Stream/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_EventsClient interface {
	Recv() (*StreamResult, error)
	grpc.ClientStream
}

type streamEventsClient struct {
	grpc.ClientStream
}

func (x *streamEventsClient) Recv() (*StreamResult, error) {
	m := new(StreamResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Stream service

type StreamServer interface {
	Events(*StreamParams, Stream_EventsServer) error
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).Events(m, &streamEventsServer{stream})
}

type Stream_EventsServer interface {
	Send(*StreamResult) error
	grpc.ServerStream
}

type streamEventsServer struct {
	grpc.ServerStream
}

func (x *streamEventsServer) Send(m *StreamResult) error {
	return x.ServerStream.SendMsg(m)
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zvelo.msg.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _Stream_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("zvelo/msg/stream.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xab, 0x2a, 0x4b, 0xcd,
	0xc9, 0xd7, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x8b, 0xeb, 0x01, 0xc5, 0xa5, 0xc4, 0x11, 0x4a, 0x52, 0x12, 0x4b,
	0x12, 0x8b, 0x53, 0x4b, 0x20, 0x6a, 0x94, 0xfc, 0xb8, 0x78, 0x82, 0xc1, 0x7a, 0x82, 0x52, 0x8b,
	0x4b, 0x73, 0x4a, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x8b, 0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x40, 0x4c, 0x21, 0x1d, 0x2e, 0x76, 0xa8, 0x16, 0x09, 0x26, 0xa0, 0x28, 0xb7, 0x91, 0x90,
	0x1e, 0xdc, 0x5c, 0x3d, 0x17, 0xa0, 0x4c, 0x70, 0x6a, 0x49, 0x10, 0x4c, 0x89, 0x12, 0x1f, 0xcc,
	0xbc, 0x80, 0xc4, 0xa2, 0xc4, 0xdc, 0x62, 0x23, 0x0f, 0x2e, 0x36, 0x08, 0x5f, 0xc8, 0x8e, 0x8b,
	0xcd, 0xb5, 0x2c, 0x35, 0xaf, 0xa4, 0x58, 0x48, 0x1c, 0xc9, 0x00, 0x64, 0xc5, 0x52, 0x98, 0x12,
	0x10, 0x57, 0x29, 0x31, 0x18, 0x30, 0x3a, 0xf1, 0x45, 0xf1, 0x40, 0x64, 0x33, 0xc1, 0xfe, 0x48,
	0x62, 0x03, 0x7b, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xbd, 0x3a, 0x9c, 0xfe, 0x00,
	0x00, 0x00,
}
