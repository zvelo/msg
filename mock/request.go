package mock

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"zvelo.io/msg"
	"zvelo.io/msg/status"
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

func WithMalicious(val ...msg.Category) ContextOption {
	return func(r *result) {
		if r.ResponseDataset == nil {
			r.ResponseDataset = &msg.DataSet{}
		}

		if r.ResponseDataset.Malicious == nil {
			r.ResponseDataset.Malicious = &msg.DataSet_Malicious{}
		}

		r.ResponseDataset.Malicious.Category = val
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
	headerCategory          = "zvelo-mock-category"
	headerMaliciousCategory = "zvelo-mock-malicious-category"
	headerCompleteAfter     = "zvelo-mock-complete-after"
	headerFetchCode         = "zvelo-mock-fetch-code"
	headerLocation          = "zvelo-mock-location"
	headerErrorCode         = "zvelo-mock-error-code"
	headerErrorMessage      = "zvelo-mock-error-message"
)

func QueryContext(ctx context.Context, opts ...ContextOption) context.Context {
	var r result
	for _, opt := range opts {
		opt(&r)
	}

	var pairs []string

	if ds := r.ResponseDataset; ds != nil {
		if c := ds.Categorization; c != nil {
			for _, cat := range c.Value {
				pairs = append(pairs, headerCategory, cat.String())
			}
		}

		if m := ds.Malicious; m != nil {
			for _, cat := range m.Category {
				pairs = append(pairs, headerMaliciousCategory, cat.String())
			}
		}
	}

	if r.CompleteAfter > 0 {
		pairs = append(pairs, headerCompleteAfter, r.CompleteAfter.String())
	}

	if qs := r.QueryStatus; qs != nil {
		if qs.FetchCode != 0 {
			pairs = append(pairs, headerFetchCode, strconv.Itoa(int(qs.FetchCode)))
		}

		if qs.Location != "" {
			pairs = append(pairs, headerLocation, qs.Location)
		}

		if e := qs.Error; e != nil {
			if e.Code != 0 {
				pairs = append(pairs, headerErrorCode, strconv.Itoa(int(e.Code)))
			}

			if e.Message != "" {
				pairs = append(pairs, headerErrorMessage, e.Message)
			}
		}
	}

	md := metadata.Pairs(pairs...)

	if cmd, ok := metadata.FromOutgoingContext(ctx); ok {
		md = metadata.Join(cmd, md)
	}

	return metadata.NewOutgoingContext(ctx, md)
}

func mdGet(md metadata.MD, key string) string {
	if md == nil {
		return ""
	}
	v := md[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func parseOpts(ctx context.Context, url string, content bool, ds []msg.DataSetType, r *result) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	for _, t := range ds {
		switch msg.DataSetType(t) {
		case msg.CATEGORIZATION:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Categorization = &msg.DataSet_Categorization{}

			if categoryNames, ok := md[headerCategory]; ok {
				categories := make([]msg.Category, len(categoryNames))
				for i, categoryName := range categoryNames {
					categories[i] = msg.ParseCategory(categoryName)
				}

				WithCategories(categories...)(r)
			}
		case msg.MALICIOUS:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Malicious = &msg.DataSet_Malicious{}

			if categoryNames, ok := md[headerMaliciousCategory]; ok {
				categories := make([]msg.Category, len(categoryNames))
				for i, categoryName := range categoryNames {
					categories[i] = msg.ParseCategory(categoryName)
				}

				WithMalicious(categories...)(r)
			}
		case msg.ECHO:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Echo = &msg.DataSet_Echo{Url: url}
		}
	}

	if s := mdGet(md, headerCompleteAfter); s != "" {
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}

		WithCompleteAfter(d)(r)
	}

	if c := mdGet(md, headerFetchCode); c != "" {
		code, err := strconv.ParseInt(c, 10, 32)
		if err != nil {
			return err
		}

		WithFetchCode(int32(code))(r)
	} else if !content {
		WithFetchCode(http.StatusOK)(r)
	}

	if l := mdGet(md, headerLocation); l != "" {
		WithLocation(l)(r)
	}

	var errorCode codes.Code
	if c := mdGet(md, headerErrorCode); c != "" {
		code, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			return err
		}

		errorCode = codes.Code(code)
	}

	errorMsg := mdGet(md, headerErrorMessage)

	if errorCode != 0 || errorMsg != "" {
		WithError(errorCode, errorMsg)(r)
	}

	return nil
}
