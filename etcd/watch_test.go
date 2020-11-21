package etcd

import (
	"context"
	"fmt"
	"testing"
)

func Test_watch_01(t *testing.T) {
	rch := cli.Watch(context.Background(), "lmh") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
