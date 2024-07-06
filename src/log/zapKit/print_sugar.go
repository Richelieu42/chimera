package zapKit

// Debugf 格式化输出的信息日志，类似于fmt.Printf，可以使用格式化字符串.
func Debugf(template string, args ...interface{}) {
	innerS.Debugf(template, args...)
}

// Debugw 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Debugw(msg string, keysAndValues ...interface{}) {
	innerS.Debugw(msg, keysAndValues...)
}

// Debugln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Debugln(args ...interface{}) {
	innerS.Debugln(args...)
}

func Infof(template string, args ...interface{}) {
	innerS.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	innerS.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	innerS.Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	innerS.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	innerS.Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	innerS.Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	innerS.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	innerS.Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	innerS.Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	innerS.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	innerS.DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	innerS.DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	innerS.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	innerS.Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	innerS.Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	innerS.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	innerS.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	innerS.Fatalln(args...)
}
