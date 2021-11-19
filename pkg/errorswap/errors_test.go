package errorswap

import (
	"errors"
	"fmt"
	"testing"
)

//errors.Is 判断是否a错误是否是b错误的后代
//errors.Unwrap 将a错误的包装剔除一层
//errors.As 将a错误一直剔除到错误类型为 B 类型为止
//fmt.Errorf("%w", err) 将err错误包装一层

func TestErrorWap(t *testing.T) {
	o := errors.New("original error")
	n := fmt.Errorf("%w wrapped error", o)
	fmt.Printf("n is o: %t\n", errors.Is(n, o)) // n error container 0 error
	errors.Unwrap(n) // 包装剔除一层,return "original error"
}
