// Code generated by protoc-gen-go.
// source: zvelo/msg/dataset.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DataSetType int32

const (
	DataSetType_CATEGORIZATION DataSetType = 0
	DataSetType_ADFRAUD        DataSetType = 1
	// 2 is reserved
	DataSetType_KEYWORD   DataSetType = 3
	DataSetType_MALICIOUS DataSetType = 4
	DataSetType_ECHO      DataSetType = 5
	DataSetType_SENTIMENT DataSetType = 6
)

var DataSetType_name = map[int32]string{
	0: "CATEGORIZATION",
	1: "ADFRAUD",
	3: "KEYWORD",
	4: "MALICIOUS",
	5: "ECHO",
	6: "SENTIMENT",
}
var DataSetType_value = map[string]int32{
	"CATEGORIZATION": 0,
	"ADFRAUD":        1,
	"KEYWORD":        3,
	"MALICIOUS":      4,
	"ECHO":           5,
	"SENTIMENT":      6,
}

func (x DataSetType) String() string {
	return proto.EnumName(DataSetType_name, int32(x))
}
func (DataSetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type DataSet struct {
	Categorization *DataSet_Categorization `protobuf:"bytes,1,opt,name=categorization" json:"categorization,omitempty"`
	Adfraud        *DataSet_AdFraud        `protobuf:"bytes,2,opt,name=adfraud" json:"adfraud,omitempty"`
	// 3 is reserved
	// We expect that only keyword or sentiment will
	// be populated at a given point in time to not
	// duplicate data. If sentiment is not nil, then
	// it contains all of the keywords in its map,
	// and as such we don't have to populate Keyword.
	Keyword   *DataSet_Keyword   `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	Malicious *DataSet_Malicious `protobuf:"bytes,5,opt,name=malicious" json:"malicious,omitempty"`
	Echo      *DataSet_Echo      `protobuf:"bytes,6,opt,name=echo" json:"echo,omitempty"`
	Sentiment *DataSet_Sentiment `protobuf:"bytes,7,opt,name=sentiment" json:"sentiment,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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

func (m *DataSet) GetKeyword() *DataSet_Keyword {
	if m != nil {
		return m.Keyword
	}
	return nil
}

func (m *DataSet) GetMalicious() *DataSet_Malicious {
	if m != nil {
		return m.Malicious
	}
	return nil
}

func (m *DataSet) GetEcho() *DataSet_Echo {
	if m != nil {
		return m.Echo
	}
	return nil
}

func (m *DataSet) GetSentiment() *DataSet_Sentiment {
	if m != nil {
		return m.Sentiment
	}
	return nil
}

type DataSet_Categorization struct {
	Values []Category `protobuf:"varint,1,rep,packed,name=values,enum=msg.Category" json:"values,omitempty"`
}

func (m *DataSet_Categorization) Reset()                    { *m = DataSet_Categorization{} }
func (m *DataSet_Categorization) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Categorization) ProtoMessage()               {}
func (*DataSet_Categorization) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type DataSet_AdFraud struct {
	Verdict   bool   `protobuf:"varint,1,opt,name=verdict" json:"verdict,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *DataSet_AdFraud) Reset()                    { *m = DataSet_AdFraud{} }
func (m *DataSet_AdFraud) String() string            { return proto.CompactTextString(m) }
func (*DataSet_AdFraud) ProtoMessage()               {}
func (*DataSet_AdFraud) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

type DataSet_Sentiment struct {
	Values map[string]float32 `protobuf:"bytes,1,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
}

func (m *DataSet_Sentiment) Reset()                    { *m = DataSet_Sentiment{} }
func (m *DataSet_Sentiment) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Sentiment) ProtoMessage()               {}
func (*DataSet_Sentiment) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func (m *DataSet_Sentiment) GetValues() map[string]float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type DataSet_Keyword struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *DataSet_Keyword) Reset()                    { *m = DataSet_Keyword{} }
func (m *DataSet_Keyword) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Keyword) ProtoMessage()               {}
func (*DataSet_Keyword) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 3} }

