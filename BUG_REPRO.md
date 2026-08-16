# BUG_REPRO

## Bug 是什么
报名名额校验写成 count > capacity 导致满员仍可报；去重查询只按 activity_id 过滤导致重复报名；报名状态默认值写成 cancelled。

## 如何触发
`cd backend && go test ./...`

## 错误信息
- service.TestSignupCapacityAndDedup 失败：第 3 人仍能报名、同一用户重复报名未拦截。
