package logrusKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"testing"
)

func TestNewLogger(t *testing.T) {
	///* 输出到控制台 */
	//logger := NewLogger()
	//logger.Info("to console")
	//
	///* 输出到文件（可rotate） */
	//path := "test.log"
	//writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath(path))
	//assert.Nil(t, err)
	//logger = NewLogger(WithWriter(writer))
	//logger.Info("to rotatable file")
	//
	//err = assertKit.ExistAndIsFile(path)
	//assert.Nil(t, err)

	writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath(path))
	/*
		process err
	*/
	logger := NewLogger(WithWriter(writer))
}
