package sqlliteKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/gorm/gormKit"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := sqlserver.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
