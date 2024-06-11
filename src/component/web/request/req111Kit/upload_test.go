package req111Kit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"net/http"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	/* server */
	go func() {
		engine := gin.Default()

		engine.POST("/upload", func(ctx *gin.Context) {
			name := ctx.PostForm("name")
			if name == "" {
				ctx.String(http.StatusOK, "name is empty")
				return
			}

			fh, err := ctx.FormFile("file")
			if err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}

			if err := ctx.SaveUploadedFile(fh, name); err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.String(http.StatusOK, "ok")
		})

		if err := engine.Run(":12345"); err != nil {
			panic(err)
		}
	}()

	// 等一会，让server先启动
	time.Sleep(time.Second * 3)

	/* client */
	{
		url := "http://127.0.0.1:12345/upload"

		client := GetDefaultClient()
		resp, err := client.R().
			/* (1) 上传文件 */
			//SetFile("file", "/Users/richelieu/Downloads/test.txt").
			/* (2) 上传字节流 */
			SetFileBytes("file", "test.txt", []byte("test content")).
			SetFormData(map[string]string{
				"name": "_" + idKit.NewULID() + ".txt",
			}).
			Post(url)
		if err != nil {
			panic(err)
		}
		respStr := resp.String()
		fmt.Println("respStr:", respStr)
		if respStr != "ok" {
			panic("not equal")
		}
	}
}
