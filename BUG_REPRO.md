# BUG_REPRO

## Bug 是什么
错误包装用 %v 丢掉哨兵；报名、活动、通知查询未命中返回新错误，导致 errors.Is 失效。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestErrorChainSentinels 失败：FindByID 未命中不再匹配 ErrNotFound。
