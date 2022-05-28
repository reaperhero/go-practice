package _select

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func breakSelect() {
	//ee: // 不会打印111，会返回
	for {
	ee: // 会打印 111，一直循环
		select {
		case <-time.After(time.Second * 1):
			for i := 0; i < 3; i++ {
				if i == 2 {
					break ee
				}
				fmt.Println(i)
			}
			fmt.Println(111)
		}
	}

}



func closeChan() error {
	errChan := make(chan error)
	exitChan := make(chan struct{})
	finishChan := make(chan struct{})
	wait := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			select {
			case <-exitChan:
				log.Println("exit chan")
				return
			default:
				time.Sleep(time.Second)
				if i == 5 {
					errChan <- fmt.Errorf("err chan reciver")
				}
				fmt.Println(i)
				return
			}
		}(i)
	}

	go func() {
		wait.Wait()
		finishChan <- struct{}{}
	}()
	select {
	case err := <-errChan:
		close(exitChan)
		log.Println(err)
		return err
	case <-finishChan:
		fmt.Println("success")
		return nil
	}
}

func TestSelect(t *testing.T) {
	closeChan()
	time.Sleep(time.Second * 3)

}

func TestSelectWait(t *testing.T) {

	wg := sync.WaitGroup{}
	waitCh := make(chan struct{})
	wg.Add(10)

	// In another go routine Wait for the wait group to finish.
	go func() {
		// Run some actions
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				fmt.Println("do some action")
			}()
		}

		wg.Wait()
		close(waitCh)
	}()

	select {
	case <-waitCh:
		fmt.Println("WaitGroup finished!")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("WaitGroup timed out..")
	}
}




func TestName(t *testing.T)  {
	//tk := time.After(time.Second*10)
	//
	//go func() {
	//	no:
	//	for  {
	//		select {
	//		case <-tk:
	//			fmt.Println(2)
	//			return
	//		case <-time.After(time.Second*2):
	//			fmt.Println(1)
	//			goto no
	//		}
	//		fmt.Println(3)
	//	}
	//}()
	//time.Sleep(time.Second*30)
}


