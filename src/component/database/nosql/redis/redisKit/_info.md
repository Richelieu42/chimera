(1) key、value、field都可以为"";

## 参考

notes/database（数据库）/Redis/Redis.wps
notes/Golang/database/golang - Redis.wps

## redis/go-redis官方文档

https://redis.uptrace.dev/zh/guide/

#### 连接池

https://redis.uptrace.dev/zh/guide/go-redis-debugging.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E5%A4%A7%E5%B0%8F
go-redis 底层维护了一个连接池，不需要手动管理。默认情况下， go-redis 连接池大小为 runtime.GOMAXPROCS * 10，
在大多数情况下默认值已经足够使用，且设置太大的连接池几乎没有什么用，可以在 配置项 中调整连接池数量.

## 命令教程

菜鸟教程:
https://www.runoob.com/redis/redis-tutorial.html
redis命令手册:
https://www.redis.net.cn/order/

## script（lua脚本）

Redis Cluster中使用Lua脚本
https://blog.csdn.net/qq_20128967/article/details/108611161
耗时12天，我整理Redis面试突击34问（含答案），助你面试“脱颖而出”（建议收藏）
https://www.bilibili.com/video/BV1XS4y1c7Tp/?buvid=Y44D4D448DC195994A5A88CED2DA982C60DF&is_story_h5=false&mid=5%2BiuUUrTqJQOdIa1r3VR0g%3D%3D&p=21&plat_id=114&share_from=ugc&share_medium=iphone&share_plat=ios&share_session_id=1F825A06-3FF6-4204-ABA8-F7FE5B30EB75&share_source=WEIXIN&share_tag=s_i&timestamp=1685414128&unique_k=RhYslZx&up_id=519608853
相关资料: 百度网盘"Redis面试资料"目录下

## 发布订阅 (pub/sub)

#### 取消订阅

在 goroutine1 中通过 PubSub.Channel()返回的只读信道ch 接收发布的数据，
过一段时间后，在 goroutine2 中调用 PubSub.Unsubscribe() 取消订阅，
此时虽然无法通过ch继续接收发布的数据，但 goroutine1 没有结束（还在从ch中读数据），直到 调用PubSub.Close() 才结束.

## TODO: key的前缀prefix

go-redis目前还不支持:
https://github.com/redis/go-redis/issues/607

在go-redis库中，你可以通过在每个键前面添加一个字符串来设置键的前缀。但是，go-redis库本身并没有提供直接设置键前缀的功能。
如果你需要在所有键前面添加一个公共的前缀，你可能需要自己实现这个功能。一种可能的方法是:
创建一个包装器函数，该函数接受一个键作为参数，然后返回一个带有前缀的键。

## Stream

### Consumer从哪里开始消费？

依据 XGroupCreate() 或 XGroupCreateMkStream() 的传参start.

### 删除消息（否则Redis占用的内存会越来越大）

![_stream.png](_stream.png)

## !!!: TongRDS的bug

将特殊的[]byte类型的数据（gob序列化的）存入redis，再取出来，结果和放进去的不一样.
一种解决方案: base64编码一下再存入Redis，取出后先base64解码.


