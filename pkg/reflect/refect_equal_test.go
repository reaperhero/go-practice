package main

import (
	"fmt"
	"reflect"
	"testing"
)

//
//DeepEqual函数用来判断两个值是否深度一致。具体比较规则如下：
//
//不同类型的值永远不会深度相等
//当两个数组的元素对应深度相等时，两个数组深度相等
//当两个相同结构体的所有字段对应深度相等的时候，两个结构体深度相等
//当两个函数都为nil时，两个函数深度相等，其他情况不相等（相同函数也不相等）
//当两个interface的真实值深度相等时，两个interface深度相等
//map的比较需要同时满足以下几个
//  -两个map都为nil或者都不为nil，并且长度要相等
//  -相同的map对象或者所有key要对应相同
//  -map对应的value也要深度相等
//指针，满足以下其一即是深度相等
//  -两个指针满足go的==操作符
//  -两个指针指向的值是深度相等的
//切片，需要同时满足以下几点才是深度相等
//  -两个切片都为nil或者都不为nil，并且长度要相等
//  -两个切片底层数据指向的第一个位置要相同或者底层的元素要深度相等
//  -注意：空的切片跟nil切片是不深度相等的
//其他类型的值（numbers, bools, strings, channels）如果满足go的==操作符，则是深度相等的。要注意不是所有的值都深度相等于自己，例如函数，以及嵌套包含这些值的结构体，数组等

type S struct {
	Name    string
	Age     int
	Address *int
	Data    []int
}

func TestRefectEqual(t *testing.T) {

	a := S{
		Name:    "aa",
		Age:     1,
		Address: new(int),
		Data:    []int{1, 2, 3},
	}
	b := S{
		Name:    "aa",
		Age:     1,
		Address: new(int),
		Data:    []int{1, 2, 3},
	}

	fmt.Println(reflect.DeepEqual(a, b)) // true
}
