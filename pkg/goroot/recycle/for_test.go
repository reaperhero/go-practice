package recycle

import (
	"fmt"
	"testing"
)

func Test_continue01(t *testing.T) {

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 { // j==4 和 j==5 的时候，没有任何输出, 标签的作用对象为外部循环，因此 i 会直接变成下一个循环的值
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func Test_continue02(t *testing.T) {
	i := 0
HERE:
	print(i) // 01234
	i++
	if i == 5 {
		return
	}
	goto HERE   // goto 必须在标签定义的后面
}
