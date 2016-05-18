// Code generated by protoc-gen-go.
// source: zvelo.io/msg/zvelo-api.proto
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
func (*QueryURLRequests) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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
	proto.RegisterType((*QueryURLRequests)(nil), "msg.QueryURLRequests")
	proto.RegisterType((*QueryContentRequests)(nil), "msg.QueryContentRequests")
	proto.RegisterType((*QueryContentRequests_URLContent)(nil), "msg.QueryContentRequests.URLContent")
	proto.RegisterType((*QueryReply)(nil), "msg.QueryReply")
}

var fileDescriptor6 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x53, 0xcf, 0xcb, 0xd3, 0x40,
	0x10, 0x25, 0x49, 0xbf, 0x7c, 0xed, 0x44, 0x62, 0x59, 0x8a, 0xc4, 0xa0, 0x50, 0xaa, 0xa0, 0x08,
	0xa6, 0x52, 0x2f, 0xea, 0xc1, 0x83, 0x3f, 0xa0, 0x82, 0x17, 0xb7, 0xf6, 0xe2, 0xa5, 0x6c, 0x93,
	0x45, 0x43, 0xd7, 0x26, 0xee, 0x6e, 0x0a, 0xf1, 0xbf, 0x13, 0xcf, 0xfe, 0x4f, 0x6e, 0x66, 0x37,
	0xfd, 0x01, 0x1e, 0x3c, 0x7e, 0xb7, 0x99, 0x37, 0x6f, 0x86, 0xf7, 0x66, 0x76, 0xe1, 0xde, 0xcf,
	0x03, 0x17, 0x55, 0x56, 0x56, 0xf3, 0xef, 0xea, 0xeb, 0x1c, 0x93, 0xa7, 0xac, 0x2e, 0xb3, 0x5a,
	0x56, 0xba, 0x22, 0x81, 0x01, 0xd3, 0xf4, 0x82, 0x52, 0x30, 0xcd, 0x14, 0xd7, 0x96, 0x90, 0xde,
	0xbd, 0xa8, 0x29, 0xcd, 0x74, 0xa3, 0x6c, 0x69, 0xf6, 0xcb, 0x83, 0xf1, 0xa7, 0x86, 0xcb, 0x76,
	0x4d, 0x3f, 0x52, 0xfe, 0xa3, 0xe1, 0x4a, 0x2b, 0x32, 0x86, 0xa0, 0x91, 0x22, 0xf1, 0xa6, 0xc1,
	0xe3, 0x11, 0xed, 0x42, 0xf2, 0x04, 0xae, 0xdd, 0xc8, 0xc4, 0x37, 0x68, 0xbc, 0x18, 0x67, 0x66,
	0x54, 0xf6, 0xce, 0x60, 0x2b, 0xae, 0x3f, 0xb7, 0x35, 0xa7, 0x3d, 0x81, 0xa4, 0x30, 0xcc, 0x99,
	0x10, 0x5b, 0x96, 0xef, 0x92, 0x60, 0xea, 0x99, 0x11, 0xc7, 0x9c, 0x10, 0x18, 0xd4, 0x95, 0x10,
	0xc9, 0xc0, 0xe0, 0x43, 0x8a, 0x31, 0x79, 0x04, 0xb7, 0x6b, 0x26, 0x75, 0xc9, 0xc4, 0x46, 0x72,
	0xd5, 0x08, 0xad, 0x92, 0x2b, 0x2c, 0xc7, 0x0e, 0xa6, 0x16, 0x25, 0x77, 0x20, 0x64, 0x79, 0xce,
	0x6b, 0x9d, 0x84, 0x38, 0xd6, 0x65, 0xb3, 0xdf, 0x01, 0x4c, 0xd0, 0xc3, 0xdb, 0x6a, 0xaf, 0xf9,
	0x5e, 0x1f, 0x7d, 0xbc, 0x86, 0xeb, 0xdc, 0x42, 0xe8, 0x25, 0x5a, 0x3c, 0x44, 0xd5, 0xff, 0xe2,
	0x66, 0xc6, 0x7f, 0x0f, 0xf5, 0x4d, 0x37, 0xda, 0x75, 0xfa, 0xc7, 0x03, 0x38, 0x89, 0x26, 0x31,
	0xf8, 0x65, 0x61, 0x6c, 0x76, 0x14, 0x13, 0xf5, 0x37, 0xf4, 0x11, 0xc0, 0x1b, 0x2e, 0x21, 0xfc,
	0xc6, 0x59, 0xc1, 0xa5, 0xd1, 0xd7, 0x2d, 0xe3, 0xd9, 0xff, 0x2c, 0x23, 0x5b, 0x62, 0xcb, 0xfb,
	0xbd, 0x96, 0x2d, 0x75, 0xfd, 0x24, 0x39, 0xed, 0x75, 0x80, 0xf3, 0xfb, 0x34, 0x7d, 0x09, 0xd1,
	0x59, 0x43, 0x27, 0x62, 0xc7, 0x5b, 0xa7, 0xaa, 0x0b, 0xc9, 0x04, 0xae, 0x0e, 0x4c, 0x34, 0xdc,
	0x09, 0xb3, 0xc9, 0x2b, 0xff, 0x85, 0x37, 0xab, 0x00, 0x50, 0x0b, 0xe5, 0xb5, 0x68, 0xc9, 0x03,
	0x08, 0xed, 0x3b, 0xc5, 0xe6, 0x68, 0x11, 0xa1, 0xd8, 0x15, 0x42, 0xd4, 0x95, 0xc8, 0x7d, 0x00,
	0x69, 0x25, 0x6f, 0x8c, 0x77, 0x1f, 0x9f, 0xeb, 0xc8, 0x21, 0x1f, 0x8a, 0xae, 0x6c, 0x89, 0x9b,
	0x6e, 0x13, 0xf6, 0x28, 0x23, 0x8b, 0xac, 0xa5, 0x78, 0x13, 0x7f, 0xb9, 0x75, 0xfe, 0x2f, 0xb6,
	0x21, 0xfe, 0x88, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xca, 0x73, 0x78, 0x39, 0x6d, 0x03,
	0x00, 0x00,
}