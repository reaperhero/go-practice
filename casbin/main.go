package main

import (
	"fmt"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 要使用自己定义的数据库rbac_db,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xormadapter.NewAdapter("mysql", "root:123456@tcp(192.168.31.125:3306)/goblog?charset=utf8", true)
	e := casbin.NewEnforcer("./casbin_auth_model.conf", a)

	//从DB加载策略
	e.LoadPolicy()

	//获取router路由对象
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok := e.AddPolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})
	//删除policy
	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("删除Policy")
		if ok := e.RemovePolicy("admin", "/api/v1/hello", "GET"); !ok {
			fmt.Println("Policy不存在")
		} else {
			fmt.Println("删除成功")
		}
	})
	//获取policy
	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("查看policy")
		list := e.GetPolicy()
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
		}
	})
	//使用自定义拦截器中间件
	r.Use(Authorize(e))
	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	r.Run(":9000") //参数为空 默认监听8080端口
}

//拦截器
func Authorize(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
