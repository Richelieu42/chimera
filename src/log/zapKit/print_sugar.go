package zapKit

func Debugf(template string, args ...interface{}) {
	getInnerS().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	getInnerS().Debugw(msg, keysAndValues...)
}

func Debugln(args ...interface{}) {
	getInnerS().Debugln(args...)
}

// Infof 格式化输出的信息日志，类似于 fmt.Printf ，可以使用格式化字符串.
func Infof(template string, args ...interface{}) {
	getInnerS().Infof(template, args...)
}

// Infow 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Infow(msg string, keysAndValues ...interface{}) {
	getInnerS().Infow(msg, keysAndValues...)
}

// Infoln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Infoln(args ...interface{}) {
	getInnerS().Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	getInnerS().Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	getInnerS().Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	getInnerS().Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	getInnerS().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	getInnerS().Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	getInnerS().Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	getInnerS().DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	getInnerS().DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	getInnerS().DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	getInnerS().Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	getInnerS().Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	getInnerS().Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	getInnerS().Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	getInnerS().Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	getInnerS().Fatalln(args...)
}
