package console

import "github.com/richelieu-yang/chimera/v3/src/log/zapKit"

var (
	Debugf func(template string, args ...interface{}) = zapKit.Debugf

	Infof func(template string, args ...interface{}) = zapKit.Infof

	Warnf func(template string, args ...interface{}) = zapKit.Warnf

	Errorf func(template string, args ...interface{}) = zapKit.Errorf

	DPanicf func(template string, args ...interface{}) = zapKit.DPanicf

	Panicf func(template string, args ...interface{}) = zapKit.Panicf

	Fatalf func(template string, args ...interface{}) = zapKit.Fatalf
)

//var (
//	innerS = zapKit.getInnerS()
//)
//
//// Debugf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
//
//	innerS.Debugf(template, args...)
//}
//
//// Infof TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
//
//	innerS.Infof(template, args...)
//}
//
//// Warnf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
//
//	innerS.Warnf(template, args...)
//}
//
//// Errorf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
// {
//	innerS.Errorf(template, args...)
//}
//
//// DPanicf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
// {
//	innerS.DPanicf(template, args...)
//}
//
//// Panicf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
// {
//	innerS.Panicf(template, args...)
//}
//
//// Fatalf TODO: 暂不将 zapKit.Debugf 赋值给函数变量，原因: GoLand不支持.
// {
//	innerS.Fatalf(template, args...)
//}
