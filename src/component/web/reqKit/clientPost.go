package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
)

// SimplePost
/*
参考:
基础POST请求 https://req.cool/zh/docs/prologue/quickstart/#%e5%9f%ba%e7%a1%80-post-%e8%af%b7%e6%b1%82
*/
func (c *Client) SimplePost(url string, queryParams map[string][]string, body interface{}) (*req.Response, error) {
	url, err := urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return nil, err
	}

	return c.Client.R().SetBody(body).Post(url)
}

func (c *Client) SimplePostWithJson(url string, queryParams map[string][]string, jsonBytes []byte) (*req.Response, error) {
	url, err := urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return nil, err
	}

	return c.Client.R().SetBodyJsonBytes(jsonBytes).Post(url)
}

func (c *Client) SimplePostWithString(url string, queryParams map[string][]string, str string) (*req.Response, error) {
	url, err := urlKit.PolyfillUrl(url, queryParams)
	if err != nil {
		return nil, err
	}

	return c.Client.R().SetBodyString(str).Post(url)
}
