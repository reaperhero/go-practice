package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/reaperhero/go-practice/copier/interface/client1"
	"github.com/reaperhero/go-practice/copier/interface/client2"
)

func main() {
	c1 := client1.ClientOne{I: client1.ClientOneImpl{}}
	c2 := client2.ClientTwo{}
	c1.I.Echo("")
	err := copier.Copy(&c2, &c1)
	if err != nil {
		fmt.Println(err)
	}
	c2.I.Echo("")
}
