package _map

import (
	"fmt"
	"testing"
)

//更新 map 中 struct 元素的字段值，有 2 个方法

// 方法一：使用局部变量
// 提取整个 struct 到局部变量中，修改字段值后再整个赋值
type data struct {
	name string
}

func Test_map_01(t *testing.T) {
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
func Test_map_02(t *testing.T) {
	m := map[string]*data{
		"x": {"Tom"},
	}

	m["x"].name = "Jerry" // 直接修改 m["x"] 中的字段
	fmt.Println(m["x"])   // &{Jerry}
}
