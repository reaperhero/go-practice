package customJson

import (
	"encoding/json"
	"fmt"
	"testing"
)


var data = []byte(`
{
    "labels": "{\"alarm_level\":\"2\"}",
    "uid": "0IASnQd4z"
  }`)

func TestName(t *testing.T) {
	var d = map[string]interface{}{}
	err := json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println(err)
	}
	c,_:= NewJson([]byte(d["labels"].(string)))
	fmt.Println(d,c)
}
