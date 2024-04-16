package ioKit

import "io"

// Copy 读取 io.Reader实例 的数据并写入 io.Writer实例.
/*
复制会持续进行直到 src 没有更多的数据可读（即 src.Read() 返回 io.EOF 错误），或者发生错误.

PS:
(1) 适用场景: 当不需要限制复制数据量时，直接使用 io.Copy 即可;
(2) 会"读"reader.

@param dst 	(1) 实现了 io.Writer 接口的对象
			(2) 一般是 *bytes.Buffer 类型
@param src 实现了 io.Reader 接口的对象，如文件、网络连接或字节缓冲
@return written	表示实际复制的数据量（以字节为单位）
@return err		在复制过程中遇到的第一个非 EOF 错误
*/
var Copy func(dst io.Writer, src io.Reader) (written int64, err error) = io.Copy

// CopyN 从 源src 向 目标dst 复制数据（仅复制指定的字节数）.
/*
此函数与 Copy 类似，但仅尝试复制指定数量 n 的字节。当复制了 n 字节后，或者源读取达到 EOF，或者发生错误时，复制停止。无论哪种情况，CopyN 都会返回已成功复制的字节数以及可能发生的任何错误。

PS:
(1) 适用场景: 如果需要精确控制复制特定数量的数据，应选用 io.CopyN;
(2) 会"读"reader.

@param n 要复制的字节数
*/
var CopyN func(dst io.Writer, src io.Reader, n int64) (written int64, err error) = io.CopyN

// CopyBuffer 从 源src 向 目标dst 复制数据（带缓冲区）.
/*
此函数与 Copy 功能相同，但是允许用户传递一个预分配的缓冲区 buf。这样，在复制数据的过程中，CopyBuffer 会使用提供的缓冲区暂存数据，避免每次系统调用时都需要临时分配新的缓冲区，从而提高了性能。对于那些不支持零拷贝或者其他优化机制的 Reader 或 Writer，手动提供缓冲区能够更有效地处理大量数据传输，尤其是当频繁进行小块数据传输时。

PS:
(1) 适用场景: 对于性能敏感且可以预估缓冲区合理大小的情况，使用 io.CopyBuffer 可以减少内存分配带来的开销;
(2) 会"读"reader.

@param buf 暂存数据的缓冲区
*/
var CopyBuffer func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) = io.CopyBuffer
