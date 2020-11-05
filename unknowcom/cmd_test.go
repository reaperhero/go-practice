package unknwon

import (
	"strings"
	"testing"

	"github.com/unknwon/com"
)

func Test_cmd_01(t *testing.T) {
	stdout, stderr, err := com.ExecCmd("go", "help", "get")
	if err != nil {
		t.Errorf("ExecCmd:\n Expect => %v\n Got => %v\n", nil, err)
	} else if len(stderr) != 0 {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "", stderr)
	} else if !strings.HasPrefix(stdout, "usage: go get") {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "usage: go get", stdout)
	}
}
