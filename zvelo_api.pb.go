// Code generated by protoc-gen-go.
// source: zvelo/msg/zvelo_api.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QueryURLRequests struct {
	Url            []string      `protobuf:"bytes,1,rep,name=url" json:"url,omitempty"`
	Dataset        []DataSetType `protobuf:"varint,2,rep,packed,name=dataset,enum=zvelo.msg.DataSetType" json:"dataset,omitempty"`
	Callback       string        `protobuf:"bytes,3,opt,name=callback" json:"callback,omitempty"`
	Poll           bool          `protobuf:"varint,4,opt,name=poll" json:"poll,omitempty"`
	PartialResults bool          `protobuf:"varint,5,opt,name=partial_results,json=partialResults" json:"partial_results,omitempty"`
	Accept         string        `protobuf:"bytes,6,opt,name=accept" json:"accept,omitempty"`
}

func (m *QueryURLRequests) Reset()                    { *m = QueryURLRequests{} }
func (m *QueryURLRequests) String() string            { return proto.CompactTextString(m) }
func (*QueryURLRequests) ProtoMessage()               {}
func (*QueryURLRequests) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type QueryContentRequests struct {
	Content        []*QueryContentRequests_URLContent `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
	Dataset        []DataSetType                      `protobuf:"varint,2,rep,packed,name=dataset,enum=zvelo.msg.DataSetType" json:"dataset,omitempty"`
	Callback       string                             `protobuf:"bytes,3,opt,name=callback" json:"callback,omitempty"`
	Poll           bool                               `protobuf:"varint,4,opt,name=poll" json:"poll,omitempty"`
	PartialResults bool                               `protobuf:"varint,5,opt,name=partial_results,json=partialResults" json:"partial_results,omitempty"`
	Accept         string                             `protobuf:"bytes,6,opt,name=accept" json:"accept,omitempty"`
}

func (m *QueryContentRequests) Reset()                    { *m = QueryContentRequests{} }
func (m *QueryContentRequests) String() string            { return proto.CompactTextString(m) }
func (*QueryContentRequests) ProtoMessage()               {}
func (*QueryContentRequests) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *QueryContentRequests) GetContent() []*QueryContentRequests_URLContent {
	if m != nil {
		return m.Content
	}
	return nil
}

type QueryContentRequests_URLContent struct {
	// Use a unique id if content has no url, to link response to request.
	Id      string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Url     string            `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Header  map[string]string `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Content string            `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *QueryContentRequests_URLContent) Reset()         { *m = QueryContentRequests_URLContent{} }
func (m *QueryContentRequests_URLContent) String() string { return proto.CompactTextString(m) }
func (*QueryContentRequests_URLContent) ProtoMessage()    {}
func (*QueryContentRequests_URLContent) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{1, 0}
}

func (m *QueryContentRequests_URLContent) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

