package hex

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestHex(t *testing.T) {
	str := "ff68b4ff"
	b, _ := hex.DecodeString(str)
	encodedStr := hex.EncodeToString(b)
	fmt.Printf("@@@@--bytes-->%02x \n", b)
	fmt.Printf("@@@@--string-->%s \n", encodedStr)
}

func TestName(t *testing.T) {
	hexStringData := hex.EncodeToString([]byte(`测试数据`)) // 将 byte 装换为 16进制的字符串
	println(hexStringData)                                // e6b58be8af95e695b0e68dae
	fmt.Printf("%x", []byte(`测试数据`))             // 字节数组转16进制可以直接使用 用fmt就能转

	// 将 16进制的字符串 转换 byte
	hexData, _ := hex.DecodeString(hexStringData)
	println(string(hexData)) // 测试数据
}


// 16进制的字符串手动转换成字节数组
func Hextob(str string)([]byte) {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1;
			slen = slen - 2;
		}
	}
	return bHex
}