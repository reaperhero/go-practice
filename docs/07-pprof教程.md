# pprof



我们可以通过标准库的代码包runtime和runtime/pprof中的程序来生成三种包含实时性数据的概要文件，分别是CPU概要文件、内存概要文件和程序阻塞概要文件

- cpu概要文件

默认情况下，Go语言的运行时系统会以100 Hz的的频率对CPU使用情况进行取样

- 内存概要文件
  
内存概要文件用于保存在用户程序执行期间的内存使用情况。这里所说的内存使用情况，其实就是程序运行过程中堆内存的分配情况。

- 程序阻塞概要文件

程序阻塞概要文件用于保存用户程序中的Goroutine阻塞事件的记录。

## 调试

```
go tool pprof http://127.0.0.1/debug/pprof/heap # web文件
go tool pprof g1.out  # 本地文件
```

```
go tool pprof -http=127.0.0.1:8999 http://127.0.0.1/debug/pprof/heap



```