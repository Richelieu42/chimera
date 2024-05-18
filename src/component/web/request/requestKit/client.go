package requestKit

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"time"
)

var defClient = NewClient(false, time.Second*15)

// NewClient
/*
重用 Client
	https://req.cool/zh/docs/tutorial/best-practices/#%e9%87%8d%e7%94%a8-client

PS:
(1) 不要每次发请求都创建 Client，造成不必要的开销，通常可以复用同一 Client 发所有请求.
*/
func NewClient(debug bool, timeout time.Duration) (client *req.Client) {
	client = req.C()

	// timeout
	client.SetTimeout(timeout)

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

	/*
		debug
		在生产环境动态开启Debug
			https://req.cool/zh/docs/examples/enable-debug-dynamically-in-production/
	*/
	if debug {
		client.EnableDumpAll()
		client.EnableDebugLog()
	} else {
		client.DisableDumpAll()
		client.DisableDebugLog()
	}

	/*
		DevMode(启用所有调试的特性(Dump, DebugLog, Trace))
		https://req.cool/zh/docs/tutorial/debugging/#devmode
	*/
	//if debug {
	//	client.DevMode()
	//}

	return
}
