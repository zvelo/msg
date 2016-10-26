// Code generated by protoc-gen-go.
// source: zvelo/msg/status.proto
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
func (Status_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

type Status struct {
	Code        Status_Code `protobuf:"varint,1,opt,name=code,enum=msg.Status_Code" json:"code,omitempty"`
	Message     string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	FetchStatus *Status     `protobuf:"bytes,3,opt,name=fetch_status,json=fetchStatus" json:"fetch_status,omitempty"`
	Location    string      `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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

func init() { proto.RegisterFile("zvelo/msg/status.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x54, 0xcb, 0x92, 0x5b, 0x35,
	0x10, 0xe5, 0x5a, 0xc3, 0x3c, 0xe4, 0xc9, 0xa4, 0xa3, 0x49, 0x06, 0x13, 0x2a, 0x15, 0x6a, 0x8a,
	0x05, 0x2b, 0xa7, 0x2a, 0x7c, 0x81, 0xe6, 0xde, 0xb6, 0xad, 0x9a, 0x7b, 0xa5, 0x8b, 0xa4, 0x3b,
	0x8e, 0xb3, 0x51, 0x0d, 0x13, 0x13, 0x52, 0x95, 0x60, 0x0a, 0x1b, 0x16, 0x7c, 0x05, 0xef, 0x37,
	0x04, 0x58, 0x27, 0x04, 0x8a, 0x2a, 0x7e, 0x80, 0x0d, 0xef, 0xaf, 0xe1, 0xb9, 0x82, 0x96, 0xc6,
	0x76, 0x9c, 0xa5, 0xba, 0x5b, 0x7d, 0x4e, 0x9f, 0x3e, 0x12, 0xdf, 0x7b, 0xe3, 0xf5, 0xf1, 0xed,
	0xc9, 0x95, 0x3b, 0xd3, 0x9b, 0x57, 0xa6, 0xb3, 0xe3, 0xd9, 0x6b, 0xd3, 0xee, 0x2b, 0xaf, 0x4e,
	0x66, 0x13, 0xc1, 0x28, 0xb2, 0xff, 0xfd, 0x16, 0x5f, 0x77, 0x29, 0x2a, 0x9e, 0xe1, 0x6b, 0x27,
	0x93, 0x1b, 0xe3, 0x4e, 0xf6, 0x74, 0xf6, 0xec, 0xce, 0x55, 0xe8, 0x52, 0xba, 0x7b, 0x9a, 0xea,
	0xe6, 0x14, 0xb7, 0x29, 0x2b, 0x3a, 0x7c, 0xe3, 0xce, 0x78, 0x3a, 0x3d, 0xbe, 0x39, 0xee, 0xb4,
	0xa8, 0x70, 0xcb, 0x2e, 0x8e, 0xa2, 0xcb, 0xb7, 0x5f, 0x1c, 0xcf, 0x4e, 0x5e, 0x0a, 0xa7, 0x28,
	0x1d, 0x46, 0xe9, 0xf6, 0xd5, 0xf6, 0x4a, 0x1f, 0xdb, 0x4e, 0x05, 0x73, 0xbc, 0x8b, 0x7c, 0xf3,
	0xf6, 0xe4, 0xe4, 0x78, 0x76, 0x6b, 0xf2, 0x72, 0x67, 0x2d, 0xb5, 0x5a, 0x9e, 0xf7, 0xef, 0x6e,
	0xf2, 0xb5, 0x08, 0x2a, 0xda, 0x7c, 0xa3, 0xd1, 0x87, 0xda, 0x0c, 0x35, 0x3c, 0x26, 0xb6, 0xf9,
	0x66, 0x6e, 0xb4, 0x57, 0xba, 0x41, 0xb8, 0x21, 0x9e, 0xe0, 0xbb, 0x6e, 0xa8, 0x7c, 0x3e, 0x50,
	0xba, 0x1f, 0x6a, 0x6b, 0xbc, 0xc9, 0x4d, 0xe9, 0x60, 0x2c, 0x36, 0x78, 0xcb, 0x1c, 0xc2, 0x8f,
	0x19, 0xd5, 0x6f, 0xe4, 0x16, 0xa5, 0xc7, 0x02, 0x7e, 0xca, 0xc4, 0x19, 0xbe, 0x29, 0xf3, 0x1c,
	0xeb, 0x78, 0xfc, 0x39, 0x13, 0x4f, 0xf1, 0x3d, 0x6d, 0x74, 0x90, 0x8d, 0x1f, 0x18, 0xab, 0xbc,
	0xf4, 0xea, 0x08, 0x83, 0xd2, 0x3d, 0x03, 0xbf, 0x64, 0xe2, 0x2c, 0xe7, 0xda, 0x84, 0x08, 0x86,
	0xda, 0xc3, 0xaf, 0x99, 0x10, 0xfc, 0x8c, 0x45, 0x87, 0x7e, 0x19, 0xfb, 0x2d, 0x13, 0xe7, 0xf9,
	0xd9, 0x5a, 0x5a, 0xaf, 0x64, 0xb9, 0x8c, 0xfe, 0x9e, 0x89, 0x0b, 0x1c, 0xaa, 0xa6, 0xf4, 0xaa,
	0x2e, 0x31, 0xe4, 0x03, 0xa3, 0x72, 0x74, 0x70, 0xaf, 0x25, 0xf6, 0xf8, 0xb9, 0xca, 0x1c, 0x61,
	0x11, 0x6a, 0xb4, 0x95, 0xd4, 0x54, 0x5c, 0x8e, 0xe0, 0x7e, 0x4b, 0x70, 0xfe, 0x78, 0xcf, 0x34,
	0xba, 0x80, 0xaf, 0x5b, 0x62, 0x87, 0x6f, 0x39, 0xc4, 0x60, 0xfc, 0x00, 0x2d, 0x3c, 0x68, 0x89,
	0x73, 0x7c, 0x5b, 0x1b, 0x1f, 0x2a, 0x53, 0xa8, 0x9e, 0x22, 0xd6, 0xdf, 0xa4, 0x92, 0xc6, 0x61,
	0x1c, 0xf7, 0xda, 0x08, 0xbe, 0x6d, 0x91, 0x08, 0xc2, 0x63, 0x55, 0x1b, 0x2b, 0xed, 0x28, 0x58,
	0x2c, 0x94, 0xc5, 0xdc, 0xc3, 0x77, 0x2d, 0x01, 0xbc, 0x7d, 0x20, 0x0b, 0x0a, 0x3d, 0xdf, 0xa0,
	0xf3, 0xf0, 0x26, 0x8b, 0xdd, 0x1a, 0x3d, 0x1f, 0xf7, 0x3a, 0x75, 0x7b, 0x8b, 0x45, 0xae, 0xb5,
	0x1c, 0x55, 0x44, 0x26, 0x15, 0xd2, 0xe5, 0x02, 0xde, 0x66, 0x11, 0xa4, 0x67, 0xec, 0x81, 0x2a,
	0x0a, 0xd4, 0xf0, 0x4e, 0x3a, 0x47, 0x1e, 0xa7, 0x3c, 0xdf, 0x65, 0x11, 0xb4, 0x42, 0xea, 0x53,
	0x84, 0x18, 0x96, 0x65, 0x69, 0x86, 0x74, 0xf1, 0x3d, 0x26, 0x76, 0xf9, 0x4e, 0x8a, 0x24, 0x99,
	0xe5, 0x41, 0x89, 0xf0, 0x3e, 0x23, 0xc7, 0xec, 0x26, 0xba, 0x49, 0xea, 0x87, 0x38, 0x1f, 0xb0,
	0x28, 0xe0, 0x9c, 0x5f, 0xf0, 0xaa, 0x42, 0xd3, 0x78, 0xf8, 0x90, 0xc5, 0x3d, 0x91, 0x9c, 0xbd,
	0x52, 0xd1, 0x20, 0x1f, 0x31, 0xb1, 0xc5, 0xd7, 0xfa, 0x46, 0x23, 0x7c, 0x9c, 0xea, 0x4b, 0xd4,
	0xfd, 0xd5, 0x2e, 0x9f, 0xcc, 0xfb, 0x23, 0x5d, 0x29, 0x94, 0x57, 0xb4, 0xd1, 0x9e, 0x54, 0x25,
	0x65, 0x3e, 0x65, 0xe2, 0x12, 0xef, 0x2c, 0xfa, 0xd3, 0x88, 0xca, 0x8f, 0x82, 0x37, 0x26, 0x94,
	0xd2, 0xf6, 0x11, 0x3e, 0x63, 0xe2, 0x49, 0x7e, 0x7e, 0x91, 0x6e, 0xac, 0x3a, 0xcd, 0x19, 0xdd,
	0x87, 0xcf, 0x59, 0x34, 0x47, 0xa3, 0x5d, 0x53, 0x93, 0xb0, 0x64, 0x97, 0x50, 0x91, 0xb0, 0x32,
	0xf8, 0x51, 0x8d, 0x70, 0x97, 0xd1, 0x43, 0xb9, 0x3c, 0xbf, 0x47, 0x29, 0x2b, 0x75, 0x1f, 0x93,
	0x0e, 0x8e, 0x2c, 0xe4, 0x7a, 0x2a, 0x8d, 0xfd, 0x45, 0x12, 0x09, 0xaf, 0xd5, 0xb4, 0x0d, 0xb9,
	0xca, 0xea, 0x4b, 0x46, 0x96, 0x5e, 0xf7, 0x28, 0x6b, 0xe3, 0xe1, 0x2b, 0x46, 0x8f, 0xe0, 0xc2,
	0x23, 0xe4, 0x97, 0x83, 0xdd, 0x63, 0xd1, 0x32, 0x91, 0x13, 0xb9, 0x65, 0xb4, 0xd8, 0xa3, 0x83,
	0xfb, 0xab, 0xf8, 0x61, 0x80, 0xb2, 0x40, 0x1b, 0xc8, 0x1d, 0x65, 0xe1, 0x56, 0xa6, 0x7b, 0xc0,
	0xc4, 0x3e, 0xbf, 0x44, 0xeb, 0x3e, 0x22, 0xd8, 0xc8, 0x88, 0x96, 0x67, 0x43, 0x89, 0x7d, 0xf2,
	0x2a, 0xbd, 0x08, 0x67, 0xb4, 0x83, 0x1f, 0x12, 0xba, 0x22, 0xe3, 0x5a, 0x4d, 0x61, 0x87, 0xf6,
	0x88, 0x5a, 0xa1, 0xb5, 0xc6, 0xc2, 0x1f, 0x49, 0xec, 0x38, 0x95, 0xaa, 0xc8, 0xc8, 0xd1, 0x21,
	0xc4, 0xe9, 0x4f, 0xb6, 0xb0, 0x55, 0x9f, 0x1e, 0xd5, 0x50, 0x8e, 0xe0, 0xaf, 0x24, 0x7f, 0xbc,
	0x4a, 0x3e, 0x0f, 0x2b, 0x78, 0xf0, 0x77, 0xea, 0x30, 0xaf, 0x5b, 0xae, 0xf7, 0x1f, 0x26, 0x2e,
	0xf3, 0x8b, 0x03, 0xef, 0xeb, 0x40, 0x60, 0x2e, 0x4e, 0x9c, 0xa4, 0x5b, 0x28, 0x0d, 0xff, 0xa6,
	0xf1, 0x34, 0xfa, 0xa1, 0xb1, 0x87, 0xc9, 0x31, 0x71, 0x73, 0xb9, 0x7c, 0x54, 0x9c, 0xff, 0xd8,
	0xc1, 0xce, 0xf5, 0xed, 0xf4, 0xaf, 0x75, 0x6f, 0xa5, 0xaf, 0xed, 0x85, 0xf5, 0xf4, 0xa9, 0x3d,
	0xf7, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x05, 0x02, 0x80, 0xee, 0x04, 0x00, 0x00,
}
