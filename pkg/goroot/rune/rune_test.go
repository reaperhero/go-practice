package rune

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	var str = "小手25是什么"

	fmt.Println(len(str)) //17
	fmt.Println(str[:8])  //”小手25“共占8个字节

	s := []rune(str)
	fmt.Println(len(s))        //长度只有 7，每字汉字当一个字节
	fmt.Println(string(s[:4])) //取前4个，取出后转成string
}
