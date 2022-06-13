package main

import (
	"fmt"
	"testing"
)

func TestLdapCon(t *testing.T)  {
	l := NewLdapConnManagent()
	l.ListUser()
	fmt.Println("--------")
	l.ListGroup()
}
