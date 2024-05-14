## richelieu-yang/chimera
Golang的工具类.
- [github]https://github.com/richelieu-yang/chimera  

#### 手动更新ip库（ip2region.xdb）
src:  https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb  
dest: resource/ip2region/ip2region.xdb

## !!!: 业务项目
#### (1) 安装此依赖  
PS: 如果安装失败（因为部分依赖 i/o timeout等原因），可以考虑更改环境变量GOPROXY，比如切换为阿里的代理（ https://mirrors.aliyun.com/goproxy/ ）.  
命令: 
go get github.com/richelieu42/chimera

#### (2) 在main()所在的.go文件中，通过"import _"导入一些包  
- logrusInitKit（最优先，应该在最上面）
- jsonKit
- 业务自己的 config 包

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





