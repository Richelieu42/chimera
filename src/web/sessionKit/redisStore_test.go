package sessionKit

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/atomicKit"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"testing"
)

var i = atomicKit.NewInt()

// TestRedisStore
/*
准备工作: 启动Redis.
访问地址: http://localhost/test
*/
func TestRedisStore(t *testing.T) {
	// cookie的name
	cookieName := "session-id"
	// Redis配置（单节点）
	redisOptions := &redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	}
	client := redis.NewClient(redisOptions)

	/* ------------------------------- store starts ------------------------------- */
	// Redis中的key的前缀（value为 string 类型）
	redisKeyPrefix := "session:"
	opts := sessions.Options{
		HttpOnly: false,
		MaxAge:   0, // 只有 >= 0 的情况下，才会将数据写到Redis中
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
	}
	keyGen := func() (string, error) {
		return idKit.NewULID(), nil
	}
	store, err := NewRedisStore(context.TODO(), client, redisKeyPrefix, keyGen, opts, timeKit.Minute*10)
	if err != nil {
		panic(err)
	}
	//defer store.Close()
	/* ------------------------------- store ends ------------------------------- */

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		/* (1) 获取session */
		session, err := store.Get(ctx.Request, cookieName)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		if session.IsNew {
			session.Values["i"] = i.Add(1)
		}

		/* (2) 保存session数据，本质上是将内存中的数据持久化到存储介质中（序列化并写到Redis中；会重置key的TTL） */
		if err := session.Save(ctx.Request, ctx.Writer); err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		logrus.WithFields(logrus.Fields{
			"IsNew":  session.IsNew,
			"ID":     session.ID,
			"Values": session.Values,
		}).Info("session")

		ctx.String(http.StatusOK, strconv.Itoa(intKit.ToInt(session.Values["i"])))
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
