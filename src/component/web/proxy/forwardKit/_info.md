## 参考

- [HTTP 请求转发在Go当中的实践](https://zhuanlan.zhihu.com/p/349020346)
- notes/Golang/WEB/proxy（代理; forward）.wps

## 建议

代理请求失败时，建议返回状态码502(http.StatusBadGateway).


