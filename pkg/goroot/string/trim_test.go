package string

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestTrim(t *testing.T) {
	sss := "root.parent01.child"
	fmt.Println(strings.TrimPrefix(sss, "ro"))
	fmt.Println(strings.TrimLeft(sss, "ro"))

	parse, err := url.Parse("http://link_ip/cluster?user.name=DtStack1024\"")
	if err != nil {
		return
	}
	fmt.Println(parse.Scheme + "://" + "172.16.82.176" + parse.Port() + "/")
	asdasdas()
	time.Sleep(time.Second * 10)
}

func asdasdas() {
	ss := 4
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println(ss)
	}()

}
