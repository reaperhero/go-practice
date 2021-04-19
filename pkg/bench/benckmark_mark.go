package bench

import (
	"math/rand"
	"testing"
	"time"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}


func BenchmarkFibTimer(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer() // 重置定时器
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}


func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()  // 使用 StopTimer 暂停计时
		nums := generateWithCap(10000)
		b.StartTimer() // 使用 StartTimer 开始计时
		bubbleSort(nums)
	}
}