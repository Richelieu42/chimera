package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/gorm/gormKit"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := clickhouse.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
