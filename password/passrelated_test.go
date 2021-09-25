package password

import (
	"log"
	"testing"
)

func Test_GeneratePassword(t *testing.T) {
	log.Println(GeneratePassword(9))
}


func Test_CheckPassword(t *testing.T) {
	log.Println(CheckPassword("a23bcdAB!@#$^%&*()"))
}
