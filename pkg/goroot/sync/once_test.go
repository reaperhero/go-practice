package sync

import (
	"log"
	"sync"
	"testing"
	"time"
)

// sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。作用与 init 函数类似，但有区别。

//原理
//第一：保证变量仅被初始化一次，需要有个标志来判断变量是否已初始化过，若没有则需要初始化。
//第二：线程安全，支持并发，无疑需要互斥锁来实现
type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		config = &Config{Server: "server01", Port: 90}
		log.Println("init config")
	})
	return config
}

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	time.Sleep(time.Second)
}
