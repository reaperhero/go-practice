package tview

import (
	"github.com/rivo/tview"
	"testing"
)

// Terminal UI library with rich, interactive widgets — written in Golang
func TestName(t *testing.T) {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
