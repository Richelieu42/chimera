## vulcand/oxy

- [github 2k](https://github.com/vulcand/oxy)

#### 缺陷

- 没有健康检查机制;
- 明知道某一服务有问题，也会继续使用它，因此需要调大重试次数（建议>=服务数量），以免找不到好的服务.
- 有它代理，Centrifugo的 sse 和 http_stream 长连接连不上. 


