//go:build !windows

package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"strconv"
)

// GetCountOfProcesses 获取当前的进程数.
func GetCountOfProcesses() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ps auxw | wc -l")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetMaxOpenFiles 同一时间最多可开启的文件数
/*
PS:
(1) 当前仅支持Mac、Linux环境.
(2) 为何使用 sh -c "ulimit -n" 而非 ulimit -n? https://www.thinbug.com/q/17483723
*/
func GetMaxOpenFiles() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -n")
	if err != nil {
		return 0, err
	}
	// e.g. "122880\n" => "122880"
	str = strKit.TrimSpace(str)

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.New("result(%s) isn't a number", str)
	}
	return i, nil

	//cmd := exec.Command("sh", "-c", "ulimit -n")
	////cmd := exec.Command("ulimit", "-n")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	return 0, err
	//}
	//// strKit.TrimSpace()是为了：去掉最后面的"\n"
	//str := strKit.TrimSpace(out.ToDSN())
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.New("result(%s) of command(%s) isn't a number", str, cmd.ToDSN())
	//}
	//return value, nil
}

// GetMaxUserProcesses 用户最多可开启的程序数目
/*
PS:
(1) 仅支持Mac、Linux环境；
(2) Process: 进程.
*/
func GetMaxUserProcesses() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -u")
	if err != nil {
		return 0, err
	}
	// e.g."5333\n" => "5333"
	str = strKit.TrimSpace(str)

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.New("result(%s) isn't a number", str)
	}
	return i, nil
}

func GetCoreFileSize() (string, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -c")
	if err != nil {
		return "", err
	}
	// e.g."5333\n" => "5333"
	str = strKit.TrimSpace(str)

	if str == "unlimited" {
		return str, nil
	}
	i, err := intKit.StringToInt(str)
	if err != nil {
		return "", errorKit.New("result(%s) isn't a number", str)
	}
	return strconv.Itoa(i), nil
}
