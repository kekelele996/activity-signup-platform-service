# BUG_REPRO

## Bug 是什么
单条已读去掉了用户过滤；全部已读去掉了用户过滤；通知列表排序写反；通知默认已读状态写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestNotificationChain 失败：别人能标记我的通知已读、列表顺序反了。
