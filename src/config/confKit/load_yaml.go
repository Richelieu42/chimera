package confKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
)

// LoadFromYamlBytes 加载 .yaml 格式的配置文件内容.
var LoadFromYamlBytes func(content []byte, v any) error = conf.LoadFromYamlBytes

func LoadFromYamlText(text string, v any) error {
	return LoadFromYamlBytes([]byte(text), v)
}

func MustLoadFromYamlBytes(content []byte, v any) {
	if err := LoadFromYamlBytes(content, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func MustLoadFromYamlText(text string, v any) {
	if err := LoadFromYamlText(text, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}
