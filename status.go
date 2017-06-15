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
	codes.ResourceExhausted:  http.StatusForbidden,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Unavailable:        http.StatusServiceUnavailable,
}

var httpToGRPCmap = map[int]codes.Code{
	http.StatusOK:                 codes.OK,
	http.StatusBadRequest:         codes.InvalidArgument,
	http.StatusRequestTimeout:     codes.DeadlineExceeded,
	http.StatusNotFound:           codes.NotFound,
	http.StatusConflict:           codes.AlreadyExists,
	http.StatusForbidden:          codes.PermissionDenied,
	http.StatusUnauthorized:       codes.Unauthenticated,
	http.StatusPreconditionFailed: codes.FailedPrecondition,
	http.StatusNotImplemented:     codes.Unimplemented,
	http.StatusServiceUnavailable: codes.Unavailable,
}

// MapGRPCToHTTPCode returns the best HTTP code for a given GRPC code.
func MapGRPCToHTTPCode(c codes.Code) int {
	if result, ok := gRPCtoHTTPmap[c]; ok {
		return result
	}
	return http.StatusInternalServerError
}

func MapHTTPCodeToGRPC(code int) codes.Code {
	if result, ok := httpToGRPCmap[code]; ok {
		return result
	}
	return codes.Unknown
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
