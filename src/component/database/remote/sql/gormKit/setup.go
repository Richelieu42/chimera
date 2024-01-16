package gormKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func MustSetUp(config *Config, logConfig *LogConfig) {
	if err := SetUp(config, logConfig); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
@param output 客户端的日志输出（nil: 输出到控制台）
*/
func SetUp(config *Config, logConfig *LogConfig) (err error) {
	if err = interfaceKit.AssertNotNil(config, "config"); err != nil {
		return
	}

	logger := NewLogger(logConfig)
	tmp, err := gorm.Open(mysql.Open(config.GetDsnString()), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return
	}

	/* verify by ping */
	sqlDB, err := tmp.DB()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = sqlDB.Close()
			return
		}
		// 成功，给 db 赋值
		db = tmp
	}()
	if err = sqlDB.Ping(); err != nil {
		err = errorKit.Wrap(err, "Fail to ping")
		return
	}

	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, NotSetupError
	}
	return db, nil
}
