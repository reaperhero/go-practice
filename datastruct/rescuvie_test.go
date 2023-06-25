package datastruct

import (
	"fmt"
	"testing"
)

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuvie(n-1)
}

func RescuvieTail(n int, a int) int {
	if n == 1 {
		return a
	}

	return RescuvieTail(n-1, a*n)
}

func TestRescuvieFront(t *testing.T) {
	//{5 * Rescuvie(4)}
	//{5 * {4 * Rescuvie(3)}}
	//{5 * {4 * {3 * Rescuvie(2)}}}
	//{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
	//{5 * {4 * {3 * {2 * 1}}}}
	//{5 * {4 * {3 * 2}}}
	//{5 * {4 * 6}}
	//{5 * 24}
	//120
	fmt.Println(Rescuvie(5))
}

func TestRescuvieEnd(t *testing.T) {
	//RescuvieTail(5, 1)
	//RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
	//RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
	//RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
	//RescuvieTail(1, 60*2)=RescuvieTail(1, 120)
	//120
	fmt.Println(RescuvieTail(5, 1))
}

