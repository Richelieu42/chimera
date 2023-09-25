package jsonKit

import (
	"github.com/richelieu-yang/chimera/v2/src/compareKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var library string

// defaultApi 默认的API
var defaultApi API = nil

// stdApi 标准的API（会对map的keys排序）
var stdApi API = nil

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}
)

func GetLibrary() string {
	return library
}

func GetDefaultApi() API {
	return defaultApi
}

func GetStdApi() API {
	return stdApi
}

func testAPI() {
	api := defaultApi

	m := map[string]interface{}{
		"0": 3.1415926,
		"1": 1,
		"2": true,
		"3": `~!@#$%^&*()_+{}|:"><?	
qwdqw强	\t\r\n无敌
qwdqwd
威尔法
496465~·《》？L:"{}|,./l;'[]\/*-的确
`,
	}
	jsonData, err := api.Marshal(m)
	if err != nil {
		logrus.WithError(err).Fatal("[%s, JSON] Fail to marshal!!!", consts.ProjectUpperName)
	}
	var m1 map[string]interface{}
	if err := api.Unmarshal(jsonData, &m1); err != nil {
		logrus.WithError(err).Fatal("[%s, JSON] Fail to unmarshal!!!", consts.ProjectUpperName)
	}
	if compareKit.Equal(m, m1) {
		logrusKit.DisableQuote(nil)
		diff := compareKit.Diff(m, m1)
		logrus.Fatalf("[%s, JSON] diff between m and m1:\n%s\n", consts.ProjectUpperName, diff)
	}
}
