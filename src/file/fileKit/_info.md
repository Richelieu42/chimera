## !!!
- 不要在本包内调用 pathKit，以防 import cycle.

## TODOs
* GoFrame的 gfile 
    https://goframe.org/pages/viewpage.action?pageId=1114225
* duke-git/lancet的 fileutil
    https://www.golancet.cn/api/packages/fileutil.html#Fileutil

## 涉及标准库：
    os
    io
    io/ioutil
    bufio
    path/filepath

## 传参flag
PS:
(1) 多个flag可以通过 "|" 组合成一个；
(2) 更多可以参考"Golang.wps".

os.O_RDONLY	    只读模式(read-only)
os.O_WRONLY	    只写模式(write-only)
os.O_RDWR	    读写模式(read-write)

os.O_CREATE	    文件不存在就创建(create a new file if none exists.)
os.O_EXCL	    与 os.O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)

os.O_SYNC	    同步方式打开，即不使用缓存，直接写入硬盘

os.O_TRUNC	    打开并清空文件
os.O_APPEND	    追加模式(append)

## mimetype
h2non/filetype（Deprecated: 最后更新时间2021/1/21）:
    https://github.com/h2non/filetype
    https://mp.weixin.qq.com/s/MIhk4jGAYSxkJnOSH1Upbg
gabriel-vasile/mimetype:
    https://github.com/gabriel-vasile/mimetype

## Go语言读取文件的几种方式
https://mp.weixin.qq.com/s/St2EtX8s-V4okM9DpEzz7g
(1) 整个文件读取
(2) 按行读取
(3) 逐个单词读取
(4) 以数据块的形式读取文件
(5) 二进制读取

## 按行读取文件
Golang 按行读取文件的三种方法
    https://blog.csdn.net/K346K346/article/details/130337527
(1) 使用 bufio.Scanner 读取文件
(2) 使用 bufio.Reader 读取文件
(3) 使用 bufio.Reader 的 ReadLine 读取文件


