package time

import (
	"fmt"
	"testing"
	"time"
)

//Add
func Test_time_ops_01(t *testing.T) {
	now := time.Now()
	later := now.Add(time.Hour)   // 当前时间加1小时后的时间
	fmt.Println(later.Sub(now))   // 1h0m0s
	fmt.Println(later.After(now)) // true
}

// 定时器
func Test_ticker_01(t *testing.T) {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}
