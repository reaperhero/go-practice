package ip

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	// 解析 IP
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// 解析地址
	ips, err := net.LookupIP("www.baidu.com")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 10 * time.Second,
			}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}
	fmt.Println(r.LookupAddr(context.Background(),"14.215.177.39"))

	ipss, _ := r.LookupHost(context.Background(), "www.google.com")
	fmt.Println(ipss)
}
