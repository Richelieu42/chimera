package rKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"time"
)

// NewClient
/*
@param insecureSkipVerify 	true:  跳过证书验证
							false: 不跳过证书验证（默认; 更加安全）
*/
func NewClient(insecureSkipVerify bool) (client *req.Client) {
	client = req.C()

	/*
		启用自动检测字符集并解码为utf-8
		（imroc/req默认: 启用）
	*/
	client.EnableAutoDecode()

	/*
		启用压缩
		（imroc/req默认: 启用）
	*/
	client.EnableCompression()

	/*
		disable sending GET method requests with body
		（imroc/req默认: 启用，即发送GET请求时附带body）
	*/
	client.DisableAllowGetMethodPayload()

	/* json序列化和反序列化 */
	client.SetJsonMarshal(jsonKit.Marshal).
		SetJsonUnmarshal(jsonKit.Unmarshal)

	/* 伪装成Chrome浏览器发起请求，主要针对: 反爬虫检测 */
	client.ImpersonateChrome()

	/*
		超时时间（发送请求的整个周期，includes connection time, any redirects, and reading the response body）
		(1) imroc/req默认: 2min
		(2) 0: no timeout
	*/
	client.SetTimeout(30 * time.Second)

	/* （imroc/req默认: 不跳过）https证书验证 */
	if insecureSkipVerify {
		client.EnableInsecureSkipVerify()
	} else {
		// 推荐正式环境使用，更安全
		client.DisableInsecureSkipVerify()
	}

	return
}
