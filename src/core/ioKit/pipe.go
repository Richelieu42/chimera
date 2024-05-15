package ioKit

import "io"

var (
	// Pipe
	/*
		pipe: 管道
		(1) 数据同步：PipeReader 和 PipeWriter 是同步的，这意味着当 PipeWriter 写入数据时，PipeReader 可以立即读取到这些数据，反之亦然。这使得在不同的协程间可以同时进行读写操作而不会丢失数据。
		(2) 非阻塞操作：Pipe 提供了一种非阻塞的方式进行读写。如果写端关闭，读端会立即知道并能检测到这个状态；同样，如果读端关闭，写端在尝试写入时也会立即得到通知。
		(3) 数据流处理：io.Pipe 常用于构建数据处理管道，例如在数据被读取后经过一系列中间处理步骤，然后被写入到最终目的地，比如网络连接或文件。
	*/
	Pipe func() (*io.PipeReader, *io.PipeWriter) = io.Pipe
)
