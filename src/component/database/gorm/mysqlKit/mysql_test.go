package mysqlKit

import (
	"testing"
)

func TestNewGormDB(t *testing.T) {
	db, err := NewGormDB("yjs:~Test123@tcp(101.32.170.155:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local")
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
