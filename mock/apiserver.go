package mock

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/segmentio/ksuid"

	msg "zvelo.io/msg/msgpb"
)

var jsonMarshaler = jsonpb.Marshaler{OrigName: true}

type apiServer struct {
	requestsLock sync.Mutex
	requests     map[string]*result

	streamsLock sync.Mutex
	streams     map[chan *msg.QueryResult]struct{}
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

func (r result) Clone() (*result, error) {
	var cpy msg.QueryResult
	if err := copyProto(&cpy, &r.QueryResult); err != nil {
		return nil, err
	}
	r.QueryResult = cpy
	return &r, nil
}

var _ msg.APIv1Server = (*apiServer)(nil)

func (s *apiServer) result(reqID string) (*result, error) {
	s.requestsLock.Lock()
	defer s.requestsLock.Unlock()

	r, ok := s.requests[reqID]
	if !ok || r == nil {
		return nil, status.Errorf(codes.NotFound, "request ID not found: %s", reqID)
	}

	r, err := r.Clone()
	if err != nil {
		return nil, err
	}

	if r.Complete() {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Complete = true
	}

	return r, nil
}

func copyProto(dst, src proto.Message) error {
	data, err := proto.Marshal(src)
	if err != nil {
		return err
	}

	return proto.Unmarshal(data, dst)
}

func (s *apiServer) store(r *result) {
	s.requestsLock.Lock()
	defer s.requestsLock.Unlock()

	r.StoredAt = time.Now()

	if s.requests == nil {
		s.requests = map[string]*result{}
	}

	s.requests[r.RequestId] = r
}

func (s *apiServer) handleResult(callbackURL, reqID string) {
	result, err := s.result(reqID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if result == nil {
		return
	}

	if !result.Complete() {
		time.AfterFunc(time.Until(result.CompleteAt()), func() {
			s.handleResult(callbackURL, reqID)
		})
		return
	}

	if callbackURL != "" {
		var buf bytes.Buffer
		if err = jsonMarshaler.Marshal(&buf, &result.QueryResult); err == nil {
			if _, err = http.Post(callbackURL, "application/json", &buf); err != nil {
				fmt.Fprintf(os.Stderr, "error posting callback: %s\n", err)
			}
		}
	}

	s.streamsLock.Lock()
	defer s.streamsLock.Unlock()

	for stream := range s.streams {
		select {
		case stream <- &result.QueryResult:
		default:
		}
	}
}

func (s *apiServer) handleResults(callbackURL string, reqIDs ...string) {
	for _, reqID := range reqIDs {
		go s.handleResult(callbackURL, reqID)
	}
}

func (s *apiServer) handleQuery(ctx context.Context, u string, content bool, ds []msg.DatasetType, out *msg.QueryReplies, reqIDs *[]string) error {
	var r result
	if err := parseOpts(ctx, u, content, ds, &r); err != nil {
		return status.Errorf(codes.Internal, "error parsing opts: %s", err)
	}

	r.RequestId = ksuid.New().String()
	*reqIDs = append(*reqIDs, r.RequestId)

	s.store(&r)

	out.Reply = append(out.Reply, &msg.QueryReply{
		RequestId: r.RequestId,
	})

	return nil
}

func (s *apiServer) Query(ctx context.Context, in *msg.QueryRequests) (*msg.QueryReplies, error) {
	if len(in.Dataset) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no datasets requested")
	}

	var out msg.QueryReplies
	var reqIDs []string

	for _, u := range in.Url {
		if u == "" {
			continue
		}

		if err := s.handleQuery(ctx, u, false, in.Dataset, &out, &reqIDs); err != nil {
			return nil, err
		}
	}

	for _, c := range in.Content {
		if c == nil || (c.Url == "" && c.Content == "") {
			continue
		}

		if err := s.handleQuery(ctx, c.Url, true, in.Dataset, &out, &reqIDs); err != nil {
			return nil, err
		}
	}

	if len(reqIDs) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no valid urls in request")
	}

	defer s.handleResults(in.Callback, reqIDs...)

	return &out, nil
}

func (s *apiServer) Result(_ context.Context, in *msg.RequestID) (*msg.QueryResult, error) {
	result, err := s.result(in.RequestId)
	if err != nil {
		return nil, err
	}

	return &result.QueryResult, nil
}

func (s *apiServer) Suggest(_ context.Context, _ *msg.Suggestion) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *apiServer) registerStream() chan *msg.QueryResult {
	ch := make(chan *msg.QueryResult)

	s.streamsLock.Lock()
	if s.streams == nil {
		s.streams = map[chan *msg.QueryResult]struct{}{}
	}
	s.streams[ch] = struct{}{}
	s.streamsLock.Unlock()

	return ch
}

func (s *apiServer) unregisterStream(ch chan *msg.QueryResult) {
	s.streamsLock.Lock()
	delete(s.streams, ch)
	s.streamsLock.Unlock()
}

func (s *apiServer) Stream(_ *empty.Empty, stream msg.APIv1_StreamServer) error {
	ch := s.registerStream()
	defer s.unregisterStream(ch)

	if err := stream.SendHeader(nil); err != nil {
		return err
	}

	for {
		select {
		case result := <-ch:
			if err := stream.Send(result); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}
