package clickhouse

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	userMap := make(map[string]string)
	userMap["name"] = "Name"

	buf, _ := xml.Marshal(StringMap(userMap))
	fmt.Println(string(buf))

	stringMap := make(map[string]string)
	err := xml.Unmarshal(buf, (*StringMap)(&stringMap))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringMap)
}
