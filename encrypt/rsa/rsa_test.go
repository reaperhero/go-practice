package rsa

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//rsa.GenerateKey()
	err := RSAGenKey(4096)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("秘钥生成成功！")
	str := "山重水复疑无路，柳暗花明又一村！"
	fmt.Println("加密之前的数据为：", string(str))
	data, err := EncyptogRSA([]byte(str), "publicKey.pem")
	data, err = DecrptogRSA(data, "privateKey.pem")
	fmt.Println("加密之后的数据为：", string(data))
}
