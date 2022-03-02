package file_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func copyDir(src string, file string) {
	var (
		dest = "C:/Users/chenqiangjun/Desktop/doc/" + file
	)
	src = strings.ReplaceAll(src, "\\", "/")
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("打开源文件错误，错误信息=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("打开目标文件错误，错误信息=%v\n", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	_, err = io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
	}

}

func getSuffixFile(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	ok := strings.HasSuffix(path, ".pdf")
	if ok {
		fileName := strings.Split(path, "\\")
		copyDir(path, fileName[len(fileName)-1])
	}
	return nil
}

func TestFileCopy(t *testing.T) {
	path := "C:/Users/chenqiangjun/Documents/2021年11月高项VIP汇总地址/01第一阶段-精讲打基础"
	err := filepath.Walk(path, getSuffixFile)
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
