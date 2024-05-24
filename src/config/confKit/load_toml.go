package confKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
)

// LoadFromTomlBytes 加载 .toml 格式的配置文件内容.
var LoadFromTomlBytes func(content []byte, v any) error = conf.LoadFromTomlBytes

func LoadFromTomlText(text string, v any) error {
	return LoadFromTomlBytes([]byte(text), v)
}

func MustLoadFromTomlBytes(content []byte, v any) {
	if err := LoadFromTomlBytes(content, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func MustLoadFromTomlText(text string, v any) {
	if err := LoadFromTomlText(text, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}
