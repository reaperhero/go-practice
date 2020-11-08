package golang_lru


import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"testing"
)

//  thread safe lru队列
func Test_lRu_01(t *testing.T)  {
	l, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
}
