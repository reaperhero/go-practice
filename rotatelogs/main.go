package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"log"
	"time"
)

func main() {
	rl, _ := rotatelogs.New(
		"/data/access.log-%Y%m%d%H%M",
		rotatelogs.WithLinkName("/data/access.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)

	log.SetOutput(rl)

	log.Printf("Hello, World!")
}
