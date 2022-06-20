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

	filepath.Abs("./log.txt")        // 输出绝对路径 F:\my\bin\log.txt
	filepath.Base("./log.txt")       // 输出 log.txt
	filepath.Ext("./a/b/c/d.jpg")    // 输出 .jpg
	filepath.Split("C:/a/b/c/d.jpg") // 输入出 C:/a/b/c/   d.jpg
}
