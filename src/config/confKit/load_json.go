package confKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
)

// LoadFromJsonBytes 加载 .json 格式的配置文件内容.
var LoadFromJsonBytes func(content []byte, v any) error = conf.LoadFromJsonBytes

func LoadFromJsonText(text string, v any) error {
	return LoadFromJsonBytes([]byte(text), v)
}

func MustLoadFromJsonBytes(content []byte, v any) {
	if err := LoadFromJsonBytes(content, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func MustLoadFromJsonText(text string, v any) {
	if err := LoadFromJsonText(text, v); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}
