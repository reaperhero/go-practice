package template

import (
	"html/template"
	"os"
	"testing"
)

// 有两个常用的传入参数的类型。
//
// 1、一个是struct，在模板内可以读取该struct域的内容来进行渲染。
// type Article struct {
//    ArticleId int
//    ArticleContent string
//}
//<p>{{.ArticleContent}}<span>{{.ArticleId}}</span></p>
//
// 2、一个是map[string]interface{}，在模板内可以使用key来进行渲染。

// 变量
// {{$article := "hello"}}            定义一个article变量，将其初始化为”hello”，
// {{$article := .ArticleContent}}    传入值的内容赋值给article
// {{$article}} 					  使用变量

// 函数
// {{funcname .arg1 .arg2}}
// {{add 1 2}}                        传递func add(left int, right int) int

// 判断
// {{if .condition1}}
// {{else if .contition2}}
// {{end}}

func TestStructTtmpalte(t *testing.T) {
	type person struct {
		Id      int
		Name    string
		Country string
	}
	tmpl := template.New("tmpl")
	tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	tmpl.Execute(os.Stdout, person{Id: 1001, Name: "liumiaocn", Country: "China"})

	// 定义struct模版参数
	type User struct {
		Name string
	}

	type Order struct {
		Id       int
		Title    string
		Customer User //嵌套字段
	}

	// 初始化模版参数
	food := Order{
		Id:    10,
		Title: "柠檬",
		Customer: User{
			Name: "李大春",
		},
	}
	tmpl = template.New("tmpl")
	tmpl.Parse("商品名: {{.Title}}用户名: {{.Customer.Name}}\n")
	tmpl.Execute(os.Stdout, food)
}
