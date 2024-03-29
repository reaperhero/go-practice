package singleflight

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量

	// 任务数，是worker的四
	task = make([]int, maxWorkers*4)
)

// Semaphore带权重的信号量，控制多个goroutine同时访问资源
// 使用场景：控制 goroutine 的阻塞和唤醒
func TestSemaphore(t *testing.T) {
	ctx := context.Background()
	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}
		// 启动worker goroutine
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}
	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		logrus.Printf("获取所有的worker失败: %v", err)
	}
	fmt.Println(maxWorkers, task)
}
