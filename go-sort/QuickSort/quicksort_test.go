package QuickSort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	quickAscendingSort(arr, 0, len(arr)-1)
	fmt.Println("quickAscendingSort:", arr)
	quickDescendingSort(arr, 0, len(arr)-1)
	fmt.Println("quickDescendingSort:", arr)
}

//升序
func quickAscendingSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickAscendingSort(arr, start, j)
		}
		if end > i {
			quickAscendingSort(arr, i, end)
		}
	}
}

//降序
func quickDescendingSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] > key {
				i++
			}
			for arr[j] < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickDescendingSort(arr, start, j)
		}
		if end > i {
			quickDescendingSort(arr, i, end)
		}
	}
}
