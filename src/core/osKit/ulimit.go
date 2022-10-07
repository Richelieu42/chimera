//go:build !windows

package osKit

import (
	"gitee.com/richelieu042/go-scales/src/cmdKit"
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/intKit"
)

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

	i, err := intKit.ParseStringToInt(str)
	if err != nil {
		return 0, errorKit.Simple("result(%s) isn't a number", str)
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
	//// strKit.Trim()是为了：去掉最后面的"\n"
	//str := strKit.Trim(out.String())
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.Simple("result(%s) of command(%s) isn't a number", str, cmd.String())
	//}
	//return value, nil
}

// GetUserMaxProcesses 用户最多可开启的程序数目
/*
PS:
仅支持Mac、Linux环境；
Process: 进程.
*/
func GetUserMaxProcesses() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -u")
	if err != nil {
		return 0, err
	}

	i, err := intKit.ParseStringToInt(str)
	if err != nil {
		return 0, errorKit.Simple("result(%s) isn't a number", str)
	}
	return i, nil

	//cmd := exec.Command("sh", "-c", "ulimit -u")
	////cmd := exec.Command("ulimit", "-u")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	return 0, err
	//}
	//// strKit.Trim()是为了：去掉最后面的"\n"
	//str := strKit.Trim(out.String())
	//
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.Simple("result(%s) of command(%s) isn't a number", str, cmd.String())
	//}
	//return value, nil
}
