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
	Code        Status_Code `protobuf:"varint,1,opt,name=code,enum=zvelo.msg.Status_Code" json:"code,omitempty"`
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
	proto.RegisterType((*Status)(nil), "zvelo.msg.Status")
	proto.RegisterEnum("zvelo.msg.Status_Code", Status_Code_name, Status_Code_value)
}

func init() { proto.RegisterFile("zvelo/msg/status.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x54, 0x4b, 0x73, 0x1c, 0x35,
	0x10, 0x66, 0x46, 0xc6, 0xf6, 0xca, 0x8e, 0x23, 0xcb, 0x89, 0x59, 0x42, 0xa5, 0x42, 0xb9, 0x38,
	0x50, 0x1c, 0x36, 0x55, 0x81, 0x3f, 0x20, 0xcf, 0xf4, 0xce, 0xaa, 0x3c, 0x23, 0x0d, 0x92, 0xc6,
	0x9b, 0xc9, 0x45, 0x65, 0x9c, 0x25, 0xa4, 0x2a, 0x61, 0x29, 0x76, 0xe1, 0xc0, 0xaf, 0xe0, 0xfd,
	0x86, 0x00, 0xe7, 0x84, 0x40, 0xf1, 0x0f, 0xe0, 0xc2, 0xfb, 0xd7, 0xf0, 0x3c, 0x41, 0x4b, 0xfb,
	0x60, 0x53, 0x1c, 0xd5, 0xdd, 0xea, 0xef, 0xeb, 0xaf, 0x3f, 0x89, 0xee, 0xbf, 0xf6, 0xea, 0xe8,
	0xd6, 0xf8, 0xf2, 0xed, 0xc9, 0x8d, 0xcb, 0x93, 0xe9, 0xc9, 0xf4, 0x95, 0x49, 0xef, 0xa5, 0x97,
	0xc7, 0xd3, 0x31, 0xef, 0xc4, 0x78, 0x0f, 0xe3, 0x07, 0xdf, 0x76, 0xe8, 0xba, 0x8d, 0x39, 0xfe,
	0x14, 0x5d, 0x3b, 0x1d, 0x5f, 0x1f, 0x75, 0x93, 0xc7, 0x93, 0x27, 0x77, 0xae, 0xec, 0xf7, 0x96,
	0x45, 0xbd, 0x59, 0x41, 0x2f, 0xc3, 0xac, 0x89, 0x35, 0xbc, 0x4b, 0x37, 0x6e, 0x8f, 0x26, 0x93,
	0x93, 0x1b, 0xa3, 0x6e, 0x8a, 0xe5, 0x1d, 0xb3, 0x38, 0xf2, 0x67, 0xe8, 0xf6, 0xf3, 0xa3, 0xe9,
	0xe9, 0x0b, 0x7e, 0x86, 0xd8, 0x25, 0x98, 0xde, 0xba, 0xb2, 0xfb, 0xbf, 0x6e, 0x66, 0x2b, 0x96,
	0xcd, 0xb1, 0x2f, 0xd0, 0xcd, 0x5b, 0xe3, 0xd3, 0x93, 0xe9, 0xcd, 0xf1, 0x8b, 0xdd, 0xb5, 0xd8,
	0x70, 0x79, 0x3e, 0xb8, 0xb3, 0x49, 0xd7, 0x02, 0x34, 0xdf, 0xa2, 0x1b, 0x8d, 0x3a, 0x52, 0x7a,
	0xa8, 0xd8, 0x43, 0x7c, 0x9b, 0x6e, 0x66, 0x5a, 0x39, 0xa9, 0x1a, 0x60, 0xd7, 0xf9, 0x23, 0x74,
	0xcf, 0x0e, 0xa5, 0xcb, 0x06, 0x52, 0x15, 0xbe, 0x36, 0xda, 0xe9, 0x4c, 0x97, 0x96, 0x8d, 0xf8,
	0x06, 0x4d, 0xf5, 0x11, 0xfb, 0x3e, 0xc1, 0xfa, 0x8d, 0xcc, 0x80, 0x70, 0x90, 0xb3, 0x1f, 0x12,
	0x7e, 0x86, 0x6e, 0x8a, 0x2c, 0x83, 0x3a, 0x1c, 0x7f, 0x4c, 0xf8, 0x63, 0x74, 0x5f, 0x69, 0xe5,
	0x45, 0xe3, 0x06, 0xda, 0x48, 0x27, 0x9c, 0x3c, 0x06, 0x2f, 0x55, 0x5f, 0xb3, 0x9f, 0x12, 0x7e,
	0x96, 0x52, 0xa5, 0x7d, 0x00, 0x03, 0xe5, 0xd8, 0xcf, 0x09, 0xe7, 0xf4, 0x8c, 0x01, 0x0b, 0x6e,
	0x19, 0xfb, 0x25, 0xe1, 0xe7, 0xe8, 0xd9, 0x5a, 0x18, 0x27, 0x45, 0xb9, 0x8c, 0xfe, 0x9a, 0xf0,
	0xf3, 0x94, 0x55, 0x4d, 0xe9, 0x64, 0x5d, 0x82, 0xcf, 0x06, 0x5a, 0x66, 0x60, 0xd9, 0xdd, 0x94,
	0xef, 0xd3, 0xdd, 0x4a, 0x1f, 0x43, 0xee, 0x6b, 0x30, 0x95, 0x50, 0x58, 0x5c, 0xb6, 0xec, 0x5e,
	0xca, 0x29, 0x7d, 0xb8, 0xaf, 0x1b, 0x95, 0xb3, 0x2f, 0x53, 0xbe, 0x43, 0x3b, 0x16, 0xc0, 0x6b,
	0x37, 0x00, 0xc3, 0xee, 0xa7, 0x7c, 0x97, 0x6e, 0x2b, 0xed, 0x7c, 0xa5, 0x73, 0xd9, 0x97, 0xc8,
	0xfa, 0xab, 0x58, 0xd2, 0x58, 0x08, 0xe3, 0x5e, 0x6d, 0xd9, 0xd7, 0x29, 0x8a, 0xc0, 0x1d, 0x54,
	0xb5, 0x36, 0xc2, 0xb4, 0xde, 0x40, 0x2e, 0x0d, 0x64, 0x8e, 0x7d, 0x93, 0x72, 0x46, 0xb7, 0x0e,
	0x45, 0x8e, 0xa1, 0x67, 0x1b, 0xb0, 0x8e, 0xbd, 0x4e, 0x42, 0xb7, 0x46, 0xcd, 0xc7, 0xbd, 0x86,
	0xdd, 0xde, 0x20, 0x81, 0x6b, 0x2d, 0xda, 0x0a, 0xc9, 0xc4, 0x42, 0xbc, 0x9c, 0xb3, 0x37, 0x49,
	0x00, 0xe9, 0x6b, 0x73, 0x28, 0xf3, 0x1c, 0x14, 0x7b, 0x2b, 0x9e, 0x03, 0x8f, 0x19, 0xcf, 0xb7,
	0x49, 0x00, 0xad, 0x00, 0xfb, 0xe4, 0x3e, 0x84, 0x45, 0x59, 0xea, 0x21, 0x5e, 0x7c, 0x87, 0xf0,
	0x3d, 0xba, 0x13, 0x23, 0x51, 0x66, 0x71, 0x58, 0x02, 0x7b, 0x97, 0xa0, 0x6f, 0xf6, 0x22, 0xdd,
	0x28, 0xf5, 0x7f, 0x38, 0xef, 0x91, 0x20, 0xe0, 0x9c, 0x9f, 0x77, 0xb2, 0x02, 0xdd, 0x38, 0xf6,
	0x3e, 0x09, 0x7b, 0x42, 0x39, 0xfb, 0xa5, 0xc4, 0x41, 0x3e, 0x20, 0xbc, 0x43, 0xd7, 0x0a, 0xad,
	0x80, 0x7d, 0x18, 0xeb, 0x4b, 0x50, 0xc5, 0x6a, 0x97, 0x8f, 0xe6, 0xfd, 0x01, 0xaf, 0xe4, 0xd2,
	0x49, 0xdc, 0x68, 0x5f, 0xc8, 0x12, 0x33, 0x1f, 0x13, 0x7e, 0x91, 0x76, 0x17, 0xfd, 0x71, 0x44,
	0xe9, 0x5a, 0xef, 0xb4, 0xf6, 0xa5, 0x30, 0x05, 0xb0, 0x4f, 0x08, 0x7f, 0x94, 0x9e, 0x5b, 0xa4,
	0x1b, 0x23, 0x67, 0x39, 0xad, 0x0a, 0xf6, 0x29, 0x09, 0xe6, 0x68, 0x94, 0x6d, 0x6a, 0x14, 0x16,
	0xed, 0xe2, 0x2b, 0x14, 0x56, 0x78, 0xd7, 0xd6, 0xc0, 0xee, 0x10, 0xfe, 0x04, 0xbd, 0x34, 0xbf,
	0x87, 0x29, 0x23, 0x54, 0x01, 0x51, 0x07, 0x8b, 0x16, 0xb2, 0x7d, 0x19, 0xc7, 0xfe, 0x2c, 0x8a,
	0x04, 0x57, 0x6b, 0xdc, 0x86, 0x58, 0x65, 0xf5, 0x39, 0x41, 0x4b, 0xaf, 0x3b, 0x10, 0xb5, 0x76,
	0xec, 0x0b, 0x82, 0x8f, 0xe0, 0xfc, 0x03, 0xe4, 0x97, 0x83, 0xdd, 0x25, 0xc1, 0x32, 0x81, 0x13,
	0xba, 0xa5, 0x5d, 0xec, 0xd1, 0xb2, 0x7b, 0xab, 0xf8, 0x7e, 0x00, 0x22, 0x07, 0xe3, 0xd1, 0x1d,
	0x65, 0x6e, 0x57, 0xa6, 0xbb, 0x4f, 0xf8, 0x01, 0xbd, 0x88, 0xeb, 0x3e, 0x46, 0xd8, 0xc0, 0x08,
	0x97, 0x67, 0x7c, 0x09, 0x05, 0x7a, 0x15, 0x5f, 0x84, 0xd5, 0xca, 0xb2, 0xef, 0x22, 0xba, 0x44,
	0xe3, 0x1a, 0x85, 0x61, 0x0b, 0xe6, 0x18, 0x5b, 0x81, 0x31, 0xda, 0xb0, 0xdf, 0xa2, 0xd8, 0x61,
	0x2a, 0x59, 0xa1, 0x91, 0x83, 0x43, 0x90, 0xd3, 0xef, 0x64, 0x61, 0xab, 0x02, 0x1f, 0xd5, 0x50,
	0xb4, 0xec, 0x8f, 0x28, 0x7f, 0xb8, 0x8a, 0x3e, 0xf7, 0x2b, 0x78, 0xec, 0xcf, 0xd8, 0x61, 0x5e,
	0xb7, 0x5c, 0xef, 0x5f, 0x84, 0x5f, 0xa2, 0x17, 0x06, 0xce, 0xd5, 0x1e, 0xc1, 0x6c, 0x98, 0x38,
	0x4a, 0xb7, 0x50, 0x9a, 0xfd, 0x1d, 0xc7, 0x53, 0xe0, 0x86, 0xda, 0x1c, 0x45, 0xc7, 0x84, 0xcd,
	0x65, 0xe2, 0x41, 0x71, 0xfe, 0x21, 0x87, 0x3b, 0xd7, 0xb6, 0x67, 0xdf, 0xcb, 0xcd, 0xf8, 0xd9,
	0x3d, 0xb7, 0x1e, 0xbf, 0xb9, 0xa7, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x9e, 0x73, 0x16,
	0x00, 0x05, 0x00, 0x00,
}
