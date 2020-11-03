package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var (
	config             = consulapi.DefaultConfig()
	consulAgentAddress = "localhost:8500"
	client, err        = consulapi.NewClient(config)
)

// 获取所有service
func queryAllservice() {
	config.Address = consulAgentAddress
	services, _ := client.Agent().Services()
	for s, service := range services {
		fmt.Println(s)
		fmt.Println(service)
	}
}

// 获取指定service
func queryByserviceId() {
	config.Address = consulAgentAddress
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	service, _, err := client.Agent().Service("337", nil)
	if err == nil {
		fmt.Println(service)
	}
}


// 查询指定服务的健康状态
func ConsulCheckHeath()  {
	// 创建连接consul服务配置
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 健康检查
	a, b, _ := client.Agent().AgentHealthServiceByID("337")
	fmt.Println(a)
	fmt.Println(b)
}

// 只获取健康的service
func queryByhealthservice() {
	serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	if err == nil {
		fmt.Println(serviceHealthy[0].Service)
	}

}

// 通过多个tag查询
func queryBy() {

}

func main() {
	queryAllservice()
	ConsulCheckHeath()
}
