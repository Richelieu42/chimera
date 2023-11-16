package main

import (
	"crypto/tls"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	httpClient := bot.Caller.Client.HTTPClient()
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {

		msg.IsSendBySelf()

		if msg.IsText() && msg.Content == "ping" {
			_, err := msg.ReplyText("pong")
			if err != nil {
				logrus.WithError(err).Error("Fail to pong")
			} else {
				logrus.Info("Manager to pong")
			}
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(friends)

	results := friends.SearchByRemarkName(1, "狗蛋妈")
	if results.Count() == 0 {
		logrus.Fatal("results.Count() == 0")
	}
	wife := results.First()

	go func() {
		send(wife)
	}()

	bot.Block()
}

// 发送问候
func send(f *openwechat.Friend) {
	cron := cronKit.NewCron()

	_, err := cron.AddFunc("0 30 6 * * *", func() {
		text := "宝贝，早安！"
		_, err := f.SendText(text)
		if err != nil {
			logrus.WithError(err).WithField("text", text).Error("Fail to greet.")
		} else {
			logrus.WithField("text", text).Error("Manager to greet.")
		}
	})
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 11 * * *", func() {
		text := "宝贝，午安！"
		_, err := f.SendText(text)
		if err != nil {
			logrus.WithError(err).WithField("text", text).Error("Fail to greet.")
		} else {
			logrus.WithField("text", text).Error("Manager to greet.")
		}
	})
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 21 * * *", func() {
		text := "宝贝，晚安！"
		_, err := f.SendText(text)
		if err != nil {
			logrus.WithError(err).WithField("text", text).Error("Fail to greet.")
		} else {
			logrus.WithField("text", text).Error("Manager to greet.")
		}
	})
	if err != nil {
		logrus.Fatal(err)
	}

	cron.Start()
}
