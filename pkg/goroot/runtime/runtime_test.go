package runtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// runtime.Gosched()
func Test_gosched_01(t *testing.T) {

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 出让CPU时间片
		fmt.Println("hello")
	}
}

// runtime.Goexit()
func Test_Goexit_02(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 以下内容不会不会执行
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		time.Sleep(1 * time.Second)
	}
}
