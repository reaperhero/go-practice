package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	//先读取客户端的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//再回写ping...ping...ping
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1 创建一个server句柄
	s := znet.NewServer()

	//2 配置路由
	s.AddRouter(0, &PingRouter{})

	//3 开启服务
	s.Serve()
}
