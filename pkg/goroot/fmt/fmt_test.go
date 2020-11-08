package fmt

import (
	"fmt"
	"os"
	"testing"
)

// Fprint
// 输出到一个io.Writer接口类型的变量w中
func Test_fmt_01(t *testing.T) {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}

// Sprint
// Sprint系列函数会把传入的数据生成并返回一个字符串。
func Test_Sprint_01(t *testing.T) {
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}

type Website struct {
	Name string
}

var site = Website{Name: "studygolang"}

// Printf
func Test_print01(t *testing.T) {
	//普通占位符
	fmt.Printf("%v", site)  // 相应值的默认格式。 {studygolang}
	fmt.Printf("%+v", site) // 在打印结构体时，“加号”标记（%+v）会添加字段名 {Name:studygolang}
	fmt.Printf("%#v", site) // 相应值的Go语法表示         main.Website{Name:"studygolang"}
	fmt.Printf("%T", site)  // 相应值的类型的Go语法表示 main.Website

	//布尔占位符
	fmt.Printf("%t", true) //  true

	// 整数占位符
	fmt.Printf("%b", 5)      // 二进制表示  101
	fmt.Printf("%c", 0x4E2D) // 相应Unicode码点所表示的字符  中
	fmt.Printf("%d", 0x12)   // 十进制表示   18
	fmt.Printf("%d", 10)     // 八进制表示   12
	fmt.Printf("%q", 0x4E2D) // 单引号围绕的字符字面值，由Go语法安全地转义  '中'
	fmt.Printf("%x", 13)     //  十六进制表示，字母形式为小写 a-f       d
	fmt.Printf("%x", 13)     // 十六进制表示，字母形式为大写 A-F         D
	fmt.Printf("%U", 0x4E2D) // Unicode格式：U+1234，等同于 "U+%04X"        U+4E2D

	// 字符串与字节切片
	fmt.Printf("%s", []byte("Go语言中文网")) // 输出字符串表示（string类型或[]byte) Go语言中文网
	fmt.Printf("%q", "Go语言中文网")         //  双引号围绕的字符串，由Go语法安全地转义   "Go语言中文网"
	fmt.Printf("%x", "golang")          // 十六进制，小写字母，每字节两个字符   676f6c616e67
	fmt.Printf("%X", "golang")          // 十六进制，大写字母，每字节两个字符    676F6C616E67

	//指针
	fmt.Printf("%p", &site)  //十六进制表示，前缀 0x  0x4f57f0
	fmt.Printf("%#p", &site) //十六进制表示，不带0x的指针  4f57f0

}

// Scan
func Test_scan_01(t *testing.T) {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
