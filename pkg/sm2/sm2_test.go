package sm2

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"log"
	"testing"
)

func TestSm2(t *testing.T) {
	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	fmt.Printf("[sm2 PubliKey]: 04%x%x\n", priv.PublicKey.X, priv.PublicKey.Y)
	fmt.Printf("[sm2 PrivateKey]: %x\n", priv.D)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("Tongji Fintech Research Institute")
	pub := &priv.PublicKey
	ciphertxt, err := pub.EncryptAsn1(msg,rand.Reader) //sm2加密
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("加密结果:%x\n",ciphertxt)
	plaintxt,err :=  priv.DecryptAsn1(ciphertxt)  //sm2解密
	if err != nil {
		log.Fatal(err)
	}
	if !bytes.Equal(msg,plaintxt){
		log.Fatal("原文不匹配")
	}

	sign,err := priv.Sign(rand.Reader, msg, nil)  //sm2签名
	if err != nil {
		log.Fatal(err)
	}
	isok := pub.Verify(msg, sign)    //sm2验签
	fmt.Printf("Verified: %v\n", isok)

}
