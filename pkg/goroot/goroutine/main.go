package main

import (
	"fmt"
	"time"
)

func hold1(s string) {
}

func hold2() string {
	time.Sleep(10 * time.Second)
	return "sdsa"
}

func main() {
	go hold1(hold2()) // 会阻塞
	fmt.Println("exec next")
}
