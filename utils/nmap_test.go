package utils

import (
	"fmt"
	ping "github.com/go-ping/ping"
	"testing"
	"time"
)

func TestNampIP(t *testing.T) {
	for i := 2; i < 254; i++ {
		addr := fmt.Sprintf("172.16.82.%d", i)
		pinger, _ := ping.NewPinger(addr)

		pinger.Count = 2

		pinger.Timeout = time.Duration(2 * time.Second)

		pinger.SetPrivileged(true)

		pinger.Run() // blocks until finished

		stats := pinger.Statistics()

		if stats.PacketsRecv >= 1 {
			fmt.Printf("%s pong\n", addr)
			continue
		}
		fmt.Printf("%s time out\n", addr)
	}

}
