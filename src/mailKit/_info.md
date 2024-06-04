详见: "Golang - 1.docx"。

TODO: imap.go中的并发问题，可以考虑使用"通用连接池"。

## jordan-wright/email 
- [github](https://github.com/jordan-wright/email)
- [Go 每日一库之email](https://darjun.github.io/2020/02/16/godailylib/email/)

## gopkg.in/gomail.v2
- [github](https://github.com/go-gomail/gomail)
- [go-gin-api mail](https://www.yuque.com/xinliangnote/go-gin-api/rwim2a)
- [腾讯云 Golang语言怎么使用gomail库发送邮件？](https://cloud.tencent.com/developer/article/1770771)

## 协议s
PS: 简单来说，SMTP负责“发”，而POP3和IMAP负责“收”（Richelieu: 个人感觉IMAP更好）。  

- SMTP（Simple Mail Transfer Protocol，简单邮件传输协议）  
- POP3（Post Office Protocol，邮局协议）  
- IMAP（Internet Message Access Protocol，互联网消息访问协议）  

![_protocols.png](_protocols.png)

#### POP3 VS IMAP
总结
- POP3适用于单设备使用、需要离线访问邮件的场景，优点是减少服务器负担，但邮件状态不同步。
- IMAP适用于多设备使用、需要实时同步邮件的场景，优点是邮件状态统一，管理更方便，但需要持续的网络连接和更多的服务器存储空间。
![img.png](_vs.png)
