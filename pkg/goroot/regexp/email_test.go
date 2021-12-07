package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestEmail(t *testing.T) {
	text := "My email is ccmouse@gmail.com"
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	fmt.Println(match)
}
