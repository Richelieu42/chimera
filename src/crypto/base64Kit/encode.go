package base64Kit

import (
	"encoding/base64"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

// Encode []byte => []byte
/*
参考: gbase64.Encode()
*/
func Encode(src []byte, options ...Base64Option) []byte {
	opts := loadOptions(options...)

	dst := make([]byte, opts.encoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

// EncodeToString []byte => string
func EncodeToString(str []byte, options ...Base64Option) string {
	return string(Encode(str, options...))
}

// EncodeString string => []byte
func EncodeString(str string, options ...Base64Option) []byte {
	return Encode([]byte(str), options...)
}

// EncodeStringToString string => string
func EncodeStringToString(src string, options ...Base64Option) string {
	return string(EncodeString(src, options...))
}

// EncodeFile file => []byte
func EncodeFile(path string, options ...Base64Option) ([]byte, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Encode(data, options...), nil
}

// EncodeFileToString file => string
func EncodeFileToString(path string, options ...Base64Option) (string, error) {
	data, err := EncodeFile(path, options...)
	return string(data), err
}
