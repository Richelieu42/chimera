package floatKit

import (
	"github.com/spf13/cast"
)

func ToFloat32(i interface{}) float32 {
	return cast.ToFloat32(i)
}

func ToFloat32E(i interface{}) (float32, error) {
	return cast.ToFloat32E(i)
}

func ToFloat64(i interface{}) float64 {
	return cast.ToFloat64(i)
}

func ToFloat64E(i interface{}) (float64, error) {
	return cast.ToFloat64E(i)
}
