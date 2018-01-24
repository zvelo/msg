package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

func marshalJSON(s *spec.Swagger) ([]byte, error) {
	b1, err := json.MarshalIndent(s.SwaggerProps, "", "  ")
	if err != nil {
		return nil, err
	}

	b2, err := json.MarshalIndent(s.VendorExtensible, "", "  ")
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2), nil
}

func Patch(file string) error {
	doc, err := loads.Spec(file)
	if err != nil {
		return err
	}

	s := doc.Spec()
	s.BasePath = "/"
	s.Host = "api.zvelo.com"
	s.Schemes = []string{"https"}
	s.Info.Version = "1"
	s.Info.Contact = &spec.ContactInfo{
		Name:  "zvelo",
		URL:   "https://zvelo.com",
		Email: "support@zvelo.com",
	}
	s.Info.License = &spec.License{
		Name: "zvelo Proprietary",
	}

	s.SecurityDefinitions = spec.SecurityDefinitions{
		"authorization": &spec.SecurityScheme{
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Type:             "oauth2",
				Flow:             "accessCode",
				AuthorizationURL: "https://auth.zvelo.com/oauth2/auth",
				TokenURL:         "https://auth.zvelo.com/oauth2/token",
				Scopes: map[string]string{
					"zvelo.dataset": "Access zvelo Datasets",
				},
			},
		},
	}

	security := map[string][]string{
		"Post:/v1/query":             {"zvelo.dataset"},
		"Get:/v1/query/{request_id}": {"zvelo.dataset"},
	}

	for p, v := range s.Paths.Paths {
		rv := reflect.ValueOf(v)
		for _, method := range []string{
			"Get", "Put", "Post", "Delete", "Options", "Head", "Patch",
		} {
			f := rv.FieldByName(method)
			if f.IsNil() {
				continue
			}

			op, ok := f.Interface().(*spec.Operation)
			if !ok {
				continue
			}

			if sec, ok := security[method+":"+p]; ok {
				op.Security = []map[string][]string{{
					"authorization": sec,
				}}
			}

			m := op.Responses.StatusCodeResponses

			for code, r := range m {
				if r.Description == "" {
					r.Description = "n/a"
					m[code] = r
				}
			}
		}
	}

	data, err := marshalJSON(s)
	if err != nil {
		return err
	}

	data = bytes.Replace(data, []byte(`"n/a"`), []byte(`""`), -1)

	return ioutil.WriteFile(file, data, 0644)
}
