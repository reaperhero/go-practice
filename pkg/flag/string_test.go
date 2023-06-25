package flag

import (
	"flag"
	"fmt"
	"testing"
)

var (
	serviceName = flag.String("name", "m  yun", "")
	cliAge      = flag.Int("age", 28, "Input Your Age")
	boolflag    = flag.Bool("boolflag", false, "bool flag value")
)

// go test -v -name="sadsada"
func Test_flag_string_01(t *testing.T) {
	flag.Parse()
	fmt.Println(*serviceName)
	last := flag.Arg(0) // 除了识别以为的第一个值
	fmt.Println(last)
}
