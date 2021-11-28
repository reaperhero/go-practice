package RadixSort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	array := []int{12, 3, 8, 5, 9, 11, 23, 36, 20, 28, 21}
	fmt.Println("before radixSort:", array)
	radixSort(array)
	fmt.Println("after radixSort:", array)
}

//获取数组的最大值
func maxValue(arr []int) (ret int) {
	ret = 1
	var key int = 10
	for i := 0; i < len(arr); i++ {
		for arr[i] >= key {
			key = key * 10
			ret++
		}
	}
	return
}
func radixSort(arr []int) {
	key := maxValue(arr)
	tmp := make([]int, len(arr), len(arr))
	count := new([10]int)
	radix := 1
	var i, j, k int
	for i = 0; i < key; i++ { //进行key次排序
		for j = 0; j < 10; j++ {
			count[j] = 0
		}
		for j = 0; j < len(arr); j++ {
			k = (arr[j] / radix) % 10
			count[k]++
		}
		for j = 1; j < 10; j++ { //将tmp中的为准依次分配给每个桶
			count[j] = count[j-1] + count[j]
		}
		for j = len(arr) - 1; j >= 0; j-- {
			k = (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}
		for j = 0; j < len(arr); j++ {
			arr[j] = tmp[j]
		}
		radix = radix * 10
	}
}
