package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"testing"
)



func TestPrint01(t *testing.T)  {
	uri := new([]string)
	toFlagStringsVar("kafka.server", "Address (host:port) of Kafka server.", "kafka:9092", uri)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	fmt.Println(uri)
}