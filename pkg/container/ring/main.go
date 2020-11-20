package main

import (
	"container/ring"
	"log"
)

func main()  {
	// 初始化环需要指定环的大小，一旦创建，长度是不可变的
	// 初始化一个大小为3的环
	r := ring.New(3)

	// 初始化
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}

	// 获取环的长度，复杂度是O(n)
	log.Print(r.Len())

	// 遍历环，需要记录第一个节点，然后就进行遍历
	p := r.Next()
	log.Println(p.Value)
	for p != r {
		p = p.Next()
		log.Print(p.Value)
	}

	// Ring提供Do方法，来遍历环中的每个元素
	var sum int
	r.Do(func(v interface{}) {
		sum += v.(int)
	})
	log.Print(sum)


	// Link能将2个环连接在一起,
	r2 := ring.New(3)
	for i := 4; i <= 6; i++ {
		r2.Value = i
		r2 = r2.Next()
	}
	r.Link(r2)

	// 当前指针移动n，n可以是负数，代表向prev的方向移动
	r.Move(2)

	// 从当前位置删除n个元素
	r.Unlink(3)

	// 最后谈一下ring的应用场景
	// ring可以用来保存固定数量的元素，例如保存最近100条日志，用户最近10次操作
	// ring的一个典型应用场景是构造定长环回队列，比如网页上的轮播；
}