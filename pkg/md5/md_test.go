package md_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

// md5不是一个可逆算法
func TestMd1(t *testing.T) {
	hasher := md5.New()
	io.WriteString(hasher, "111")
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
}