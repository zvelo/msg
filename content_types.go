package msg

/* Constants indicating the HTTP Content-type for each proto/json type in this package */
const (
	QueryContentType            string = "application/vnd.zvelo.query"
	QueryContentTypeJSON        string = "application/vnd.zvelo.query+json"
	ContentQueryContentType     string = "application/vnd.zvelo.content-query"
	ContentQueryContentTypeJSON string = "application/vnd.zvelo.content-query+json"
	QueryReplyContentType       string = "application/vnd.zvelo.query-reply"
	QueryReplyContentTypeJSON   string = "application/vnd.zvelo.query-reply+json"
	QueryResultContentType      string = "application/vnd.zvelo.query-result"
	QueryResultContentTypeJSON  string = "application/vnd.zvelo.query-result+json"

	StreamResultsContentType     string = "application/vnd.zvelo.stream-results"
	StreamResultsContentTypeJSON string = "application/vnd.zvelo.stream-results+json"
	StreamReplyContentType       string = "application/vnd.zvelo.stream-reply"
	StreamReplyContentTypeJSON   string = "application/vnd.zvelo.stream-reply+json"
	StreamsReplyContentType      string = "application/vnd.zvelo.streams-reply"
	StreamsReplyContentTypeJSON  string = "application/vnd.zvelo.streams-reply+json"
	StreamRequestContentType     string = "application/vnd.zvelo.stream-request"
	StreamRequestContentTypeJSON string = "application/vnd.zvelo.stream-request+json"

	SeedResultsContentType     string = "application/vnd.zvelo.seed-results"
	SeedResultsContentTypeJSON string = "application/vnd.zvelo.seed-results+json"
	SeedContentType            string = "application/vnd.zvelo.seed"
	SeedContentTypeJSON        string = "application/vnd.zvelo.seed+json"
)
