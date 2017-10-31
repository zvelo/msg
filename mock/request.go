package mock

import (
	"context"
	"net/http"
	"net/textproto"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"zvelo.io/msg"
)

type ContextOption func(*result)

func WithCategories(val ...msg.Category) ContextOption {
	return func(r *result) {
		if len(val) == 0 {
			return
		}

		if r.ResponseDataset == nil {
			r.ResponseDataset = &msg.DataSet{}
		}

		if r.ResponseDataset.Categorization == nil {
			r.ResponseDataset.Categorization = &msg.DataSet_Categorization{}
		}

		r.ResponseDataset.Categorization.Value = val
	}
}

func WithMalicious(verdict msg.DataSet_Malicious_Verdict, cat msg.Category) ContextOption {
	return func(r *result) {
		if r.ResponseDataset == nil {
			r.ResponseDataset = &msg.DataSet{}
		}

		if r.ResponseDataset.Malicious == nil {
			r.ResponseDataset.Malicious = &msg.DataSet_Malicious{}
		}

		r.ResponseDataset.Malicious.Verdict = verdict
		r.ResponseDataset.Malicious.Category = cat
	}
}

func WithCompleteAfter(val time.Duration) ContextOption {
	return func(r *result) {
		r.CompleteAfter = val
	}
}

func WithFetchCode(val int32) ContextOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.FetchCode = val
	}
}

func WithLocation(val string) ContextOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Location = val
	}
}

func WithError(c codes.Code, str string) ContextOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Error = status.New(c, str).Proto()
	}
}

const (
	headerCategory          = "Zvelo-Mock-Category"
	headerMaliciousVerdict  = "Zvelo-Mock-Malicious-Verdict"
	headerMaliciousCategory = "Zvelo-Mock-Malicious-Category"
	headerCompleteAfter     = "Zvelo-Mock-Complete-After"
	headerFetchCode         = "Zvelo-Mock-Fetch-Code"
	headerLocation          = "Zvelo-Mock-Location"
	headerErrorCode         = "Zvelo-Mock-Error-Code"
	headerErrorMessage      = "Zvelo-Mock-Error-Message"
)

func QueryContext(ctx context.Context, opts ...ContextOption) context.Context {
	var r result
	for _, opt := range opts {
		opt(&r)
	}

	h := http.Header{}
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		h = http.Header(md.Copy())
	}

	if ds := r.ResponseDataset; ds != nil {
		if c := ds.Categorization; c != nil {
			for _, cat := range c.Value {
				h.Add(headerCategory, cat.String())
			}
		}

		if m := ds.Malicious; m != nil {
			if m.Category != 0 {
				h.Set(headerMaliciousCategory, m.Category.String())
			}

			if m.Verdict != msg.VERDICT_UNKNOWN {
				h.Set(headerMaliciousVerdict, m.Verdict.String())
			}
		}
	}

	if r.CompleteAfter > 0 {
		h.Set(headerCompleteAfter, r.CompleteAfter.String())
	}

	if qs := r.QueryStatus; qs != nil {
		if qs.FetchCode != 0 {
			h.Set(headerFetchCode, strconv.Itoa(int(qs.FetchCode)))
		}

		if qs.Location != "" {
			h.Set(headerLocation, qs.Location)
		}

		if e := qs.Error; e != nil {
			if e.Code != 0 {
				h.Set(headerErrorCode, strconv.Itoa(int(e.Code)))
			}

			if e.Message != "" {
				h.Set(headerErrorMessage, e.Message)
			}
		}
	}

	return metadata.NewOutgoingContext(ctx, metadata.MD(h))
}

func parseOpts(ctx context.Context, url string, content bool, ds []msg.DataSetType, r *result) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	h := http.Header(md)

	for _, t := range ds {
		switch msg.DataSetType(t) {
		case msg.CATEGORIZATION:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Categorization = &msg.DataSet_Categorization{}

			if categoryNames, ok := h[textproto.CanonicalMIMEHeaderKey(headerCategory)]; ok {
				categories := make([]msg.Category, len(categoryNames))
				for i, categoryName := range categoryNames {
					if categoryID, ok := msg.Category_value[categoryName]; ok {
						categories[i] = msg.Category(categoryID)
					}
				}

				WithCategories(categories...)(r)
			}
		case msg.MALICIOUS:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Malicious = &msg.DataSet_Malicious{}

			if v := h.Get(headerMaliciousVerdict); v != "" {
				if verdict, ok := msg.DataSet_Malicious_Verdict_value[v]; ok {
					WithMalicious(msg.DataSet_Malicious_Verdict(verdict), msg.UNKNOWN_CATEGORY)(r)
				}
			}

			if categoryName := h.Get(headerMaliciousCategory); categoryName != "" {
				if categoryID, ok := msg.Category_value[categoryName]; ok {
					WithMalicious(msg.VERDICT_MALICIOUS, msg.Category(categoryID))(r)
				}
			}
		case msg.ECHO:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Echo = &msg.DataSet_Echo{Url: url}
		}
	}

	if s := h.Get(headerCompleteAfter); s != "" {
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}

		WithCompleteAfter(d)(r)
	}

	if c := h.Get(headerFetchCode); c != "" {
		code, err := strconv.ParseInt(c, 10, 32)
		if err != nil {
			return err
		}

		WithFetchCode(int32(code))(r)
	} else if !content {
		WithFetchCode(http.StatusOK)(r)
	}

	if l := h.Get(headerLocation); l != "" {
		WithLocation(l)(r)
	}

	var errorCode codes.Code
	if c := h.Get(headerErrorCode); c != "" {
		code, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			return err
		}

		errorCode = codes.Code(code)
	}

	errorMsg := h.Get(headerErrorMessage)

	if errorCode != 0 || errorMsg != "" {
		WithError(errorCode, errorMsg)(r)
	}

	return nil
}
