package logKit

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

// 输出的日志不带日期
func TestNewLogger(t *testing.T) {
	//flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	flag := log.Ltime | log.Lmicroseconds | log.Lshortfile
	logger := NewLogger(os.Stdout, "", flag)
	logger.Println("hello")
	logger.Println("world")

	/*
		13:58:00.063515 logger_test.go:13: hello
		13:58:00.063693 logger_test.go:14: world
	*/
}

// 输出到: 指定io.Writer实例
func TestNewLogger1(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := NewLogger(buf, "", log.Ltime|log.Lmicroseconds|log.Lshortfile)

	logger.Println("hello")
	logger.Println("world")
	fmt.Println("---")
	fmt.Println(buf.String())
}

// 输出到: 控制台
func TestNewStdoutLogger(t *testing.T) {
	logger0 := NewStdoutLogger("")
	logger0.Println(0)
	logger0.Println(1)

	logger1 := NewStdoutLogger("[TEST] ")
	logger1.Println(0)
	logger1.Println(1)
	/*
		output:
		2024/06/08 13:48:38.940927 logger_test.go:9: 0
		2024/06/08 13:48:38.941157 logger_test.go:10: 1
		2024/06/08 13:48:38.941162 logger_test.go:13: [TEST] 0
		2024/06/08 13:48:38.941166 logger_test.go:14: [TEST] 1
	*/
}

// 输出到: 文件
func TestNewFileLogger(t *testing.T) {
	logger, err := NewFileLogger("_test.log", "[TEST] ", 0644)
	//logger, err := NewFileLogger("_test.log", "", 0644)
	if err != nil {
		panic(err)
	}

	logger.Println("0")
	logger.Println("1")
	logger.Println("2")
	logger.Printf("---")
	logger.Println("4")
}
