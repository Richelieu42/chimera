## 参考

- [在线Cron表达式生成器](https://cron.qqe2.com/)

## 注意点!!!

- *cron.Cron实例需要手动启动（通过 Start || Run）;
  (a) Run()    "会阻塞"调用此方法的goroutine
  (b) Start()  "不会阻塞"调用此方法的goroutine
- 希望停止任务，可以调用 cron.Cron 的 Stop() || Remove(id EntryID).
