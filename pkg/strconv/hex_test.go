package stand

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHex(t *testing.T)  {
	bs, _ := hex.DecodeString("0fa8")
	num := binary.BigEndian.Uint16(bs[:2])
	fmt.Println(num)
}

