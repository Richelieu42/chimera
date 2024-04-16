package brotliKit

type (
	Config struct {
		Level int `json:"level" yaml:"level" validate:"min=0,max=11"`

		CompressThreshold int `json:"compressThreshold" yaml:"compressThreshold"`
	}
)
