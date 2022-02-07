package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

// 反射获取interface类型信息
func TestGetType_01(tt *testing.T) {
	a := "3.2"
	t := reflect.TypeOf(a) //反射获取interface类型信息
	switch t.Kind() {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

// 反射修改值
func Test_updateValue_01(t *testing.T) {
	var x float64 = 3.4
	// 反射认为下面是指针类型，不是float类型
	v := reflect.ValueOf(&x)
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
	fmt.Println("main:", x)
}

// 环境变量tag
func TestReflectsEnv(t *testing.T) {
	type Config struct {
		Name    string `env:"server-name"` // CONFIG_SERVER_NAME
		IP      string `env:"server-ip"`   // CONFIG_SERVER_IP
		URL     string `env:"server-url"`  // CONFIG_SERVER_URL
		Timeout string `env:"timeout"`     // CONFIG_TIMEOUT
	}
	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "geektutu.com")

	// read from xxx.json，省略
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("env"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}

	fmt.Printf("%+v", config)
}

// mapindex
func TestRefecttGetValue(tt *testing.T) {
	data := map[string]interface{}{"First": "firstValue"}
	cdata := reflect.ValueOf(data).MapIndex(reflect.ValueOf("First"))

	fmt.Printf("Value:%+v \n", cdata.Interface())
	fmt.Printf("Kind:%+v \n", cdata.Kind())

	type Test struct {
		Data string
	}

	d := map[string]Test{"Geeks": Test{Data: "data test"}}
	mydata := reflect.ValueOf(d).MapIndex(reflect.ValueOf("Geeks"))
	fmt.Println(reflect.ValueOf(mydata.Interface()))
}


