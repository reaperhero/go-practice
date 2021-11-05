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
	fmt.Println(seq[:index], seq[index+1:]) // [a b c] [e f g]
	seq = append(seq[:index], seq[index+1:]...) // [a b c e f g]
	fmt.Println(seq)
}
