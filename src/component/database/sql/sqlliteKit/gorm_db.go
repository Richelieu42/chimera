package sqlliteKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn	DSN（Data Source Name），它是用于指定数据库文件路径以及可能的其他选项的字符串.
			e.g. 	内存数据库
				":memory:"
				// "file::memory:"	表示SQLite应创建一个仅存在于内存中的临时数据库。这意味着当SQLite连接关闭时，这个数据库的内容将不会被保存到磁盘上，它只在当前进程的上下文中存在。
				// "?cache=shared"	是SQLite的一个特定选项，它指定了内存数据库应该启用共享缓存模式。在默认情况下，每个SQLite连接都会有自己的私有缓存，而设置为shared后，多个连接可以共享同一份缓存数据。这意味着对于同一个数据库的不同连接，读取的数据会保持一致，并且能减少内存使用量。这对于多线程或多进程环境下使用内存数据库非常有用，各个进程可以通过SQLite API安全地共享和访问相同的内存数据库内容。
				"file::memory:?cache=shared"
			e.g.1	本地SQLite数据库文件
				"./my_database.db"
				"/path/to/your/database.sqlite"
				"file:/path/to/your/database.sqlite?_loc=Local"
*/
func NewGormDB(dsn string, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	if err := strKit.AssertNotEmpty(dsn, "dsn"); err != nil {
		return nil, err
	}
	var path string
	if strKit.StartWith(dsn, "file:") {
		if strKit.StartWith(dsn, "file::memory:") {
			// (1) 内存数据库
			path = ""
		} else {
			// (2) 本地SQLite数据库文件
			path = strKit.RemovePrefixIfExists(dsn, "file:")
		}
	} else if strKit.StartWith(dsn, ":memory:") {
		// (3) 内存数据库
		path = ""
	} else {
		// (4) 本地SQLite数据库文件
		path = dsn
	}

	if err := fileKit.AssertNotExistOrIsFile(path); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(path); err != nil {
		return nil, err
	}

	dialector := sqlite.Open(dsn)
	return gormKit.Open(dialector, poolConfig, opts...)
}
