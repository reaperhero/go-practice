package cron

import (
	"fmt"
	"github.com/jakecoffman/cron"
	"testing"
)

func TestJake(t *testing.T) {
	cronMap := cron.New()
	cronMap.AddFunc("0 1 1 * * 0,5,6", func() {
		fmt.Println(111)
	}, "")
	cronMap.Start()
}
