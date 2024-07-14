package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
)

// LoadBalancerClient http客户端的负载均衡
type LoadBalancerClient struct {
	client *req.Client

	urls []string
}

func (lbc *LoadBalancerClient) Get() (*req.Response, error) {
	/* Richelieu: 不能每次都从0开始，否则第一个url压力太大 */
	index := randomKit.Int(0, len(lbc.urls))
	startUrl := lbc.urls[index]
	console.Infof("start url: %s", startUrl)

	r := lbc.client.Get(startUrl)
	r.SetRetryHook(func(resp *req.Response, err error) {
		index++
		index = index % len(lbc.urls)
		retryUrl := lbc.urls[index]
		console.Infof("retry url: %s", retryUrl)

		r.SetURL(retryUrl)
	})

	resp := r.Do()
	return resp, resp.Err
}
