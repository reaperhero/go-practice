package string

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func Test_builder(T *testing.T) {
	s := strings.Builder{}
	s.Grow(128)
	s.WriteString("a")
	s.WriteString("b")
	fmt.Println(s.String()) // 调用string的时候才实际创建字符串
}

func TestStrings(t *testing.T) {
	fmt.Println(strings.Contains("oasdsadk", "ok"))    // 需要连续的ok字符串
	fmt.Println(strings.ContainsAny("oasdsadk", "ok")) // 不需要连续的ok字符串
	fmt.Println(strings.Compare("okk", "ok"))          // 比较大小 if a==b, -1 if a < b, and +1 if a > b
}

func TestSplit(t *testing.T) {
	u,_ := url.Parse( "http://172.16.82.74:20080/DTStream_1654767287483332156/html")
	u.Path = "aaa.zip"
	fmt.Println(u.String())
}
