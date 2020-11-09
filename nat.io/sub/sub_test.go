package sub

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

var (
	natsConn *nats.Conn
)

func Test_nat_sub_01(t *testing.T) {
	natsConn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		fmt.Println(err)
		return
	}
	subj := "foo"
	natsConn.Subscribe(subj, func(msg *nats.Msg) {
		data := string(msg.Data)
		fmt.Println(data)
	})
	select {}
}
