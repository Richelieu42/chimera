## redis_store.go
实际上是修改了 https://github.com/gorilla/sessions + github.com/rbcervilla/redisstore/v9 v9.0.0.（因为该依赖有bug，详见 go - session.wps）
PS:
(1) 想看修改了哪些，可以"在文件中搜索关键词Richelieu"，或者"看文件的git记录"（推荐）.
(2) session.IsNew == true 的可能情况:  (a) session.ID == "": 浏览器端"无"记录session id的cookie && Redis中没有对应数据;
                                      (b) session.ID != "": 浏览器端"有"记录session id的cookie && Redis中没有对应数据.

!!!:
(1) 如果要将session存储到 Redis 中，需要使用 "分布式唯一id"!!! 否则如果生成的id一致，后面的会覆盖前面的.

## 3种方案
探索分布式 Session 管理
    https://mp.weixin.qq.com/s/yRN7nsM31K9SmX-GibctQg
* session复制（仅适用于小规模应用）
* session复制
* Session共享（将会话信息存储在可共享的第三方缓存中，如 Redis 或 Memcached）
