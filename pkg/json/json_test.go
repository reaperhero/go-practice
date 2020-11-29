package stand

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 解析到结构体
// 字段需要导出
// 不区分大小写，不匹配的字段会被忽略
func Test_json2struct(t *testing.T) {
	type Server struct {
		ServerName string
		ServerIP   string
	}

	type Serverslice struct {
		Servers []Server
	}
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

// 解析到interface
// 解析未知结构的JSON
func Test_json2interface(t *testing.T) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	//这个时候f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里
	//f = map[string]interface{}{
	//	"Name": "Wednesday",
	//	"Age":  6,
	//	"Parents": []interface{}{
	//		"Gomez",
	//		"Morticia",
	//	},
	//}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}

// 生成JSON
// 字段的tag是"-"，那么这个字段不会输出到JSON
// tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
// 如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
func Test_struct2json(t *testing.T) {
	type Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`

		// ServerName2 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`

		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP string `json:"serverIP,omitempty"`
	}

	type Serverslice struct {
		Servers []Server `json:"servers"`
	}
	var s Serverslice
	s.Servers = append(s.Servers, Server{ID: 1, ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: ""})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func Test_json4interface(t *testing.T) {

}

func Test_json5interface(t *testing.T) {

}