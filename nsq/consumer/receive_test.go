package consumer

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
	"testing"
)

func Test_receive_01(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1000)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("test", "consumer02", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", message.Body)
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("192.168.40.32:4150")
	if err != nil {
		log.Panic(err)
	}
	wg.Wait()
}
