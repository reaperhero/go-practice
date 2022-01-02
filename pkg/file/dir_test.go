package file

import (
	"fmt"
	"strings"
	"testing"
)

func TestDirCheck(t *testing.T) {
	bo := strings.ContainsAny("/data/easymanager/", " #$%^&*()_+}{\":?><\"⌘")

	fmt.Println(bo)
}
