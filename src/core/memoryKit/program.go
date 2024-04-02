package memoryKit

import (
	"runtime"
)

// GetProgramMemoryStats 获取有关Go程序的实时内存统计信息.
/*
runtime.MemStats结构体的的字段:
(1) Alloc（单位: 字节）
	Alloc is bytes of allocated heap objects.
	当前分配给 heap 的内存量（不包括尚未释放的对象）.

(2) TotalAlloc（单位: 字节）
	TotalAlloc is cumulative bytes allocated for heap objects.
	自程序启动以来，堆上分配的总内存量.

(3) Sys（单位: 字节）
	Sys is the total bytes of memory obtained from the OS.
	程序向操作系统申请的总内存量（包括未使用的页）.

(4) NumGC
	NumGC is the number of completed GC cycles.
	它表示自程序启动以来垃圾回收器运行的次数。每次垃圾回收器运行时，NumGC 的值都会增加。这个字段可以用来监控程序的垃圾回收情况。
	!!!: 如果你发现 NumGC 的值增长得非常快，那么可能意味着你的程序存在内存分配问题。

(5) EnableGC
	EnableGC indicates that GC is enabled. It is always true, even if GOGC=off.
	表示是否允许垃圾回收。它是一个 bool 类型，如果为 true，则允许垃圾回收；如果为 false，则禁止垃圾回收。

(6) DebugGC
	DebugGC is currently unused.
	表示是否启用调试垃圾回收。它是一个 bool 类型，如果为 true，则启用调试垃圾回收；如果为 false，则禁用调试垃圾回收。
*/
func GetProgramMemoryStats() *runtime.MemStats {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	return stats
}
