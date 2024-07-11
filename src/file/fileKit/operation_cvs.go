package fileKit

import "github.com/duke-git/lancet/v2/fileutil"

// ReadCsvFile 读取csv文件内容到切片.
func ReadCsvFile(filePath string) ([][]string, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	return fileutil.ReadCsvFile(filePath)
}

// WriteCsvFile 向csv文件写入内容.
func WriteCsvFile(filePath string, records [][]string, append bool) error {
	if err := AssertNotExistOrIsFile(filePath, true); err != nil {
		return err
	}

	return fileutil.WriteCsvFile(filePath, records, append)
}
