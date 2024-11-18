package gormKit

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	defaultLoggerOnce sync.Once
	defaultLogger     logger.Interface
)

// GetDefaultLogger 默认的Logger.
/*
PS: 参考了 logger.Default.
*/
func GetDefaultLogger() logger.Interface {
	defaultLoggerOnce.Do(func() {
		writer := log.New(os.Stdout, "\r\n", log.LstdFlags|log.Lmicroseconds)
		config := logger.Config{
			SlowThreshold:             time.Millisecond * 200, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: false,                  // true: Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                  // true: Don't include params in the SQL log
			Colorful:                  true,                   // false: Disable color
		}
		defaultLogger = logger.New(writer, config)
	})

	return defaultLogger
}
