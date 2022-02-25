
https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch4-rpc/ch4-07-pbgo.html

## 环境安装
```
go get -u github.com/golang/protobuf/protoc-gen-go

go get github.com/gogo/protobuf/protoc-gen-gofast


go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/jsonpb
go get github.com/gogo/protobuf/gogoproto
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```


## 命令生成
```
protoc --go_out=plugins=grpc:. proto/service-1.proto
```

## 参数含义
```
-I: 指定proto文件所在的源目录，如果不指定，则为运行程序的当前目录
--proto_path=${GOPATH}/src/github.com/google/protobuf/src   寻找proto依赖
```


- enum

```
enum Phase {
	Pending = 1;
    Running = 0;
}
message Status {
    Phase status = 1;
}

-->

type Phase int32
const (
	Phase_Pending    Phase = 0
	Phase_Running    Phase = 1
}

```