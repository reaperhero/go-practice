package string

import (
	"fmt"
	"strings"
	"testing"
)

func Test_builder(T *testing.T) {
	s := strings.Builder{}
	s.Grow(128)
	s.WriteString("a")
	s.WriteString("b")
	fmt.Println(s.String()) // 􏰃􏰄􏰆􏱆􏰶􏰅􏰉􏰚􏰁􏱋􏱮􏱕􏱗􏱸􏱊􏰙􏱅􏰎􏰳􏱁􏱙􏰇􏰊􏰻调用string的时候才实际创建字符串
}
