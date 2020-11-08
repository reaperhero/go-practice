package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好


//读取操作
//func LoadInt32(addr int32) (val int32)
//func LoadInt64(addr `int64) (val int64)
//func LoadUint32(addruint32) (val uint32)
//func LoadUint64(addruint64) (val uint64)
//func LoadUintptr(addruintptr) (val uintptr)
//func LoadPointer(addrunsafe.Pointer`) (val unsafe.Pointer)
//写入操作
//func StoreInt32(addr *int32, val int32)
//func StoreInt64(addr *int64, val int64)
//func StoreUint32(addr *uint32, val uint32)
//func StoreUint64(addr *uint64, val uint64)
//func StoreUintptr(addr *uintptr, val uintptr)
//func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
//修改操作
//func AddInt32(addr *int32, delta int32) (new int32)
//func AddInt64(addr *int64, delta int64) (new int64)
//func AddUint32(addr *uint32, delta uint32) (new uint32)
//func AddUint64(addr *uint64, delta uint64) (new uint64)
//func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
//交换操作
//func SwapInt32(addr *int32, new int32) (old int32)
//func SwapInt64(addr *int64, new int64) (old int64)
//func SwapUint32(addr *uint32, new uint32) (old uint32)
//func SwapUint64(addr *uint64, new uint64) (old uint64)
//func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
//func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
//比较并交换操作
//func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
//func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
//func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
//func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
//func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
//func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)


var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func add() {
	// x = x + 1
	x++ // 等价于上面的操作
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func Test_atomic_01(T *testing.T) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}