package jsonKit

var defaultAPI API = nil

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}
)

func GetDefaultAPI() API {
	return defaultAPI
}

func SetDefaultAPI(api API) {
	defaultAPI = api
}
