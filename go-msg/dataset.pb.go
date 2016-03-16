// Code generated by protoc-gen-go.
// source: dataset.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	dataset.proto
	query_result.proto
	status.proto
	zvelo-api.proto

It has these top-level messages:
	DataSet
	QueryResult
	Status
	QueryURLRequests
	QueryContentRequests
	QueryReply
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type DataSetType int32

const (
	DataSetType_CATEGORIZATION DataSetType = 0
	DataSetType_ADFRAUD        DataSetType = 1
	// 2 and 3 are reserved
	DataSetType_MALICIOUS DataSetType = 4
)

var DataSetType_name = map[int32]string{
	0: "CATEGORIZATION",
	1: "ADFRAUD",
	4: "MALICIOUS",
}
var DataSetType_value = map[string]int32{
	"CATEGORIZATION": 0,
	"ADFRAUD":        1,
	"MALICIOUS":      4,
}

func (x DataSetType) String() string {
	return proto.EnumName(DataSetType_name, int32(x))
}
func (DataSetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DataSet struct {
	Categorization *DataSet_Categorization `protobuf:"bytes,1,opt,name=categorization" json:"categorization,omitempty"`
	Adfraud        *DataSet_AdFraud        `protobuf:"bytes,2,opt,name=adfraud" json:"adfraud,omitempty"`
	// 3 and 4 are reserved
	Malicious *DataSet_Malicious `protobuf:"bytes,5,opt,name=malicious" json:"malicious,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DataSet) GetCategorization() *DataSet_Categorization {
	if m != nil {
		return m.Categorization
	}
	return nil
}

func (m *DataSet) GetAdfraud() *DataSet_AdFraud {
	if m != nil {
		return m.Adfraud
	}
	return nil
}

func (m *DataSet) GetMalicious() *DataSet_Malicious {
	if m != nil {
		return m.Malicious
	}
	return nil
}

type DataSet_Categorization struct {
	Id []int32 `protobuf:"varint,1,rep,name=id" json:"id,omitempty"`
}

func (m *DataSet_Categorization) Reset()                    { *m = DataSet_Categorization{} }
func (m *DataSet_Categorization) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Categorization) ProtoMessage()               {}
func (*DataSet_Categorization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type DataSet_AdFraud struct {
	Fraud     bool   `protobuf:"varint,1,opt,name=fraud" json:"fraud,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *DataSet_AdFraud) Reset()                    { *m = DataSet_AdFraud{} }
func (m *DataSet_AdFraud) String() string            { return proto.CompactTextString(m) }
func (*DataSet_AdFraud) ProtoMessage()               {}
func (*DataSet_AdFraud) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type DataSet_Malicious struct {
	Category  int32  `protobuf:"varint,1,opt,name=category" json:"category,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Verdict   bool   `protobuf:"varint,3,opt,name=verdict" json:"verdict,omitempty"`
}

func (m *DataSet_Malicious) Reset()                    { *m = DataSet_Malicious{} }
func (m *DataSet_Malicious) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Malicious) ProtoMessage()               {}
func (*DataSet_Malicious) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func init() {
	proto.RegisterType((*DataSet)(nil), "msg.DataSet")
	proto.RegisterType((*DataSet_Categorization)(nil), "msg.DataSet.Categorization")
	proto.RegisterType((*DataSet_AdFraud)(nil), "msg.DataSet.AdFraud")
	proto.RegisterType((*DataSet_Malicious)(nil), "msg.DataSet.Malicious")
	proto.RegisterEnum("msg.DataSetType", DataSetType_name, DataSetType_value)
}

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x24, 0x5f, 0x4c, 0x33, 0xa1, 0x35, 0x0e, 0x22, 0x21, 0x7a, 0x10, 0x41, 0xa8,
	0x0a, 0x39, 0xd8, 0xb3, 0x87, 0xa5, 0xb1, 0x12, 0xb0, 0x06, 0xda, 0xf4, 0xe2, 0x6d, 0xcd, 0xae,
	0x61, 0xc1, 0x36, 0x25, 0xd9, 0x08, 0xf5, 0xc7, 0xf8, 0x5b, 0xdd, 0xac, 0xa9, 0x10, 0xf1, 0xba,
	0xf3, 0xcc, 0xfb, 0x3e, 0x3b, 0x30, 0x64, 0x54, 0xd2, 0x9a, 0xcb, 0x68, 0x5b, 0x95, 0xb2, 0x44,
	0x6b, 0x5d, 0x17, 0x17, 0x9f, 0x26, 0x38, 0xb1, 0x7a, 0x5e, 0x72, 0x89, 0x13, 0x18, 0xe5, 0x54,
	0xf2, 0xa2, 0xac, 0xc4, 0x07, 0x95, 0xa2, 0xdc, 0x04, 0xc6, 0xb9, 0x31, 0xf6, 0x6e, 0x4f, 0x23,
	0x45, 0x46, 0x1d, 0x15, 0x4d, 0x7b, 0x08, 0x5e, 0x82, 0x43, 0xd9, 0x6b, 0x45, 0x1b, 0x16, 0x98,
	0x9a, 0x3e, 0xee, 0xd1, 0x84, 0xcd, 0xda, 0x19, 0x5e, 0x81, 0xbb, 0xa6, 0x6f, 0x22, 0x17, 0x65,
	0x53, 0x07, 0xb6, 0x06, 0x4f, 0x7a, 0xe0, 0x7c, 0x3f, 0x0d, 0xcf, 0x60, 0xf4, 0xab, 0x03, 0xc0,
	0x14, 0x4c, 0xc9, 0x58, 0x63, 0x3b, 0xbc, 0x01, 0x67, 0x9f, 0x39, 0x04, 0xfb, 0xbb, 0xb8, 0xd5,
	0x1c, 0xe0, 0x11, 0xb8, 0xb5, 0x28, 0x36, 0x54, 0x36, 0x15, 0xd7, 0x2e, 0x6e, 0x48, 0xc0, 0xfd,
	0xc9, 0x45, 0x1f, 0x06, 0xdd, 0xf7, 0x76, 0x7a, 0xc3, 0xfe, 0x63, 0x03, 0x0f, 0xc1, 0x79, 0xe7,
	0x15, 0x13, 0xb9, 0x0c, 0xac, 0x36, 0xf5, 0xfa, 0x0e, 0xbc, 0x4e, 0x31, 0xdb, 0x6d, 0x39, 0xa2,
	0x92, 0x23, 0xd9, 0xfd, 0x43, 0xba, 0x48, 0x9e, 0x49, 0x96, 0xa4, 0x4f, 0xfe, 0x3f, 0xf4, 0x94,
	0x52, 0x3c, 0x5b, 0x90, 0x55, 0xec, 0x1b, 0x4a, 0xca, 0x9d, 0x93, 0xc7, 0x64, 0x9a, 0xa4, 0xab,
	0xa5, 0xff, 0xff, 0xe5, 0x40, 0xdf, 0x7a, 0xf2, 0x15, 0x00, 0x00, 0xff, 0xff, 0x80, 0x5b, 0x96,
	0x8e, 0x7c, 0x01, 0x00, 0x00,
}
