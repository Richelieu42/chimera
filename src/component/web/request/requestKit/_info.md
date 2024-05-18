## !!!
(1) 请求的url地址，建议先使用 urlKit.PolyfillUrl() 对其进行处理.
e.g.
    url := "http://127.0.0.1/测试.wps"
    url, err := urlKit.PolyfillUrl(url, nil)
    if err != nil {
        panic(err)
    }

## imroc/req
- [快速开始（中文）](https://req.cool/zh/docs/prologue/quickstart/)
- [GO语言最好用的HTTP请求库](https://www.bilibili.com/video/BV1Mu4y197dA/)
- [github](https://github.com/imroc/req)

#### Close
不需要手动调用 resp.Body.Close()（无论是 普通GET、POST请求 还是 下载请求）.
e.g.
    普通GET请求，调用 Do() 时就已经将 resp.Body 关闭了，不管你后续读不读 resp.Body 的内容.

#### 处理 Response
https://req.cool/zh/docs/tutorial/handle-response/
判断响应状态码:
(1) resp.IsSuccessState()
(2) resp.IsErrorState()
Unmarshal结构体: resp.Unmarshal 或 resp.Into

#### 请求超时（timeout; 2选1，推荐使用第1种）
(1) 调用 Client.SetTimeout() 设置超时时间（imroc/req默认: 2 * time.Minute）
(2) 调用 Request.Do() 时传参（context.Context类型）
