# BUG_REPRO

## Bug 是什么
签到状态检查把 Registered 写成 CheckedIn；签到记录列表排序写反；签到率漏乘 100；签到方式默认值写错。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestCheckinAndStats 失败：正常报名状态签到被拒、签到率不是 100。
