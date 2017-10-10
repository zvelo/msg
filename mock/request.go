package mock

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zvelo.io/msg"
)

type URLOption func(*result)

func WithCategories(val ...msg.Category) URLOption {
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

func WithMalicious(verdict msg.DataSet_Malicious_Verdict, cat msg.Category) URLOption {
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

func WithCompleteAfter(val time.Duration) URLOption {
	return func(r *result) {
		r.CompleteAfter = val
	}
}

func WithFetchCode(val int32) URLOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.FetchCode = val
	}
}

func WithLocation(val string) URLOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Location = val
	}
}

func WithError(c codes.Code, str string) URLOption {
	return func(r *result) {
		if r.QueryStatus == nil {
			r.QueryStatus = &msg.QueryStatus{}
		}

		r.QueryStatus.Error = status.New(c, str).Proto()
	}
}

func QueryURL(rawurl string, opts ...URLOption) (string, error) {
	if !strings.Contains(rawurl, "://") {
		rawurl = "http://" + rawurl
	}

	var r result
	for _, opt := range opts {
		opt(&r)
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	query := u.Query()

	if ds := r.ResponseDataset; ds != nil {
		if c := ds.Categorization; c != nil {
			for _, cat := range c.Value {
				query.Add("zvelo_cat", strconv.Itoa(int(cat)))
			}
		}

		if m := ds.Malicious; m != nil {
			if m.Category != 0 {
				query.Set("zvelo_malicious_category", strconv.Itoa(int(m.Category)))
			}

			if m.Verdict != msg.VERDICT_UNKNOWN {
				query.Set("zvelo_malicious_verdict", strconv.Itoa(int(m.Verdict)))
			}
		}
	}

	if r.CompleteAfter > 0 {
		query.Set("zvelo_complete_after", r.CompleteAfter.String())
	}

	if qs := r.QueryStatus; qs != nil {
		if qs.FetchCode != 0 {
			query.Set("zvelo_fetchcode", strconv.Itoa(int(qs.FetchCode)))
		}

		if qs.Location != "" {
			query.Set("zvelo_location", qs.Location)
		}

		if e := qs.Error; e != nil {
			if e.Code != 0 {
				query.Set("zvelo_errorcode", strconv.Itoa(int(e.Code)))
			}

			if e.Message != "" {
				query.Set("zvelo_errormsg", e.Message)
			}
		}
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

func parseURL(rawurl string, ds []msg.DataSetType, r *result) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	for _, t := range ds {
		switch msg.DataSetType(t) {
		case msg.CATEGORIZATION:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Categorization = &msg.DataSet_Categorization{}

			if catIDs, ok := u.Query()["zvelo_cat"]; ok {
				cats := make([]msg.Category, len(catIDs))
				for i, catID := range catIDs {
					cat, err := strconv.Atoi(catID)
					if err != nil {
						return err
					}
					cats[i] = msg.Category(cat)
				}

				WithCategories(cats...)(r)
			}
		case msg.MALICIOUS:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Malicious = &msg.DataSet_Malicious{}

			if v := u.Query().Get("zvelo_malicious_verdict"); v != "" {
				verdict, err := strconv.Atoi(v)
				if err != nil {
					return err
				}

				WithMalicious(msg.DataSet_Malicious_Verdict(verdict), msg.UNKNOWN_CATEGORY)(r)
			}

			if catID := u.Query().Get("zvelo_malicious_category"); catID != "" {
				cat, err := strconv.Atoi(catID)
				if err != nil {
					return err
				}

				WithMalicious(msg.VERDICT_MALICIOUS, msg.Category(cat))(r)
			}
		case msg.ECHO:
			if r.ResponseDataset == nil {
				r.ResponseDataset = &msg.DataSet{}
			}

			r.ResponseDataset.Echo = &msg.DataSet_Echo{Url: rawurl}
		}
	}

	if s := u.Query().Get("zvelo_complete_after"); s != "" {
		d, err := time.ParseDuration(s)
		if err != nil {
			return err
		}

		WithCompleteAfter(d)(r)
	}

	if c := u.Query().Get("zvelo_fetchcode"); c != "" {
		code, err := strconv.ParseInt(c, 10, 32)
		if err != nil {
			return err
		}

		WithFetchCode(int32(code))(r)
	} else if u != nil && u.Host != "" {
		WithFetchCode(http.StatusOK)(r)
	}

	if l := u.Query().Get("zvelo_location"); l != "" {
		WithLocation(l)(r)
	}

	var errorCode codes.Code
	if c := u.Query().Get("zvelo_errorcode"); c != "" {
		code, err := strconv.ParseUint(c, 10, 32)
		if err != nil {
			return err
		}

		errorCode = codes.Code(code)
	}

	errorMsg := u.Query().Get("zvelo_errormsg")

	if errorCode != 0 || errorMsg != "" {
		WithError(errorCode, errorMsg)(r)
	}

	return nil
}
