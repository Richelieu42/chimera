package mysqlKit

import (
	"fmt"
	"time"
)

type (
	Config struct {
	}

	PoolConfig struct {
		MaxIdleConns    int           `json:"maxIdleConns"`
		MaxOpenConns    int           `json:"maxOpenConns"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime"`
	}

	DsnConfig struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host"`
		DBName string `json:"dbName"`
	}
)

func (c *DsnConfig) ToDSN() string {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.DBName,
	)
}
