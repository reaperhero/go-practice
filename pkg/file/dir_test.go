package file

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestDirCheck(t *testing.T) {
	strings.ContainsAny("/data/easymanager/", " #$%^&*()_+}{\":?><\"⌘")
	fmt.Println(filepath.IsAbs("data/easymanager/"))
}

