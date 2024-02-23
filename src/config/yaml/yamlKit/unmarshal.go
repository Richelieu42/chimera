package yamlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"gopkg.in/yaml.v3"
)

// Unmarshal
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
var Unmarshal func(in []byte, out interface{}) (err error) = yaml.Unmarshal

// UnmarshalFromString
/*
PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
func UnmarshalFromString(in string, out interface{}) error {
	return Unmarshal([]byte(in), out)
}

func UnmarshalFromFile(path string, out interface{}) error {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return err
	}

	return Unmarshal(data, out)
}
