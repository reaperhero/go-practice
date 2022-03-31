package url

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

// QueryUnescape函数用于将QueryEscape转码的字符串还原。它会把%AB改为字节0xAB，将'+'改为' '。如果有某个%后面未跟两个十六进制数字，本函数会返回错误。

// scheme://[userinfo@]host/path[?query][#fragment]

func TestUrl(t *testing.T) {
	u := url.URL{Host: "example.com", Path: "foo"}
	fmt.Println(u.IsAbs()) //false
	u.Scheme = "http"
	fmt.Println(u.IsAbs()) //true
	u.RawQuery = "x=1&y=2"
	fmt.Println(u.Query()) //map[x:[1] y:[2]]
}

func TestUrlString(t *testing.T) {
	u, err := url.Parse("https://example.org/path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Query()) // map[foo:[bar]]
	u = &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.User.Username()) //me
	password, b := u.User.Password()
	if b == false {
		log.Fatal("can not get password")
	}
	fmt.Println(password)        //pass
	fmt.Println(u.User.String()) //me:pass

}
