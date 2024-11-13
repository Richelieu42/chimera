package gormKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"gorm.io/gorm"
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
	return db, nil
}
