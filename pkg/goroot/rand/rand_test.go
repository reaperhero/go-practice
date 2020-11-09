package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(1000)) // 随机数
}