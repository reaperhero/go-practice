package regexp

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	for _, strings := range result {
		fmt.Println(strings[1])
	}
}

func TestArgs(t *testing.T) {
	arg := "${args}"
	reg1 := regexp.MustCompile(`\${\w+}`)
	reg2 := regexp.MustCompile(`\w+`)
	vaule := reg1.FindString(arg)
	k := reg2.FindString(vaule)
	fmt.Println(k)
}

func TestRexReplace(t *testing.T) {
	var ts = []string{
		"AA && \\ 		\n",
		"AA && aa \n",
		"AA && a ",
		" && a ",
		"&& a ",
	}
	re3, _ := regexp.Compile(`&&[ ]+\\[ ]+`)

	for _, s := range ts {
		a := re3.ReplaceAllString(s, `&& \`)
		fmt.Println(a)

	}

}

func TestNameSQl(t *testing.T) {
	re := regexp.MustCompile("`db:\"(.*)\" ")

	var str = "type UserInfo struct {\n\t\tID                  int               `db:\"id\" json:\"id\"`\n\t\tUserName            string            `db:\"username\" json:\"username\"`\n\t\tPassWord            string            `db:\"password\" json:\"password\"`\n\t\tCompany             string            `db:\"company\" json:\"company\"`\n\t\tFullName            string            `db:\"full_name\" json:\"full_name\"`\n\t\tEmail               string            `db:\"email\" json:\"email\"`\n\t\tPhone               string            `db:\"phone\" json:\"phone\"`\n\t\tStatus              int               `db:\"status\" json:\"status\"`\n\t\tResetPasswordStatus int               `db:\"reset_password_status\" json:\"reset_password_status\"`\n\t\tUpdateTime          dbhelper.NullTime `db:\"update_time\" json:\"update_time\"`\n\t\tCreateTime          dbhelper.NullTime `db:\"create_time\" json:\"create_time\"`\n\t\tIsDeleted           int               `db:\"is_deleted\" json:\"is_deleted\"`\n\t}"
	splits := strings.Split(str, "\n")
	for i, s := range splits {
		result := re.FindAllStringSubmatch(s, -1)
		for _, is := range result {
			if is[1] == "id" {
				old := is[0]
				new := is[0] + fmt.Sprintf(`orm:"column(%s);pk;auto" `, is[1])
				result := strings.Replace(s, old, new, -1)
				//fmt.Println(old, new, result)
				splits[i] = result
			} else {
				old := is[0]
				new := is[0] + fmt.Sprintf(`orm:"column(%s)" `, is[1])
				result := strings.Replace(s, old, new, -1)
				//fmt.Println(old, new, result)
				splits[i] = result
			}

		}
	}
	sss := strings.Join(splits, "\n")
	fmt.Println(sss)
}

func TestNameSQlD(t *testing.T) {
	re := regexp.MustCompile("`db:\"([A-Za-z_]+)\"")
	replaceS := func(old string) string {
		splits := strings.Split(old, "\n")
		for i, s := range splits {
			result := re.FindAllStringSubmatch(s, -1)
			for _, is := range result {
				if is[1] == "id" {
					old := is[0]
					new := is[0] + fmt.Sprintf(` orm:"column(%s);pk;auto"`, is[1])
					result := strings.Replace(s, old, new, -1)
					splits[i] = result
				} else {
					old := is[0]
					new := is[0] + fmt.Sprintf(` orm:"column(%s)"`, is[1])
					result := strings.Replace(s, old, new, -1)
					//fmt.Println(old, new, result)
					splits[i] = result
				}
			}
		}
		sss := strings.Join(splits, "\n")
		return sss
	}
	filepath.Walk("/Users/chenqiangjun/gitlab/easymatrix/matrix/model", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}
		//if path == "/Users/chenqiangjun/gitlab/easymatrix/matrix/model/exec_shell_list.go" {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		s := replaceS(string(file))
		err = ioutil.WriteFile(path, []byte(s), info.Mode())
		if err != nil {
			panic(err)
		}
		//}

		return nil
	})

}

func TestNameModel(t *testing.T) {
	readStruct := func(fileName string) []string {
		var sss []string
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		splits := strings.Split(string(file), "\n")
		for i, split := range splits {
			if strings.Contains(split, "pk;auto") {
				for ii := 1; ii < 3; ii++ {
					i2 := strings.Split(splits[i-ii], " ")
					if len(i2) < 2 {
						continue
					}
					sss = append(sss, i2[1])
				}
			}
		}
		return sss
	}
	var sPls []string
	filepath.WalkDir("/Users/chenqiangjun/go/daishuyun/easymatrix/matrix/model", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		fileName := strings.TrimPrefix(path, "/Users/chenqiangjun/go/daishuyun/easymatrix/matrix/model/")
		if strings.Contains(fileName, "/") {
			return err
		}
		sss := readStruct(path)

		sPls = append(sPls, sss...)
		return err
	})
	join := strings.Join(sPls, "), new(model.")

	fmt.Println(join)
}

func bytesContain(source []byte, search byte) bool {
	for _, b := range source {
		if b == search {
			return true
		}
	}
	return false
}
func TestNameExec(t *testing.T) {
	sss := []byte{'>', '=', '<'}
	var sql = "sadsad>=<asdasda"
	var first int = -1
	for i, i2 := range sql {

		if first != -1 && !bytesContain(sss, byte(i2)) {
			fmt.Println(sql[first:i])
		}
		if first == -1 && bytesContain(sss, byte(i2)) {
			first = i
		}
	}
}
