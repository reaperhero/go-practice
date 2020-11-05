package flag

import (
	"flag"
	"fmt"
	"testing"
)

var (
	serviceName = flag.String("name", "myun", "")
)

// go test -v -name="sadsada"
func Test_flag_string_01(t *testing.T) {
	flag.Parse()
	fmt.Println(*serviceName)
}
