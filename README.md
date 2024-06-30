## richelieu-yang/chimera

Golang的工具类  
Tools of Golang

- [github](https://github.com/richelieu-yang/chimera)

## 业务项目

#### (1) 安装此依赖

命令: go get github.com/richelieu42/chimera/v3

#### (2) 在main()所在的.go文件中，通过 "import _" 导入一些包

- jsonKit
- 业务自己的 config 包

## 应用进程的退出(exit)

#### 主动退出

不要使用 os.Exit，建议使用 appKit.Exit.

#### 监听退出信号

建议使用 signalKit.MonitorExitSignals，其内部会调用 appKit.Exit.

## 参考

- duke-git/lancet  
  官方API说明  
  https://www.golancet.cn/api/overview.html  
  支持300+常用功能的开源GO语言工具函数库  
  https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247498172&idx=1&sn=461d8429c094189f4e10732d00805339  
  github(3.2k): https://github.com/duke-git/lancet/blob/main/README_zh-CN.md
- samber/lo  
  Golang.wps  
  github(14.2k): https://github.com/samber/lo
- GoFramev2  
  https://goframe.org/pages/viewpage.action?pageId=1114859  
  github(10.4k): https://github.com/gogf/gf

### TODOs





