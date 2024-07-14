package main

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
	"net/http"
	"time"
)

// LbClient 负载均衡的http客户端
type LbClient struct {
	client *req.Client

	urls []string
}

func (c *LbClient) Get() (*req.Response, error) {
	// Richelieu: 不能每次都从0开始，否则第一个url压力太大
	index := randomKit.Int(0, len(c.urls))

	r := c.client.Get(c.urls[index])
	r.SetRetryHook(func(resp *req.Response, err error) {
		index++
		index = index % len(c.urls)

		newUrl := c.urls[index]
		r.SetURL(newUrl)
	})

	resp := r.Do()
	return resp, resp.Err
}

// NewLbClient
/*
@param commonRetryInterval 	重试周期
@param commonRetryCondition	重试条件
*/
func NewLbClient(urls []string, commonRetryInterval time.Duration, commonRetryCondition req.RetryConditionFunc) (*LbClient, error) {
	if err := sliceKit.AssertNotEmpty(urls, "urls"); err != nil {
		return nil, err
	}

	if commonRetryInterval <= 0 {
		commonRetryInterval = time.Millisecond * 100
	}
	if commonRetryCondition == nil {
		commonRetryCondition = func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode != http.StatusOK
		}
	}

	client := reqKit.NewClient()
	client.SetCommonRetryCount(len(urls) - 1).
		SetCommonRetryFixedInterval(commonRetryInterval).
		SetCommonRetryCondition(commonRetryCondition)

	return &LbClient{
		client: client,
		urls:   urls,
	}, nil
}

func main() {
	urls := []string{
		"http://127.0.0.1:8000/test",
		"http://127.0.0.1:8001/test",
		"http://127.0.0.1:8002/test",
	}

	c := reqKit.NewClient()
	c.SetCommonRetryCount(len(urls) - 1).
		SetCommonRetryFixedInterval(time.Millisecond * 100).
		SetCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode != http.StatusOK
		})

	i := randomKit.Int(0, len(urls))
	url := urls[i]

	r := c.Get(url)
	r.SetRetryHook(func(resp *req.Response, err error) {
		r.SetURL()
	})

}
