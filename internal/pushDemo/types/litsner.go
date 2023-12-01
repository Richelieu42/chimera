package types

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Listener struct {
	pushKit.Listener
}

func (l *Listener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("failureInfo", failureInfo).Error("OnFailure")
}

func (l *Listener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {
	bsid := httpKit.ObtainGetParam(r, "bsid")
	user := httpKit.ObtainGetParam(r, "user")
	group := httpKit.ObtainGetParam(r, "group")

	text := fmt.Sprintf("Hello, your id is [%s].", channel.GetId())
	if err := channel.Push([]byte(text)); err != nil {
		logrus.Error(err)
	}

	channel.BindBsid(bsid)
	channel.BindUser(user)
	channel.BindGroup(group)

	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
		"bsid":     channel.GetBsid(),
		"user":     channel.GetUser(),
		"group":    channel.GetGroup(),
	}).Info("OnHandshake")

	//go func() {
	//	time.Sleep(time.Second * 3)
	//	_ = channel.Close("测试")
	//}()
}

func (l *Listener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
		"bsid":     channel.GetBsid(),
		"user":     channel.GetUser(),
		"group":    channel.GetGroup(),

		"messageType": messageType,
		"text":        string(data),
	}).Info("OnMessage")
}

func (l *Listener) BeforeClosedByBackend(channel pushKit.Channel, closeInfo string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
		"bsid":     channel.GetBsid(),
		"user":     channel.GetUser(),
		"group":    channel.GetGroup(),

		"closeInfo": closeInfo,
	}).Info("BeforeClosedByBackend")
}

func (l *Listener) OnClose(channel pushKit.Channel, closeInfo string, bsid, user, group string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
		"bsid":     bsid,
		"user":     user,
		"group":    group,

		"closeInfo": closeInfo,
	}).Info("OnClose")
}
