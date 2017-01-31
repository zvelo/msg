package msg

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// NewContextError builds a correct error from a context
func NewContextError(ctx context.Context) error {
	if ctx.Err() == nil {
		return nil
	}

	switch ctx.Err() {
	case context.Canceled:
		return NewError(codes.Canceled, ctx.Err().Error())
	case context.DeadlineExceeded:
		return NewError(codes.DeadlineExceeded, ctx.Err().Error())
	}

	return NewError(codes.Unknown, ctx.Err().Error())
}

// NewError is used as a way to prevent go vet errors regarding Errorf
func NewError(c codes.Code, format string, a ...interface{}) error {
	return grpc.Errorf(c, format, a...)
}

// Any GRPC code not in this map is returned as an Internal Server Error.
var gRPCtoHTTPmap = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusRequestTimeout,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusRequestTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Unavailable:        http.StatusServiceUnavailable,
}

// MapGRPCToHTTPCode returns the best HTTP code for a given GRPC code.
func MapGRPCToHTTPCode(c codes.Code) int {
	result, ok := gRPCtoHTTPmap[c]
	if ok {
		return result
	}
	return http.StatusInternalServerError
}

// ErrorDesc returns the simple error message of a gRPC error
func ErrorDesc(in error) string {
	if in == nil {
		return ""
	}

	var code codes.Code
	var desc string

	// the scan format was taken from grpc
	// it is probably pretty fickle...
	if num, err := fmt.Sscanf(in.Error(), "rpc error: code = %d desc = %q", &code, &desc); err != nil || num != 2 {
		return in.Error()
	}

	return desc
}

// NewStatus creates a new Status object
func NewStatus(err error) *Status {
	if err == nil {
		return &Status{Code: STATUS_OK}
	}

	rpcCode := grpc.Code(err)
	statusCode := MapGRPCToHTTPCode(rpcCode)

	return &Status{
		Code:    Status_Code(statusCode),
		Message: ErrorDesc(err),
	}
}
