package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

func MarshalJSON(s *spec.Swagger) ([]byte, error) {
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

func main() {
	doc, err := loads.Spec("swagger.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
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
					"offline":           "Offline Access",
					"openid":            "Your Identity (including email)",
					"hydra":             "Manage All of Hydra",
					"hydra.clients":     "Manage Hydra Clients",
					"hydra.groups":      "Manage Warden Groups",
					"hydra.keys.create": "Create JSON Web Keys",
					"hydra.keys.delete": "Delete JSON Web Keys",
					"hydra.keys.get":    "Fetch JSON Web Keys",
					"hydra.keys.update": "Get JSON Web Keys",
					"hydra.policies":    "Manage Access Control Policies",
					"hydra.warden":      "Make Access Control Inquiries",
					"zvelo":             "Manage All of zvelo",
					"zvelo.dataset":     "Access zvelo Datasets",
					"zvelo.override":    "Manage zvelo Overrides",
					"zvelo.websocket":   "Request a zvelo WebSocket",
				},
			},
		},
	}

	security := map[string][]string{
		"Post:/v1/queries/content":             {"zvelo.dataset"},
		"Get:/v1/queries/content/{request_id}": {"zvelo.dataset"},
		"Post:/v1/queries/url":                 {"zvelo.dataset"},
		"Get:/v1/queries/url/{request_id}":     {"zvelo.dataset"},
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

	data, err := MarshalJSON(s)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	data = bytes.Replace(data, []byte(`"n/a"`), []byte(`""`), -1)

	if err = ioutil.WriteFile("swagger.json", data, 0644); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
