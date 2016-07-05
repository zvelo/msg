// Code generated by protoc-gen-go.
// source: zvelo.io/msg/status.proto
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
	Status_UNKNOWN Status_Code = 0
	// 100s
	Status_CONTINUE            Status_Code = 100
	Status_SWITCHING_PROTOCOLS Status_Code = 101
	// 200s
	Status_OK                     Status_Code = 200
	Status_CREATED                Status_Code = 201
	Status_ACCEPTED               Status_Code = 202
	Status_NON_AUTHORITATIVE_INFO Status_Code = 203
	Status_NO_CONTENT             Status_Code = 204
	Status_RESET_CONTENT          Status_Code = 205
	Status_PARTIAL_CONTENT        Status_Code = 206
	// 300s
	Status_MULTIPLE_CHOICES   Status_Code = 300
	Status_MOVED_PERMANENTLY  Status_Code = 301
	Status_FOUND              Status_Code = 302
	Status_SEE_OTHER          Status_Code = 303
	Status_NOT_MODIFIED       Status_Code = 304
	Status_USE_PROXY          Status_Code = 305
	Status_TEMPORARY_REDIRECT Status_Code = 307
	// 400s
	Status_BAD_REQUEST                     Status_Code = 400
	Status_UNAUTHORIZED                    Status_Code = 401
	Status_PAYMENT_REQUIRED                Status_Code = 402
	Status_FORBIDDEN                       Status_Code = 403
	Status_NOT_FOUND                       Status_Code = 404
	Status_METHOD_NOT_ALLOWED              Status_Code = 405
	Status_NOT_ACCEPTABLE                  Status_Code = 406
	Status_PROXY_AUTH_REQUIRED             Status_Code = 407
	Status_REQUEST_TIMEOUT                 Status_Code = 408
	Status_CONFLICT                        Status_Code = 409
	Status_GONE                            Status_Code = 410
	Status_LENGTH_REQUIRED                 Status_Code = 411
	Status_PRECONDITION_FAILED             Status_Code = 412
	Status_REQUEST_ENTITY_TOO_LARGE        Status_Code = 413
	Status_REQUEST_URI_TOO_LONG            Status_Code = 414
	Status_UNSUPPORTED_MEDIA_TYPE          Status_Code = 415
	Status_REQUESTED_RANGE_NOT_SATISFIABLE Status_Code = 416
	Status_EXPECTATION_FAILED              Status_Code = 417
	Status_TEAPOT                          Status_Code = 418
	Status_PRECONDITION_REQUIRED           Status_Code = 428
	Status_TOO_MANY_REQUESTS               Status_Code = 429
	Status_REQUEST_HEADER_FIELDS_TOO_LARGE Status_Code = 431
	Status_UNAVAILABLE_FOR_LEGAL_REASONS   Status_Code = 451
	// 500s
	Status_INTERNAL_SERVER_ERROR           Status_Code = 500
	Status_NOT_IMPLEMENTED                 Status_Code = 501
	Status_BAD_GATEWAY                     Status_Code = 502
	Status_SERVICE_UNAVAILABLE             Status_Code = 503
	Status_GATEWAY_TIMEOUT                 Status_Code = 504
	Status_HTTP_VERSION_NOT_SUPPORTED      Status_Code = 505
	Status_NETWORK_AUTHENTICATION_REQUIRED Status_Code = 511
)

