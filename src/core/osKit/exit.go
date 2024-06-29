package osKit

import (
	"os"
)

// Exit 退出程序
/*
PS: 无论是在main程还是子程中，只要调用os.Exit()，程序就会终止.

@param code (1) == 0: 正常退出
			(2) != 0（一般用1）: 非正常退出
*/
func Exit(code int) {
	RunExitHandlers()

	os.Exit(code)
}
