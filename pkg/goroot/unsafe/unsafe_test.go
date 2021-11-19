package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Demo struct {
	s  string
	i  int
	f  float64
	bs []byte
}

//Alignof: 输出给定类型内存对齐的大小
//Offsetof: 输出给定结构体具体属性相对于结构体其实内存位置的偏移量
//Sizeof: 输出给定类型所占内存的大小
func TestUnsafe(t *testing.T) {
	d := Demo{}

	fmt.Println("Alignof:")
	fmt.Println(unsafe.Alignof(d.s))
	fmt.Println(unsafe.Alignof(d.i))
	fmt.Println(unsafe.Alignof(d.f))
	fmt.Println(unsafe.Alignof(d.bs))

	fmt.Println("Offsetof:")
	fmt.Println(unsafe.Offsetof(d.s))
	fmt.Println(unsafe.Offsetof(d.i))
	fmt.Println(unsafe.Offsetof(d.f))
	fmt.Println(unsafe.Offsetof(d.bs))

	fmt.Println("Sizeof:")
	fmt.Println(unsafe.Sizeof(d.s))
	fmt.Println(unsafe.Sizeof(d.i))
	fmt.Println(unsafe.Sizeof(d.f))
	fmt.Println(unsafe.Sizeof(d.bs))
}

func TestUnsafeString(t *testing.T) {
	// b2s converts byte slice to a string without memory allocation.
	b2s := func(b []byte) string {
		return *(*string)(unsafe.Pointer(&b))
	}
	// s2b converts string to a byte slice without memory allocation.
	s2b := func(s string) []byte {
		sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
		bh := reflect.SliceHeader{
			Data: sh.Data,
			Len:  sh.Len,
			Cap:  sh.Len,
		}
		return *(*[]byte)(unsafe.Pointer(&bh))
	}

	fmt.Println(b2s([]byte("dsads")))
	fmt.Println(s2b(("dsads")))
}
