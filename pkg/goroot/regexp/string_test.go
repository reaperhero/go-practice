package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

// http://c.biancheng.net/view/5124.html

func Test_regexp_string(t *testing.T) {
	buf := "abc azc a7c aac 888 a9c  tac"
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func Test_regexp_html(t *testing.T) {
	// 原生字符串
	buf := `
    
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>C语言中文网 | Go语言入门教程</title>
</head>
<body>
    <div>Go语言简介</div>
    <div>Go语言基本语法
    Go语言变量的声明
    Go语言教程简明版
    </div>
    <div>Go语言容器</div>
    <div>Go语言函数</div>
</body>
</html>
    `
	//解释正则表达式
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	//过滤<></>
	fmt.Println(result)
}

func TestArgs(t *testing.T) {
	arg := "${args}"
	reg1 := regexp.MustCompile(`\${\w+}`)
	reg2 := regexp.MustCompile(`\w+`)
	vaule := reg1.FindString(arg)
	k := reg2.FindString(vaule)
	fmt.Println(k)
}

func TestRexReplace(t *testing.T)  {
	var ts = []string{
		"AA && \\ 		\n",
		"AA && aa \n",
		"AA && a ",
		" && a ",
		"&& a ",
	}
	re3, _ := regexp.Compile(`&&[ ]+\\[ ]+`)

	for _, s := range ts {
		a := re3.ReplaceAllString(s ,`&& \`)
		fmt.Println(a)

	}

}

