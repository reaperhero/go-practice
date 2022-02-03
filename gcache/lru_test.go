package gcache

import (
	"fmt"
	"github.com/bluele/gcache"
	"testing"
	"time"
)

//Discards the least recently used items first.

func TestLruWithExpire(t *testing.T) {
	gc := gcache.New(20).
		LRU().
		Build()
	if err := gc.SetWithExpire("key", "ok", time.Second*10); err != nil {
		panic(err)
	}
	value, _ := gc.Get("key")
	fmt.Println("Get:", value)

	// Wait for value to expire
	time.Sleep(time.Second * 8)
	fmt.Println("Get:", value)

	// flush cache
	_ = gc.SetWithExpire("key", "ok", time.Second*5)
	time.Sleep(time.Second * 3)
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
