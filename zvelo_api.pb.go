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
func (*QueryURLRequests) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
func (*QueryContentRequests) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

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
	return fileDescriptor3, []int{1, 0}
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
func (*QueryReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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

func init() { proto.RegisterFile("zvelo/msg/zvelo_api.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xd4, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x92, 0x6e, 0x76, 0x33, 0x45, 0x61, 0xb1, 0x56, 0x25, 0x44, 0x42, 0xaa, 0x7a, 0xe1,
	0xe3, 0x90, 0xa2, 0x22, 0x21, 0xe0, 0x08, 0x45, 0x02, 0x09, 0x21, 0xe1, 0xd2, 0x0b, 0x97, 0xc8,
	0x4d, 0x2c, 0x88, 0x6a, 0x9a, 0x60, 0x3b, 0x95, 0xc2, 0x1f, 0xe4, 0x07, 0x20, 0xfe, 0x0f, 0xce,
	0xd8, 0x69, 0x7a, 0xe0, 0xc0, 0x95, 0xdb, 0xcc, 0x9b, 0x37, 0xe3, 0x37, 0xcf, 0x36, 0xdc, 0xfb,
	0x71, 0xe4, 0xa2, 0x5e, 0x7e, 0x53, 0x5f, 0x96, 0x18, 0xe5, 0xac, 0xa9, 0xb2, 0x46, 0xd6, 0xba,
	0x26, 0x11, 0x02, 0x99, 0x29, 0xa5, 0x77, 0x47, 0x56, 0xc9, 0x34, 0x53, 0x5c, 0x5b, 0x4e, 0x3a,
	0x1b, 0x0b, 0x4a, 0x33, 0xdd, 0x2a, 0x8b, 0x2f, 0x7e, 0x7a, 0x70, 0xfd, 0xb1, 0xe5, 0xb2, 0xdb,
	0xd2, 0xf7, 0x94, 0x7f, 0x6f, 0xb9, 0xd2, 0x8a, 0x5c, 0x43, 0xd0, 0x4a, 0x91, 0x78, 0xf3, 0xe0,
	0x61, 0x44, 0xfb, 0x90, 0x3c, 0x81, 0x4b, 0x37, 0x2f, 0xf1, 0x0d, 0x1a, 0xaf, 0x66, 0xd9, 0xe9,
	0xd0, 0x6c, 0x6d, 0x2a, 0x1b, 0xae, 0x3f, 0x75, 0x0d, 0xa7, 0x03, 0x8d, 0xa4, 0x70, 0x55, 0x30,
	0x21, 0x76, 0xac, 0xd8, 0x27, 0xc1, 0xdc, 0x33, 0x83, 0x4e, 0x39, 0x21, 0x30, 0x69, 0x6a, 0x21,
	0x92, 0x89, 0xc1, 0xaf, 0x28, 0xc6, 0xe4, 0x01, 0xdc, 0x6e, 0x98, 0xd4, 0x15, 0x13, 0xb9, 0xe4,
	0xaa, 0x15, 0x5a, 0x25, 0x17, 0x58, 0x8e, 0x1d, 0x4c, 0x2d, 0x4a, 0x66, 0x10, 0xb2, 0xa2, 0xe0,
	0x8d, 0x4e, 0x42, 0x1c, 0xeb, 0xb2, 0xc5, 0xef, 0x00, 0x6e, 0x70, 0x93, 0xd7, 0xf5, 0x41, 0xf3,
	0x83, 0x3e, 0x6d, 0xb3, 0x86, 0xcb, 0xc2, 0x42, 0xb8, 0xd1, 0x74, 0xf5, 0xf8, 0x4c, 0xfb, 0xdf,
	0x3a, 0x32, 0xe3, 0xc5, 0x00, 0x0d, 0xad, 0xff, 0x81, 0x03, 0xe9, 0x2f, 0x0f, 0x60, 0x94, 0x4e,
	0x62, 0xf0, 0xab, 0xd2, 0xac, 0xdc, 0x53, 0x4c, 0x34, 0xdc, 0xaa, 0x8f, 0x00, 0xde, 0xea, 0x07,
	0x08, 0xbf, 0x72, 0x56, 0x72, 0x69, 0xf4, 0xf5, 0xc6, 0x3c, 0xfb, 0x77, 0x63, 0xb2, 0xb7, 0xd8,
	0xf8, 0xe6, 0xa0, 0x65, 0x47, 0xdd, 0x14, 0x92, 0x8c, 0x4e, 0x4f, 0xf0, 0x94, 0x21, 0x4d, 0x5f,
	0xc0, 0xf4, 0xac, 0xa1, 0x97, 0xb2, 0xe7, 0x9d, 0xd3, 0xd6, 0x87, 0xe4, 0x06, 0x2e, 0x8e, 0x4c,
	0xb4, 0xdc, 0xc9, 0xb3, 0xc9, 0x4b, 0xff, 0xb9, 0xb7, 0x68, 0x01, 0x50, 0x0b, 0xe5, 0x8d, 0xe8,
	0xc8, 0x23, 0x08, 0xed, 0xfb, 0xc5, 0xe6, 0xe9, 0xea, 0xce, 0x99, 0xe4, 0x0d, 0x16, 0xa8, 0x23,
	0x90, 0xfb, 0x00, 0xd2, 0x0a, 0xcf, 0x8d, 0x0f, 0x3e, 0x3e, 0xe6, 0xc8, 0x21, 0xef, 0xca, 0xbe,
	0x6c, 0x89, 0x79, 0xef, 0x8a, 0xbd, 0xa0, 0xc8, 0x22, 0x5b, 0x29, 0x5e, 0xc5, 0x9f, 0x6f, 0xd9,
	0xc9, 0x15, 0xfe, 0x9a, 0x5d, 0x88, 0xff, 0xe5, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81,
	0xba, 0x01, 0xf8, 0x88, 0x03, 0x00, 0x00,
}
