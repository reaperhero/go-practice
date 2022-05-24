package _map

import (
	"fmt"
	"testing"
	"time"
)

//更新 map 中 struct 元素的字段值，有 2 个方法

// 方法一：使用局部变量
// 提取整个 struct 到局部变量中，修改字段值后再整个赋值
type data struct {
	name string
}

func TestMap_01(t *testing.T) {
	m := map[string]data{
		"x": {"Tom"},
	}
	r := m["x"]
	r.name = "Jerry"
	m["x"] = r
	fmt.Println(m) // map[x:{Jerry}]
}

// 方法二：使用指向元素的 map 指针
// 使用指向元素的 map 指针
func TestMap_02(t *testing.T) {
	m := map[string]*data{
		"x": {"Tom"},
	}

	m["x"].name = "Jerry" // 直接修改 m["x"] 中的字段
	fmt.Println(m["x"])   // &{Jerry}
}

func TestMap_03(t *testing.T) {
	var a map[string]string
	a["s"] = "sad"
	fmt.Println(a)
}

func TestMap04(t *testing.T) {
	type a struct {
		Name string
		H    map[string]string
		B    map[string][]struct{}
	}
	b := a{B: make(map[string][]struct{})} // []struct不需要做初始化
	bB := b.B["1"]
	bB = append(bB, struct{}{})
}

func TestMap05(t *testing.T) {
	res := make(map[int]int)
	for i := 0; i < 100; i++ {
		i := i
		go func() {
			time.Sleep(time.Millisecond*time.Duration(i))
			res[i] = i
		}()
	}
	for i := range res {
		fmt.Println(i)
	}
	fmt.Println(111)

}
