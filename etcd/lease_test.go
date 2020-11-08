package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
)

func Test_lease_01(t *testing.T)  {
	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, lmh2 这个key就会被移除
	_, err = cli.Put(context.TODO(), "lmh2", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
