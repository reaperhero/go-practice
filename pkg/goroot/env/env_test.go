package env

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	// 获取所有环境变量
	for _, env := range os.Environ() {
		t := strings.SplitN(env, "=", 2)
		fmt.Println(t[0] + "=" + t[1])
	}
}
