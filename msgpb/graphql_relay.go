package msgpb

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"

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

	for key, vals := range req.Header {
		for _, val := range vals {
			pairs = append(pairs, key, val)
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

	if timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	}
	if len(pairs) == 0 {
		return ctx, cancel, nil
	}
	md := metadata.Pairs(pairs...)
	return metadata.NewOutgoingContext(ctx, md), cancel, nil
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

	response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(responseJSON) // #nosec
}
