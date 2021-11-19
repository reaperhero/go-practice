package 冒泡排序_

import (
	"fmt"
	"testing"
)

func Test_bubble_sort(t *testing.T) {
	var arr = []int{9, 10, 11, 5, 3, 4, 27, 2, 1, 3, 20}
	//升序
	bubbleAscendingSort(arr)
	//降序
	bubbleDescendingSort(arr)
}

//升序
func bubbleAscendingSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("bubbleAscendingSort:", arr)
}

//降序
func bubbleDescendingSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("bubbleDescendingSort:", arr)
}
