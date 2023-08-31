package client2

import (
	"fmt"
	i "github.com/reaperhero/go-practice/copier/interface/interface"
)

type ClientTwo struct {
	I i.UseCase
}

type ClientTwoImpl struct {
}

func (c ClientTwoImpl) Echo(name string) string {
	fmt.Println("client 2")
	return "client02"
}
