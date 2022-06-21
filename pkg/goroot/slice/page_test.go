package slice

import (
	"fmt"
	"reflect"
	"testing"
)

// 直接切片
func handlePaginate(item interface{}, page, size int) (int, error) {
	if page < 0 || size < 0 {
		return 0, fmt.Errorf("page size not legal")
	}
	offset := page * size
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Slice {
		return 0, fmt.Errorf("item not slice")
	}

	if v.Len() == 0 {
		return 0, fmt.Errorf("item len is zero")
	}
	count := v.Len()

	switch v.Index(0).Kind() {
	case reflect.String:
		var list []string
		for i := 0; i < count; i++ {
			list = append(list, v.Index(i).Interface().(string))
		}
		if count < offset {
			return count, nil
		}
		if count-offset <= size {
			list = list[offset:count]
		}
		if count-offset > size {
			list = list[offset : offset+size]
		}
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(list))
	}
	return count, nil
}

func TestSlicePage(t *testing.T) {
	list := []string{"1", "2", "3"}
	count, err := handlePaginate(&list, 0, 1)
	fmt.Println(list, count, err)
}
