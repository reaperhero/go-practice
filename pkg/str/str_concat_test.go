package str

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// +
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// printf
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// build
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func builderGrowConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))  // string.Builder 也提供了预分配内存的方式 Grow
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// buffer
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

// byte
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// pre buffer
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}

func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }

//strings.Builder、bytes.Buffer 和 []byte 的性能差距不大
//性能最好且消耗内存最小的是 preByteConcat，这种方式预分配了内存，在字符串拼接的过程中，不需要进行字符串的拷贝，也不需要分配新的内存，因此性能最好，且内存消耗最小
//BenchmarkPlusConcat-12                32          36388704 ns/op        530999197 B/op     10038 allocs/op
//BenchmarkSprintfConcat-12             18          65778852 ns/op        834132716 B/op     37468 allocs/op
//BenchmarkBuilderConcat-12          12040             91288 ns/op          522226 B/op         23 allocs/op
//BenchmarkBufferConcat-12           12217            103099 ns/op          423538 B/op         13 allocs/op
//BenchmarkByteConcat-12             12608             94052 ns/op          628723 B/op         24 allocs/op
//BenchmarkPreByteConcat-12          22354             49220 ns/op          212992 B/op          2 allocs/op
