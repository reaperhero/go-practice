# benchmark

go test 命令默认不运行 benchmark 用例的，如果我们想运行 benchmark 用例，则需要加上 -bench 参数


benchmark 是如何工作的？

benchmark 用例的参数 b *testing.B，有个属性 b.N 表示这个用例需要运行的次数。b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快。最终确认一秒的操作数



```
go test -bench='Fib$' .           # 只运行以 Fib 结尾的 benchmark 用例
go test -bench='Fib$' -cpu=2,4 .  # -cpu 参数改变 GOMAXPROCS
go test -bench='Fib$' -benchtime=5s . # 压测5s。benchmark 的默认时间是 1s，，实际执行会大于5s
go test -bench='Fib$' -benchtime=50x . # 执行 50 次
go test -bench='Fib$' -benchtime=5s -count=3 . # 执行 5s, 并进行 3 轮 benchmark
go test -bench="Fib$" -cpuprofile=cpu.pprof .  # 支持生成 CPU(-cpuprofile)、memory(-memprofile) 和 block(-blockprofile) 的 profile 文件
```