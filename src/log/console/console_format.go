package console

import "github.com/richelieu-yang/chimera/v3/src/log/zapKit"

var (
	innerS = zapKit.GetInnerSugaredLogger()
)

func Debugf(template string, args ...interface{}) {
	innerS.Debugf(template, args...)
}

// Infof 格式化输出的信息日志，类似于 fmt.Printf ，可以使用格式化字符串.
func Infof(template string, args ...interface{}) {
	innerS.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	innerS.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	innerS.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	innerS.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	innerS.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	innerS.Fatalf(template, args...)
}
