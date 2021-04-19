package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力

// 反射获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a) //反射获取interface类型信息
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

func Test_getType_01(t *testing.T) {
	var x float64 = 3.4
	reflect_type(x)
}

// 反射获取interface值信息
func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

func Test_getValue_01(t *testing.T) {
	var x float64 = 3.4
	reflect_value(x)
}

//反射修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func Test_updateValue_01(t *testing.T) {
	var x float64 = 3.4
	// 反射认为下面是指针类型，不是float类型
	reflect_set_value(&x)
	fmt.Println("main:", x)
}

//type User struct {
//	Id   int
//	Name string
//	Age  int
//}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

// 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

// 结构体与反射
func Test_struct_01(t *testing.T) {
	u := User{1, "zs", 20}
	Poni(u)
}
