package redisKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/config/confKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"go.uber.org/zap"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 测试: Redis客户端通用订阅与发布.
func TestClient_Publish(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		console.Infof("working dir: %s", wd)
	}

	type config struct {
		Redis Config `json:"redis"`
	}
	c := &config{}
	path := "_chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(&c.Redis)
	client, err := GetClient()
	if err != nil {
		panic(err)
	}
	client = client

	{
		var wg sync.WaitGroup
		channel := "测试"
		pubSub := client.Subscribe(context.TODO(), channel)
		defer pubSub.Close()

		wg.Add(1)
		go func() {
			defer wg.Done()

			ch := pubSub.Channel()
			for msg := range ch {
				console.Info("Receive a message.", zap.String("channel", msg.Channel), zap.String("payLoad", msg.Payload))
			}
			console.Info("Receive ends.......................")
		}()
		// 3s后取消订阅
		go func() {
			time.Sleep(time.Second * 3)

			console.Info("Unsubscribe starts.")
			if err := pubSub.Unsubscribe(context.TODO(), channel); err != nil {
				panic(err)
			}
			console.Info("Unsubscribe ends.")

			console.Info("Close starts.")
			if err := pubSub.Close(); err != nil {
				panic(err)
			}
			console.Info("Close ends.")
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := 0; i < 4; i++ {
				time.Sleep(time.Second)
				_, err := client.Publish(context.TODO(), channel, strconv.Itoa(i))
				if err != nil {
					panic(err)
				}
			}
		}()

		wg.Wait()
		console.Info("===")
	}
}

// 测试: Redis客户端监听某一db中key的超时.
/*
!!!: 需要先配置Redis.
*/
func TestSubscribeExpired(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		console.Infof("working dir: %s", wd)
	}

	type config struct {
		Redis Config `json:"redis"`
	}
	c := &config{}
	path := "_chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(&c.Redis)
	client, err := GetClient()
	if err != nil {
		panic(err)
	}
	client = client

	{
		flag := atomicKit.NewBool(false)
		id := idKit.NewULID()

		/* pubSub使用方法1 */
		//go func() {
		//	pubSub := client.Subscribe(context.TODO(), "__keyevent@0__:expired")
		//	defer pubSub.Close()
		//
		//	for {
		//		msg, err := pubSub.ReceiveMessage(context.TODO())
		//		if err != nil {
		//			panic(err)
		//		}
		//
		//		// msg.Payload: 过期的键（key）
		//		console.Info("Receive a message.", zap.String("channel", msg.Channel), zap.String("payLoad", msg.Payload))
		//	}
		//}()

		/* pubSub使用方法2 */
		go func() {
			pubSub := client.Subscribe(context.TODO(), "__keyevent@0__:expired")
			defer pubSub.Close()

			ch := pubSub.Channel()
			for msg := range ch {
				// msg.Payload: 过期的键（key）
				console.Info("Receive a message.", zap.String("channel", msg.Channel), zap.String("payLoad", msg.Payload))
				if msg.Payload == id {
					flag.CompareAndSwap(false, true)
				}
			}
		}()

		go func() {
			_, err := client.Set(context.TODO(), id, "test", time.Second*3)
			if err != nil {
				panic(err)
			}
		}()

		time.Sleep(time.Second * 5)
		if !flag.Load() {
			panic("value of param flag is [false]")
		}
	}
}
