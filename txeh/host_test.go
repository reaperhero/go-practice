package txeh

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/txn2/txeh"
	"testing"
)

func TestName(t *testing.T) {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		logrus.Panic(err)
	}
	//
	//hosts.AddHost("127.100.100.100", "test")
	//hosts.AddHost("127.100.100.101", "logstash")
	//hosts.AddHosts("127.100.100.102", []string{"a", "b", "c"})
	//
	//hosts.RemoveHosts([]string{"example", "example.machine", "example.machine.example.com"})
	//hosts.RemoveHosts(strings.Fields("example2 example.machine2 example.machine.example.com2"))
	//
	//hosts.RemoveAddress("127.1.27.1")

	//removeList := []string{
	//	"127.1.27.15",
	//	"127.1.27.14",
	//	"127.1.27.13",
	//}
	//
	//hosts.RemoveAddresses(removeList)

	hfData := hosts.RenderHostsFile()

	// if you like to see what the outcome will
	// look like
	fmt.Println(hfData)

	//hosts.Save()
	// or hosts.SaveAs("./test.hosts")
}
