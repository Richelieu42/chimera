package mysqlKit

import (
	"fmt"
	"testing"
)

func TestNewGormDB(t *testing.T) {
	user := "yjs"
	password := "Test123/@"
	dsn := fmt.Sprintf("%s:%s@tcp(101.32.170.155:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local", user, password)
	fmt.Printf("dsn: %s\n", dsn)

	db, err := NewGormDB(dsn)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	{

	}
}
