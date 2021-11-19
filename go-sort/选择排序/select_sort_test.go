package 选择排序

import (
	"fmt"
	"testing"
)

func Test_select_sort(t *testing.T) {
	var arr = []int{19,28,17,5,13,4,6,7,9,3,10}
	//升序
	selectAscendingSort(arr)
	//降序
	selectDescendingSort(arr)
}

//升序
func selectAscendingSort(arr []int) {
	l := len(arr)
	m := len(arr) - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i+1; j < l; j++ {
			if arr[k] > arr[j] {
				k = j
			}
		}
		if k != i {
			arr[k],arr[i] = arr[i],arr[k]
		}
	}
	fmt.Println("selectAscendingSort:",arr)
}

//降序
func selectDescendingSort(arr []int) {
	l := len(arr)
	m := len(arr) - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i+1; j < l; j++ {
			if arr[k] < arr[j] {
				k = j
			}
		}
		if k != i {
			arr[k],arr[i] = arr[i],arr[k]
		}
	}
	fmt.Println("selectDescendingSort:",arr)
}