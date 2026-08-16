# BUG_REPRO

## Bug 是什么
座位判空改成了反向状态判断导致 nil 解引用；存储初始化未建 map；订座不再校验查询结果导致空指针；统计把未预订座位误算为已预订。

## 如何触发
`go test ./...`

## 错误信息
- model.TestIsBookedNilSafe 触发 nil 指针 panic。
- store.TestNewInitializesMap 失败（map 未初始化）。
- service.TestBookMissingReturnsError 触发 nil 指针 panic。
- worker.TestCountBooked 失败（统计反向）。
