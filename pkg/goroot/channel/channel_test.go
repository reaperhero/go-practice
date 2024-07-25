package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel_01(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok { // close: data,ok = 0 false
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

// 如何判断一个通道是否被关闭?
func Test_channel_02(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 已关闭的 channel 接收数据是安全的
func Test_channel_04(t *testing.T) {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done: // 关闭后会立即返回
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	fmt.Println("Result: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}

func TestName(t *testing.T) {
	ch := make(chan int, 1024)
	go func(ch chan int) {
		for {
			val := <-ch
			fmt.Printf("val:%d\n", val)
		}
	}(ch)

	tick := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case <-tick.C:
			fmt.Printf("%d: case <-tick.C\n", i)
		}

		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
	tick.Stop()
}

func TestNam1e(t *testing.T) {
	dChan := make(chan int)
	var ids []int
	go func() {
		for i := 0; i < 10; i++ {
			dChan <- i
			time.Sleep(time.Second)
		}
		close(dChan)
	}()
	for {
		select {
		case v, ok := <-dChan:
			ids = append(ids, v)
			if !ok {
				fmt.Println(ids)
				return
			}
		}
		fmt.Println(ids)
	}
}

func TestNam2e(t *testing.T) {
	var cc = make(chan struct{}, 1)
	go func() {
		close(cc)
	}()
	time.Sleep(time.Second * 1)

	select {
	case <-cc:
		fmt.Println(1) // 被执行
	case <-time.After(time.Second * 3):
		fmt.Println(2)
	}
	cc = nil
	select {
	case <-cc: // 不会选中
		fmt.Println(1)
	case <-time.After(time.Second * 3):
		fmt.Println(2) // 被执行
	}
}
