package aKit

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// NewLogger
/*
参考: logger.Default
*/
func NewLogger(logConfig *LogConfig) logger.Interface {
	if logConfig == nil {
		logConfig = &LogConfig{
			Output:                    os.Stdout,
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		}
	}
	if logConfig.Output == nil {
		logConfig.Output = os.Stdout
	}

	prefix := "\r\n"
	writer := log.New(logConfig.Output, prefix, log.LstdFlags|log.Lmicroseconds)
	return logger.New(writer, logger.Config{
		// 慢SQL阈值
		SlowThreshold: logConfig.SlowThreshold,
		// 日志级别
		LogLevel: logConfig.LogLevel,
		// 彩色打印？
		Colorful: logConfig.Colorful,
		// 忽略 logger.ErrRecordNotFound（记录未找到错误） ？
		IgnoreRecordNotFoundError: logConfig.IgnoreRecordNotFoundError,
	})
}

//func NewLogger1(writer logger.Writer, config logger.Config) logger.Interface {
//	return logger.Sign(writer, config)
//}
