package client1

import (
	"fmt"
	i "github.com/reaperhero/go-practice/copier/interface/interface"
)

type ClientOne struct {
	I i.UseCase
}

type ClientOneImpl struct {
}

func NewClientOneImpl() i.UseCase {
	return ClientOneImpl{}
}

func (c ClientOneImpl) Echo(name string) string {
	fmt.Println("client 1")
	return "client01"
}
