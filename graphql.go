package msg

import (
	"context"
	"net/http"

	grpc "google.golang.org/grpc"

	"github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"

	"zvelo.io/msg/internal/static"
)

func GraphQLHandler(client APIv1Client, opts ...graphql.SchemaOpt) (http.Handler, error) {
	schemaFile, err := static.FSString(false, "/schema.graphql")
	if err != nil {
		return nil, err
	}

	schema, err := graphql.ParseSchema(schemaFile, &graphQLResolver{client: client}, opts...)
	if err != nil {
		return nil, err
	}

	return relay{Schema: schema}, nil
}

type graphQLResolver struct {
	client APIv1Client
}

func (r graphQLResolver) SuggestURL(ctx context.Context, args graphQLSuggestURL) error {
	s := Suggestion{
		Url: args.URL,
	}

	if args.DataSet != nil {
		s.Dataset = &DataSet{}

		if cs := args.DataSet.Categorization; cs != nil {
			s.Dataset.Categorization = &DataSet_Categorization{}

			for _, cn := range *cs {
				if c := ParseCategory(cn); c != UNKNOWN_CATEGORY {
					s.Dataset.Categorization.Value = append(s.Dataset.Categorization.Value, c)
				}
			}
		}

		if ms := args.DataSet.Malicious; ms != nil {
			s.Dataset.Malicious = &DataSet_Malicious{}

			for _, mn := range *ms {
				if m := ParseCategory(mn); m != UNKNOWN_CATEGORY {
					s.Dataset.Malicious.Category = append(s.Dataset.Malicious.Category, m)
				}
			}
		}
	}

	md := serverMetadataFromContext(ctx)
	md.Lock()
	defer md.Unlock()

	_, err := r.client.Suggest(ctx, &s, grpc.Header(&md.Header))
	return err
}

func (r graphQLResolver) QueryURL(ctx context.Context, args graphQLQueryURL) (*graphQLQueryReply, error) {
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
		req.Dataset = append(req.Dataset, DataSetType(id))
	}

	md := serverMetadataFromContext(ctx)
	md.Lock()
	defer md.Unlock()

	replies, err := r.client.Query(ctx, &req, grpc.Header(&md.Header))
	if err != nil {
		return nil, err
	}

	if len(replies.Reply) == 0 {
		return nil, errors.New("didn't get a reply")
	}

	return &graphQLQueryReply{replies.Reply[0]}, nil
}

func (r graphQLResolver) QueryContent(ctx context.Context, args graphQLQueryContent) (*graphQLQueryReply, error) {
	content := URLContent{
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
		Content: []*URLContent{&content},
	}

	if args.Callback != nil {
		req.Callback = *args.Callback
	}

	for _, name := range args.DataSet {
		id, ok := DataSetType_value[name]
		if !ok {
			return nil, errors.Errorf("invalid dataset type: %s", name)
		}
		req.Dataset = append(req.Dataset, DataSetType(id))
	}

	md := serverMetadataFromContext(ctx)
	md.Lock()
	defer md.Unlock()

	replies, err := r.client.Query(ctx, &req, grpc.Header(&md.Header))
	if err != nil {
		return nil, err
	}

	if len(replies.Reply) == 0 {
		return nil, errors.New("didn't get a reply")
	}

	return &graphQLQueryReply{replies.Reply[0]}, nil
}

func (r graphQLResolver) Result(ctx context.Context, args struct{ RequestID graphql.ID }) (*graphQLQueryResult, error) {
	md := serverMetadataFromContext(ctx)
	md.Lock()
	defer md.Unlock()

	result, err := r.client.Result(ctx, &RequestID{RequestId: string(args.RequestID)}, grpc.Header(&md.Header))
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

type graphQLSuggestURL struct {
	URL     string
	DataSet *graphQLDataSetInput
}

type graphQLDataSetInput struct {
	Categorization *[]string
	Malicious      *[]string
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
	if r.msg == nil || r.msg.Error == nil {
		return nil
	}

	return &graphQLStatus{r.msg.Error}
}

type graphQLStatus struct {
	msg *Status
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
	if r.msg == nil || r.msg.ResponseDataset == nil {
		return nil
	}
	return &graphQLDataSet{r.msg.ResponseDataset}
}

func (r graphQLQueryResult) QueryStatus() *graphQLQueryStatus {
	if r.msg == nil || r.msg.QueryStatus == nil {
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
	if s.msg == nil || s.msg.Error == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}

func (s graphQLQueryStatus) FetchCode() *int32 {
	if s.msg == nil {
		return nil
	}
	code := s.msg.FetchCode
	return &code
}

func (s graphQLQueryStatus) Location() *string {
	if s.msg == nil {
		return nil
	}
	loc := s.msg.Location
	return &loc
}

type graphQLDataSet struct {
	msg *DataSet
}

func (s graphQLDataSet) Categorization() *graphQLDataSetCategorization {
	if s.msg == nil || s.msg.Categorization == nil {
		return nil
	}

	return &graphQLDataSetCategorization{s.msg.Categorization}
}

func (s graphQLDataSet) Malicious() *graphQLDataSetMalicious {
	if s.msg == nil || s.msg.Malicious == nil {
		return nil
	}

	return &graphQLDataSetMalicious{s.msg.Malicious}
}

func (s graphQLDataSet) Echo() *graphQLDataSetEcho {
	if s.msg == nil || s.msg.Echo == nil {
		return nil
	}

	return &graphQLDataSetEcho{s.msg.Echo}
}

func (s graphQLDataSet) Language() *graphQLDataSetLanguage {
	if s.msg == nil || s.msg.Language == nil {
		return nil
	}

	return &graphQLDataSetLanguage{s.msg.Language}
}

type graphQLDataSetCategorization struct {
	msg *DataSet_Categorization
}

func (s graphQLDataSetCategorization) Value() *[]string {
	if s.msg == nil || len(s.msg.Value) == 0 {
		return nil
	}

	var ret []string
	for _, c := range s.msg.Value {
		ret = append(ret, c.String())
	}
	return &ret
}

func (s graphQLDataSetCategorization) Error() *graphQLStatus {
	if s.msg == nil || s.msg.Error == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}

type graphQLDataSetMalicious struct {
	msg *DataSet_Malicious
}

func (s graphQLDataSetMalicious) Category() *[]string {
	if s.msg == nil || len(s.msg.Category) == 0 {
		return nil
	}

	var ret []string
	for _, c := range s.msg.Category {
		ret = append(ret, c.String())
	}
	return &ret
}

func (s graphQLDataSetMalicious) Error() *graphQLStatus {
	if s.msg == nil || s.msg.Error == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}

type graphQLDataSetEcho struct {
	msg *DataSet_Echo
}

func (s graphQLDataSetEcho) URL() *string {
	if s.msg == nil {
		return nil
	}
	return &s.msg.Url
}

func (s graphQLDataSetEcho) Error() *graphQLStatus {
	if s.msg == nil || s.msg.Error == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}

type graphQLDataSetLanguage struct {
	msg *DataSet_Language
}

func (s graphQLDataSetLanguage) Code() *string {
	if s.msg == nil {
		return nil
	}
	return &s.msg.Code
}

func (s graphQLDataSetLanguage) Error() *graphQLStatus {
	if s.msg == nil || s.msg.Error == nil {
		return nil
	}
	return &graphQLStatus{s.msg.Error}
}
