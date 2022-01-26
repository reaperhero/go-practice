package maintest

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSingal(t *testing.T) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	select {
	case <-signals:
		fmt.Println("signal")
		os.Exit(1)
	case <-ctx.Done():
		fmt.Println("time out")
	}
	time.Sleep(time.Second * 10)
	cancel()
}
