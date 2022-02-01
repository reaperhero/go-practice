package avl

import (
	"log"
	"testing"
)

func TestNewAVLTreeRoot(t *testing.T) {
	array := []int{5, 3, 1, 8, 9, 10, 11, 2, 4, 7, 6, 12}
	var root *AVLTreeNode
	for _, v := range array {
		root = root.InsertNode(v)
		log.Println("root=> key:", root.value, "high:", root.high)
	}
	root.PrintSortTree()
}
