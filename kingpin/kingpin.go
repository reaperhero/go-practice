package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

// go run kingpin.go to.go  = [kafka:9092]
// go run kingpin.go to.go --kafka.server=kafka:9091,kafka:9093 = [kafka:9091,kafka:9093]
func main() {
	uri := new([]string)
	sasl := new(bool)
	work := new(int)
	toFlagStringsVar("kafka.server", "Address (host:port) of Kafka server.", "kafka:9092", uri)
	toFlagBoolVar("sasl.disable-PA-FX-FAST", "Configure the Kerberos client to not use PA_FX_FAST.", false, "false", sasl)
	toFlagIntVar("topic.workers", "Number of topic workers", 100, "100", work)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	fmt.Println(*uri,*sasl,*work)
}
