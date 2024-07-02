package zapKit

import (
	"os"
)

// Exit 退出程序（应用）.
/*
PS: 无论是在main程还是子程中，只要调用os.Exit()，程序就会终止.

@param codes	(1) 可以不传（默认code: 1）
				(2) 为了可移植性，推荐范围: [0, 125]
					For portability, the status code should be in the range [0, 125].
*/
func Exit(codes ...int) {
	/*
		(1) == 0: 正常退出
		(2) != 0（一般用1）: 非正常退出
	*/
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 1
	}

	// (1) 运行退出处理函数（毁尸灭迹）
	RunExitHandlers()

	// (2) 刷新所有日志缓冲
	Sync()

	os.Exit(code)
}
