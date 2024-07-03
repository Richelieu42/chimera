## 参考

- [HTTP 请求转发在Go当中的实践](https://zhuanlan.zhihu.com/p/349020346)
- notes/Golang/WEB/proxy（代理; forward）.wps

## 建议

代理请求失败时，建议返回状态码502(http.StatusBadGateway, 网关错误).
![_502.png](_502.png)

## !!!请求转发的可能原因

- linux服务器 请求转发给 Windows机器，报错: 超时timeout，原因: Windows机器的域防火墙（入站规则）拦截了请求.


