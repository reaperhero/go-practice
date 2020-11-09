package pub

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

var (
	natsConn *nats.Conn
)

func Test_nat_pub_01(t *testing.T) {
	opts := []nats.Option{
		nats.Name("NATS Sample Pub"),
		nats.ReconnectWait(time.Second),
		nats.MaxReconnects(1000),
	}
	natsConn, err := nats.Connect("nats://127.0.0.1:4222", opts...)
	if err != nil {
		fmt.Println(err)
		return
	}
	subj := "foo"
	data := "data"
	natsConn.Publish(subj, []byte(data))
	natsConn.Flush()
}