var Status_Code_name = map[int32]string{
	0:   "UNKNOWN",
	100: "CONTINUE",
	101: "SWITCHING_PROTOCOLS",
	200: "OK",
	201: "CREATED",
	202: "ACCEPTED",
	203: "NON_AUTHORITATIVE_INFO",
	204: "NO_CONTENT",
	205: "RESET_CONTENT",
	206: "PARTIAL_CONTENT",
	300: "MULTIPLE_CHOICES",
	301: "MOVED_PERMANENTLY",
	302: "FOUND",
	303: "SEE_OTHER",
	304: "NOT_MODIFIED",
	305: "USE_PROXY",
	307: "TEMPORARY_REDIRECT",
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	402: "PAYMENT_REQUIRED",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	405: "METHOD_NOT_ALLOWED",
	406: "NOT_ACCEPTABLE",
	407: "PROXY_AUTH_REQUIRED",
	408: "REQUEST_TIMEOUT",
	409: "CONFLICT",
	410: "GONE",
	411: "LENGTH_REQUIRED",
	412: "PRECONDITION_FAILED",
	413: "REQUEST_ENTITY_TOO_LARGE",
	414: "REQUEST_URI_TOO_LONG",
	415: "UNSUPPORTED_MEDIA_TYPE",
	416: "REQUESTED_RANGE_NOT_SATISFIABLE",
	417: "EXPECTATION_FAILED",
	418: "TEAPOT",
	428: "PRECONDITION_REQUIRED",
	429: "TOO_MANY_REQUESTS",
	431: "REQUEST_HEADER_FIELDS_TOO_LARGE",
	451: "UNAVAILABLE_FOR_LEGAL_REASONS",
	500: "INTERNAL_SERVER_ERROR",
	501: "NOT_IMPLEMENTED",
	502: "BAD_GATEWAY",
	503: "SERVICE_UNAVAILABLE",
	504: "GATEWAY_TIMEOUT",
	505: "HTTP_VERSION_NOT_SUPPORTED",
	511: "NETWORK_AUTHENTICATION_REQUIRED",
}
var Status_Code_value = map[string]int32{
	"UNKNOWN":             0,
	"CONTINUE":            100,
	"SWITCHING_PROTOCOLS": 101,
	"OK":                              200,
	"CREATED":                         201,
	"ACCEPTED":                        202,
	"NON_AUTHORITATIVE_INFO":          203,
	"NO_CONTENT":                      204,
	"RESET_CONTENT":                   205,
	"PARTIAL_CONTENT":                 206,
	"MULTIPLE_CHOICES":                300,
	"MOVED_PERMANENTLY":               301,
	"FOUND":                           302,
	"SEE_OTHER":                       303,
	"NOT_MODIFIED":                    304,
	"USE_PROXY":                       305,
	"TEMPORARY_REDIRECT":              307,
	"BAD_REQUEST":                     400,
	"UNAUTHORIZED":                    401,
	"PAYMENT_REQUIRED":                402,
	"FORBIDDEN":                       403,
	"NOT_FOUND":                       404,
	"METHOD_NOT_ALLOWED":              405,
	"NOT_ACCEPTABLE":                  406,
	"PROXY_AUTH_REQUIRED":             407,
	"REQUEST_TIMEOUT":                 408,
	"CONFLICT":                        409,
	"GONE":                            410,
	"LENGTH_REQUIRED":                 411,
	"PRECONDITION_FAILED":             412,
	"REQUEST_ENTITY_TOO_LARGE":        413,
	"REQUEST_URI_TOO_LONG":            414,
	"UNSUPPORTED_MEDIA_TYPE":          415,
	"REQUESTED_RANGE_NOT_SATISFIABLE": 416,
	"EXPECTATION_FAILED":              417,
	"TEAPOT":                          418,
	"PRECONDITION_REQUIRED":           428,
	"TOO_MANY_REQUESTS":               429,
	"REQUEST_HEADER_FIELDS_TOO_LARGE": 431,
	"UNAVAILABLE_FOR_LEGAL_REASONS":   451,
	"INTERNAL_SERVER_ERROR":           500,
	"NOT_IMPLEMENTED":                 501,
	"BAD_GATEWAY":                     502,
	"SERVICE_UNAVAILABLE":             503,
	"GATEWAY_TIMEOUT":                 504,
	"HTTP_VERSION_NOT_SUPPORTED":      505,
	"NETWORK_AUTHENTICATION_REQUIRED": 511,
}

