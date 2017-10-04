package msg

import (
	"context"
	"net/http"

	google_rpc "google.golang.org/genproto/googleapis/rpc/status"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/pkg/errors"

	igraphql "zvelo.io/msg/internal/graphql"
)

func GraphQLHandler(client APIClient, opts ...graphql.SchemaOpt) (http.Handler, error) {
	schemaFile, err := igraphql.FSString(false, "/schema.graphql")
	if err != nil {
		return nil, err
	}

	schema, err := graphql.ParseSchema(schemaFile, &graphQLResolver{client: client}, opts...)
	if err != nil {
		return nil, err
	}
	return &relay.Handler{Schema: schema}, nil
}

type graphQLResolver struct {
	client APIClient
}

func (r graphQLResolver) URL(ctx context.Context, args graphQLQueryURL) (*graphQLQueryReply, error) {
	req := QueryRequests{
		Url: []string{args.URL},
	}

	if args.Callback != nil {
		req.Callback = *args.Callback
	}

	for _, name := range args.DataSet {
		id, ok := DataSetType_value[name]
		if !ok {
			return nil, errors.Errorf("invalid dataset type: %s", name)
		}
		req.Dataset = append(req.Dataset, uint32(id))
	}

	replies, err := r.client.QueryV1(ctx, &req)
	if err != nil {
		return nil, err
	}

	if len(replies.Reply) == 0 {
		return nil, errors.New("didn't get a reply")
	}

	return &graphQLQueryReply{replies.Reply[0]}, nil
}

func (r graphQLResolver) Content(ctx context.Context, args graphQLQueryContent) (*graphQLQueryReply, error) {
	content := QueryRequests_URLContent{
		Content: args.Content.Content,
	}

	if args.Content.URL != nil {
		content.Url = *args.Content.URL
	}

	if args.Content.Header != nil {
		for _, h := range *args.Content.Header {
			content.Header[h.Name] = h.Value
		}
	}

	req := QueryRequests{
		Content: []*QueryRequests_URLContent{&content},
	}

	if args.Callback != nil {
		req.Callback = *args.Callback
	}

	for _, name := range args.DataSet {
		id, ok := DataSetType_value[name]
		if !ok {
			return nil, errors.Errorf("invalid dataset type: %s", name)
		}
		req.Dataset = append(req.Dataset, uint32(id))
	}

	replies, err := r.client.QueryV1(ctx, &req)
	if err != nil {
		return nil, err
	}

	if len(replies.Reply) == 0 {
		return nil, errors.New("didn't get a reply")
	}

	return &graphQLQueryReply{replies.Reply[0]}, nil
}

func (r graphQLResolver) Result(ctx context.Context, args struct{ RequestID graphql.ID }) (*graphQLQueryResult, error) {
	result, err := r.client.QueryResultV1(ctx, &QueryPollRequest{RequestId: string(args.RequestID)})
	if err != nil {
		return nil, err
	}

	return &graphQLQueryResult{result}, nil
}

type graphQLHeader struct {
	Name  string
	Value string
}

type graphQLURLContent struct {
	URL     *string
	Header  *[]graphQLHeader
	Content string
}

type graphQLQueryURL struct {
	URL      string
	Callback *string
	DataSet  []string
}

type graphQLQueryContent struct {
	Content  graphQLURLContent
	Callback *string
	DataSet  []string
}

type graphQLQueryReply struct {
	msg *QueryReply
}

func (r graphQLQueryReply) RequestID() *graphql.ID {
	if r.msg == nil {
		return nil
	}

	id := graphql.ID(r.msg.RequestId)
	return &id
}

func (r graphQLQueryReply) Error() *graphQLStatus {
	if r.msg == nil {
		return nil
	}

	return &graphQLStatus{r.msg.Error}
}

type graphQLStatus struct {
	msg *google_rpc.Status
}

func (s graphQLStatus) Code() int32 {
	if s.msg == nil {
		return 0
	}
	return s.msg.Code
}

func (s graphQLStatus) Message() string {
	if s.msg == nil {
		return ""
	}
	return s.msg.Message
}

type graphQLQueryResult struct {
	msg *QueryResult
}

func (r graphQLQueryResult) RequestID() *graphql.ID {
	if r.msg == nil {
		return nil
	}
	id := graphql.ID(r.msg.RequestId)
	return &id
}

func (r graphQLQueryResult) ResponseDataSet() *graphQLDataSet {
	if r.msg == nil {
		return nil
	}
	return &graphQLDataSet{r.msg.ResponseDataset}
}

func (r graphQLQueryResult) QueryStatus() *graphQLQueryStatus {
	if r.msg == nil {
		return nil
	}
	return &graphQLQueryStatus{r.msg.QueryStatus}
}

type graphQLQueryStatus struct {
	msg *QueryStatus
}

func (s graphQLQueryStatus) Complete() bool {
	if s.msg == nil {
		return false
	}

	return s.msg.Complete
}

func (s graphQLQueryStatus) Error() *graphQLStatus {
	if s.msg == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}

func (s graphQLQueryStatus) FetchCode() *int32 {
	if s.msg == nil {
		return nil
	}
	return &s.msg.FetchCode
}

func (s graphQLQueryStatus) Location() *string {
	if s.msg == nil {
		return nil
	}
	return &s.msg.Location
}

type graphQLDataSet struct {
	msg *DataSet
}

func (s graphQLDataSet) Categorization() *[]string {
	if s.msg == nil || s.msg.Categorization == nil {
		return nil
	}

	categories := make([]string, len(s.msg.Categorization.Value))
	for i, id := range s.msg.Categorization.Value {
		categories[i] = Category(id).String()
	}

	return &categories
}

func (s graphQLDataSet) Malicious() *graphQLDataSetMalicious {
	if s.msg == nil {
		return nil
	}

	return &graphQLDataSetMalicious{s.msg.Malicious}
}

func (s graphQLDataSet) Echo() *string {
	if s.msg == nil || s.msg.Echo == nil {
		return nil
	}

	return &s.msg.Echo.Url
}

type graphQLDataSetMalicious struct {
	msg *DataSet_Malicious
}

func (s graphQLDataSetMalicious) Category() *string {
	if s.msg == nil {
		return nil
	}
	str := Category(s.msg.Category).String()
	return &str
}

func (s graphQLDataSetMalicious) Verdict() *string {
	if s.msg == nil {
		return nil
	}
	str := DataSet_Malicious_Verdict(s.msg.Verdict).String()
	return &str
}
