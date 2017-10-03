package mock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/segmentio/ksuid"

	"zvelo.io/msg"
)

type apiServer struct {
	mu       sync.Mutex
	requests map[string]*result
}

type result struct {
	msg.QueryResult
	StoredAt      time.Time
	CompleteAfter time.Duration
}

func (r result) CompleteAt() time.Time {
	return r.StoredAt.Add(r.CompleteAfter)
}

func (r result) Complete() bool {
	return time.Now().After(r.CompleteAt())
}

var _ msg.APIServer = (*apiServer)(nil)

func (s *apiServer) result(reqID string) *result {
	s.mu.Lock()
	defer s.mu.Unlock()

	r, ok := s.requests[reqID]
	if !ok {
		return nil
	}

	if r.Complete() {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Complete = true
	}

	return r
}

func (s *apiServer) store(r *result) {
	s.mu.Lock()
	defer s.mu.Unlock()

	r.StoredAt = time.Now()

	if s.requests == nil {
		s.requests = map[string]*result{}
	}

	s.requests[r.RequestId] = r
}

func (s *apiServer) handleCallback(u, reqID string) {
	result := s.result(reqID)
	if result == nil {
		return
	}

	if !result.Complete() {
		time.AfterFunc(time.Until(result.CompleteAt()), func() {
			s.handleCallback(u, reqID)
		})
		return
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(s.result(reqID)); err == nil {
		if _, err = http.Post(u, "application/json", &buf); err != nil {
			fmt.Fprintf(os.Stderr, "error posting callback: %s\n", err)
		}
	}
}

func (s *apiServer) postCallbacks(u string, reqIDs ...string) {
	if u == "" {
		return
	}

	for _, reqID := range reqIDs {
		go s.handleCallback(u, reqID)
	}
}

func (s *apiServer) handleQuery(u string, ds []uint32, out *msg.QueryReplies, reqIDs *[]string) error {
	var r result
	if err := parseURL(u, ds, &r); err != nil {
		return status.Errorf(codes.Internal, "error parsing url %s: %s", u, err)
	}

	r.RequestId = ksuid.New().String()
	*reqIDs = append(*reqIDs, r.RequestId)

	s.store(&r)

	out.Reply = append(out.Reply, &msg.QueryReply{
		RequestId: r.RequestId,
	})

	return nil
}

func (s *apiServer) QueryV1(_ context.Context, in *msg.QueryRequests) (*msg.QueryReplies, error) {
	if len(in.Dataset) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no datasets requested")
	}

	var out msg.QueryReplies
	var reqIDs []string

	for _, u := range in.Url {
		if u == "" {
			continue
		}

		if err := s.handleQuery(u, in.Dataset, &out, &reqIDs); err != nil {
			return nil, err
		}
	}

	for _, c := range in.Content {
		if c == nil || (c.Url == "" && c.Content == "") {
			continue
		}

		if err := s.handleQuery(c.Url, in.Dataset, &out, &reqIDs); err != nil {
			return nil, err
		}
	}

	if len(reqIDs) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no valid urls in request")
	}

	defer s.postCallbacks(in.Callback, reqIDs...)

	return &out, nil
}

func (s *apiServer) QueryResultV1(_ context.Context, in *msg.QueryPollRequest) (*msg.QueryResult, error) {
	result := s.result(in.RequestId)
	if result == nil {
		return nil, status.Errorf(codes.NotFound, "request ID not found: %s", in.RequestId)
	}

	return &result.QueryResult, nil
}
