package slice

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_copy_01(t *testing.T) {

	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0]) // 10000 10000 0xc420080000
	res := make([]byte, 3)
	copy(res, raw[:3])
	fmt.Println(len(res), cap(res), &res[0]) // 3 3 0xc4200160b8
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

func TestSliceDelete(t *testing.T) {
	{
		// 删除第一个
		seq := []string{"a", "b", "c", "d", "e", "f", "g"}
		index := 0
		fmt.Println(seq[:index], seq[index+1:]) // [] [b c d e f g]
	}
	{
		// 删除最后一个
		seq := []string{"a", "b", "c", "d", "e", "f", "g"}
		index := 6
		fmt.Println(seq[:index], seq[index+1:]) // [a b c d e f] []
	}
}

func TestDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 1, 5, 1}
	for i := 0; i < len(a); i++ {
		if a[i] > 3 {
			a = append(a[:i], a[i+1:]...) // [1 2 3 1 1]
			i--
		}
	}
	fmt.Println(a)
}

func TestModifySlice1(t *testing.T) {
	var s []int
	modify1 := func(slice []int) {
		slice = append(slice, 1)
	}
	modify1(s) // 不会对s进行修改，因为已经slice扩容，底层数组改变
}

func TestModifySlice2(t *testing.T) {
	s := make([]int, 0, 10)
	modify1 := func(slice *[]int) {
		*slice = append(*slice, 1)
	}
	modify1(&s) // 会对s进行修改
}


func TestName(t *testing.T)  {
	fmt.Println(regexp.MustCompile(`(?i).*password.*|^sensitive_`).MatchString("sensitive_password"))
}
