package singleflight

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

// 模拟从数据库读取
func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article: %d", id), nil
}

// 模拟优先读缓存，缓存不存在读取数据库，并且只有一个请求读取数据库，其它请求等待
func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})

	return v.(string), err
}

var count int32

func TestSingleflight(t *testing.T) {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetArticle(sg, 1)
			if res != "article: 1" {
				logrus.Panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s\n", n, time.Since(now))
}
