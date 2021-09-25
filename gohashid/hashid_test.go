package gohashid

import (
	"log"
	"testing"
)

func TestEncodeHashID(t *testing.T) {
	result, _ := HashIDEncodeBySalt(1)
	num, _ := HashIDDecodeBySalt(result)
	log.Println(num)
}
