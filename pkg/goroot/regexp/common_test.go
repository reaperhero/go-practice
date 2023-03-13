package regexp

import (
	"fmt"
	"log"
	"regexp"
	"testing"
)

func TestRegxIp(t *testing.T) {
	str := "1.2.3"
	ipReg := `^((0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])\.){3}(0|[1-9]\d?|1\d\d|2[0-4]\d|25[0-5])$`
	r, _ := regexp.Compile(ipReg)
	match := r.MatchString(str)
	if match {
		fmt.Printf("%s is a legal ipv4 address\n", str)
	} else {
		fmt.Printf("%s is not a legal ipv4 address\n", str)
	}
}

func TestExpr(t *testing.T)  {
	expr := "floor(delta(jvm_gc_collection_seconds_count(java_lang_OperatingSystem_OpenFileDescriptorCount{product_name=\\\"DTApi\\\",service_name=\\\"Api\\\"}/java_lang_OperatingSystem_MaxFileDescriptorCount{product_name=\\\"DTApi\\\",service_name=\\\"Api\\\"}"
	g,err := regexp.Compile("\\((.*?){")
	if err!=nil{
		log.Fatal(err)
	}
	list := g.FindStringSubmatch(expr)
	fmt.Println(list[1])
}
