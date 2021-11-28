package ShellSort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	var length = 15
	var list []int

	// 以时间戳为种子生成随机数，保证每次运行数据不重复
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	fmt.Println(list)

	// 这里就以n/2为增量z
	gap := 2
	step := length / gap // 7

	for step >= 1 {
		// 这里按步长开始每个分组的排序
		for i := step; i < length; i++ {
			// 将按步长分组的子队列用直接插入排序算法进行排序
			insertionSortByStep(list, step)
		}
		// 完成一轮后再次缩小增量
		step /= gap

		// 输出每轮缩小增量各组排序后的结果
		fmt.Println(list)
	}
}

// 这里把上篇直接选择排序的算法抽出来，并将步长从1改成step
func insertionSortByStep(tree []int, step int) { // 7...15
	for i := step; i < len(tree); i++ {
		for j := i; j >= step && tree[j] < tree[j-step]; j -= step {
			tree[j], tree[j-step] = tree[j-step], tree[j]
		}
	}
}
