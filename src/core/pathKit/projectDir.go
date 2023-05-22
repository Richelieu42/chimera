package pathKit

import (
	"github.com/richelieu42/chimera/v2/src/assert/fileAssert"
	"github.com/richelieu42/chimera/v2/src/core/file/fileKit"
)

// projectDir 项目目录（绝对路径）
var projectDir string

func GetProjectDir() string {
	return projectDir
}

func SetProjectDir(dir string) error {
	if err := fileKit.MkDirs(dir); err != nil {
		return err
	}
	if err := fileAssert.AssertExistAndIsDir(dir); err != nil {
		return err
	}
	projectDir = dir
	return nil
}
