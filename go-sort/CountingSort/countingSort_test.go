package CountingSort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	array := []int{69,16,48,2,3,32,10,27,17,42,29,8,28,12,9,}
	countingSort(array,array[0])
	fmt.Println("BucketSort:",array)
}
func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	sortedIndex := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return arr
}