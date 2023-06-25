package maptest

import (
	"encoding/json"
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
		fmt.Println(bb.front.end1)  //重复打印
		fmt.Println(*bb.front.end2) //重复打印
	}(a)
}

func TestMapread(t *testing.T) {
	testMap := make(map[string]alert)
	for i := 0; i < 50; i++ {
		testMap[fmt.Sprintf("%d", i)] = alert{
			front: struct {
				end1 int
				end2 *int
			}{end1: i, end2: &i}}
	}
	for _, i := range testMap {
		i.print() // i 的内存地址一直没变，值在变，将会导致print时候会打印重复值
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

func TestM(t *testing.T) {
	var (
		a = make(map[int]int)
	)
	a[1] = 2
	fmt.Println(a)
}



func TestName(t *testing.T) {
	j := make(map[string]string)
	j["sada"]="sad"
	go func() {
		for i := 0; i < 100; i++ {
			_,ok:=j["sada"]
			if ok{
			}
			json.Marshal(j)
		}
	}()
	for i := 0; i < 100; i++ {
		json.Marshal(j)
	}

}


