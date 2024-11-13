package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/gorm/gormKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB
/*
参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL

@param dsn "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}

// NewGormDBOptionally 自定义驱动.
/*
参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL
*/
func NewGormDBOptionally(mysqlConfig *mysql.Config, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(mysqlConfig, "mysqlConfig"); err != nil {
		return nil, err
	}

	dialector := mysql.New(*mysqlConfig)
	return gormKit.NewDB(dialector, opts...)
}
