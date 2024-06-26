package zapKit

// Debugf 格式化输出的信息日志，类似于fmt.Printf，可以使用格式化字符串.
func Debugf(template string, args ...interface{}) {
	sl.Debugf(template, args...)
}

// Debugw 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Debugw(msg string, keysAndValues ...interface{}) {
	sl.Debugw(msg, keysAndValues...)
}

// Debugln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Debugln(args ...interface{}) {
	sl.Debugln(args...)
}

func Infof(template string, args ...interface{}) {
	sl.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sl.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	sl.Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	sl.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sl.Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	sl.Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	sl.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sl.Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	sl.Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	sl.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	sl.DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	sl.DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	sl.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	sl.Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	sl.Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	sl.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sl.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	sl.Fatalln(args...)
}
