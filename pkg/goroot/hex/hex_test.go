package hex

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"
)


// hex(16进制)：0、1、2、3、4、5、6、7、8、9、A、B、C、D、E、F。其中A-F分别表示十进制数字10-15

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




func TestName1(t *testing.T) {
	str :="abcdefijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"

	b, err := hex.DecodeString(str)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Decoded bytes %v	", b)
}
