package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// consul提供的是http接口,无法支持通知功能

const (
	consulAddress = "localhost:8500"
	localIp       = "192.168.40.32"
)

var (
	config       = consulapi.DefaultConfig()
	client, err  = consulapi.NewClient(config)
	registration = new(consulapi.AgentServiceRegistration)
)

// 注册
func consulRegister() {
	config.Address = consulAddress
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration.ID = "337"
	registration.Name = "service337"
	registration.Port = 8080
	registration.Tags = []string{"tag1"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

func consulDeregister() {
	client.Agent().ServiceDeregister("337")
}

func ConsulKVTest() {
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// KV, put值
	values := "test"
	key := "go-consul-test/172.16.242.129:8100"
	client.KV().Put(&consulapi.KVPair{Key: key, Flags: 0, Value: []byte(values)}, nil)

	// KV get值
	data, _, _ := client.KV().Get(key, nil)
	fmt.Println(string(data.Value))

	// KV list
	datas, _, _ := client.KV().List("go", nil)
	for _, value := range datas {
		fmt.Println(value)
	}
	keys, _, _ := client.KV().Keys("go", "", nil)
	fmt.Println(keys)
}

func main() {
	consulRegister()
	//定义一个http接口
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("you are visiting health check api"))
	})
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	ConsulKVTest()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("Got signal:", s)

	consulDeregister()
}
