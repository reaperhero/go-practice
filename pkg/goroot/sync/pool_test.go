package sync

import (
	"encoding/json"
	"sync"
	"testing"
)


// sync.Pool 保存和复用临时对象，减少内存分配，降低 GC 压力。
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "chenqiangjun", Age: 25})

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
