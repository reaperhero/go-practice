package xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)


//XMLName 字段，因为前面提到的原因，会被忽略
//带有 “-” 标签的字段会被忽略
//带有 “name,attr” 标签的字段会成为 XML 元素的属性， 其中属性的名字为这里给定的 name
//带有 ”,attr” 标签的字段会成为 XML 元素的属性， 其中属性的名字为字段的名字
//带有 ”,chardata” 标签的字段将会被封装为字符数据而不是 XML 元素。
//带有 ”,cdata” 标签的字段将会被封装为字符数据而不是 XML 元素， 并且这些数据还会被一个或多个 <![CDATA[ ... ]]> 标签包围。
//带有 ”,innerxml” 标签的字段无需进行任何封装， 它会以原样进行输出。
//带有 ”,comment” 标签的字段无需进行任何封装， 它会直接输出为 XML 注释。 这个字段内部不能包含 “–” 字符串。
//如果字段的标签中包含 “omitempty” 选项， 那么在字段的值为空时， 这个字段将被忽略。 空值指的是 false ， 0 ，为 nil 的指针、接口值、数组、切片、map ，以及长度为 0 的字符串。
//匿名结构字段会被看作是外层结构的其中一部分来处理。



type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func TestXml(t *testing.T) {
	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}

	//序列化到文件中
	encoder := xml.NewEncoder(os.Stdout)
	err := encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
	}
}
