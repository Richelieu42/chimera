package httpClientKit

import (
	"context"
	"errors"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"net/url"
)

// IsTimeoutError 是否是请求超时错误?
func IsTimeoutError(err error) (flag bool) {
	if err != nil {
		/* (1) context.Context + http.NewRequestWithContext() */
		if errors.Is(err, context.DeadlineExceeded) {
			flag = true
			return
		}

		/* (2) http.Client 结构体的 Timeout 字段 */
		tmp := &url.Error{}
		// !!!: 由于url.Error{}结构体的Error()方法绑在 结构体指针 上，tmp前面要多个"&"
		if errors.As(err, &tmp) {
			if tmp.Timeout() {
				flag = true
				return
			}
		}

		/* (3) 判断错误文本 */
		if strKit.Contains(err.Error(), context.DeadlineExceeded.Error()) {
			flag = true
			return
		}
	}
	return
}
