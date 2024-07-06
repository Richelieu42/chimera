package zapKit

// Debugf 格式化输出的信息日志，类似于fmt.Printf，可以使用格式化字符串.
func Debugf(template string, args ...interface{}) {
	s.Debugf(template, args...)
}

// Debugw 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Debugw(msg string, keysAndValues ...interface{}) {
	s.Debugw(msg, keysAndValues...)
}

// Debugln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Debugln(args ...interface{}) {
	s.Debugln(args...)
}

func Infof(template string, args ...interface{}) {
	s.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	s.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	s.Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	s.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	s.Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	s.Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	s.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	s.Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	s.Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	s.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	s.DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	s.DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	s.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	s.Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	s.Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	s.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	s.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	s.Fatalln(args...)
}
