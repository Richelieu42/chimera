package netKit

import (
	"context"
	"net"
	"time"
)

var (
	// Dial Deprecated: Use DialTimeout instead.
	/*
		不推荐使用的原因: net.Dial如果连接的是未开放的端口,一个端口可能就是20s+.
	*/
	Dial func(network, address string) (net.Conn, error) = net.Dial

	// DialTimeout 建立一个到TCP服务器的连接.
	/*
		PS:
		(1) 可设置超时时间（如果超过 timeout 的指定的时间，连接没有完成，会返回超时错误）;
		(2) 如果你需要更多的控制并且需要超时，你可以组合使用 net.Dialer结构体、Dialer.Dial、Dialer.DialContext.

		e.g.
			net.DialTimeout("tcp", "example.com:80", 5*time.Second)
	*/
	DialTimeout func(network, address string, timeout time.Duration) (net.Conn, error) = net.DialTimeout

	// DialTCP 建立一个到TCP服务器的连接.
	/*
		@param network 	网络类型，可以是"tcp"、“tcp4”（仅IPv4）、“tcp6”（仅IPv6）
		@param laddr	本地的TCP地址（通常设为nil，系统会自动选择一个本地地址和端口来建立连接）
		@param raddr	远程的TCP地址

		PS:
		(1) 相较于 net.DialTimeout ，此函数 提供了更多的控制，但不直接支持超时;
		(2) 如果你需要更多的控制并且需要超时，你可以组合使用 net.Dialer结构体、Dialer.Dial、Dialer.DialContext.
	*/
	DialTCP func(network string, laddr, raddr *net.TCPAddr) (*net.TCPConn, error) = net.DialTCP
)

func DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	d := &net.Dialer{}
	return d.DialContext(ctx, network, address)
}

//func DialContextWithTimeout(ctx context.Context, network, address string, timeout time.Duration) (net.Conn, error) {
//	d := &net.Dialer{
//		Timeout: timeout,
//	}
//	return d.DialContext(ctx, network, address)
//}
