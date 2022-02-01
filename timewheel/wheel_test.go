package timewheel

import (
	"fmt"
	"github.com/wuYin/timewheel"
	"testing"
	"time"
)

func TestTaskWheel(t *testing.T)  {
	tw := timewheel.NewTimeWheel(100*time.Millisecond, 600) // 周期为一分钟

	// 执行定时任务
	tid, _ := tw.After(5*time.Second, func() interface{} {
		fmt.Println("after 5 seconds, task1 executed")
		return nil
	})

	// 执行指定次数的重复任务
	_, allDone := tw.Repeat(1*time.Second, 3, func() interface {}{
		fmt.Println("per 1 second, task2 executed")
		return nil
	})
	<-allDone

	// 中途取消任务
	tw.Cancel(tid)
}
