// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: zvelo/msg/apiv1.proto

/*
Package msg is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package msg

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_APIv1_Query_0(ctx context.Context, marshaler runtime.Marshaler, client APIv1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq QueryRequests
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.Query(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_APIv1_Result_0(ctx context.Context, marshaler runtime.Marshaler, client APIv1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RequestID
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["request_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "request_id")
	}

	protoReq.RequestId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "request_id", err)
	}

	msg, err := client.Result(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_APIv1_Suggest_0(ctx context.Context, marshaler runtime.Marshaler, client APIv1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Suggestion
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.Suggest(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_APIv1_Stream_0(ctx context.Context, marshaler runtime.Marshaler, client APIv1Client, req *http.Request, pathParams map[string]string) (APIv1_StreamClient, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	stream, err := client.Stream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterAPIv1HandlerFromEndpoint is same as RegisterAPIv1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAPIv1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAPIv1Handler(ctx, mux, conn)
}

// RegisterAPIv1Handler registers the http handlers for service APIv1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAPIv1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAPIv1HandlerClient(ctx, mux, NewAPIv1Client(conn))
}

// RegisterAPIv1Handler registers the http handlers for service APIv1 to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "APIv1Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "APIv1Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "APIv1Client" to call the correct interceptors.
func RegisterAPIv1HandlerClient(ctx context.Context, mux *runtime.ServeMux, client APIv1Client) error {

	mux.Handle("POST", pattern_APIv1_Query_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_APIv1_Query_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_APIv1_Query_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_APIv1_Result_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_APIv1_Result_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_APIv1_Result_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_APIv1_Suggest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_APIv1_Suggest_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_APIv1_Suggest_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_APIv1_Stream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_APIv1_Stream_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_APIv1_Stream_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_APIv1_Query_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "query"}, ""))

	pattern_APIv1_Result_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"v1", "query", "request_id"}, ""))

	pattern_APIv1_Suggest_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "suggest"}, ""))

	pattern_APIv1_Stream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "stream"}, ""))
)

var (
	forward_APIv1_Query_0 = runtime.ForwardResponseMessage

	forward_APIv1_Result_0 = runtime.ForwardResponseMessage

	forward_APIv1_Suggest_0 = runtime.ForwardResponseMessage

	forward_APIv1_Stream_0 = runtime.ForwardResponseStream
)
