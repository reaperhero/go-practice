package markdown

import (
	"fmt"
	"github.com/russross/blackfriday"
	"testing"
)

func TestMarkdown_01(t *testing.T) {
	result := blackfriday.MarkdownCommon([]byte("# title"))
	fmt.Println(string(result))
}
