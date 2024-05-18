package reqKit

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"time"
)

type Client struct {
	*req.Client

	// maxRetryTimes 最大重试次数
	maxRetryTimes int
}

var defaultClient = NewClient(3)

// GetDefaultClient
/*
重用Client https://req.cool/zh/docs/tutorial/best-practices/#%e9%87%8d%e7%94%a8-client
	不要每次发请求都创建 Client，造成不必要的开销，通常可以复用 同一Client 发所有请求.

!!!:
(1) 不修改返回值的话，可以调用此方法；否则调用 NewClient;
(2) 最大重试次数 == 3.
*/
func GetDefaultClient() *Client {
	return defaultClient
}

// NewClient
/*
参考:
使用req封装SDK https://req.cool/zh/docs/prologue/quickstart/#%e4%bd%bf%e7%94%a8-req-%e5%b0%81%e8%a3%85-sdk
*/
func NewClient(maxRetryTimes int) *Client {
	if maxRetryTimes <= 0 {
		maxRetryTimes = 3
	}

	client := req.C()

	// timeout
	client.SetTimeout(time.Second * 10)

	// 自动探测字符集并解码到 utf-8（默认就是启用）
	client.EnableAutoDecode()

	// 不验证非法的证书（默认验证）
	client.EnableInsecureSkipVerify()

	// 自定义 Marshal 和 Unmarshal
	api := jsonKit.GetDefaultApi()
	client.SetJsonMarshal(api.Marshal).
		SetJsonUnmarshal(api.Unmarshal)

	client.EnableDumpEachRequest().
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
				return nil
			}
			if !resp.IsSuccessState() {
				// Neither a success response nor a error response, record details to help troubleshooting
				resp.Err = fmt.Errorf("bad status: %s\nraw content:\n%s", resp.Status, resp.Dump())
				return nil
			}
			return nil
		})

	return &Client{
		Client:        client,
		maxRetryTimes: maxRetryTimes,
	}
}
