package slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	type a []struct {
	}
	var b a
	// b = make([]struct{},0) 这个初始化不做也可以
	b = append(b, struct{}{})

	c := make([]struct{}, 10)
	c = append(c, struct{}{})
	fmt.Println(len(c)) // 11
}
