package maptest

import (
	"fmt"
	"testing"
	"time"
)

type alert struct {
	front struct {
		end1 int
		end2 *int
	}
}

func (a *alert) print() {
	go func(bb *alert) {
		//fmt.Println(bb.front.end1)  //不会重复
		fmt.Println(*bb.front.end2) //重复打印
	}(a)
}

func TestMapread(t *testing.T) {
	testMap := make(map[string]alert)
	for i := 0; i < 5; i++ {
		testMap[fmt.Sprintf("%d", i)] = alert{
			front: struct {
				end1 int
				end2 *int
			}{end1: i, end2: &i}}
	}
	for _, i := range testMap {
		i.print()
	}
	time.Sleep(time.Second * 2)
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 3, 4, 5, 6, 7, 788, 9}
	for _, i := range s {
		//不正常
		//go func() {
		//	fmt.Println(i)
		//}()
		//正常
		go func(a int) {
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}
