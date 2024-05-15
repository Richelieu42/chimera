package ioKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	// 创建一个管道
	pr, pw := Pipe()

	// 新建一个协程写入数据
	go func() {
		for i := 0; i < 3; i++ {
			_, err := pw.Write([]byte(strconv.Itoa(i) + "\n"))
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}

		logrus.Info("sleep starts...")
		time.Sleep(time.Second * 3)
		logrus.Info("sleep ends...")
		_ = pw.Close()
	}()

	// 在主协程中读取管道数据（会在此处阻塞至pw关闭）
	_, err := io.Copy(os.Stdout, pr) // 从PipeReader复制数据到标准输出
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Info("---")
}
