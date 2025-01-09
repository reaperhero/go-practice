package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

var restartChan = make(chan bool)

var watchPath = "/root/worker/end"

// 当前进程的pid
var curPid int

func main() {

	curPid = os.Getpid()
	go watch()
	go restartWeb()
	web()
}

// 监听文件改变
func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if !strings.Contains(event.Name, "end_less_new") {
					continue
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("监听到新版本文件")
					restartChan <- true
				}
			case <-watcher.Errors:
				fmt.Println("监听失败")
			}
		}
	}()

	watcher.Add(watchPath)

	<-make(chan bool)

}

// 开始重启
func restartWeb() {
	for {
		<-restartChan
		time.Sleep(time.Second * 10)
		err := exec.Command("sh", "-c", fmt.Sprintf("cd /root/worker/end && kill -9 %d && \\mv end_less_new end_less && ./end_less", curPid)).Run()
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

// 启动web
func web() {
	e := echo.New()
	e.GET("/api/pid", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%d", curPid))
	})
	e.GET("/api/version", func(c echo.Context) error {
		return c.String(http.StatusOK, "web2")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
