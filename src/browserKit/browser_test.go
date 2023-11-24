package browserKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"testing"
	"time"
)

// 让默认浏览器打开 "https://www.baidu.com"
func TestOpenURL(t *testing.T) {
	if err := OpenURL("https://www.baidu.com"); err != nil {
		panic(err)
	}
}

// 打开文件.
func TestOpenFileOrDirectory(t *testing.T) {
	temp, err := pathKit.GetExclusiveTempDir()
	if err != nil {
		panic(err)
	}
	path := pathKit.Join(temp, fmt.Sprintf("%s.wps", idKit.NewXid()))

	f, err := fileKit.Create(path)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(`份额额服务强
无敌
aas
~!@#$%^&*()_+{}:"|<>?"
群无
多`)
	if err != nil {
		panic(err)
	}
	_ = f.Close()

	if err := OpenFileOrDirectory(path); err != nil {
		panic(err)
	}

	// Richelieu: 过一会再删，防止打开后内容为空
	time.Sleep(time.Second * 3)
	_ = fileKit.Delete(path)
}

// 打开目录.
func TestOpenFileOrDirectory1(t *testing.T) {
	dir := "/Users/richelieu/Downloads"
	if err := OpenFileOrDirectory(dir); err != nil {
		panic(err)
	}
}
