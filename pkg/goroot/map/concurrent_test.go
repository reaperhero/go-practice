package _map

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 常规map并发操作
func TestMap01(t *testing.T) {
	a := make(map[int]string)
	lock := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			a[i] = fmt.Sprintf("%d", i)
		}()
	}
	time.Sleep(time.Second * 10)
}

func TestSyncMap01(t *testing.T) {
	var scene sync.Map
	for i := 0; i < 100; i++ {
		go func() {
			scene.Store(i, fmt.Sprintf("%d", i))
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			if value, ok := scene.Load(i); ok {
				v := value.(string)
				fmt.Println(v)
			}
		}()
	}
	time.Sleep(time.Second * 10)
}

func TestSyncMap02(t *testing.T) {
	var scene sync.Map
	type m struct {
		name string
	}
	for i := 0; i < 100; i++ {
		go func() {
			mv := m{name: fmt.Sprintf("%d", i)}
			scene.Store(i, &mv)
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			if value, ok := scene.LoadAndDelete(i); ok {
				v, ok := value.(*m)
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}()
	}
	time.Sleep(time.Second * 5)
}
