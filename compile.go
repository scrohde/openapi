package openapi

import (
	"github.com/scrohde/openapi/objects"
	"gopkg.in/yaml.v3"
	"net/url"
	"os"
)

func Compile(rawURL string) (*objects.OpenAPI, error) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if URL.Scheme == "file" {
		b, err := os.ReadFile(URL.Path)
		if err != nil {
			return nil, err
		}
		openAPI := &objects.OpenAPI{}
		err = yaml.Unmarshal(b, openAPI)
		if err != nil {
			return nil, err
		}
		return openAPI, nil
	}
	panic("URL not yet supported")
}
