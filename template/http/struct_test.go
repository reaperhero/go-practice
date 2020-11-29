package http

import (
	"html/template"
	"net/http"
	"os"
	"testing"
	"time"
)

// 在template中，点"."代表当前作用域的当前对象,表示顶级作用域
// range、with、if等内置action都有自己的本地作用域
// "- "来去除它前面的空白(包括换行符、制表符、空格等)
// with定义新的作用域


type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Test_Struct_01(t *testing.T) {
	tmpl := template.Must(template.ParseFiles("struct_test.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":85", nil)
	time.Sleep(100 * time.Second)
}

func Test_Struct_02(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"longshuai", 23}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}

func Test_Struct_03(t *testing.T) {
	type Friend struct {
		Fname string
	}
	type Person struct {
		UserName string
		Emails   []string
		Friends  []*Friend
	}
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}

	tmpl, _ := template.New("test").Parse(`hello {{.UserName}}!
{{ range .Emails }}
	an email {{ . }}
{{- end }}

{{ with .Friends }}
	{{- range . }}
		my friend name is {{.Fname}}
	{{- end }}
{{ end }}`)
	p := Person{
		UserName: "chenqiangjun",
		Emails:   []string{"a1@qq.com", "a2@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	tmpl.Execute(os.Stdout, p)
}
