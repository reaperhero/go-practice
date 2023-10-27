package main

import "fmt"

//类型形参 (Type parameter)
//类型实参(Type argument)
//类型形参列表( Type parameter list)
//类型约束(Type constraint)
//实例化(Instantiations)
//泛型类型(Generic type)
//泛型接收器(Generic receiver)
//泛型函数(Generic function)

// type Slice[T int|float32|float64 ] []T
//
//	T 就是上面介绍过的类型形参
//	int|float32|float64 这部分被称为类型约束

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	//当调用泛型函数的时候, 我们经常可以使用类型推断。 注意，当调用 MapKeys 的时候， 我们不需要为 K 和 V 指定类型 - 编译器会进行自动推断

	fmt.Println("keys m:", MapKeys(m))
	//… 虽然我们也可以明确指定这些类型。

	mk := MapKeys[int, string](m)
	fmt.Println(mk)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
