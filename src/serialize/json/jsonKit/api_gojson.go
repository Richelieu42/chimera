//go:build !((linux || windows || darwin) && sonic && avx && go1.17 && amd64)

package jsonKit

import "github.com/goccy/go-json"

type gojsonImpl struct {
}

func (i *gojsonImpl) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (i *gojsonImpl) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (i *gojsonImpl) MarshalToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func (i *gojsonImpl) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (i *gojsonImpl) UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

func init() {
	library = "goccy/go-json"

	tmp := &gojsonImpl{}
	defaultApi = tmp
	stdApi = tmp
}
