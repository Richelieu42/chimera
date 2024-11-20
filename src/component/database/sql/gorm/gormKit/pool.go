package gormKit

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// ConfigurePool 连接池
func ConfigurePool(db *gorm.DB, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	ConfigurePoolWithSqlDB(sqlDB, maxIdleConns, maxOpenConns, connMaxLifetime)
	return nil
}

// ConfigurePoolWithSqlDB 连接池
func ConfigurePoolWithSqlDB(sqlDB *sql.DB, maxIdleConns, maxOpenConns int, connMaxLifetime time.Duration) {
	/*
		SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		== 0: 	defaultMaxIdleConns（目前是2）
		< 0:	0
	*/
	sqlDB.SetMaxIdleConns(maxIdleConns)
	/*
		SetMaxOpenConns 设置打开数据库连接的最大数量。
		<= 0: unlimited
	*/
	sqlDB.SetMaxOpenConns(maxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(connMaxLifetime)
}
