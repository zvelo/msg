// Code generated by protoc-gen-go.
// source: zvelo-api.proto
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
	Dataset        []DataSetType `protobuf:"varint,2,rep,name=dataset,enum=msg.DataSetType" json:"dataset,omitempty"`
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
	Dataset        []DataSetType                      `protobuf:"varint,2,rep,name=dataset,enum=msg.DataSetType" json:"dataset,omitempty"`
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
	proto.RegisterType((*QueryURLRequests)(nil), "msg.QueryURLRequests")
	proto.RegisterType((*QueryContentRequests)(nil), "msg.QueryContentRequests")
	proto.RegisterType((*QueryContentRequests_URLContent)(nil), "msg.QueryContentRequests.URLContent")
	proto.RegisterType((*QueryReply)(nil), "msg.QueryReply")
}

var fileDescriptor3 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x93, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0x64, 0xcb, 0xf6, 0xa8, 0xb5, 0xcd, 0x62, 0x8a, 0x10, 0x14, 0x8c, 0x5b, 0x68,
	0x29, 0x54, 0x14, 0xf7, 0xd2, 0xf6, 0xd0, 0x4b, 0x5b, 0x70, 0x21, 0x97, 0xac, 0xe3, 0xb3, 0x59,
	0x4b, 0x4b, 0x62, 0xbc, 0xb1, 0x94, 0xdd, 0x95, 0x41, 0x79, 0xbb, 0x90, 0x73, 0xde, 0x29, 0xab,
	0xd9, 0x95, 0x9d, 0x43, 0x0e, 0x39, 0xe6, 0x36, 0xf3, 0xed, 0xcc, 0xf0, 0xff, 0x33, 0x12, 0x8c,
	0x6e, 0x0f, 0x5c, 0x14, 0x5f, 0x59, 0xb9, 0x4d, 0x4b, 0x59, 0xe8, 0x82, 0x04, 0xd7, 0xea, 0x32,
	0x79, 0x9b, 0x33, 0xcd, 0x14, 0xd7, 0x96, 0x25, 0x6f, 0x94, 0x66, 0xba, 0x52, 0x36, 0x9b, 0xdd,
	0x79, 0x30, 0x3e, 0xaf, 0xb8, 0xac, 0x57, 0xf4, 0x8c, 0xf2, 0x9b, 0x8a, 0x2b, 0xad, 0xc8, 0x18,
	0x82, 0x4a, 0x8a, 0xd8, 0x9b, 0x06, 0x9f, 0x07, 0xb4, 0x09, 0xc9, 0x17, 0xe8, 0xb9, 0x29, 0xb1,
	0x6f, 0xe8, 0x70, 0x3e, 0x4e, 0xcd, 0xe8, 0xf4, 0xaf, 0x61, 0x4b, 0xae, 0x2f, 0xea, 0x92, 0xd3,
	0xb6, 0x80, 0x24, 0xd0, 0xcf, 0x98, 0x10, 0x1b, 0x96, 0xed, 0xe2, 0x60, 0xea, 0x99, 0x11, 0xc7,
	0x9c, 0x10, 0xe8, 0x94, 0x85, 0x10, 0x71, 0xc7, 0xf0, 0x3e, 0xc5, 0x98, 0x7c, 0x82, 0x51, 0xc9,
	0xa4, 0xde, 0x32, 0xb1, 0x96, 0x5c, 0x55, 0x42, 0xab, 0xb8, 0x8b, 0xcf, 0x43, 0x87, 0xa9, 0xa5,
	0xe4, 0x1d, 0x84, 0x2c, 0xcb, 0x78, 0xa9, 0xe3, 0x10, 0xc7, 0xba, 0x6c, 0x76, 0x1f, 0xc0, 0x04,
	0x3d, 0xfc, 0x29, 0xf6, 0x9a, 0xef, 0xf5, 0xd1, 0xc7, 0x6f, 0xe8, 0x65, 0x16, 0xa1, 0x97, 0x68,
	0xfe, 0x11, 0x55, 0x3f, 0x57, 0x9b, 0x1a, 0xff, 0x2d, 0x6a, 0x9b, 0x5e, 0xb5, 0xeb, 0xe4, 0xc1,
	0x03, 0x38, 0x89, 0x26, 0x43, 0xf0, 0xb7, 0xb9, 0xb1, 0xd9, 0x94, 0x98, 0xa8, 0xbd, 0xa1, 0x8f,
	0x00, 0x6f, 0xb8, 0x80, 0xf0, 0x8a, 0xb3, 0x9c, 0x4b, 0xa3, 0xaf, 0x59, 0xc6, 0xb7, 0x97, 0x2c,
	0x23, 0x5d, 0x60, 0xcb, 0xbf, 0xbd, 0x96, 0x35, 0x75, 0xfd, 0x24, 0x3e, 0xed, 0xb5, 0x83, 0xf3,
	0xdb, 0x34, 0xf9, 0x09, 0xd1, 0x93, 0x86, 0x46, 0xc4, 0x8e, 0xd7, 0x4e, 0x55, 0x13, 0x92, 0x09,
	0x74, 0x0f, 0x4c, 0x54, 0xdc, 0x09, 0xb3, 0xc9, 0x2f, 0xff, 0x87, 0x37, 0x2b, 0x00, 0x50, 0x0b,
	0xe5, 0xa5, 0xa8, 0xc9, 0x07, 0x08, 0xed, 0x77, 0x8a, 0xcd, 0xd1, 0x3c, 0x42, 0xb1, 0x4b, 0x44,
	0xd4, 0x3d, 0x91, 0xf7, 0x00, 0xd2, 0x4a, 0x5e, 0x1b, 0xef, 0x3e, 0x7e, 0xae, 0x03, 0x47, 0xfe,
	0xe7, 0xcd, 0xb3, 0x2d, 0x5c, 0x37, 0x9b, 0xb0, 0x47, 0x19, 0x58, 0xb2, 0x92, 0x62, 0x13, 0xe2,
	0x1f, 0xf0, 0xfd, 0x31, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xb0, 0x49, 0x26, 0x36, 0x03, 0x00, 0x00,
}
