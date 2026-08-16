# BUG_REPRO

## Bug 是什么
报名和活动查询在未命中时返回 nil,nil，上层拿到 nil 后直接解引用，导致空指针 panic。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestMissingRegistrationAndActivity 失败：取消报名/统计缺失记录时发生 nil pointer dereference。
