package console

import "github.com/richelieu-yang/chimera/v3/src/log/zapKit"

var (
	innerS = zapKit.GetInnerSugaredLogger()
)

// Debugf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Debugf(template string, args ...interface{}) {
	zapKit.Debugf(template, args...)
	innerS.Debugf(template, args...)
}

// Infof TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Infof(template string, args ...interface{}) {
	innerS.Infof(template, args...)
}

// Warnf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Warnf(template string, args ...interface{}) {
	innerS.Warnf(template, args...)
}

// Errorf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Errorf(template string, args ...interface{}) {
	innerS.Errorf(template, args...)
}

// DPanicf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func DPanicf(template string, args ...interface{}) {
	innerS.DPanicf(template, args...)
}

// Panicf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Panicf(template string, args ...interface{}) {
	innerS.Panicf(template, args...)
}

// Fatalf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
func Fatalf(template string, args ...interface{}) {
	innerS.Fatalf(template, args...)
}
