package _select

import (
	"fmt"
	"testing"
	"time"
)

func breakSelect() {

	for {
	ee:
		select {
		case <-time.After(time.Second * 3):
			for i := 0; i < 10; i++ {
				if i == 11{
					break ee
				}
				fmt.Println(i)
			}
			fmt.Println(111)
			return
		}
	}

}
func TestName(t *testing.T) {
	breakSelect()
}
