package time

import (
	"fmt"
	"testing"
	"time"
)

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func parseStringtime() {
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(time.Now()))
}

func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 基于时间对象获取时间戳
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

// 将时间戳转为时间格式
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间操作
func timestampDemo3() {
	now := time.Now()
	now.Add(time.Hour)               // 返回一个小时之后的时间
	time.Now().Add(-time.Minute * 5) // 返回五分钟前的时间
}

// func (t Time) Sub(u Time) Duration   求两个时间之间的差值
// func (t Time) Before(u Time) bool    如果t代表的时间点在u之前，返回真；否则返回假。
// func (t Time) After(u Time) bool     如果t代表的时间点在u之后，返回真；否则返回假。

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器,本质上是一个通道（channel）
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func Test_ticker_timeout(t *testing.T) {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C // 2秒后可以读到值
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C // 被取消后，下面不会执行
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func Test_ticker_timeout_02(t *testing.T) {
	var msg string
	ch := make(chan string, 1)
	defer close(ch)

	go func() {
		//time.Sleep(1 * time.Microsecond)   // uncomment to timeout
		ch <- "hi"
	}()

	select {
	case msg = <-ch:
		fmt.Println("Read from ch:", msg)
	case <-time.After(1 * time.Microsecond):
		fmt.Println("Timed out")
	}
}

func Test_ticker_timeout_03(t *testing.T) {
	var msg string
	ch := make(chan string, 1)
	defer close(ch)

	timer := time.NewTimer(1 * time.Microsecond)
	defer timer.Stop()

	go func() {
		//time.Sleep(1 * time.Microsecond) // uncomment to timeout
		ch <- "hi"
	}()

	select {
	case msg = <-ch:
		fmt.Println("Read from ch:", msg)
	case <-timer.C:
		fmt.Println("Timed out")
	}
}
