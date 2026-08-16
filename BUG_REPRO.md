# BUG_REPRO

## Bug 是什么
座位判空缺失导致 nil 指针解引用；存储初始化未建 map 导致写入 nil map 崩溃；查询缺失座位返回 (nil,true) 使上层空指针；统计把未预订座位误算为已预订。

## 如何触发
`go test ./...`

## 错误信息
- model.TestIsBookedNilSafe 触发 nil 指针 panic。
- service.TestBookMissingReturnsError 触发 nil 指针 panic。
- store.TestNewInitializesMap / TestGetSeatMissing 失败（nil map、ok=true）。
- worker.TestCountBooked 触发 assignment to entry in nil map panic。
