package third

import (
	"fmt"
	"github.com/bitly/go-simplejson"
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
	js, _ := simplejson.NewJson(bytes)
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr, i, ms) // [1 2 3] 10 simplejson
}
