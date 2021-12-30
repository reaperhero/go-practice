package md_test

import (
	"crypto/md5"
	"encoding/base64"
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

func TestBase64(t *testing.T) {
	decod, _ := base64.StdEncoding.DecodeString("aaad12334asddsad")
	fmt.Println(decod)
	encode:=base64.StdEncoding.EncodeToString(decod)
	fmt.Println(encode)
}
