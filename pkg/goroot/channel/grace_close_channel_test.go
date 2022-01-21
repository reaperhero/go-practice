package channel

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

//channel 问题
//在不改变 channel 自身状态的情况下，无法获知一个 channel 是否关闭。
//关闭一个 closed channel 会导致 panic。所以，如果关闭 channel 的一方在不知道 channel 是否处于关闭状态时就去贸然关闭 channel 是很危险的事情。
//向一个 closed channel 发送数据会导致 panic。所以，如果向 channel 发送数据的一方不知道 channel 是否处于关闭状态时就去贸然向 channel 发送数据是很危险的事情

//根据 sender 和 receiver 的个数，分下面几种情况：
//1、一个 sender，一个 receiver: close by sender
//2、一个 sender， M 个 receiver: close by sender

//3、N 个 sender，一个 reciver: 解决方案就是增加一个传递关闭信号的 channel，receiver 通过信号 channel 下达关闭数据 channel 指令。senders 监听到关闭信号后，停止发送数据
func TestGraceClose03(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const Max = 1000
	const NumSenders = 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	// the receiver
	go func() {
		for value := range dataCh {
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)
				return
			}
			fmt.Println(value)
		}
	}()

	time.Sleep(time.Second*4)

}

//4、N 个 sender， M 个 receiver: 需要增加一个中间人，M 个 receiver 都向它发送关闭 dataCh 的“请求”，中间人收到第一个请求后，就会直接下达关闭 dataCh 的指令（通过关闭 stopCh，这时就不会发生重复关闭的情况，因为 stopCh 的发送方只有中间人一个）。另外，这里的 N 个 sender 也可以向中间人发送关闭 dataCh 的请求。
func TestGraceClose04(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	const Max = 1000
	const NumReceivers = 10
	const NumSenders = 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// It must be a buffered channel.
	toStop := make(chan string, 1)


	// moderator
	go func() {
		_ = <-toStop
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				select {
				case <- stopCh:
					return
				case dataCh <- value:
				}
				fmt.Println(value)
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <- stopCh:
					return
				case value := <-dataCh:
					if value == Max-1 {
						select {
						case toStop <- "receiver#" + id:
						default:
						}
						return
					}
					fmt.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	time.Sleep(time.Second*4)
}