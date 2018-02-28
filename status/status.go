package status

import (
	fmt "fmt"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zvelo.io/msg"
)

// Status represents an RPC status code, message, and details.  It is immutable
// and should be created with New, Newf, or FromProto.
type Status struct {
	s *msg.Status
}

// Code returns the status code contained in s.
func (s *Status) Code() codes.Code {
	if s == nil || s.s == nil {
		return codes.OK
	}
	return codes.Code(s.s.Code)
}

// Message returns the message contained in s.
func (s *Status) Message() string {
	if s == nil || s.s == nil {
		return ""
	}
	return s.s.Message
}

// Proto returns s's status as an msg.Status proto message.
func (s *Status) Proto() *msg.Status {
	if s == nil {
		return nil
	}
	return proto.Clone(s.s).(*msg.Status)
}

// Err returns an immutable error representing s; returns nil if s.Code() is
// OK.
func (s *Status) Err() error {
	if s.Code() == codes.OK {
		return nil
	}
	return status.Error(s.Code(), s.Message())
}

// New returns a Status representing c and msg.
func New(c codes.Code, message string) *Status {
	return &Status{s: &msg.Status{Code: int32(c), Message: message}}
}

// Newf returns New(c, fmt.Sprintf(format, a...)).
func Newf(c codes.Code, format string, a ...interface{}) *Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Error returns an error representing c and msg.  If c is OK, returns nil.
func Error(c codes.Code, msg string) error {
	return status.Error(c, msg)
}

// Errorf returns Error(c, fmt.Sprintf(format, a...)).
func Errorf(c codes.Code, format string, a ...interface{}) error {
	return Error(c, fmt.Sprintf(format, a...))
}

// ErrorProto returns an error representing s.  If s.Code is OK, returns nil.
func ErrorProto(s *msg.Status) error {
	if s == nil {
		return nil
	}
	return status.Error(codes.Code(s.Code), s.Message)
}

// FromError returns a Status representing err if it was produced from this
// package, otherwise it returns nil, false.
func FromError(err error) (s *Status, ok bool) {
	st, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	return New(st.Code(), st.Message()), true
}

// Convert is a convenience function which removes the need to handle the
// boolean return value from FromError.
func Convert(err error) *Status {
	s, _ := FromError(err)
	return s
}
