package rocksdbKit

import (
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/sirupsen/logrus"
	"testing"
)

// TestOpenDB 测试打开数据库
/*
PS: 想直接在IDE中Run或Debug，建议配置Go Modules: CGO_CFLAGS=-I/opt/homebrew/Cellar/rocksdb/8.9.1/include;CGO_LDFLAGS=-L/opt/homebrew/Cellar/rocksdb/8.9.1 -lrocksdb -lsnappy
*/
func TestOpenDB(t *testing.T) {
	dirPath := "_testDir"
	id := idKit.NewULID()
	keyStr := []byte("key")
	valueStr := []byte(id)
	logrus.Infof("id: [%s]", id)

	db, err := OpenDB(dirPath, NewDefaultDBOptions())
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	writeOpts := NewDefaultWriteOptions()
	err = db.Put(writeOpts, keyStr, valueStr)
	if err != nil {
		logrus.Fatal(err)
	}

	readOpts := NewDefaultReadOptions()
	value, err := db.Get(readOpts, keyStr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer value.Free()
	logrus.Infof("Read value from database: [%s]", string(value.Data()))
}
