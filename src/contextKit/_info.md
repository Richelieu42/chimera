## 注意点

- 如果涉及 http.Request.WithContext ，不要使用 gorilla/context库 !!!  
  does not play well with the shallow copying of the request that http.Request.WithContext (added to net/http Go 1.7
  onwards) performs

- 请求转发（代理请求），A服务将http请求代理给B服务时，A在 request context 里面塞了数据，  
  (a) 代理成功，B收到请求后无法从 request context 里面取得对应数据.  
  (b) 代理失败，A能从 request context 里面取得对应数据（理所当然，就是A塞进去的）.  


