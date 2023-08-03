package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`     //该XML文件的根元素为person
	Id        int      `xml:"id,attr"`    //该值会作为person元素的属性
	FirstName string   `xml:"name>first"` //first为name的子元素
	LastName  string   `xml:"name>last"`  //last
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"` //含omitempty选项的字段如果为空值会省略
	Married   bool     //默认为false
	Address            //匿名字段（其标签无效）会被处理为其字段是外层结构体的字段，所以没有Address这个元素，而是直接显示City, State这两个元素
	Comment   string   `xml:",comment"` //注释
}

func TestNameEnocde(t *testing.T) {

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
}
