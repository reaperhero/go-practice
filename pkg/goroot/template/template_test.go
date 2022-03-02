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
	return v.Value.(string)
}

func TestTemplate(t *testing.T) {
	tpl, err := template.ParseFiles("./test.html")
	if err != nil {
		return
	}

	v := VisualConfig{
		Default: "48",
		Desc:    "sads",
		Type:    "dsada",
		Value:   "90",
	}
	buf := &bytes.Buffer{}
	m := map[string]interface{}{
		"Name": &v,
	}
	if err = tpl.Option("missingkey=error").Execute(buf, m); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}

func TestString(t *testing.T) {
	var v interface{} = &VisualConfig{
		Default: "48",
		Desc:    "sads",
		Type:    "dsada",
		Value:   "90",
	}
	fmt.Printf("%s",v)
}