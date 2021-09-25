package aes

import (
	"log"
	"testing"
)

func TestAes(t *testing.T) {
	// 加密
	var originData = `{"aaa": "bbb","ccc": "ddd"}`
	var key = "9D0988CE-98B5-45B3-91F7-AD7F775AD81B"

	cipher, err := GCMEncrypt(originData, key)
	if err != nil {
		t.Fatal("encrypt error:", err)
	}
	// 解密
	result, err := GCMDecrypt(cipher, key)
	log.Println(string(result)) // {"aaa": "bbb","ccc": "ddd"}
}

func TestCBCEncrypt(t *testing.T) {
	key := []byte("9wxh7u2p6agc8lmb0nst5qf34zkemvd9")
	src := []byte("dasadad")
	mod := 16
	// 加密
	cbcEncode, _ := AESCBCEncrypt(key, src, mod)

	// 解密
	result, _ := AESCBCDecrypt(key, cbcEncode, mod)
	log.Println(string(result)) // dasadad
}
