package gzipKit

type (
	Config struct {
		Level int `json:"level" yaml:"level" validate:"min=-2,max=9"`

		CompressThreshold int `json:"compressThreshold" yaml:"compressThreshold"`
	}
)