func (x Status_Code) String() string {
	return proto.EnumName(Status_Code_name, int32(x))
}
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type Status struct {
	Code        Status_Code `protobuf:"varint,1,opt,name=code,enum=msg.Status_Code" json:"code,omitempty"`
	Message     string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	FetchStatus *Status     `protobuf:"bytes,3,opt,name=fetch_status,json=fetchStatus" json:"fetch_status,omitempty"`
	Location    string      `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

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

func init() { proto.RegisterFile("zvelo.io/msg/status.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x54, 0xcb, 0x92, 0x5b, 0x35,
	0x10, 0xe5, 0x5a, 0xc3, 0x3c, 0xe4, 0xc9, 0xa4, 0xa3, 0x49, 0x82, 0x13, 0x2a, 0x15, 0x6a, 0x8a,
	0x05, 0x2b, 0xa7, 0x2a, 0x7c, 0x81, 0xe6, 0xde, 0xb6, 0xad, 0x9a, 0x7b, 0xa5, 0x8b, 0xa4, 0x3b,
	0x8e, 0xb3, 0x51, 0x0d, 0x13, 0x13, 0x52, 0x95, 0x60, 0x0a, 0x1b, 0x16, 0x7c, 0x05, 0xef, 0x37,
	0x04, 0x58, 0x27, 0x04, 0x8a, 0x2a, 0x7e, 0x80, 0x0d, 0xef, 0xaf, 0xe1, 0xb9, 0x82, 0x96, 0xc6,
	0x36, 0x9e, 0xa5, 0xba, 0x5b, 0x7d, 0x4e, 0x9f, 0x3e, 0x12, 0xbf, 0xf4, 0xfa, 0x6b, 0xe3, 0xbb,
	0x93, 0xee, 0x9d, 0xc9, 0xb5, 0x7b, 0xd3, 0xdb, 0xd7, 0xa6, 0xb3, 0xa3, 0xd9, 0xab, 0xd3, 0xee,
	0xcb, 0xaf, 0x4c, 0x66, 0x13, 0xc1, 0x28, 0xb2, 0xf7, 0xdd, 0x16, 0x5f, 0x77, 0x29, 0x2a, 0x9e,
	0xe6, 0x6b, 0xc7, 0x93, 0x5b, 0xe3, 0x4e, 0xf6, 0x54, 0xf6, 0xcc, 0xce, 0x75, 0xe8, 0x52, 0xba,
	0x7b, 0x92, 0xea, 0xe6, 0x14, 0xb7, 0x29, 0x2b, 0x3a, 0x7c, 0xe3, 0xde, 0x78, 0x3a, 0x3d, 0xba,
	0x3d, 0xee, 0xb4, 0xa8, 0x70, 0xcb, 0x2e, 0x8e, 0xa2, 0xcb, 0xb7, 0x5f, 0x18, 0xcf, 0x8e, 0x5f,
	0x0c, 0x27, 0x28, 0x1d, 0x46, 0xe9, 0xf6, 0xf5, 0xf6, 0x4a, 0x1f, 0xdb, 0x4e, 0x05, 0x73, 0xbc,
	0xcb, 0x7c, 0xf3, 0xee, 0xe4, 0xf8, 0x68, 0x76, 0x67, 0xf2, 0x52, 0x67, 0x2d, 0xb5, 0x5a, 0x9e,
	0xf7, 0xee, 0x6f, 0xf2, 0xb5, 0x08, 0x2a, 0xda, 0x7c, 0xa3, 0xd1, 0x07, 0xda, 0x0c, 0x35, 0x3c,
	0x26, 0xb6, 0xf9, 0x66, 0x6e, 0xb4, 0x57, 0xba, 0x41, 0xb8, 0x25, 0x9e, 0xe0, 0xbb, 0x6e, 0xa8,
	0x7c, 0x3e, 0x50, 0xba, 0x1f, 0x6a, 0x6b, 0xbc, 0xc9, 0x4d, 0xe9, 0x60, 0x2c, 0x36, 0x78, 0xcb,
	0x1c, 0xc0, 0x0f, 0x19, 0xd5, 0x6f, 0xe4, 0x16, 0xa5, 0xc7, 0x02, 0x7e, 0xcc, 0xc4, 0x19, 0xbe,
	0x29, 0xf3, 0x1c, 0xeb, 0x78, 0xfc, 0x29, 0x13, 0x4f, 0xf2, 0x8b, 0xda, 0xe8, 0x20, 0x1b, 0x3f,
	0x30, 0x56, 0x79, 0xe9, 0xd5, 0x21, 0x06, 0xa5, 0x7b, 0x06, 0x7e, 0xce, 0xc4, 0x59, 0xce, 0xb5,
	0x09, 0x11, 0x0c, 0xb5, 0x87, 0x5f, 0x32, 0x21, 0xf8, 0x19, 0x8b, 0x0e, 0xfd, 0x32, 0xf6, 0x6b,
	0x26, 0xce, 0xf3, 0xb3, 0xb5, 0xb4, 0x5e, 0xc9, 0x72, 0x19, 0xfd, 0x2d, 0x13, 0x17, 0x38, 0x54,
	0x4d, 0xe9, 0x55, 0x5d, 0x62, 0xc8, 0x07, 0x46, 0xe5, 0xe8, 0xe0, 0x41, 0x4b, 0x5c, 0xe4, 0xe7,
	0x2a, 0x73, 0x88, 0x45, 0xa8, 0xd1, 0x56, 0x52, 0x53, 0x71, 0x39, 0x82, 0x87, 0x2d, 0xc1, 0xf9,
	0xe3, 0x3d, 0xd3, 0xe8, 0x02, 0xbe, 0x6a, 0x89, 0x1d, 0xbe, 0xe5, 0x10, 0x83, 0xf1, 0x03, 0xb4,
	0xf0, 0xa8, 0x25, 0xce, 0xf1, 0x6d, 0x6d, 0x7c, 0xa8, 0x4c, 0xa1, 0x7a, 0x8a, 0x58, 0x7f, 0x9d,
	0x4a, 0x1a, 0x87, 0x71, 0xdc, 0x1b, 0x23, 0xf8, 0xa6, 0x45, 0x22, 0x08, 0x8f, 0x55, 0x6d, 0xac,
	0xb4, 0xa3, 0x60, 0xb1, 0x50, 0x16, 0x73, 0x0f, 0xdf, 0xb6, 0x04, 0xf0, 0xf6, 0xbe, 0x2c, 0x28,
	0xf4, 0x5c, 0x83, 0xce, 0xc3, 0x1b, 0x2c, 0x76, 0x6b, 0xf4, 0x7c, 0xdc, 0x9b, 0xd4, 0xed, 0x4d,
	0x16, 0xb9, 0xd6, 0x72, 0x54, 0x11, 0x99, 0x54, 0x48, 0x97, 0x0b, 0x78, 0x8b, 0x45, 0x90, 0x9e,
	0xb1, 0xfb, 0xaa, 0x28, 0x50, 0xc3, 0xdb, 0xe9, 0x1c, 0x79, 0x9c, 0xf0, 0x7c, 0x87, 0x45, 0xd0,
	0x0a, 0xa9, 0x4f, 0x11, 0x62, 0x58, 0x96, 0xa5, 0x19, 0xd2, 0xc5, 0x77, 0x99, 0xd8, 0xe5, 0x3b,
	0x29, 0x92, 0x64, 0x96, 0xfb, 0x25, 0xc2, 0x7b, 0x8c, 0x1c, 0xb3, 0x9b, 0xe8, 0x26, 0xa9, 0xff,
	0xc7, 0x79, 0x9f, 0x45, 0x01, 0xe7, 0xfc, 0x82, 0x57, 0x15, 0x9a, 0xc6, 0xc3, 0x07, 0x2c, 0xee,
	0x89, 0xe4, 0xec, 0x95, 0x8a, 0x06, 0xf9, 0x90, 0x89, 0x2d, 0xbe, 0xd6, 0x37, 0x1a, 0xe1, 0xa3,
	0x54, 0x5f, 0xa2, 0xee, 0xaf, 0x76, 0xf9, 0x78, 0xde, 0x1f, 0xe9, 0x4a, 0xa1, 0xbc, 0xa2, 0x8d,
	0xf6, 0xa4, 0x2a, 0x29, 0xf3, 0x09, 0x13, 0x57, 0x78, 0x67, 0xd1, 0x9f, 0x46, 0x54, 0x7e, 0x14,
	0xbc, 0x31, 0xa1, 0x94, 0xb6, 0x8f, 0xf0, 0x29, 0x13, 0x97, 0xf8, 0xf9, 0x45, 0xba, 0xb1, 0xea,
	0x24, 0x67, 0x74, 0x1f, 0x3e, 0x63, 0xd1, 0x1c, 0x8d, 0x76, 0x4d, 0x4d, 0xc2, 0x92, 0x5d, 0x42,
	0x45, 0xc2, 0xca, 0xe0, 0x47, 0x35, 0xc2, 0x7d, 0x46, 0x0f, 0xe5, 0xea, 0xfc, 0x1e, 0xa5, 0xac,
	0xd4, 0x7d, 0x4c, 0x3a, 0x38, 0xb2, 0x90, 0xeb, 0xa9, 0x34, 0xf6, 0xe7, 0x49, 0x24, 0xbc, 0x51,
	0xd3, 0x36, 0xe4, 0x2a, 0xab, 0x2f, 0x18, 0x59, 0x7a, 0xdd, 0xa3, 0xac, 0x8d, 0x87, 0x2f, 0x19,
	0x3d, 0x82, 0x0b, 0xa7, 0xc8, 0x2f, 0x07, 0x7b, 0xc0, 0xa2, 0x65, 0x22, 0x27, 0x72, 0xcb, 0x68,
	0xb1, 0x47, 0x07, 0x0f, 0x57, 0xf1, 0xc3, 0x00, 0x65, 0x81, 0x36, 0x90, 0x3b, 0xca, 0xc2, 0xad,
	0x4c, 0xf7, 0x88, 0x89, 0x3d, 0x7e, 0x85, 0xd6, 0x7d, 0x48, 0xb0, 0x91, 0x11, 0x2d, 0xcf, 0x86,
	0x12, 0xfb, 0xe4, 0x55, 0x7a, 0x11, 0xce, 0x68, 0x07, 0xdf, 0x27, 0x74, 0x45, 0xc6, 0xb5, 0x9a,
	0xc2, 0x0e, 0xed, 0x21, 0xb5, 0x42, 0x6b, 0x8d, 0x85, 0xdf, 0x93, 0xd8, 0x71, 0x2a, 0x55, 0x91,
	0x91, 0xa3, 0x43, 0x88, 0xd3, 0x1f, 0x6c, 0x61, 0xab, 0x3e, 0x3d, 0xaa, 0xa1, 0x1c, 0xc1, 0x9f,
	0x49, 0xfe, 0x78, 0x95, 0x7c, 0x1e, 0x56, 0xf0, 0xe0, 0xaf, 0xd4, 0x61, 0x5e, 0xb7, 0x5c, 0xef,
	0xdf, 0x4c, 0x5c, 0xe5, 0x97, 0x07, 0xde, 0xd7, 0x81, 0xc0, 0x5c, 0x9c, 0x38, 0x49, 0xb7, 0x50,
	0x1a, 0xfe, 0x49, 0xe3, 0x69, 0xf4, 0x43, 0x63, 0x0f, 0x92, 0x63, 0xe2, 0xe6, 0x72, 0x79, 0x5a,
	0x9c, 0x7f, 0xd9, 0xfe, 0xce, 0xcd, 0xed, 0xd5, 0xaf, 0xed, 0xf9, 0xf5, 0xf4, 0xa9, 0x3d, 0xfb,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xe7, 0x89, 0xdf, 0xf1, 0x04, 0x00, 0x00,
}
