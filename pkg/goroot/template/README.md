# readme


模板语法

内置函数：[hdr-Functions](https://pkg.go.dev/text/template#hdr-Functions)

参考地址：[gotpl](https://blog.gmem.cc/gotpl)

```
// var letter = "{{}}"
// result := template.Must(template.New("").Parse(letter))
// err := result.Execute(os.Stdout, r)

// // t := template.Must(template.New("").Funcs(template.FuncMap(templateFunc)).Parse(templateText))

```

- 值
```
type Recipient struct {
		Name       string
		Attended   bool
		Friends    []string
        Customer struct {
            Addr string
        }
}
```
- template
```
{{ .Customer.Addr }}  // 取值
{{index .Friends 0}}

{{- range $key, $val := .Friends -}}
    {{$key}} = {{ $val }}
{{ end -}}

{ {range .Items}}
  <div class="item">
    <h3 class="name">{{.Name}}</h3>
    <span class="price">${{.Price}}</span>
  </div>
{ {end}}

{{if .Attended}}     // bool
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}



{{with .Gift -}}   // 是否为默认值
Thank you for the lovely {{.}}.s
{{end}}

```
