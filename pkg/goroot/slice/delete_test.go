package slice

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 1, 5, 1}
	for i := 0; i < len(a); i++ {
		if a[i] > 3 {
			a = append(a[:i], a[i+1:]...) // [1 2 3 1 1]
			i--
		}
	}
	fmt.Println(a)
}

func TestDelete1(t *testing.T) {
	a := []int{1, 2, 3, 4, 2, 5, 2}
	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	fmt.Println(a)
}
