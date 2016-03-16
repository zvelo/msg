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
	Status_SERVICE_UNAVAILABLE             Status_Code = 503
)

var Status_Code_name = map[int32]string{
	0:   "UNKNOWN",
	200: "OK",
	201: "CREATED",
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
	503: "SERVICE_UNAVAILABLE",
}
var Status_Code_value = map[string]int32{
	"UNKNOWN":                         0,
	"OK":                              200,
	"CREATED":                         201,
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
	"SERVICE_UNAVAILABLE":             503,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Status struct {
	Code    Status_Code `protobuf:"varint,1,opt,name=code,enum=msg.Status_Code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*Status)(nil), "msg.Status")
	proto.RegisterEnum("msg.Status_Code", Status_Code_name, Status_Code_value)
}

var fileDescriptor2 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x2c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xff, 0xc4, 0x55, 0xf3, 0xf7, 0xb6, 0xb4, 0xc6, 0x80, 0x54, 0x31, 0x00, 0xaa, 0x18,
	0x98, 0x32, 0xc0, 0x13, 0xb8, 0xc9, 0x0d, 0x58, 0x4d, 0xed, 0xe2, 0x38, 0x45, 0x62, 0xb1, 0x0a,
	0x54, 0x9d, 0xaa, 0x22, 0x52, 0xde, 0x03, 0x04, 0x03, 0x48, 0x0c, 0x7d, 0x1c, 0x78, 0x07, 0x1e,
	0x03, 0xc4, 0x8a, 0x13, 0x3a, 0xfa, 0x9c, 0x73, 0xbf, 0x7b, 0x8f, 0xa1, 0x55, 0x2c, 0x27, 0xcb,
	0xfb, 0x22, 0xbc, 0xbd, 0x5b, 0x2c, 0x17, 0x8c, 0xcc, 0x8b, 0x59, 0xef, 0xd3, 0x87, 0x7a, 0x56,
	0xa9, 0x6c, 0x0f, 0x6a, 0xd7, 0x8b, 0x9b, 0x69, 0xd7, 0x3b, 0xf0, 0x8e, 0xda, 0xc7, 0x34, 0x74,
	0x76, 0xf8, 0x67, 0x85, 0x91, 0xd3, 0x59, 0x07, 0x82, 0xf9, 0xb4, 0x28, 0x26, 0xb3, 0x69, 0xd7,
	0x77, 0x91, 0x46, 0x6f, 0xe5, 0x43, 0xad, 0x72, 0x9a, 0x10, 0xe4, 0x72, 0x20, 0xd5, 0x85, 0xa4,
	0xff, 0x58, 0x00, 0xbe, 0x1a, 0xd0, 0x77, 0x8f, 0xb5, 0x20, 0x88, 0x34, 0x72, 0x83, 0x31, 0xfd,
	0xf0, 0x18, 0x85, 0x66, 0x9f, 0xc7, 0x56, 0xe3, 0x79, 0x8e, 0x99, 0xa1, 0x0f, 0x84, 0x6d, 0x42,
	0x2b, 0x97, 0x3c, 0x37, 0x67, 0x4a, 0x8b, 0x4b, 0x17, 0x7a, 0x24, 0xac, 0x0d, 0x8d, 0x44, 0xe9,
	0xbe, 0x88, 0x63, 0x94, 0xf4, 0xa9, 0x7a, 0x4b, 0x65, 0x6c, 0xa2, 0x72, 0x19, 0xd3, 0x67, 0xc2,
	0xb6, 0xa1, 0xb3, 0x06, 0x58, 0x23, 0x86, 0xa8, 0x72, 0x43, 0x5f, 0x08, 0xdb, 0x80, 0xff, 0x91,
	0x92, 0x49, 0x2a, 0x22, 0x43, 0x5f, 0x09, 0xeb, 0xc2, 0xd6, 0x48, 0xa3, 0x53, 0x62, 0x61, 0x84,
	0x92, 0x36, 0xe1, 0x22, 0x75, 0xf8, 0x37, 0xc2, 0x0e, 0x61, 0x7f, 0x3d, 0x8e, 0xee, 0x12, 0x2e,
	0x4f, 0xd1, 0x96, 0xf8, 0x8c, 0x1b, 0x91, 0x25, 0x82, 0xf7, 0x53, 0xa4, 0x2b, 0xc2, 0x76, 0x61,
	0x47, 0x48, 0x83, 0x5a, 0xf2, 0xd4, 0x66, 0xa8, 0xc7, 0xa8, 0x2d, 0x6a, 0xad, 0x34, 0xfd, 0xaa,
	0x0e, 0x28, 0x27, 0xc4, 0x70, 0x94, 0xe2, 0x10, 0x65, 0xd9, 0xed, 0xbb, 0xda, 0x58, 0x06, 0x45,
	0x84, 0xd6, 0x35, 0x1a, 0xbb, 0x7d, 0x15, 0xeb, 0x87, 0x5c, 0xd5, 0xab, 0xaf, 0x3e, 0xf9, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xae, 0x1e, 0x44, 0x23, 0x7a, 0x01, 0x00, 0x00,
}
