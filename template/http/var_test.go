package http

import (
	"html/template"
	"os"
	"testing"
)

// 44 333 444
// 55 333 444
func Test_var_01(t *testing.T) {
	tx := template.Must(template.New("hh").Parse(
		`{{range $x := . -}}
{{$y := 333}}
{{- if (gt $x 33)}}{{println $x $y ($z := 444)}}{{- end}}
{{- end}}
`))
	s := []int{11, 22, 33, 44, 55}
	_ = tx.Execute(os.Stdout, s)
}

// 条件判断

// range...end迭代,可以迭代slice、数组、map或channel
// {{range pipeline}} T1 {{end}}
// {{range pipeline}} T1 {{else}} T0 {{end}}
func Test_range_02(t *testing.T) {
	tx := template.Must(template.New("hh").Parse(
		`{{range $x := . -}}
{{println $x}}
{{- end}}
`))
	s := []int{11, 22, 33, 44, 55}
	_ = tx.Execute(os.Stdout, s)
}
