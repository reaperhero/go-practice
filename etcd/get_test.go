package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_get_01(t *testing.T)  {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cli.Close()
	resp, err := cli.Get(ctx, "lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}