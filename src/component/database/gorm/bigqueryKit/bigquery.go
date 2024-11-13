package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/gorm/gormKit"
	"gorm.io/driver/bigquery"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."bigquery://projectid/location/dataset"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := bigquery.Open(dsn)
	return gormKit.NewDB(dialector, opts...)
}
