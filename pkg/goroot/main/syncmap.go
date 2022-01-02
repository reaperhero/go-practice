package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	waitDone, _ := m.LoadOrStore("1", make(chan interface{}, 1))
	c := waitDone.(chan interface{})
	select {
	case aaa := <-c:
		fmt.Println(aaa)
	default:

	}

}
