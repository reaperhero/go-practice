package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_put_01(t *testing.T)  {
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}