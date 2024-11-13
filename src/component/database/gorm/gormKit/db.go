package gormKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"gorm.io/gorm"
	"time"
)

// NewDB
/*
@param dialector 方言（针对不同的数据库; 不能为nil!）
*/
func NewDB(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(dialector, "dialector"); err != nil {
		return nil, err
	}
	if len(opts) == 0 {
		opts = []gorm.Option{&gorm.Config{}}
	}

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	/* ping */
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	/* 连接池（pool） */
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(256)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(4096)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
