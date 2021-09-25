package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)


const (
	nonceLength       = 16
	nonceBase64Length = 24
)

// AES 加密
func GCMEncrypt(origData, key string) (string, error) {
	keyBytes, err := SHA1PRNG([]byte(key), 128)
	if err != nil {
		return "", fmt.Errorf("SHA1PRNG process key error: %v", err)
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("new cipher error. key: %s, error: %v", key, err)
	}
	aead, err := cipher.NewGCMWithNonceSize(block, nonceLength)
	if err != nil {
		return "", fmt.Errorf("new GCM error. error: %v", err)
	}
	iv := getIV(nonceLength)
	cipherWithoutIVByte := aead.Seal(nil, iv, []byte(origData), nil)
	result := base64.StdEncoding.EncodeToString(iv) + base64.StdEncoding.EncodeToString(cipherWithoutIVByte)
	return result, nil
}

// AES 解密
func GCMDecrypt(cipherStr, key string) ([]byte, error) {
	// get iv and content
	iv, err := base64.StdEncoding.DecodeString(cipherStr[:nonceBase64Length])
	if err != nil {
		return nil, fmt.Errorf("base64 decode iv error: %v", err)
	}
	cipherWithoutIVByte, err := base64.StdEncoding.DecodeString(cipherStr[nonceBase64Length:])
	if err != nil {
		return nil, fmt.Errorf("base64 decode cipher error: %v", err)
	}
	// get key
	keyBytes, err := SHA1PRNG([]byte(key), 128)
	if err != nil {
		return nil, fmt.Errorf("SHA1PRNG process key error: %v", err)
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("new cipher error: %v", err)
	}
	aead, err := cipher.NewGCMWithNonceSize(block, nonceLength)
	if err != nil {
		return nil, fmt.Errorf("new GCM error: %v", err)
	}
	result, err := aead.Open(nil, iv, cipherWithoutIVByte, nil)
	if err != nil {
		return nil, fmt.Errorf("aead open error: %v", err)
	}
	return result, nil
}

// cbc 加密
func AESCBCEncrypt(key, src []byte, mod int) (data []byte, err error) {
	var (
		block cipher.Block
	)
	if block, err = aes.NewCipher(key[:mod]); err != nil {
		return data, err
	}
	origData := PKCS7Padding(src, block.BlockSize())
	cipherText := make([]byte, len(origData))
	cbc := cipher.NewCBCEncrypter(block, key[mod:])
	cbc.CryptBlocks(cipherText, origData)
	return cipherText, nil
}

// cbc 解密
func AESCBCDecrypt(key, src []byte, mod int) (data []byte, err error) {
	var (
		block cipher.Block
	)
	if block, err = aes.NewCipher(key[:mod]); err != nil {
		return data, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key[mod:])
	cipherText := make([]byte, len(src))
	blockModel.CryptBlocks(cipherText, src)
	cipherText = PKCS7UnPadding(cipherText)
	return cipherText, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}

func getIV(length int) []byte {
	iv := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		iv = append(iv, byte(r.Intn(127)))
	}
	return iv
}

func SHA1PRNG(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := SHA1(SHA1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, fmt.Errorf("Not Support %d, Only Support Lower then %d [% x]", realLen, maxLen, hashs)
	}

	return hashs[0:realLen], nil
}

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
