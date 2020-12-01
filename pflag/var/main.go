package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"strings"
)

var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = flag.IntP("age", "a", 22, "Input Your Age")
var cliGender = flag.StringP("gender", "g", "male", "Input Your Gender")
var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {

	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc) // 设置标准化参数名称的函数

	// go run main.go  -age  不设置时候默认值
	flag.Lookup("age").NoOptDefVal = "25"

	flag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}
