package math

import (
	"math/big"
	"strings"
	"testing"
)



func Trunc(a string, prec int) string {
	newn := strings.Split(a, ".")

	if prec <= 0 {
		return newn[0]
	}

	if len(newn) < 2 || prec >= len(newn[1]) {
		return a
	}

	return newn[0] + "." + newn[1][:prec]
}

func TestName(t *testing.T)  {

	a := new(big.Float).SetFloat64(2.0)
	b := new(big.Float).SetFloat64(3.0)
	c := new(big.Float).Quo(a, b)
	println(Trunc(c.Text('g', 512), 1))
}
