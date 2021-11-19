package 堆排序

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	array := []int{52, 16, 37, 2, 3, 32, 12, 27, 19, 42, 29, 13, 37, 12, 9}
	HeapSort(array)
	fmt.Println("HeapSort:", array)
}

func HeapSort(array []int) {
	m := len(array)           // 15
	s := m / 2                // 7
	for i := s; i > -1; i-- { // i:= 7 i > -1 i--
		heap(array, i, m-1)   // 7,14
	}
	for i := m - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}
}
func heap(array []int, i, end int) {
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && array[r] > array[l] {
		n = r
	}
	if array[i] > array[n] {
		return
	}
	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}
