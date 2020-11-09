package pub

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

var (
	natsConn *nats.Conn
)

func Test_nat_pub_01(t *testing.T) {
	natsConn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		fmt.Println(err)
		return
	}
	subj := "foo"
	data := "data"
	natsConn.Publish(subj, []byte(data))
	natsConn.Flush()

}
