package gcache

import (
	"fmt"
	"github.com/bluele/gcache"
	"testing"
	"time"
)

// Discards the least frequently used items first.

func TestLfu(t *testing.T) {
	// size: 10
	gc := gcache.New(100).
		LFU().
		Expiration(time.Second * 3).
		Build()
	for i := 0; i < 100; i++ {
		gc.Set(fmt.Sprintf("key%d", i), "value")
	}
	go func() {
		time.Sleep(time.Second*2)
		gc.Set("key85", "value")
	}()
	time.Sleep(time.Second*4)
	for k, v := range gc.GetALL(true) {
		fmt.Println(k, v)
	}
	//key99 value
	//key97 value
	//key85 value
}