type QueryReply struct {
	Status    *Status  `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	RequestId []string `protobuf:"bytes,2,rep,name=request_id,json=requestId" json:"request_id,omitempty"`
	// the target of QueryStatus, to ensure that the QueryStatus call
	// goes to the same cluster as the original call.
	// Will only be set if user set p=true in request.
	StatusUrl string `protobuf:"bytes,3,opt,name=status_url,json=statusUrl" json:"status_url,omitempty"`
}

func (m *QueryReply) Reset()                    { *m = QueryReply{} }
func (m *QueryReply) String() string            { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()               {}
func (*QueryReply) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *QueryReply) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryURLRequests)(nil), "zvelo.msg.QueryURLRequests")
	proto.RegisterType((*QueryContentRequests)(nil), "zvelo.msg.QueryContentRequests")
	proto.RegisterType((*QueryContentRequests_URLContent)(nil), "zvelo.msg.QueryContentRequests.URLContent")
	proto.RegisterType((*QueryReply)(nil), "zvelo.msg.QueryReply")
}

func init() { proto.RegisterFile("zvelo/msg/zvelo_api.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x53, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x25, 0xc9, 0x6c, 0x76, 0x53, 0x23, 0x71, 0x6d, 0x96, 0xb1, 0x0d, 0x08, 0x61, 0x2e, 0x46,
	0x0f, 0x59, 0x19, 0x41, 0xd4, 0xa3, 0xae, 0xa0, 0xb0, 0x08, 0xf6, 0x3a, 0x17, 0x2f, 0xa1, 0x37,
	0x69, 0xd6, 0x30, 0x3d, 0x93, 0xd8, 0xdd, 0x19, 0x88, 0x3f, 0xe8, 0x07, 0x88, 0xff, 0x23, 0xa9,
	0xee, 0x4c, 0xe6, 0xe0, 0xc1, 0xab, 0xb7, 0xaa, 0x57, 0xaf, 0x2a, 0xaf, 0x5e, 0xa5, 0xe1, 0xd1,
	0x8f, 0xbd, 0x90, 0xcd, 0xe5, 0x56, 0xdf, 0x5d, 0x62, 0x54, 0xf0, 0xb6, 0xce, 0x5b, 0xd5, 0x98,
	0x86, 0x44, 0x08, 0xe4, 0x5b, 0x7d, 0x97, 0x3c, 0x9c, 0x58, 0x15, 0x37, 0x5c, 0x0b, 0x63, 0x39,
	0xc9, 0x62, 0x2a, 0x68, 0xc3, 0x4d, 0xa7, 0x2d, 0xbe, 0xfc, 0xe9, 0xc1, 0xf9, 0xe7, 0x4e, 0xa8,
	0x7e, 0xcd, 0xae, 0x99, 0xf8, 0xde, 0x09, 0x6d, 0x34, 0x39, 0x87, 0xa0, 0x53, 0x92, 0x7a, 0x69,
	0x90, 0x45, 0x6c, 0x08, 0xc9, 0x73, 0x38, 0x75, 0xf3, 0xa8, 0x9f, 0x06, 0x59, 0xbc, 0x5a, 0xe4,
	0x87, 0x8f, 0xe6, 0x57, 0xdc, 0xf0, 0x1b, 0x61, 0xbe, 0xf4, 0xad, 0x60, 0x23, 0x8d, 0x24, 0x70,
	0x56, 0x72, 0x29, 0x6f, 0x79, 0xb9, 0xa1, 0x41, 0xea, 0x65, 0x11, 0x3b, 0xe4, 0x84, 0xc0, 0xac,
	0x6d, 0xa4, 0xa4, 0xb3, 0xd4, 0xcb, 0xce, 0x18, 0xc6, 0xe4, 0x09, 0xdc, 0x6f, 0xb9, 0x32, 0x35,
	0x97, 0x85, 0x12, 0xba, 0x93, 0x46, 0xd3, 0x13, 0x2c, 0xc7, 0x0e, 0x66, 0x16, 0x25, 0x0b, 0x08,
	0x79, 0x59, 0x8a, 0xd6, 0xd0, 0x10, 0xc7, 0xba, 0x6c, 0xf9, 0x3b, 0x80, 0x0b, 0xdc, 0xe4, 0x5d,
	0xb3, 0x33, 0x62, 0x67, 0x0e, 0xdb, 0x5c, 0xc1, 0x69, 0x69, 0x21, 0xdc, 0x68, 0xbe, 0x7a, 0x76,
	0xa4, 0xfd, 0x6f, 0x1d, 0xf9, 0x9a, 0x5d, 0x8f, 0xd0, 0xd8, 0xfa, 0x1f, 0x38, 0x90, 0xfc, 0xf2,
	0x00, 0x26, 0xe9, 0x24, 0x06, 0xbf, 0xae, 0xa8, 0x87, 0x14, 0xbf, 0xae, 0xc6, 0xab, 0xfa, 0x08,
	0xe0, 0x55, 0x3f, 0x41, 0xf8, 0x4d, 0xf0, 0x4a, 0x28, 0x1a, 0xa0, 0x31, 0x2f, 0xff, 0xdd, 0x98,
	0xfc, 0x03, 0x36, 0xbe, 0xdf, 0x19, 0xd5, 0x33, 0x37, 0x85, 0xd0, 0xc9, 0xe9, 0x19, 0x7e, 0x65,
	0x4c, 0x93, 0xd7, 0x30, 0x3f, 0x6a, 0x18, 0xa4, 0x6c, 0x44, 0xef, 0xb4, 0x0d, 0x21, 0xb9, 0x80,
	0x93, 0x3d, 0x97, 0x9d, 0x70, 0xf2, 0x6c, 0xf2, 0xc6, 0x7f, 0xe5, 0x2d, 0x3b, 0x00, 0xd4, 0xc2,
	0x44, 0x2b, 0x7b, 0xf2, 0x14, 0x42, 0xfb, 0xff, 0x62, 0xf3, 0x7c, 0xf5, 0xe0, 0x48, 0xf2, 0x0d,
	0x16, 0x98, 0x23, 0x90, 0xc7, 0x00, 0xca, 0x0a, 0x2f, 0xea, 0x0a, 0x8f, 0x16, 0xb1, 0xc8, 0x21,
	0x1f, 0xab, 0xa1, 0x6c, 0x89, 0xc5, 0xe0, 0x8a, 0x3d, 0x50, 0x64, 0x91, 0xb5, 0x92, 0x6f, 0xe3,
	0xaf, 0xf7, 0xec, 0xe4, 0x1a, 0x5f, 0xcd, 0x6d, 0x88, 0xef, 0xe5, 0xc5, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x81, 0xba, 0x01, 0xf8, 0x88, 0x03, 0x00, 0x00,
}
