package reentrantMutex

import (
	"fmt"
	"testing"
)


func TestName(t *testing.T) {
	var mutex = NewReentrantLock()
	mutex.Lock()
	mutex.Lock()
	fmt.Println(111)
	mutex.Unlock()
	mutex.Unlock()
}
