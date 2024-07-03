package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	cancel()

	GetNetworkTime(ctx)

	select {}
}

var sources = []string{
	"https://www.google.com/",
	"https://www.tencent.com/",
	"https://github.com/",
	"https://www.bilibili.com/",
	"https://www.baidu.com/",
	"https://cn.bing.com/",
	"http://www.ntsc.ac.cn/",
	"https://www.taobao.com/",
	"https://www.360.cn/",
	"https://www.kingsoft.com/",
	"https://www.yozosoft.com/",
}

func GetNetworkTime(ctx context.Context) (t time.Time, source string, err error) {
	type bean struct {
		source string
		time   time.Time
	}

	ch := make(chan *bean, len(sources))
	for _, source := range sources {
		go func(url string) {
			t, err = GetNetworkTimeByUrl(ctx, url)
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
	case <-ctx.Done():
		err = ctx.Err()
		return
	}
}

func GetNetworkTimeByUrl(ctx context.Context, url string) (t time.Time, err error) {
	resp := reqKit.SimpleGet(ctx, url)
	if resp.Err != nil {
		err = resp.Err
		return
	}

	str := resp.GetHeader(httpKit.HeaderDate)
	if str == "" {
		err = errorKit.Newf("empty header: %s", httpKit.HeaderDate)
		return
	}
	t, err = timeKit.Parse(timeKit.FormatRFC1123, str)
	if err != nil {
		err = errorKit.Wrapf(err, "fail to prase with timeKit.FormatRFC1123")
		return
	}
	return
}
