package main

import (
	"container/list"
	"log"
)

func main() {
	// 初始化一个list链表
	l := list.New()
	// 向链表末尾添加元素
	l.PushBack(1)
	// 向链表头部添加元素
	l.PushFront(2)

	e := l.PushFront(3)
	// 在某个元前插入元素
	l.InsertBefore(4, e)
	// 在某个元素后插入元素
	l.InsertAfter(5, e)
	// 移动元素到最后
	l.MoveToBack(e)
	// 移动元素到最前
	l.MoveToFront(e)

	// 遍历链表
	for v := l.Front(); v != nil; v = v.Next() {
		log.Println(v.Value)
	}

	// 最后谈一下list的应用场景
	// list可以作为queue和 stack 的基础数据结构，list的一个典型应用场景是构造FIFO队列；

}
