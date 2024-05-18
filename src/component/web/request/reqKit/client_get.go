package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
)

// SimpleGet
/*
!!!:
(1) 第2个返回值==nil的情况下，不需要手动关闭第1个返回值;
(2) 最大重试次数，参考了 eatmoreapple/openwechat 中的 client.go.
*/
func (c *Client) SimpleGet(url string, queryParams map[string][]string) (resp *req.Response, err error) {
	url, err = urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return
	}

	for i := 0; i < c.maxRetryTimes; i++ {
		resp, err = c.Client.R().Get(url)
		if err != nil {
			break
		}
	}
	return resp, err
}

func (c *Client) Get(url string, queryParams map[string][]string) (status int, data []byte, err error) {
	var resp *req.Response
	resp, err = c.SimpleGet(url, queryParams)
	if err != nil {
		return
	}
	// 不需要手动关闭 resp
	//defer resp.Body.Close()

	status = resp.StatusCode
	data = resp.Bytes()
	return
}
