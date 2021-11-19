package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	a := sync.Map{}
	a.Store("key", "v")
	v, _ := a.Load("key")
	fmt.Println(v)
}

