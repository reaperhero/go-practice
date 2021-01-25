# pprof

我们可以通过标准库的代码包runtime和runtime/pprof中的程序来生成三种包含实时性数据的概要文件，分别是CPU概要文件、内存概要文件和程序阻塞概要文件

- cpu概要文件

默认情况下，Go语言的运行时系统会以100 Hz的的频率对CPU使用情况进行取样

- 内存概要文件
  
内存概要文件用于保存在用户程序执行期间的内存使用情况。这里所说的内存使用情况，其实就是程序运行过程中堆内存的分配情况。

- 程序阻塞概要文件

程序阻塞概要文件用于保存用户程序中的Goroutine阻塞事件的记录。

## 调试

- 配置
```
go tool pprof http://127.0.0.1/debug/pprof/block   # 查看导致阻塞同步的堆栈跟踪
go tool pprof http://127.0.0.1/debug/pprof/profile # cpu
go tool pprof http://127.0.0.1/debug/pprof/goroutine   # 查看当前所有运行的 goroutines 堆栈跟踪
go tool pprof http://127.0.0.1/debug/pprof/heap # 查看活动对象的内存分配情况
go tool pprof g1.out  # 本地文件
   参数：
     ?seconds=60
   
   命令行：
     web 打开浏览器
```

- 结果
```
cpu
    (pprof) top10
    Showing nodes accounting for 25.92s, 97.63% of 26.55s total
    Dropped 85 nodes (cum <= 0.13s)
    Showing top 10 nodes out of 21
          flat  flat%   sum%        cum   cum%
        23.28s 87.68% 87.68%     23.29s 87.72%  syscall.Syscall
         0.77s  2.90% 90.58%      0.77s  2.90%  runtime.memmove
         0.58s  2.18% 92.77%      0.58s  2.18%  runtime.freedefer
         0.53s  2.00% 94.76%      1.42s  5.35%  runtime.scanobject
         0.36s  1.36% 96.12%      0.39s  1.47%  runtime.heapBitsForObject
         0.35s  1.32% 97.44%      0.45s  1.69%  runtime.greyobject
         0.02s 0.075% 97.51%     24.96s 94.01%  main.main.func1
         0.01s 0.038% 97.55%     23.91s 90.06%  os.(*File).Write
         0.01s 0.038% 97.59%      0.19s  0.72%  runtime.mallocgc
         0.01s 0.038% 97.63%     23.30s 87.76%  syscall.Write

     flat：给定函数上运行耗时
     flat%：同上的 CPU 运行耗时总比例
     sum%：给定函数累积使用 CPU 总比例
     cum：当前函数加上它之上的调用运行总耗时
     cum%：同上的 CPU 运行耗时总比例

heap
    (pprof) top
    Showing nodes accounting for 837.48MB, 100% of 837.48MB total
          flat  flat%   sum%        cum   cum%
      837.48MB   100%   100%   837.48MB   100%  main.main.func1

    -inuse_space：分析应用程序的常驻内存占用情况
    -alloc_objects：分析应用程序的内存临时分配情况
```


- 火焰图

可以查看所有信息
```
go tool pprof -http=127.0.0.1:8999 http://127.0.0.1/debug/pprof/profile
```