type DataSet_Malicious struct {
	Category MaliciousCategory `protobuf:"varint,1,opt,name=category,enum=msg.MaliciousCategory" json:"category,omitempty"`
	// 2 reserved for signature
	Verdict bool `protobuf:"varint,3,opt,name=verdict" json:"verdict,omitempty"`
}

func (m *DataSet_Malicious) Reset()                    { *m = DataSet_Malicious{} }
func (m *DataSet_Malicious) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Malicious) ProtoMessage()               {}
func (*DataSet_Malicious) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 4} }

type DataSet_Echo struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *DataSet_Echo) Reset()                    { *m = DataSet_Echo{} }
func (m *DataSet_Echo) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Echo) ProtoMessage()               {}
func (*DataSet_Echo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 5} }

func init() {
	proto.RegisterType((*DataSet)(nil), "msg.DataSet")
	proto.RegisterType((*DataSet_Categorization)(nil), "msg.DataSet.Categorization")
	proto.RegisterType((*DataSet_AdFraud)(nil), "msg.DataSet.AdFraud")
	proto.RegisterType((*DataSet_Sentiment)(nil), "msg.DataSet.Sentiment")
	proto.RegisterType((*DataSet_Keyword)(nil), "msg.DataSet.Keyword")
	proto.RegisterType((*DataSet_Malicious)(nil), "msg.DataSet.Malicious")
	proto.RegisterType((*DataSet_Echo)(nil), "msg.DataSet.Echo")
	proto.RegisterEnum("msg.DataSetType", DataSetType_name, DataSetType_value)
}

