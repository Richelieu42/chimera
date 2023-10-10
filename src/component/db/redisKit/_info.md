(1) key、value、field都可以为"";


## redis/go-redis官方文档
https://redis.uptrace.dev/zh/guide/
#### 连接池
https://redis.uptrace.dev/zh/guide/go-redis-debugging.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0%E5%A4%A7%E5%B0%8F
go-redis 底层维护了一个连接池，不需要手动管理。默认情况下， go-redis 连接池大小为 runtime.GOMAXPROCS * 10，
在大多数情况下默认值已经足够使用，且设置太大的连接池几乎没有什么用，可以在 配置项 中调整连接池数量.
#### TODO: 分布式锁
bsm/redislock(1.1k star): 
    https://github.com/bsm/redislock

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

## Stream（Redis5.0新增）
#### 参考
Redis Stream | 菜鸟教程
    https://www.runoob.com/redis/redis-stream.html

#### VS 发布订阅 (pub/sub)
* 发布订阅 (pub/sub) 有个缺点就是消息无法持久化，如果出现网络断开、Redis 宕机等，消息就会被丢弃.
* Redis Stream 提供了消息的持久化和主备复制功能，可以让任何客户端访问任何时刻的数据，并且能记住每一个客户端的访问位置，还能保证消息不丢失.
