# casbin

[casbin教程](http://topgoer.com/gin%E6%A1%86%E6%9E%B6/%E5%85%B6%E4%BB%96/%E6%9D%83%E9%99%90%E7%AE%A1%E7%90%86.html)


[参考地址](https://github.com/zupzup/casbin-http-role-example/blob/main/README.md)


- example

```
GET http://localhost:9000/api/v1/hello
很遗憾,权限验证没有通过

POST http://localhost:9000/api/v1/add
增加Policy
增加成功

GET http://localhost:9000/api/v1/hello
恭喜您,权限验证通过
Hello 接收到GET请求..
```
