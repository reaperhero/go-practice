package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
)

func Test_lease_01(t *testing.T) {
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

// 租约机制（自动过期）
func Test_lease_02(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect err:", err)
	}

	defer client.Close()

	// 获取etcd读写对象
	kv := clientv3.NewKV(cli)

	// 申请一个10秒租约
	lease := clientv3.NewLease(cli)
	leaseR, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println("lease err:", err)
		return
	}

	// 使用该租约put一个kv
	putR, err := kv.Put(context.TODO(), "/cron/lock/job1", "10001", clientv3.WithLease(leaseR.ID))
	if err != nil {
		fmt.Println("put err:", err)
		return
	}
	fmt.Println("写入成功：", putR.Header.Revision)

	// 定时查看key是否过期
	for {
		getR, err := kv.Get(context.TODO(), "/cron/lock/job1")
		if err != nil {
			fmt.Println("get err:", err)
			return
		}
		if getR.Count == 0 {
			fmt.Println("key过期")
			break
		} else {
			fmt.Println("还未过期")
			time.Sleep(2 * time.Second)
		}
	}
}

// 租约续租
func Test_lease_03(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect err:", err)
	}
	// 连接
	defer client.Close()
	if err != nil {
		return
	}

	// 获取etcd读写对象
	kv := clientv3.NewKV(cli)

	// 申请一个10秒租约
	lease := clientv3.NewLease(cli)
	leaseR, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println("lease err:", err)
		return
	}

	// 自动续租 返回值是个只读的chan，因为写入只能是etcd实现
	// context.TODO()也可以设置成context.WithTimeout(context.TODO(), 5 * time.Second)，（5 + 10)s后自动会过期
	keepChan, err := lease.KeepAlive(context.TODO(), leaseR.ID)
	if err != nil {
		fmt.Println("keep err:", err)
		return
	}
	// 启动一个协程去消费chan的应答
	go func() {
		for {
			select {
			case keepR := <-keepChan:
				if keepChan == nil { // 此时系统异常或者主动取消context
					fmt.Println("租约失效")
					goto END
				} else { // 每秒续租一次
					fmt.Println("收到自动续租应答：", keepR.ID)
				}
			}
		}
	END:
	}()

	// 使用该租约put一个kv
	putR, err := kv.Put(context.TODO(), "/cron/lock/job1", "10001", clientv3.WithLease(leaseR.ID))
	if err != nil {
		fmt.Println("put err:", err)
		return
	}
	fmt.Println("写入成功：", putR.Header.Revision)

	// 定时查看key是否过期
	for {
		getR, err := kv.Get(context.TODO(), "/cron/lock/job1")
		if err != nil {
			fmt.Println("get err:", err)
			return
		}
		if getR.Count == 0 {
			fmt.Println("key过期")
			break
		} else {
			fmt.Println("还未过期")
			time.Sleep(2 * time.Second)
		}
	}

}
