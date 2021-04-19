# benchmark

```
go test -bench='Fib$' .           # 只运行以 Fib 结尾的 benchmark 用例
go test -bench='Fib$' -cpu=2,4 .  # -cpu 参数改变 GOMAXPROCS
go test -bench='Fib$' -benchtime=5s . # 压测5s。benchmark 的默认时间是 1s，，实际执行会大于5s
go test -bench='Fib$' -benchtime=50x . # 执行 50 次
go test -bench='Fib$' -benchtime=5s -count=3 . # 执行 50 次, 并进行 3 轮 benchmark
go test -bench="Fib$" -cpuprofile=cpu.pprof .  # 支持生成 CPU(-cpuprofile)、memory(-memprofile) 和 block(-blockprofile) 的 profile 文件
```