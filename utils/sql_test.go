package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNAame(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		return
	}
	for _, v := range []string{"2022122914sda0200_v5.3.x.sql"} {
		if bytes.Contains(dataBytes, []byte(v)) {
			continue
		}

		fmt.Println(1)
		dataBytes = append(dataBytes, []byte(v+"\n")...)
	}
	err = ioutil.WriteFile("./1.txt", dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
