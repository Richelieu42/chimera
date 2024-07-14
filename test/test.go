package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
	"net/http"
	"time"
)

// LbClient http客户端的负载均衡
type LbClient struct {
	client *req.Client

	urls []string
}

func (c *LbClient) Get() (*req.Response, error) {
	// Richelieu: 不能每次都从0开始，否则第一个url压力太大
	index := randomKit.Int(0, len(c.urls))
	startUrl := c.urls[index]
	console.Infof("start url: %s", startUrl)

	r := c.client.Get(startUrl)
	r.SetRetryHook(func(resp *req.Response, err error) {
		index++
		index = index % len(c.urls)
		newUrl := c.urls[index]
		console.Infof("new url: %s", newUrl)
		r.SetURL(newUrl)
	})

	resp := r.Do()
	return resp, resp.Err
}

// NewLbClient
/*
@param baseClient !!!: 本函数会修改此传参
@param commonRetryInterval 	重试周期
@param commonRetryCondition	重试条件
*/
func NewLbClient(baseClient *req.Client, urls []string, commonRetryInterval time.Duration, commonRetryCondition req.RetryConditionFunc) (*LbClient, error) {
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

	if baseClient == nil {
		baseClient = reqKit.NewClient()
	}
	baseClient.SetCommonRetryCount(len(urls) - 1).
		SetCommonRetryFixedInterval(commonRetryInterval).
		SetCommonRetryCondition(commonRetryCondition)

	return &LbClient{
		client: baseClient,
		urls:   urls,
	}, nil
}

func main() {
	urls := []string{
		"http://127.0.0.1:8000/test",
		"http://127.0.0.1:8001/test",
		"http://127.0.0.1:8002/test",
	}

	c := reqKit.NewClient(reqKit.WithDev())
	lbc, err := NewLbClient(c, urls, time.Millisecond*100, nil)
	if err != nil {
		panic(err)
	}
	resp, err := lbc.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())
}
