// Code generated by protoc-gen-go.
// source: status.proto
// DO NOT EDIT!

package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Status_Code int32

const (
	Status_UNKNOWN                         Status_Code = 0
	Status_OK                              Status_Code = 200
	Status_CREATED                         Status_Code = 201
	Status_ACCEPTED                        Status_Code = 202
	Status_BAD_REQUEST                     Status_Code = 400
	Status_UNAUTHORIZED                    Status_Code = 401
	Status_FORBIDDEN                       Status_Code = 403
	Status_NOT_FOUND                       Status_Code = 404
	Status_REQUEST_TIMEOUT                 Status_Code = 408
	Status_CONFLICT                        Status_Code = 409
	Status_PRECONDITION_FAILED             Status_Code = 412
	Status_REQUESTED_RANGE_NOT_SATISFIABLE Status_Code = 416
	Status_INTERNAL_SERVER_ERROR           Status_Code = 500
	Status_NOT_IMPLEMENTED                 Status_Code = 501
	Status_BAD_GATEWAY                     Status_Code = 502
	Status_SERVICE_UNAVAILABLE             Status_Code = 503
	Status_GATEWAY_TIMEOUT                 Status_Code = 504
)

var Status_Code_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "CREATED",
	202: "ACCEPTED",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	408: "REQUEST_TIMEOUT",
	409: "CONFLICT",
	412: "PRECONDITION_FAILED",
	416: "REQUESTED_RANGE_NOT_SATISFIABLE",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	502: "BAD_GATEWAY",
	503: "SERVICE_UNAVAILABLE",
	504: "GATEWAY_TIMEOUT",
}
var Status_Code_value = map[string]int32{
	"UNKNOWN":                         0,
	"OK":                              200,
	"CREATED":                         201,
	"ACCEPTED":                        202,
	"BAD_REQUEST":                     400,
	"UNAUTHORIZED":                    401,
	"FORBIDDEN":                       403,
	"NOT_FOUND":                       404,
	"REQUEST_TIMEOUT":                 408,
	"CONFLICT":                        409,
	"PRECONDITION_FAILED":             412,
	"REQUESTED_RANGE_NOT_SATISFIABLE": 416,
	"INTERNAL_SERVER_ERROR":           500,
	"NOT_IMPLEMENTED":                 501,
	"BAD_GATEWAY":                     502,
	"SERVICE_UNAVAILABLE":             503,
	"GATEWAY_TIMEOUT":                 504,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Status struct {
	Code        Status_Code `protobuf:"varint,1,opt,name=code,enum=msg.Status_Code" json:"code,omitempty"`
	Message     string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	FetchStatus *Status     `protobuf:"bytes,3,opt,name=fetch_status,json=fetchStatus" json:"fetch_status,omitempty"`
	Location    string      `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Status) GetFetchStatus() *Status {
	if m != nil {
		return m.FetchStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "msg.Status")
	proto.RegisterEnum("msg.Status_Code", Status_Code_name, Status_Code_value)
}

var fileDescriptor2 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x51, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0x26, 0x75, 0x75, 0x73, 0xaf, 0x13, 0x2e, 0xc6, 0x80, 0x14, 0xb1, 0x80, 0x2a, 0x06, 0xa6,
	0x0c, 0xe5, 0x09, 0xdc, 0xe4, 0xa4, 0x58, 0x4d, 0xed, 0xe2, 0x38, 0xad, 0x60, 0xb1, 0x4a, 0x09,
	0x05, 0x89, 0x12, 0x44, 0xc2, 0x7b, 0x80, 0x60, 0x00, 0x89, 0x81, 0xc7, 0x01, 0x5e, 0x80, 0xb7,
	0x60, 0xe2, 0x6f, 0xc4, 0x71, 0xd3, 0xea, 0x8e, 0xc7, 0xdf, 0x77, 0xbe, 0x9f, 0x63, 0x1c, 0x36,
	0xed, 0xba, 0x7d, 0xd3, 0xc4, 0xaf, 0x5e, 0xd7, 0x6d, 0x4d, 0xd1, 0xae, 0xd9, 0x8e, 0x7e, 0x22,
	0x7c, 0x52, 0xb8, 0x57, 0x7a, 0x07, 0x0f, 0x37, 0xf5, 0x93, 0x2a, 0xf2, 0x6e, 0x7b, 0x77, 0xcf,
	0xc7, 0x24, 0xb6, 0x70, 0xbc, 0x87, 0xe2, 0xc4, 0xbe, 0x2b, 0x87, 0xd2, 0x08, 0xfb, 0xbb, 0xaa,
	0x69, 0xd6, 0xdb, 0x2a, 0x1a, 0x58, 0xe2, 0x99, 0x3a, 0x8c, 0x34, 0xc6, 0xe1, 0xd3, 0xaa, 0xdd,
	0x3c, 0x33, 0x7b, 0x97, 0x08, 0x59, 0x38, 0x18, 0x07, 0x17, 0x74, 0x54, 0xe0, 0x08, 0xbd, 0xdf,
	0x4d, 0x7c, 0xfa, 0xa2, 0xde, 0xac, 0xdb, 0xe7, 0xf5, 0xcb, 0x68, 0xe8, 0xa4, 0x8e, 0xf3, 0xe8,
	0xc7, 0x00, 0x0f, 0x3b, 0x53, 0x1a, 0x60, 0xbf, 0x14, 0x33, 0x21, 0x57, 0x82, 0x5c, 0xa2, 0x3e,
	0x1e, 0xc8, 0x19, 0xf9, 0xea, 0xd1, 0x10, 0xfb, 0x89, 0x02, 0xa6, 0x21, 0x25, 0xdf, 0x3c, 0x7a,
	0x19, 0x9f, 0xb2, 0x24, 0x81, 0x45, 0x37, 0x7e, 0xf7, 0x28, 0xc1, 0xc1, 0x84, 0xa5, 0x46, 0xc1,
	0x83, 0x12, 0x0a, 0x4d, 0xde, 0x22, 0x7a, 0x15, 0x87, 0xa5, 0x60, 0xa5, 0xbe, 0x2f, 0x15, 0x7f,
	0x64, 0x49, 0xef, 0x10, 0x3d, 0xc7, 0x67, 0x99, 0x54, 0x13, 0x9e, 0xa6, 0x20, 0xc8, 0x7b, 0x37,
	0x0b, 0xa9, 0x4d, 0x26, 0x4b, 0x91, 0x92, 0x0f, 0x88, 0x5e, 0xc7, 0x57, 0x7a, 0x01, 0xa3, 0xf9,
	0x1c, 0x64, 0xa9, 0xc9, 0x47, 0xd4, 0x39, 0x25, 0x52, 0x64, 0x39, 0x4f, 0x34, 0xf9, 0x84, 0xec,
	0x2d, 0xae, 0x2d, 0x14, 0xd8, 0x97, 0x94, 0x6b, 0x2e, 0x85, 0xc9, 0x18, 0xcf, 0xad, 0xfc, 0x67,
	0x64, 0x6f, 0x79, 0xab, 0x5f, 0x07, 0x9b, 0x84, 0x89, 0x29, 0x98, 0x4e, 0xbe, 0x60, 0x9a, 0x17,
	0x19, 0x67, 0x93, 0x1c, 0xc8, 0x17, 0x64, 0x2f, 0x70, 0x83, 0x0b, 0x0d, 0x4a, 0xb0, 0xdc, 0x14,
	0xa0, 0x96, 0xa0, 0x0c, 0x28, 0x25, 0x15, 0xf9, 0xe5, 0x02, 0x74, 0x1b, 0x7c, 0xbe, 0xc8, 0x61,
	0x0e, 0xa2, 0xeb, 0xf6, 0x1b, 0x1d, 0xba, 0x4d, 0x6d, 0xf5, 0x15, 0x7b, 0x48, 0xfe, 0xb8, 0x0c,
	0xdd, 0x2a, 0x4f, 0xc0, 0xd8, 0x8e, 0x4b, 0x9b, 0xc0, 0xa9, 0xff, 0x75, 0x0a, 0x3d, 0xef, 0x58,
	0xe1, 0x1f, 0x7a, 0x7c, 0xe2, 0x3e, 0xff, 0xde, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x32,
	0xc7, 0x42, 0x0c, 0x02, 0x00, 0x00,
}
