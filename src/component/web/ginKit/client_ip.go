package ginKit

import "github.com/gin-gonic/gin"

// GetRemoteIP 获取客户端IP地址（客户端的远程IP地址）.
/*
e.g.
当客户端通过代理服务器连接时，RemoteIP() 返回代理服务器的 IP 地址
*/
func GetRemoteIP(ctx *gin.Context) string {
	return ctx.RemoteIP()
}

// GetClientIp 获取客户端IP地址（客户端的真实IP地址，但结果并不总是可靠的）.
/*
PS:
(1) ClientIP 方法首先检查 HTTP 请求头中的 X-Forwarded-For 和 X-Real-Ip 字段。如果这些字段存在，则 ClientIP() 会返回其中的值作为客户端的真实 IP 地址。
	如果这些字段不存在，则 ClientIP() 会返回与 RemoteIP() 相同的结果。
(2) ctx.ClientIP() VS ctx.RemoteIP():
	(2.1) 当客户端直接连接到服务器时，RemoteIP 和 ClientIP 方法返回的结果相同;
	(2.2) 当客户端通过代理服务器连接时，	(a) RemoteIP() 返回代理服务器的 IP 地址;
							 		(b) ClientIP() 尝试从 HTTP 请求头中获取客户端的真实 IP 地址.
(3) 想要更加精确地获取客户端的真实 IP 地址，可以，通过 Engine.TrustedPlatform 和 Engine.SetTrustedProxies()，参考同目录下的 _info.md.

总结:
	在使用Gin框架时，如果您希望获取客户端的真实 IP 地址，应该使用 ClientIP() 而不是 RemoteIP()。但是，请注意，由于客户端可以伪造 HTTP 请求头中的 X-Forwarded-For 和 X-Real-Ip 字段，因此 ClientIP 方法返回的结果并不总是可靠的。在某些情况下，您可能需要使用其他方法来验证客户端的真实 IP 地址。
*/
func GetClientIp(ctx *gin.Context) string {
	return ctx.ClientIP()
}
