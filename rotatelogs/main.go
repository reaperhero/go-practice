package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"log"
	"time"
)

// https://github.com/natefinch/lumberjack/blob/v2.0/example_test.go


func main() {
	rl, _ := rotatelogs.New(
		"/data/access.log-%Y%m%d%H%M",
		rotatelogs.WithLinkName("/data/access.log"),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(7),
	)

	log.SetOutput(rl)

	log.Printf("Hello, World!")
}
