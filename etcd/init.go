package etcd

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)

func init() {
	cli, _ = clientv3.New(clientv3.Config{
		Endpoints:   []string{"122.112.218.250:32328"},
		DialTimeout: 5 * time.Second,
	})
	fmt.Println("connect to etcd success")
}
