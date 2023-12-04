package k8sYamlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"os"
	"sigs.k8s.io/yaml"
)

var Marshal func(o interface{}) ([]byte, error) = yaml.Marshal

func MarshalToString(in interface{}) (string, error) {
	data, err := Marshal(in)
	return string(data), err
}

func MarshalToFile(in interface{}, filePath string, perm os.FileMode) error {
	data, err := Marshal(in)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(data, filePath, perm)
}
