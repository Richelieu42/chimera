## 建议
业务中，如果同时存在Consumer和Producer，建议把他们的初始化放在不同的包内（以防: import cycle）.

## tag
tag可以为nil.

e.g.
Producer发出消息的tag为nil，Consumer收到对应消息的tag也为nil.


