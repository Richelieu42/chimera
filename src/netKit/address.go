// Package netKit
/*
PS: Golang对于host的定义 不同于 js.

e.g. "127.0.0.1:8888"
host in Golang:	"127.0.0.1"
host in js: 	"127.0.0.1:8888"
*/
package netKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"strconv"
)

type (
	Address struct {
		// Hostname 包括ip
		Hostname string `json:"hostname" yaml:"hostname"`
		Port     int    `json:"port" yaml:"port"`
	}
)

func (addr *Address) String() string {
	return JoinHostnameAndPort(addr.Hostname, addr.Port)
}

// ParseToAddress
/*
e.g.
	("https://127.0.0.1") 		=> "127.0.0.1:443", nil
	("http://127.0.0.1:8888") 	=> "127.0.0.1:8888", nil
	("https://blog.csdn.net/weixin_52428496/article/details/110159938") => "blog.csdn.net:443", nil
*/
func ParseToAddress(str string) (*Address, error) {
	tmp := strKit.TrimSpace(str)
	tmp = strKit.ReplaceAll(tmp, "\\", "/")
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Newf("invalid str(%s)", str)
	}

	var scheme string
	index := strKit.Index(tmp, "://")
	if index != -1 {
		scheme = strKit.SubBefore(tmp, index)
		scheme = strKit.ToLower(scheme)
		tmp = strKit.SubAfter(tmp, index+3)
	}
	index = strKit.Index(tmp, "/")
	if index != -1 {
		tmp = strKit.Substring(tmp, 0, index)
	}
	if strKit.IsEmpty(tmp) {
		return nil, errorKit.Newf("invalid str(%s)", str)
	}

	arr := strKit.Split(tmp, ":")
	switch len(arr) {
	case 1:
		hostname := arr[0]
		port, err := getDefaultPortByScheme(scheme)
		if err != nil {
			return nil, errorKit.Wrapf(err, "invalid str(%s)", str)
		}
		return &Address{
			Hostname: hostname,
			Port:     port,
		}, nil
	case 2:
		hostname := arr[0]
		portStr := arr[1]
		if strKit.IsAllEmpty(hostname, portStr) {
			return nil, errorKit.Newf("invalid str(%s)", str)
		}

		var port int
		var err error
		if strKit.IsEmpty(portStr) {
			port, err = getDefaultPortByScheme(scheme)
		} else {
			port, err = strconv.Atoi(portStr)
		}
		if err != nil {
			return nil, errorKit.Wrapf(err, "invalid str(%s)", str)
		}

		return &Address{
			Hostname: hostname,
			Port:     port,
		}, nil
	default:
		return nil, errorKit.Newf("invalid str(%s)", str)
	}
}

func getDefaultPortByScheme(scheme string) (int, error) {
	switch scheme {
	case "http":
		return 80, nil
	case "https":
		return 443, nil
	default:
		return -1, errorKit.Newf("invalid scheme(%s)", scheme)
	}
}
