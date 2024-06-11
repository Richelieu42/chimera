package timeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/req111Kit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"time"
)

// networkTimeSources 网络时间的来源s
var networkTimeSources = []string{
	"https://github.com/",
	"https://www.bilibili.com/",
	"https://www.baidu.com/",
	"https://cn.bing.com/",
	"http://www.ntsc.ac.cn/",
	"https://www.taobao.com/",
	"https://www.360.cn/",
	"https://www.google.com/",
	"https://www.kingsoft.com/",
	"https://www.yozosoft.com/",
}

// GetNetworkTime 获取网络时间.
/*
PS: 获取不到的话，返回机器时间.
*/
func GetNetworkTime() (time.Time, string, error) {
	type bean struct {
		source string
		time   time.Time
	}

	// 超时时间设置的短一点，以防内网环境启动服务耗时太长
	var timeout = time.Second * 3
	var ch = make(chan *bean, len(networkTimeSources))

	// 共用一个client
	client := req111Kit.NewClient(3)
	client.SetTimeout(timeout)
	client.SetTLSHandshakeTimeout(timeout)

	// 起多个goroutine同时获取网络时间，只要有一个成功获取到，此方法就返回值
	for _, source := range networkTimeSources {
		go func(url string) {
			t, err := getNetworkTimeBySource(client, url)
			if err != nil {
				return
			}
			ch <- &bean{
				source: url,
				time:   t,
			}
		}(source)
	}
	select {
	case b := <-ch:
		return b.time, b.source, nil
	case <-time.After(timeout):
		return time.Time{}, "", errorKit.Newf("timeout(%s)", timeout)
	}
}

func getNetworkTimeBySource(client *req111Kit.Client, url string) (time.Time, error) {
	resp, err := client.SimpleGet(url, nil)
	if err != nil {
		return time.Time{}, err
	}

	// e.g."Fri, 18 Aug 2023 07:15:26 GMT"
	dateStr := resp.Header.Get("Date")
	if err := strKit.AssertNotEmpty(dateStr, "dateStr"); err != nil {
		return time.Time{}, err
	}

	t, err := Parse(string(FormatRFC1123), dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return ConvertLocation(t, time.Local), nil
}
