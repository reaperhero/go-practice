# echo 

## 常用方法

Create a Cookie
```
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
```


Read a Cookie
```
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}
```

自定义请求绑定
```
type User struct {
  Name  string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}
func(c echo.Context) (err error) {
  u := new(User)
  if err = c.Bind(u); err != nil {
    return
  }
  return c.JSON(http.StatusOK, u)
}
curl -X POST http:localhost:1323/users  -H 'Content-Type: application/json'  -d '{"name":"Joe","email":"joe@labstack"}'
curl -X POST http:localhost:1323/users  -d 'name=Joe' -d 'email=joe@labstack.com'
curl -X GET http:localhost:1323/users\?name\=Joe\&email\=joe@labstack.com
```

 表单
 ```
 func(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}
curl -X POST http:localhost:1323 -d 'name=Joe'
```

get参数
```
func(c echo.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, name)
})
curl -X GET http:localhost:1323\?name\=Joe
```

 路径参数
 ```
e.GET("/users/:name", func(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
})
```


 参数校验
 ```
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

CustomValidator struct {
	validator *validator.Validate
}
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```


## echo响应

Send JSON
```
u := &User{
Name:  "Jon",
Email: "jon@labstack.com",
}
return c.JSON(http.StatusOK, u)
```

Send File
```
return c.File("<PATH_TO_YOUR_FILE>")

return c.Attachment("file.txt", "ok")
```

Send Stream
```
func(c echo.Context) error {
  f, err := os.Open("<PATH_TO_IMAGE>")
  if err != nil {
    return err
  }
  return c.Stream(http.StatusOK, "image/png", f)
}
```



 JWT
 ```
https:echo.labstack.com/cookbook/jwt
```

websocket
```
https:echo.labstack.com/cookbook/websocket
```