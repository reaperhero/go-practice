package MergeSort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	array := []int{55, 94, 87, 12, 4, 32, 11, 8, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println("before MergeSort", array)
	array = MergeSort(array)
	fmt.Println("after MergeSort:", array)
}

func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}
func merge(left []int, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
