package main

import (
	"sort"
)

func main() {

	// heap包提供了堆相关的操作，并不想list和ring一样提供New()的创建方法
	// heap包中定义了一个Interface接口，只要实现了这个接口的struct都可以当作是堆
	type Interface interface {
		sort.Interface
		Push(x interface{}) // add x as element Len()
		Pop() interface{}   // remove and return element Len() - 1.
	}

	//heap.Fix()

	// 应用场景
	// heap可以用来排序。游戏编程中是一种高效的定时器实现方案
	// heap的一个典型应用场景是构造优先级队列
}
