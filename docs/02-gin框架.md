# gin 文档


#####

- 路径参数
```
name := c.Param("name")          # /user/:name/action
```

- URL参数
```
c.DefaultQuery("name", "枯藤")       # ?后的参数不存在，返回默认值
Query("name")                       # ?后的参数不存在，返回空串
c.DefaultPostForm("type", "post")   # 表单参数不存在，返回默认值
c.PostForm("username")              # 表单参数不存在，返回空串
c.FormFile("file")                  # 上传文件，c.SaveUploadedFile(file, file.Filename) 保存
```


- 参数校验
```
type Login struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   Age      int       `form:"age" binding:"required,gt=10"`   # 大于10
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`  
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
   Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}


err := c.ShouldBindJSON(&)  # 将request的body中的数据，自动按照json格式解析到结构体
err := c.Bind(&form)        # 默认解析并绑定form格式
err := c.ShouldBindUri(&)   # ?后的uri绑定
err := c.ShouldBind(&)      # 默认绑定
```



- 中间价

```
// 定义中间
func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        fmt.Println("中间件开始执行了")
        // 设置变量到Context的key中，可以通过Get()取
        c.Set("request", "中间件")
        // 执行函数
        c.Next()
        // 中间件执行完后续的一些事情
        status := c.Writer.Status()
        fmt.Println("中间件执行完毕", status)
        t2 := time.Since(t)
        fmt.Println("time:", t2)
    }
}

Use(MiddleWare())
```



- cookie

```
cookie, err := c.Cookie("key_cookie")
c.SetCookie("key_cookie", "value_cookie", 60)  // 60秒
```


- session

[session教程](http://topgoer.com/gin%E6%A1%86%E6%9E%B6/%E4%BC%9A%E8%AF%9D%E6%8E%A7%E5%88%B6/Sessions.html)