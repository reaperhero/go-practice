package template

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

type Recipient struct {
	Name    string
	Friends []string
}

func (r Recipient) GetName(src string) string {
	return src
}

func (r Recipient) ChangeName() string  {
	return "New" + r.Name
}


func TestTTplFunc(t *testing.T) {
	const templateText = `
    Nrd   Friend : {{index .Friends (sub (len .Friends) 1)}}  # 获取Friends最后一个值
    Nrd   Friend : {{.GetName .ChangeName}}  # 后面的函数会先运行,值作为前一个函数的参数
    Nrd   Friend : {{.ChangeName | .GetName }}  # 上面的也可以这么写

`

	templateFunc := map[string]interface{}{
		"sub": func(a, b int) int { return a - b },
	}

	recipient := Recipient{
		Name:    "Jack",
		Friends: []string{"Bob", "Json", "Tom"},
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap(templateFunc)).Parse(templateText))
	err := tpl.Execute(os.Stdout, recipient)
	if err != nil {
		fmt.Println("Executing template:", err)
	}
}
