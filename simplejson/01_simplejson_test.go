package third

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
	"testing"
)

func Test_simplejson_01(t *testing.T) {
	bytes := []byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`)
	js, err := simplejson.NewJson(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr, i, ms) // [1 2 3] 10 simplejson
}


func TestName(t *testing.T)  {
f := "/Users/chenqiangjun/gitlab/easymatrix/addons/alertchannel/old_alert_channel.json"
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	j,_ :=simplejson.NewFromReader(file)
	s,err := simplejson.NewJson([]byte(j.Get("Settings").MustString()))
	fmt.Println(s,err)
}
