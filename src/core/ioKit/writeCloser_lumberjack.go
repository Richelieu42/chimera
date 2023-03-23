package ioKit

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

// NewLumberjackWriteCloser
/*
PS:
(1) lumberjack: 日志切割组件.一般情况下，lumberjack配合其他日志库，实现日志的滚动(rolling)记录.
(2) 根据 MaxBackups、MaxAge 删除过期文件，根据 Compress 决定是否压缩哪些未压缩的旧日志文件。

Golang 语言三方库 lumberjack 日志切割组件怎么使用？
	https://mp.weixin.qq.com/s/gGnovwzS1ucW3Afxcytp_Q
go语言的日志滚动(rolling)记录器——lumberjack
	https://zhuanlan.zhihu.com/p/430224518

@param filePath		日志路径，归档日志也会保存在对应目录下（若该值为空，则日志会保存到os.TempDir()目录下，日志文件名为<processname>-lumberjack.log）
@param maxSize		（单位: MB，默认: 100MB）日志大小到达此值，就开始backup
@param maxBackups	旧日志保存的最大数量，默认保存所有旧日志文件
@param maxAge		（单位: days）旧日志保存的最大天数，默认保存所有旧日志文件
@param localTime	（默认使用UTC时间）是否使用本地时间戳？
@param compress		（默认: false）对backup的日志是否进行压缩
*/
func NewLumberjackWriteCloser(filePath string, maxSize, maxBackups, maxAge int, localTime, compress bool) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
	}
}
