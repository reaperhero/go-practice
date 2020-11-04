package signal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Test_Kill9(t *testing.T) {
	go setupCloseHandler()
	noBufch := make(chan struct{})
	<-noBufch
}

func setupCloseHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	func() {
		<-c
		fmt.Println("handle kiil - 9")
		os.Exit(0)
	}()
}
