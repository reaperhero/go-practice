package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
)

// GIN是如何绑定参数并且校验的

func mapping(dst interface{}, m map[string][]string) {
	typ := reflect.TypeOf(dst)

	// 首先判断传入参数的类型
	if !(typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct) {
		log.Printf("Should pass ptr to destination struct object. Usage: mapping(&someStruct, m)")
		return
	}

	// 拿到指针所指向的元素的类型
	typ = typ.Elem()
	// 拿到指针所指向的元素的值
	value := reflect.ValueOf(dst).Elem()

	// 遍历每一个字段
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 忽略非导出字段
		if !field.IsExported() {
			log.Printf("field %s is not exported, ignore", field.Name)
			continue
		}

		// 判断是否设置了这个tag
		formTag := field.Tag.Get("form")
		if formTag == "" {
			log.Printf("tag `form` not exist in field, ignore")
			continue
		}

		// 查看是否有取值
		vs := m[formTag]
		if len(vs) == 0 {
			log.Printf("vs by formTag %s not found, ignore", formTag)
			continue
		}
		v := vs[len(vs)-1]

		// 根据类型来设置值
		switch fieldType := field.Type.Kind(); fieldType {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			typedV, _ := strconv.ParseInt(v, 10, 64)
			value.Field(i).SetInt(typedV)
		case reflect.String:
			value.Field(i).SetString(v)
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			typedV, _ := strconv.ParseUint(v, 10, 64)
			value.Field(i).SetUint(typedV)
		case reflect.Bool:
			value.Field(i).SetBool(v == "true")
		default:
			log.Printf("field type %s not support yet", fieldType)
		}
	}
}

func TestStructMap(t *testing.T) {
	m := map[string][]string{
		"name":  {"jhony"},
		"age":   {"1"},
		"money": {"10010010"},
	}

	type Person struct {
		Name     string `form:"name"`
		Age      uint   `form:"age"`
		Money    int64  `form:"money"`
		unexport string `form:"unexport"`
		NotFound bool   `form:"not_found"`
		NoTag    int8
	}
	p := Person{}
	mapping(&p, m)

	fmt.Printf("%v\n", p)
}
