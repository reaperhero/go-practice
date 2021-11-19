package 插入排序

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	time := time.Now()
	var arr = []int{19, 13, 27, 15, 3, 4, 26, 12, 1, 0}
	insertSort(arr)
	fmt.Println(time.Sub(time))
	fmt.Println("insertSort:", arr)
}

func insertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
