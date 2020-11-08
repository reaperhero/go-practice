package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
)

func Test_keepAlive(t *testing.T) {
	defer cli.Close()
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Put(context.TODO(), "foo", "lmh", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// 刷新ttl值
	ch, kaerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
