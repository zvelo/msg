package msg

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NOTE(jrubin) this file combines github.com/neelance/graphql-go/relay with the
// AnnotateContext func of github.com/grpc-ecosystem/grpc-gateway/runtime in
// order to proxy headers received via graphql to the APIClient grpc service
//
// github.com/grpc-ecosystem/grpc-gateway is licensed under BSD-3-Clause
// github.com/neelance/graphql-go/relay is licensed under BSD-2-Clause

type relay struct {
	Schema *graphql.Schema
}

const (
	metadataGrpcTimeout   = "Grpc-Timeout"
	xForwardedFor         = "X-Forwarded-For"
	xForwardedHost        = "X-Forwarded-Host"
	defaultContextTimeout = 0 * time.Second
)

func timeoutUnitToDuration(u uint8) (d time.Duration, ok bool) {
	switch u {
	case 'H':
		return time.Hour, true
	case 'M':
		return time.Minute, true
	case 'S':
		return time.Second, true
	case 'm':
		return time.Millisecond, true
	case 'u':
		return time.Microsecond, true
	case 'n':
		return time.Nanosecond, true
	default:
	}
	return
}

func timeoutDecode(s string) (time.Duration, error) {
	size := len(s)
	if size < 2 {
		return 0, fmt.Errorf("timeout string is too short: %q", s)
	}
	d, ok := timeoutUnitToDuration(s[size-1])
	if !ok {
		return 0, fmt.Errorf("timeout unit is not recognized: %q", s)
	}
	t, err := strconv.ParseInt(s[:size-1], 10, 64)
	if err != nil {
		return 0, err
	}
	return d * time.Duration(t), nil
}

var incomingHeaders = map[string]string{
	"zvelo-debug-id": "zvelo-debug-id",
	"uber-trace-id":  "uber-trace-id",
}

var outgoingHeaders = map[string]string{
	"zvelo-trace-id": "zvelo-trace-id",
	"server-timing":  "server-timing",
	"content-type":   "",
	"trailer":        "",
}

func IncomingHeaderMatcher(name string) (string, bool) {
	if n, ok := runtime.DefaultHeaderMatcher(name); ok {
		return n, ok
	}

	if n, ok := incomingHeaders[strings.ToLower(name)]; ok {
		if n == "" {
			return n, false
		}

		return n, true
	}

	return "", false
}

func OutgoingHeaderMatcher(name string) (string, bool) {
	if n, ok := outgoingHeaders[strings.ToLower(name)]; ok {
		if n == "" {
			return n, false
		}

		return n, true
	}

	return fmt.Sprintf("%s%s", runtime.MetadataHeaderPrefix, name), true
}

func annotateContext(req *http.Request) (context.Context, context.CancelFunc, error) {
	var pairs []string
	ctx := req.Context()
	cancel := func() {}

	timeout := defaultContextTimeout

	if tm := req.Header.Get(metadataGrpcTimeout); tm != "" {
		var err error
		timeout, err = timeoutDecode(tm)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "invalid grpc-timeout: %s", tm)
		}
	}

	if timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	}

	for k, vs := range req.Header {
		var ok bool
		if k, ok = IncomingHeaderMatcher(k); ok {
			for _, v := range vs {
				pairs = append(pairs, k, v)
			}
		}
	}

	if host := req.Header.Get(xForwardedHost); host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), host)
	} else if req.Host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), req.Host)
	}

	if addr := req.RemoteAddr; addr != "" {
		if remoteIP, _, err := net.SplitHostPort(addr); err == nil {
			if fwd := req.Header.Get(xForwardedFor); fwd == "" {
				pairs = append(pairs, strings.ToLower(xForwardedFor), remoteIP)
			} else {
				pairs = append(pairs, strings.ToLower(xForwardedFor), fmt.Sprintf("%s, %s", fwd, remoteIP))
			}
		}
	}

	if len(pairs) == 0 {
		return ctx, cancel, nil
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs(pairs...)), cancel, nil
}

type serverMetadataKey struct{}

type serverMetadata struct {
	sync.Mutex
	Header metadata.MD
}

func serverMetadataFromContext(ctx context.Context) *serverMetadata {
	if md, ok := ctx.Value(serverMetadataKey{}).(*serverMetadata); ok {
		return md
	}
	return &serverMetadata{}
}

func newServerMetadataContext(ctx context.Context, md *serverMetadata) context.Context {
	return context.WithValue(ctx, serverMetadataKey{}, md)
}

func (h relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel, err := annotateContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cancel()

	var md serverMetadata
	ctx = newServerMetadataContext(ctx, &md)

	response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	md.Lock()
	defer md.Unlock()

	for k, vs := range md.Header {
		var ok bool
		if k, ok = OutgoingHeaderMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
	}

	_, _ = w.Write(responseJSON) // #nosec
}
