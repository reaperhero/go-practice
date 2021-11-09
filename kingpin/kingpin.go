package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

// go run kingpin.go to.go  = [kafka:9092]
// go run kingpin.go to.go --kafka.server=kafka:9091,kafka:9093 = [kafka:9091,kafka:9093]
func main() {
	uri := new([]string)
	toFlagStringsVar("kafka.server", "Address (host:port) of Kafka server.", "kafka:9092", uri)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	fmt.Println(*uri)
}
