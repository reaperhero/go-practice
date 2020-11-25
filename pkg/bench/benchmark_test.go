package bench

import "testing"

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithPreAlloc()
	}
}

//> go test -bench=.
//goos: darwin
//goarch: amd64
//pkg: go-example-demo/pkg/bench
//BenchmarkMakeSliceWithoutAlloc-12           3660            334327 ns/op
//BenchmarkMakeSliceWithPreAlloc-12          12146             92262 ns/op
//PASS
//ok      go-example-demo/pkg/bench       5.692s