func init() { proto.RegisterFile("zvelo/msg/dataset.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0x5d, 0x6f, 0xd3, 0x3c,
	0x1c, 0xc5, 0x9f, 0xf4, 0x2d, 0xcd, 0xbf, 0xcf, 0xa2, 0x60, 0x4d, 0xc3, 0x0a, 0x5c, 0x8c, 0x4a,
	0x93, 0x26, 0x2e, 0x32, 0xa9, 0x20, 0x01, 0xbb, 0x0b, 0x6d, 0x06, 0xd5, 0x68, 0x23, 0xb9, 0x19,
	0x68, 0xbb, 0x41, 0x26, 0x31, 0x59, 0xb4, 0xbc, 0x4c, 0x89, 0x53, 0x94, 0x5d, 0xf2, 0x31, 0xf9,
	0x34, 0x28, 0x6e, 0x5e, 0x1a, 0xd4, 0xbb, 0x3a, 0xe7, 0x77, 0xfe, 0x3e, 0x3e, 0xae, 0xe1, 0xf9,
	0xd3, 0x96, 0x85, 0xc9, 0x45, 0x94, 0xf9, 0x17, 0x1e, 0xe5, 0x34, 0x63, 0xdc, 0x78, 0x4c, 0x13,
	0x9e, 0xa0, 0x7e, 0x94, 0xf9, 0x3a, 0x6e, 0x55, 0x97, 0x72, 0xe6, 0x27, 0x69, 0xb1, 0x93, 0xf5,
	0x69, 0xab, 0x44, 0x34, 0x0c, 0xdc, 0x20, 0xc9, 0xb3, 0xef, 0x5d, 0x66, 0xfa, 0x67, 0x08, 0xf2,
	0x82, 0x72, 0xba, 0x61, 0x1c, 0xcd, 0x41, 0xad, 0xd4, 0xe0, 0x89, 0xf2, 0x20, 0x89, 0xb1, 0x74,
	0x2a, 0x9d, 0x4f, 0x66, 0x2f, 0x8c, 0x28, 0xf3, 0x8d, 0x8a, 0x32, 0xe6, 0x1d, 0x84, 0xfc, 0x63,
	0x41, 0x06, 0xc8, 0xd4, 0xfb, 0x99, 0xd2, 0xdc, 0xc3, 0x3d, 0xe1, 0x3e, 0xee, 0xb8, 0x4d, 0xef,
	0xaa, 0xd4, 0x48, 0x0d, 0x95, 0xfc, 0x03, 0x2b, 0x7e, 0x25, 0xa9, 0x87, 0x07, 0x07, 0xf8, 0xeb,
	0x9d, 0x46, 0x6a, 0x08, 0xbd, 0x05, 0xa5, 0x39, 0x0c, 0x1e, 0x0a, 0xc7, 0x49, 0xc7, 0xb1, 0xaa,
	0x55, 0xd2, 0x82, 0xe8, 0x0c, 0x06, 0xcc, 0xbd, 0x4f, 0xf0, 0x48, 0x18, 0x9e, 0x75, 0x0c, 0x96,
	0x7b, 0x9f, 0x10, 0x21, 0x97, 0xc3, 0x33, 0x16, 0xf3, 0x20, 0x62, 0x31, 0xc7, 0xf2, 0x81, 0xe1,
	0x9b, 0x5a, 0x25, 0x2d, 0xa8, 0xbf, 0x03, 0xb5, 0x5b, 0x0a, 0x3a, 0x83, 0xd1, 0x96, 0x86, 0x39,
	0xcb, 0xb0, 0x74, 0xda, 0x3f, 0x57, 0x67, 0x47, 0x62, 0x48, 0x05, 0x15, 0xa4, 0x12, 0x75, 0x13,
	0xe4, 0xaa, 0x0f, 0x84, 0x41, 0xde, 0xb2, 0xd4, 0x0b, 0x5c, 0x2e, 0x4a, 0x1f, 0x93, 0x7a, 0x89,
	0x5e, 0x82, 0x92, 0x05, 0x7e, 0x4c, 0x79, 0x9e, 0x32, 0x51, 0xa9, 0x42, 0xda, 0x0f, 0xfa, 0x6f,
	0x09, 0x94, 0x26, 0x14, 0xba, 0xec, 0xec, 0x3b, 0x99, 0x4d, 0x0f, 0x87, 0x37, 0xbe, 0x0a, 0xc8,
	0x8a, 0xf9, 0x5e, 0x98, 0x0f, 0x30, 0xd9, 0xfb, 0x8c, 0x34, 0xe8, 0x3f, 0xb0, 0x42, 0x84, 0x51,
	0x48, 0xf9, 0x13, 0x1d, 0xc3, 0x50, 0xa0, 0x22, 0x44, 0x8f, 0xec, 0x16, 0x97, 0xbd, 0xf7, 0x92,
	0xfe, 0x0a, 0xe4, 0xea, 0x9e, 0xd0, 0x49, 0x27, 0x81, 0xd2, 0x4c, 0xbf, 0x05, 0xa5, 0xb9, 0x18,
	0x34, 0x83, 0x71, 0xfd, 0x37, 0x14, 0x1b, 0xa8, 0x55, 0xcb, 0x0d, 0xd1, 0x34, 0xd5, 0x70, 0xfb,
	0x05, 0xf5, 0x3b, 0x05, 0xe9, 0x18, 0x06, 0xe5, 0x15, 0x96, 0x89, 0xf3, 0x34, 0xac, 0x13, 0xe7,
	0x69, 0xf8, 0xda, 0x85, 0x49, 0x75, 0x76, 0xa7, 0x78, 0x64, 0x08, 0x81, 0x3a, 0x37, 0x1d, 0xeb,
	0x93, 0x4d, 0x96, 0x77, 0xa6, 0xb3, 0xb4, 0xd7, 0xda, 0x7f, 0x68, 0x02, 0xb2, 0xb9, 0xb8, 0x22,
	0xe6, 0xcd, 0x42, 0x93, 0xca, 0xc5, 0xb5, 0x75, 0xfb, 0xcd, 0x26, 0x0b, 0xad, 0x8f, 0x8e, 0x40,
	0x59, 0x99, 0x5f, 0x96, 0xf3, 0xa5, 0x7d, 0xb3, 0xd1, 0x06, 0x68, 0x0c, 0x03, 0x6b, 0xfe, 0xd9,
	0xd6, 0x86, 0xa5, 0xb0, 0xb1, 0xd6, 0xce, 0x72, 0x65, 0xad, 0x1d, 0x6d, 0xf4, 0x51, 0xbd, 0xfb,
	0x5f, 0xbc, 0x33, 0x23, 0x10, 0x4f, 0xed, 0xc7, 0x48, 0x3c, 0xac, 0x37, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0xa1, 0xa5, 0x90, 0xb6, 0x03, 0x00, 0x00,
}
