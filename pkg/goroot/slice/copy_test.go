package slice

import (
	"fmt"
	"testing"
)

func Test_copy_01(t *testing.T) {

	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 0xc420080000
	res := make([]byte, 3)
	copy(res, raw[:3])
	fmt.Println(len(res), cap(res), &res[0]) // 3 3 0xc4200160b8
}

func TestSliceDelete(t *testing.T) {
	seq := []string{"a", "b", "c", "d", "e", "f", "g"}
	index := 6
	fmt.Println(seq[:index], seq[index+1:])     // [a b c] [e f g]
	seq = append(seq[:index], seq[index+1:]...) // [a b c e f g]
	fmt.Println(seq)
}

// copy( destSlice, srcSlice []T) int   返回值表示实际发生复制的元素个数
func TestSliceCopy01(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice2)  // [1 2 3]
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)  //[5 4 3 4 5]
}
