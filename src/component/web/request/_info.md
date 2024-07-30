## imroc/req VS go-resty/resty

### 参考

- [imroc/req 4k](https://github.com/imroc/req)
- [go-resty/resty 9.4k](https://github.com/go-resty/resty)

### 技术选型

![img.png](img.png)

### case: 短时间内发大量http请求

此种情况下，建议使用 valyala/fasthttp 库.
valyala/fasthttp 库的适用场景: 服务器/客户端需要每秒处理数千个中小型请求，并且需要一致的低毫秒级响应时间.

- [valyala/fasthttp](https://github.com/valyala/fasthttp)


