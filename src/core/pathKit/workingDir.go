package pathKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
	"os"
)

// GetWorkingDir 获取 当前工作目录的绝对路径
func GetWorkingDir() string {
	//return gfile.Pwd()
	path, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Fatal("fail to get working directory")
	}
	return path
}

// ChangeWorkingDir 设置 当前工作目录的绝对路径
func ChangeWorkingDir(dir string) error {
	if err := fileKit.AssertNotExistOrIsDir(dir, true); err != nil {
		return err
	}

	return gfile.Chdir(dir)
}

// ReviseWorkingDirInTestMode 适用于: 测试模式(_test.go)
/*
@return 修改后的WorkingDir + error
*/
func ReviseWorkingDirInTestMode(projectName string) (string, error) {
	if strKit.IsEmpty(projectName) {
		projectName = consts.ProjectName
	}

	wd := GetWorkingDir()
	index := strKit.Index(wd, projectName)
	if index == -1 {
		return "", errorKit.New("invalid projectName(%s)", projectName)
	}
	wd1 := strKit.SubBefore(wd, index+len(projectName))
	if err := ChangeWorkingDir(wd1); err != nil {
		return "", err
	}
	return GetWorkingDir(), nil
}
