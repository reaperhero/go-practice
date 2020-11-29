package stand

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 参考地址
// https://www.liwenzhou.com/posts/Go/go_context/

// WithCancel
func Test_WithCancel(t *testing.T) {
	// gen在单独的goroutine中生成整数
	// 将它们发送到返回的频道。
	// gen的调用者需要取消一次该上下文
	// 他们完成消耗生成的整数不泄漏
	// 内部goroutine由gen开始。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回不要泄漏goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们完成消耗整数时取消

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// WithDeadline
func Test_WithDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 即使ctx过期，最好还是调用
	// 取消功能无论如何。 如果不这样做可能会保留
	// 上下文及其父级的活动时间超过必要时间
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}

// WithTimeout
func Test_WithTimeout(t *testing.T) {
	// 传递带超时的上下文告诉阻塞函数
	// 超时过后应该放弃它的工作。
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}

}

// WithValue
func Test_WithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go") // 键值对

	f(ctx, k)                      // found value: Go
	f(ctx, favContextKey("color")) // key not found: color

}
