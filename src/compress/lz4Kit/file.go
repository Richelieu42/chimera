package lz4Kit

import (
	"github.com/pierrec/lz4/v4"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"os"
)

// CompressFile
/*
@param dest (1) 建议文件后缀为 ".lz4"
			(2) 如果已经存在 && 是个文件，会覆盖.
*/
func CompressFile(src, dest string) error {
	if err := fileKit.AssertExistAndIsFile(src); err != nil {
		return err
	}
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}

	// 读取原始文件内容
	inputData, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// 使用lz4.NewWriter创建压缩器，写入到一个新的文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	lz4Writer := lz4.NewWriter(destFile)
	defer lz4Writer.Close()

	// 写入压缩数据
	_, err = lz4Writer.Write(inputData)
	return err
}
