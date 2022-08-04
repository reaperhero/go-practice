package template

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

type VisualConfig struct {
	// must sort by field!!!
	Default interface{} `config:"default"`
	Desc    string      `config:"desc" validate:"required"`
	Type    string      `config:"type" validate:"required"`
	Value   interface{} `config:",ignore"`
}

func (v *VisualConfig) String() string {
	if v.Value.(string) != "" {
		return v.Value.(string)
	}
	return v.Default.(string)
}

func TestTemplate(t *testing.T) {
	var templateFunc = map[string]interface{}{
		"add": func() int {
			return 1 + 2
		},
	}
	tpl := template.Must(template.New("").Funcs(template.FuncMap(templateFunc)).Parse(`
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Go Web</title>
</head>
{{add}}
<body>
{{ .Name }}
</body>
</html>
`))

	v := VisualConfig{
		Default: "48",
		Desc:    "sads",
		Type:    "dsada",
		Value:   "",
	}
	buf := &bytes.Buffer{}
	m := map[string]interface{}{
		"Name": &v,
	}
	if err := tpl.Option("missingkey=error").Execute(buf, m); err != nil { // 会调用String方法
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}
