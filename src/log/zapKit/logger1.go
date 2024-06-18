package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

// NewLogger1
/*
@param writer	可以是: ioKit.NewRotatableWriteCloser()
@param level 	e.g. zap.InfoLevel
*/
func NewLogger1(writer io.Writer, level zapcore.Level) (*zap.Logger, error) {
	encoder := newEncoder()
	writeSyncer := zapcore.AddSync(writer)
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
	return zap.New(core, zap.AddCaller()), nil
}

func newEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//// NewLogger1
///*
//@param level e.g. zap.InfoLevel
//*/
//func NewLogger1(logPath string, maxFileSize, maxFileIndex int, compress bool, level zapcore.Level) (*zap.Logger, error) {
//	// 由于调用输出时才会尝试生成日志文件（包括创建父目录），失败则会在控制台有错误输出:
//	// e.g. 2022-12-07 11:03:01.742597 +0800 CST m=+0.000902876 write error: can't make directories for new logfile: mkdir /logs: read-only file system
//	// 因此，还不如在初始化logger时先把该做的都做了.
//	if err := fileKit.MkParentDirs(logPath); err != nil {
//		return nil, err
//	}
//
//	encoder := newEncoder()
//	writeSyncer := getLogWriter(logPath, maxFileSize, maxFileIndex, compress)
//
//	atomicLevel := zap.NewAtomicLevel()
//	atomicLevel.SetLevel(level)
//
//	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
//	return zap.Sign(core, zap.AddCaller()), nil
//}

//func getLogWriter(logPath string, maxFileSize, maxFileIndex int, compress bool) zapcore.WriteSyncer {
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   logPath,
//		MaxSize:    maxFileSize,
//		MaxBackups: maxFileIndex,
//		Compress:   compress,
//	}
//	return zapcore.AddSync(lumberJackLogger)
//}
