package rKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

// NewClient
/*
@param insecureSkipVerify 	true:  跳过证书验证
							false: 不跳过证书验证（默认; 更加安全）
*/
func NewClient(insecureSkipVerify bool) (client *req.Client) {
	client = req.C()

	/* （默认启用）启用自动检测字符集并解码为utf-8 */
	client.EnableAutoDecode()

	/* json序列化和反序列化 */
	client.SetJsonMarshal(jsonKit.Marshal)
	client.SetJsonUnmarshal(jsonKit.Unmarshal)

	/* 伪装成Chrome浏览器发起请求，主要针对: 反爬虫检测 */
	client.ImpersonateChrome()

	/* 启用压缩 */
	client.EnableCompression()

	/* https证书验证 */
	if insecureSkipVerify {
		client.EnableInsecureSkipVerify()
	} else {
		// 推荐正式环境使用，更安全
		client.DisableInsecureSkipVerify()
	}

	return
}

func NewRetryClient(insecureSkipVerify bool) (client *req.Client) {
	client = NewClient(insecureSkipVerify)

	// TODO:

	return
}
