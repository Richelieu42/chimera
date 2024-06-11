package rKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

func NewClient() (client *req.Client) {
	client = req.C()

	/* json序列化和反序列化 */
	client.SetJsonMarshal(jsonKit.Marshal)
	client.SetJsonUnmarshal(jsonKit.Unmarshal)

	/* 伪装成Chrome浏览器发起请求，主要针对: 反爬虫检测 */
	client.ImpersonateChrome()

	return
}

func NewRetryClient() (client *req.Client) {
	client = NewClient()

	// TODO:

	return
}
